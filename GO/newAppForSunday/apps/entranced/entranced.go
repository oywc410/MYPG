package main

import (
	envalus "appInternal/values"
	"entrance"
	"flag"
	"github.com/judwhite/go-svc/svc"
	"log"
	"os"
	"path/filepath"
)

type entranced struct {
	entrance *entrance.Entrance
}

func main() {
	prg := &entranced{}
	//注册运行windows服务svc.Service接口
	if err := svc.Run(prg); err != nil {
		log.Fatal(err)
	}
}

func (ed *entranced) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	log.Println("init")
	return nil
}

func (ed *entranced) Start() error {
	log.Println("start")

	flagSet := entrancedFlagset()
	flagSet.Parse(os.Args[1:])
	err := verificationConfig(flagSet)
	if err != nil {
		log.Fatalln(err)
	}

	ed.entrance = entrance.MakeEntrance()
	ed.loadConfig(flagSet)
	err = ed.entrance.Main()

	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func (ed *entranced) Stop() error {
	log.Println("stop")
	return nil
}

func entrancedFlagset() *flag.FlagSet {
	flagSet := flag.NewFlagSet("entranced", flag.ExitOnError)
	flagSet.String("http-address", "0.0.0.0:5001", "<addr>:<port> to listen on for HTTP clients")
	flagSet.Int("channel-mux", 1000, "channel Max")
	flagSet.Int("out-channel-mux", 1000, "out channel Max")
	nsqdAddrs := envalus.StringArray{}
	flagSet.Var(&nsqdAddrs, "nsqd-tcp-address", "nsqd TCP address (may be given multiple times)")

	return flagSet
}

func verificationConfig(flagSet *flag.FlagSet) error {
	return nil
}

func (ed *entranced) loadConfig(flagSet *flag.FlagSet) {
	ed.entrance.Configdata = &entrance.ConfigData{
		HttpAddr:   flagSet.Lookup("http-address").Value.(flag.Getter).Get().(string),
		ChannelMux: flagSet.Lookup("channel-mux").Value.(flag.Getter).Get().(int),
		OutChangelMux: flagSet.Lookup("out-channel-mux").Value.(flag.Getter).Get().(int),
		NsqAddrs:   ([]string)(*flagSet.Lookup("nsqd-tcp-address").Value.(*envalus.StringArray)),
	}
}
