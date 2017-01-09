package nsqd

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/nsqio/nsq/internal/clusterinfo"
	"github.com/nsqio/nsq/internal/dirlock"
	"github.com/nsqio/nsq/internal/http_api"
	"github.com/nsqio/nsq/internal/protocol"
	"github.com/nsqio/nsq/internal/statsd"
	"github.com/nsqio/nsq/internal/util"
	"github.com/nsqio/nsq/internal/version"
)

const (
	TLSNotRequired = iota
	TLSRequiredExceptHTTP
	TLSRequired
)

type errStore struct {
	err error
}

type NSQD struct {
	// 64bit atomic vars need to be first for proper alignment on 32bit platforms
	clientIDSequence int64

	sync.RWMutex

	opts atomic.Value

	dl        *dirlock.DirLock
	isLoading int32
	errValue  atomic.Value
	startTime time.Time

	topicMap map[string]*Topic

	lookupPeers atomic.Value

	tcpListener   net.Listener
	httpListener  net.Listener
	httpsListener net.Listener
	tlsConfig     *tls.Config

	poolSize int

	idChan               chan MessageID
	notifyChan           chan interface{}
	optsNotificationChan chan struct{}
	exitChan             chan int
	waitGroup            util.WaitGroupWrapper

	ci *clusterinfo.ClusterInfo
}

func New(opts *Options) *NSQD {

	//缓存消息的磁盘路径
	dataPath := opts.DataPath
	if opts.DataPath == "" {
		cwd, _ := os.Getwd()
		dataPath = cwd
	}

	n := &NSQD{
		startTime:            time.Now(),
		topicMap:             make(map[string]*Topic),
		idChan:               make(chan MessageID, 4096),
		exitChan:             make(chan int),
		notifyChan:           make(chan interface{}),
		optsNotificationChan: make(chan struct{}, 1),
		//WEB接听
		ci:                   clusterinfo.New(opts.Logger, http_api.NewClient(nil)),
		dl:                   dirlock.New(dataPath),
	}
	// 设置原子操作对象(配置信息) --???--是否还可以正常方式读写 (sync/atomic) (原子操作的效率要比锁来的高)
	n.swapOpts(opts)
	n.errValue.Store(errStore{})

	//阻止数据库文件的外来操作
	err := n.dl.Lock()
	if err != nil {
		n.logf("FATAL: --data-path=%s in use (possibly by another instance of nsqd)", dataPath)
		os.Exit(1)
	}

	//-max-deflate-level=6: 最大的压缩比率等级
	if opts.MaxDeflateLevel < 1 || opts.MaxDeflateLevel > 9 {
		n.logf("FATAL: --max-deflate-level must be [1,9]")
		os.Exit(1)
	}

	//-worker-id=0: 进程的唯一码(默认是主机名的哈希值)
	if opts.ID < 0 || opts.ID >= 1024 {
		n.logf("FATAL: --worker-id must be [0,1024)")
		os.Exit(1)
	}

	//-statsd-prefix="nsq.%s": 发送给统计keys 的前缀(%s for host replacement)
	if opts.StatsdPrefix != "" {//生成显示前缀
		var port string
		//-http-address="0.0.0.0:4151": 为 HTTP 客户端监听 <addr>:<port>

		_, port, err = net.SplitHostPort(opts.HTTPAddress)
		if err != nil {
			n.logf("ERROR: failed to parse HTTP address (%s) - %s", opts.HTTPAddress, err)
			os.Exit(1)
		}
		statsdHostKey := statsd.HostKey(net.JoinHostPort(opts.BroadcastAddress, port))
		prefixWithHost := strings.Replace(opts.StatsdPrefix, "%s", statsdHostKey, -1)
		if prefixWithHost[len(prefixWithHost)-1] != '.' {
			prefixWithHost += "."
		}
		opts.StatsdPrefix = prefixWithHost
	}

	//TLS安全文件传输协议
	if opts.TLSClientAuthPolicy != "" && opts.TLSRequired == TLSNotRequired {
		opts.TLSRequired = TLSRequired
	}

	tlsConfig, err := buildTLSConfig(opts)
	if err != nil {
		n.logf("FATAL: failed to build TLS config - %s", err)
		os.Exit(1)
	}
	if tlsConfig == nil && opts.TLSRequired != TLSNotRequired {
		n.logf("FATAL: cannot require TLS client connections without TLS key and cert")
		os.Exit(1)
	}
	n.tlsConfig = tlsConfig

	n.logf(version.String("nsqd"))
	n.logf("ID: %d", opts.ID)

	return n
}

