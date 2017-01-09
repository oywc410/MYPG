package main
import (
	"net"
	"time"
	"io"
	"github.com/matryer/go-oauth/oauth"
	"github.com/joeshaw/envdecode"
	"log"
	"sync"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"encoding/json"
	"github.com/bitly/go-nsq"
)

var conn net.Conn

func dial(netw, addr string) (net.Conn, error) {
	if conn != nil {
		conn.Close()
		conn = nil
	}

	//设置超时
	netc, err := net.DialTimeout(netw, addr, 5*time.Second)

	if err != nil {
		return nil, err
	}

	conn = netc

	return netc, nil
}

var reader io.ReadCloser

func closeConn() {

	if conn != nil {
		conn.Close()
	}

	if reader != nil {
		reader.Close()
	}
}

var (
	authClient *oauth.Client
	creds *oauth.Credentials
)

//执行认证
func setupTwitterAuth() {

	var ts struct {
		ConsumerKey    string `env:"SP_TWITTER_KEY,required"`
		ConsumerSecret string `env:"SP_TWITTER_SECRET,required"`
		AccessToken    string `env:"SP_TWITTER_ACCESSTOKEN,required"`
		AccessSecret   string `env:"SP_TWITTER_ACCESSSECRET,required"`
	}

	//envdecode 解析获取 环境变量
	if err := envdecode.Decode(&ts); err != nil {
		log.Fatalln(err)
	}

	creds = &oauth.Credentials{
		Token: ts.AccessToken,
		Secret: ts.AccessSecret,
	}

	authClient = &oauth.Client{
		Credentials: oauth.Credentials{
			Token: ts.ConsumerKey,
			Secret: ts.ConsumerSecret,
		},
	}
}

var (
	authSetupOnce sync.Once
	httpClient *http.Client
)

//发送HTTP请求(长链接)
func makeRequest(req *http.Request, params url.Values) (*http.Response, error) {

	//单列模式(只执行一次)
	authSetupOnce.Do(func() {
		setupTwitterAuth()
		httpClient = &http.Client{
			Transport: &http.Transport{
				Dial: dial,
			},
		}
	})

	formEnc := params.Encode()
	req.Header.Set("Authorization", authClient.AuthorizationHeader(creds, "POST", req.URL, params))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(formEnc)))

	return httpClient.Do(req)

}

type tweet struct {
	Text string
}

//解析返回值
func readFromTwitter(votes chan <- string) {
	options, err := loadOptions()
	if err != nil {
		log.Println("选项读取失败")
		return
	}
	u, err := url.Parse("https://stream.twitter.com/1.1/statuses/filter.json")
	if err != nil {
		log.Println("URL解析失败")
		return
	}
	query := make(url.Values)
	query.Set("track", strings.Join(options, ","))
	//生成HTTP请求内容
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(query.Encode()))
	if err != nil {
		log.Println("检索请求生成失败:", err)
		return
	}

	resp, err := makeRequest(req, query);
	if err != nil {
		log.Println("检索请求失败", err)
		return
	}
	reader = resp.Body
	decoder := json.NewDecoder(reader)

	//堵塞不断读取信息 直到失败
	for {
		var tweet tweet
		if err := decoder.Decode(&tweet); err != nil {
			break
		}
		for _, option := range options {
			if strings.Contains(strings.ToLower(tweet.Text), strings.ToLower(option)) {
				log.Println("投票:", option)
				votes <- option
			}
		}
	}
}

//发起twitter链接 (失败时将尝试重新链接)
func startTwitterStream(stopchan <- chan struct{}, votes chan <- string) <- chan struct{} {
	stoppedchan := make(chan struct{}, 1)
	go func() {
		defer func() {
			stoppedchan <- struct{}{}
		}()
		for {
			select {
			case <- stopchan:
				log.Println("与Twitter的链接结束")
				return
			default:
				log.Println("与Twitter链接")
				readFromTwitter(votes)
				log.Println("(待机中)")
				time.Sleep(10 * time.Second)
			}
		}
	}()

	return stoppedchan
}

//向NSQ中发布
func publishVotes(votes <- chan string) <- chan struct{} {
	stopchan := make(chan struct{}, 1)
	pub, _ := nsq.NewProducer("localhost:6150", nsq.NewConfig())
	go func() {
		//通道被关闭时跳出循环 停止程序
		for vote := range votes {
			pub.Publish("votes", []byte(vote))
		}
		log.Println("Publisher:终止中")
		pub.Stop()
		log.Println("Pubilsher:终止")
		stopchan <- struct{}{}
	}()
	return stopchan
}