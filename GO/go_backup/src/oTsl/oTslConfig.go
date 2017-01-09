package oTsl

import (
	"crypto/rand"
	"crypto/tls"
	"io"
)

type oTslConfig struct {
	crtPath  string
	keyPath  string
	rand     io.Reader
	listenT  string
	listenIp string
	command  *Command
}

func NewConfig() *oTslConfig {
	return &oTslConfig{}
}

func (obj *oTslConfig) SetListenT(str string) {
	obj.listenT = str
}

func (obj *oTslConfig) GetListenT() string {
	return obj.listenT
}

func (obj *oTslConfig) SetListenIp(str string) {
	obj.listenIp = str
}

func (obj *oTslConfig) GetListenIp() string {
	return obj.listenIp
}

func (obj *oTslConfig) SetCrtPath(path string) {
	obj.crtPath = path
}

func (obj *oTslConfig) GetCrtPath() string {
	return obj.crtPath
}

func (obj *oTslConfig) SetKeyPath(path string) {
	obj.keyPath = path
}

func (obj *oTslConfig) GetKeyPath() string {
	return obj.keyPath
}

func (obj *oTslConfig) GetRank() io.Reader {
	if obj.rand == nil {
		obj.rand = rand.Reader
	}

	return obj.rand
}

func (obj *oTslConfig) SetCommand(command *Command) {
	obj.command = command
}

func (obj *oTslConfig) GetCommand() *Command {
	return obj.command
}

/*サーバ側config対象を作成する*/
func (obj *oTslConfig) ToTslConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(obj.GetCrtPath(), obj.GetKeyPath())

	if err != nil {
		return nil, err
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = obj.GetRank()
	return config, nil
}

/*フロント側config対象を作成する*/
func (obj *oTslConfig) ToClientTslConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(obj.GetCrtPath(), obj.GetKeyPath())

	if err != nil {
		return nil, err
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	return config, err
}