func (n *NSQD) logf(f string, args ...interface{}) {
	if n.getOpts().Logger == nil {
		return
	}
	n.getOpts().Logger.Output(2, fmt.Sprintf(f, args...))
}

func (n *NSQD) getOpts() *Options {
	return n.opts.Load().(*Options)
}

func (n *NSQD) swapOpts(opts *Options) {
	n.opts.Store(opts)
}

func (n *NSQD) triggerOptsNotification() {
	select {
	case n.optsNotificationChan <- struct{}{}:
	default:
	}
}

func (n *NSQD) RealTCPAddr() *net.TCPAddr {
	n.RLock()
	defer n.RUnlock()
	return n.tcpListener.Addr().(*net.TCPAddr)
}

func (n *NSQD) RealHTTPAddr() *net.TCPAddr {
	n.RLock()
	defer n.RUnlock()
	return n.httpListener.Addr().(*net.TCPAddr)
}

func (n *NSQD) RealHTTPSAddr() *net.TCPAddr {
	n.RLock()
	defer n.RUnlock()
	return n.httpsListener.Addr().(*net.TCPAddr)
}

func (n *NSQD) SetHealth(err error) {
	n.errValue.Store(errStore{err: err})
}

func (n *NSQD) IsHealthy() bool {
	return n.GetError() == nil
}

func (n *NSQD) GetError() error {
	errValue := n.errValue.Load()
	return errValue.(errStore).err
}

func (n *NSQD) GetHealth() string {
	err := n.GetError()
	if err != nil {
		return fmt.Sprintf("NOK - %s", err)
	}
	return "OK"
}

func (n *NSQD) GetStartTime() time.Time {
	return n.startTime
}

func (n *NSQD) Main() {
	var httpListener net.Listener
	var httpsListener net.Listener

	ctx := &context{n}

	//监听请求(0.0.0.0:4150) TCP
	tcpListener, err := net.Listen("tcp", n.getOpts().TCPAddress)
	if err != nil {
		n.logf("FATAL: listen (%s) failed - %s", n.getOpts().TCPAddress, err)
		os.Exit(1)
	}
	n.Lock()
	n.tcpListener = tcpListener
	n.Unlock()
	tcpServer := &tcpServer{ctx: ctx}
	n.waitGroup.Wrap(func() {
		//调用监听
		protocol.TCPServer(n.tcpListener, tcpServer, n.getOpts().Logger)
	})

	//监听请求 HTTPS
	if n.tlsConfig != nil && n.getOpts().HTTPSAddress != "" {
		httpsListener, err = tls.Listen("tcp", n.getOpts().HTTPSAddress, n.tlsConfig)
		if err != nil {
			n.logf("FATAL: listen (%s) failed - %s", n.getOpts().HTTPSAddress, err)
			os.Exit(1)
		}
		n.Lock()
		n.httpsListener = httpsListener
		n.Unlock()
		httpsServer := newHTTPServer(ctx, true, true)
		n.waitGroup.Wrap(func() {
			http_api.Serve(n.httpsListener, httpsServer, "HTTPS", n.getOpts().Logger)
		})
	}
	//监听请求 HTTP
	httpListener, err = net.Listen("tcp", n.getOpts().HTTPAddress)
	if err != nil {
		n.logf("FATAL: listen (%s) failed - %s", n.getOpts().HTTPAddress, err)
		os.Exit(1)
	}
	n.Lock()
	n.httpListener = httpListener
	n.Unlock()
	httpServer := newHTTPServer(ctx, false, n.getOpts().TLSRequired == TLSRequired)
	n.waitGroup.Wrap(func() {
		http_api.Serve(n.httpListener, httpServer, "HTTP", n.getOpts().Logger)
	})

	//新建Ticker并且随时调整大小 包括channels
	n.waitGroup.Wrap(func() { n.queueScanLoop() })
	//--??--
	n.waitGroup.Wrap(func() { n.idPump() })
	//保持与lookupd通信
	n.waitGroup.Wrap(func() { n.lookupLoop() })
	//统计进程的 UDP <addr>:<port>
	if n.getOpts().StatsdAddress != "" {
		//发送统计数据
		n.waitGroup.Wrap(func() { n.statsdLoop() })
	}
}

