package entrance

import (
	enutil "appInternal/util"
	"github.com/bitly/go-nsq"
	"log"
)

type Entrance struct {
	Log        *log.Logger
	Configdata *ConfigData
	waitGroup  enutil.WaitGroupWrapper
	message      *message
	stopChan	chan struct{}
}

func MakeEntrance() *Entrance {
	return &Entrance{}
}

func (entrance *Entrance) Init() {
	if entrance.Configdata.ChannelMux == 0 {
		entrance.Configdata.ChannelMux = 1000
	}
}

func (entrance *Entrance) Main() error {

	entrance.message = NewMessage(*entrance.Configdata)

	go func() {
		err := entrance.StartHttp()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	if len(entrance.Configdata.NsqAddrs) > 0 {
		entrance.publishVotes(entrance.Configdata.NsqAddrs[0], entrance.message)
	}

	return nil
}

func (entrance *Entrance) StartHttp() error {
	httpServer := NewHttpServer(*entrance.Configdata, *entrance.message)
	return httpServer.Start()
}

func (entrance *Entrance) publishVotes(host string, message *message) <-chan struct{} {
	stopchan := make(chan struct{}, 1)
	pub, _ := nsq.NewProducer(host, nsq.NewConfig())
	go func() {
		//通道被关闭时跳出循环 停止程序
		for messageStr := range message.getMessageChan() {
			err := pub.Publish("votes", []byte(messageStr))
			if err != nil {
				entrance.message.addOutMessage(messageStr)
				log.Println(err)
			}
		}
		pub.Stop()
		stopchan <- struct{}{}
	}()
	return stopchan
}

func (entrance *Entrance) Stop() error {
	entrance.stopChan <- struct{}{}
	entrance.message.Close()

	return nil
}