//初始化nsqd的LoadMetadata函数使用atomic包中的方法来保证方法执行前和执行后isLoading值的改变
//解析元数据的文件得到系统中的存在的topic列表，遍历topic列表中的每个topic
func (n *NSQD) LoadMetadata() {
	//原子操作 设置开始运行
	atomic.StoreInt32(&n.isLoading, 1)
	defer atomic.StoreInt32(&n.isLoading, 0)

	fn := fmt.Sprintf(path.Join(n.getOpts().DataPath, "nsqd.%d.dat"), n.getOpts().ID)
	//启动时从磁盘中读取文件
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		if !os.IsNotExist(err) {
			n.logf("ERROR: failed to read channel metadata from %s - %s", fn, err)
		}
		return
	}

	//解析文件数据
	js, err := simplejson.NewJson(data)
	if err != nil {
		n.logf("ERROR: failed to parse metadata - %s", err)
		return
	}

	//获取话题信息
	topics, err := js.Get("topics").Array()
	if err != nil {
		n.logf("ERROR: failed to parse metadata - %s", err)
		return
	}

	for ti := range topics {
		topicJs := js.Get("topics").GetIndex(ti)

		topicName, err := topicJs.Get("name").String()
		if err != nil {
			n.logf("ERROR: failed to parse metadata - %s", err)
			return
		}
		if !protocol.IsValidTopicName(topicName) {
			n.logf("WARNING: skipping creation of invalid topic %s", topicName)
			continue
		}
		//获取话题对象 不存在则创建
		topic := n.GetTopic(topicName)

		//话题 暂停设置
		paused, _ := topicJs.Get("paused").Bool()
		if paused {
			topic.Pause()
		}

		channels, err := topicJs.Get("channels").Array()
		if err != nil {
			n.logf("ERROR: failed to parse metadata - %s", err)
			return
		}

		//获取通道信息
		for ci := range channels {
			channelJs := topicJs.Get("channels").GetIndex(ci)

			channelName, err := channelJs.Get("name").String()
			if err != nil {
				n.logf("ERROR: failed to parse metadata - %s", err)
				return
			}
			if !protocol.IsValidChannelName(channelName) {
				n.logf("WARNING: skipping creation of invalid channel %s", channelName)
				continue
			}
			channel := topic.GetChannel(channelName)
			//通道 暂停设置
			paused, _ = channelJs.Get("paused").Bool()
			if paused {
				channel.Pause()
			}
		}
	}
}

//PersistMetadata将当前的topic和channel信息写入*nsqd.%d.dat*文件中, 主要步骤是忽略#ephemeral结尾的topic和channel后将topic和channel列表json序列化后写回文件中
func (n *NSQD) PersistMetadata() error {
	// persist metadata about what topics/channels we have
	// so that upon restart we can get back to the same state
	fileName := fmt.Sprintf(path.Join(n.getOpts().DataPath, "nsqd.%d.dat"), n.getOpts().ID)
	n.logf("NSQ: persisting topic/channel metadata to %s", fileName)

	js := make(map[string]interface{})
	topics := []interface{}{}
	for _, topic := range n.topicMap {
		if topic.ephemeral {
			continue
		}
		topicData := make(map[string]interface{})
		topicData["name"] = topic.name
		topicData["paused"] = topic.IsPaused()
		channels := []interface{}{}
		topic.Lock()
		for _, channel := range topic.channelMap {
			channel.Lock()
			if channel.ephemeral {
				channel.Unlock()
				continue
			}
			channelData := make(map[string]interface{})
			channelData["name"] = channel.name
			channelData["paused"] = channel.IsPaused()
			channels = append(channels, channelData)
			channel.Unlock()
		}
		topic.Unlock()
		topicData["channels"] = channels
		topics = append(topics, topicData)
	}
	js["version"] = version.Binary
	js["topics"] = topics

	data, err := json.Marshal(&js)
	if err != nil {
		return err
	}

	tmpFileName := fmt.Sprintf("%s.%d.tmp", fileName, rand.Int())
	f, err := os.OpenFile(tmpFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		f.Close()
		return err
	}
	f.Sync()
	f.Close()

	//移动文件--???-- (为什么使用sys)
	err = atomicRename(tmpFileName, fileName)
	if err != nil {
		return err
	}

	return nil
}

func (n *NSQD) Exit() {
	if n.tcpListener != nil {
		n.tcpListener.Close()
	}

	if n.httpListener != nil {
		n.httpListener.Close()
	}

	if n.httpsListener != nil {
		n.httpsListener.Close()
	}

	n.Lock()
	err := n.PersistMetadata()
	if err != nil {
		n.logf("ERROR: failed to persist metadata - %s", err)
	}
	n.logf("NSQ: closing topics")
	for _, topic := range n.topicMap {
		topic.Close()
	}
	n.Unlock()

	// we want to do this last as it closes the idPump (if closed first it
	// could potentially starve items in process and deadlock)
	close(n.exitChan)
	n.waitGroup.Wait()

	n.dl.Unlock()
}

// GetTopic performs a thread safe operation
// to return a pointer to a Topic object (potentially new)
func (n *NSQD) GetTopic(topicName string) *Topic {
	// most likely, we already have this topic, so try read lock first.
	n.RLock()
	t, ok := n.topicMap[topicName]
	n.RUnlock()
	if ok {
		return t
	}

	n.Lock()

	t, ok = n.topicMap[topicName]
	if ok {
		n.Unlock()
		return t
	}
	deleteCallback := func(t *Topic) {
		n.DeleteExistingTopic(t.name)
	}
	t = NewTopic(topicName, &context{n}, deleteCallback)
	n.topicMap[topicName] = t

	n.logf("TOPIC(%s): created", t.name)

	// release our global nsqd lock, and switch to a more granular topic lock while we init our
	// channels from lookupd. This blocks concurrent PutMessages to this topic.
	//转换为小密度锁
	t.Lock()
	n.Unlock()

	// if using lookupd, make a blocking call to get the topics, and immediately create them.
	// this makes sure that any message received is buffered to the right channels
	// 如果启用了nsqlookupd则需要从lookupd中获取该topic的所有channel
	lookupdHTTPAddrs := n.lookupdHTTPAddrs()
	if len(lookupdHTTPAddrs) > 0 {
		channelNames, _ := n.ci.GetLookupdTopicChannels(t.name, lookupdHTTPAddrs)
		for _, channelName := range channelNames {
			if strings.HasSuffix(channelName, "#ephemeral") {
				// we don't want to pre-create ephemeral channels
				// because there isn't a client connected
				continue
			}
			t.getOrCreateChannel(channelName)
		}
	}

	t.Unlock()

	// NOTE: I would prefer for this to only happen in topic.GetChannel() but we're special
	// casing the code above so that we can control the locks such that it is impossible
	// for a message to be written to a (new) topic while we're looking up channels
	// from lookupd...
	//
	// update messagePump state
	select {
	case t.channelUpdateChan <- 1:
	case <-t.exitChan:
	}
	return t
}

// GetExistingTopic gets a topic only if it exists
func (n *NSQD) GetExistingTopic(topicName string) (*Topic, error) {
	n.RLock()
	defer n.RUnlock()
	topic, ok := n.topicMap[topicName]
	if !ok {
		return nil, errors.New("topic does not exist")
	}
	return topic, nil
}

// DeleteExistingTopic removes a topic only if it exists
func (n *NSQD) DeleteExistingTopic(topicName string) error {
	n.RLock()
	topic, ok := n.topicMap[topicName]
	if !ok {
		n.RUnlock()
		return errors.New("topic does not exist")
	}
	n.RUnlock()

	// delete empties all channels and the topic itself before closing
	// (so that we dont leave any messages around)
	//
	// we do this before removing the topic from map below (with no lock)
	// so that any incoming writes will error and not create a new topic
	// to enforce ordering
	topic.Delete()

	n.Lock()
	delete(n.topicMap, topicName)
	n.Unlock()

	return nil
}

func (n *NSQD) idPump() {
	factory := &guidFactory{}
	lastError := time.Unix(0, 0)
	//进程的唯一码(默认是主机名的哈希值)
	workerID := n.getOpts().ID
	for {
		id, err := factory.NewGUID(workerID)
		if err != nil {//触发GC
			now := time.Now()
			if now.Sub(lastError) > time.Second {
				// only print the error once/second
				n.logf("ERROR: %s", err)
				lastError = now
			}
			runtime.Gosched()
			continue
		}
		select {
		case n.idChan <- id.Hex():
		case <-n.exitChan:
			goto exit
		}
	}

exit:
	n.logf("ID: closing")
}

func (n *NSQD) Notify(v interface{}) {
	// since the in-memory metadata is incomplete,
	// should not persist metadata while loading it.
	// nsqd will call `PersistMetadata` it after loading
	persist := atomic.LoadInt32(&n.isLoading) == 0
	n.waitGroup.Wrap(func() {
		// by selecting on exitChan we guarantee that
		// we do not block exit, see issue #123
		select {
		case <-n.exitChan:
		case n.notifyChan <- v:
			if !persist {
				return
			}
			n.Lock()
			err := n.PersistMetadata()
			if err != nil {
				n.logf("ERROR: failed to persist metadata - %s", err)
			}
			n.Unlock()
		}
	})
}

// channels returns a flat slice of all channels in all topics
func (n *NSQD) channels() []*Channel {
	var channels []*Channel
	n.RLock()
	for _, t := range n.topicMap {
		t.RLock()
		for _, c := range t.channelMap {
			channels = append(channels, c)
		}
		t.RUnlock()
	}
	n.RUnlock()
	return channels
}

// resizePool adjusts the size of the pool of queueScanWorker goroutines
//
// 	1 <= pool <= min(num * 0.25, QueueScanWorkerPoolMax)
//
func (n *NSQD) resizePool(num int, workCh chan *Channel, responseCh chan bool, closeCh chan int) {
	idealPoolSize := int(float64(num) * 0.25)
	if idealPoolSize < 1 {
		idealPoolSize = 1
	} else if idealPoolSize > n.getOpts().QueueScanWorkerPoolMax {
		idealPoolSize = n.getOpts().QueueScanWorkerPoolMax
	}
	for {
		if idealPoolSize == n.poolSize {
			break
		} else if idealPoolSize < n.poolSize {
			// contract
			//缩小
			closeCh <- 1
			n.poolSize--
		} else {
			// expand
			//扩大
			n.waitGroup.Wrap(func() {
				n.queueScanWorker(workCh, responseCh, closeCh)
			})
			n.poolSize++
		}
	}
}

// queueScanWorker receives work (in the form of a channel) from queueScanLoop
// and processes the deferred and in-flight queues
//处理消息队列 (客户端发送信息,超时判断,删除队列) --???--
func (n *NSQD) queueScanWorker(workCh chan *Channel, responseCh chan bool, closeCh chan int) {
	for {
		select {
		case c := <-workCh:
			now := time.Now().UnixNano()
			dirty := false
			if c.processInFlightQueue(now) {
				dirty = true
			}
			if c.processDeferredQueue(now) {
				dirty = true
			}
			responseCh <- dirty
		case <-closeCh:
			return
		}
	}
}

// queueScanLoop runs in a single goroutine to process in-flight and deferred
// priority queues. It manages a pool of queueScanWorker (configurable max of
// QueueScanWorkerPoolMax (default: 4)) that process channels concurrently.
//
// It copies Redis's probabilistic expiration algorithm: it wakes up every
// QueueScanInterval (default: 100ms) to select a random QueueScanSelectionCount
// (default: 20) channels from a locally cached list (refreshed every
// QueueScanRefreshInterval (default: 5s)).
//
// If either of the queues had work to do the channel is considered "dirty".
//
// If QueueScanDirtyPercent (default: 25%) of the selected channels were dirty,
// the loop continues without sleep.
func (n *NSQD) queueScanLoop() {
	workCh := make(chan *Channel, n.getOpts().QueueScanSelectionCount)
	responseCh := make(chan bool, n.getOpts().QueueScanSelectionCount)
	closeCh := make(chan int)

	workTicker := time.NewTicker(n.getOpts().QueueScanInterval)
	refreshTicker := time.NewTicker(n.getOpts().QueueScanRefreshInterval)

	//获取所有通道
	channels := n.channels()
	//调整地址池大小
	n.resizePool(len(channels), workCh, responseCh, closeCh)

	for {
		select {
		case <-workTicker.C://定时触发
			if len(channels) == 0 {
				continue
			}
		case <-refreshTicker.C://定时触发
			channels = n.channels()//获取所有通道
			//--???-- 调整大小 并处理通道中的信息
			n.resizePool(len(channels), workCh, responseCh, closeCh)
			continue
		case <-n.exitChan:
			goto exit
		}

		num := n.getOpts().QueueScanSelectionCount
		if num > len(channels) {
			num = len(channels)
		}

	loop:
		for _, i := range util.UniqRands(num, len(channels)) {
			workCh <- channels[i]
		}

		numDirty := 0
		for i := 0; i < num; i++ {
			if <-responseCh {
				numDirty++
			}
		}

		if float64(numDirty)/float64(num) > n.getOpts().QueueScanDirtyPercent {
			goto loop
		}
	}

exit:
	n.logf("QUEUESCAN: closing")
	close(closeCh)
	workTicker.Stop()
	refreshTicker.Stop()
}

func buildTLSConfig(opts *Options) (*tls.Config, error) {
	var tlsConfig *tls.Config

	if opts.TLSCert == "" && opts.TLSKey == "" {
		return nil, nil
	}

	tlsClientAuthPolicy := tls.VerifyClientCertIfGiven

	cert, err := tls.LoadX509KeyPair(opts.TLSCert, opts.TLSKey)
	if err != nil {
		return nil, err
	}
	switch opts.TLSClientAuthPolicy {
	case "require":
		tlsClientAuthPolicy = tls.RequireAnyClientCert
	case "require-verify":
		tlsClientAuthPolicy = tls.RequireAndVerifyClientCert
	default:
		tlsClientAuthPolicy = tls.NoClientCert
	}

	tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tlsClientAuthPolicy,
		MinVersion:   opts.TLSMinVersion,
		MaxVersion:   tls.VersionTLS12, // enable TLS_FALLBACK_SCSV prior to Go 1.5: https://go-review.googlesource.com/#/c/1776/
	}

	if opts.TLSRootCAFile != "" {
		tlsCertPool := x509.NewCertPool()
		caCertFile, err := ioutil.ReadFile(opts.TLSRootCAFile)
		if err != nil {
			return nil, err
		}
		if !tlsCertPool.AppendCertsFromPEM(caCertFile) {
			return nil, errors.New("failed to append certificate to pool")
		}
		tlsConfig.ClientCAs = tlsCertPool
	}

	tlsConfig.BuildNameToCertificate()

	return tlsConfig, nil
}

func (n *NSQD) IsAuthEnabled() bool {
	return len(n.getOpts().AuthHTTPAddresses) != 0
}
