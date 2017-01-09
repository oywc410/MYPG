package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"../trace"
	"github.com/stretchr/objx"
)

type room struct {
	//广播房间内用户使用
	forward chan *message
	//进入房间时使用的通道
	join chan *client
	//推出房间时使用的通道
	leave chan *client
	//正在房间里的用户
	clients map[*client]bool
	//使用trace包来记录日志
	tracer trace.Tracer
}

func newRoom() *room {
	return &room {
		forward:make(chan *message),
		join:make(chan *client),
		leave:make(chan *client),
		clients:make(map[*client]bool),
		tracer:trace.Off(),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			//参加
			r.clients[client] = true
			r.tracer.Trace("新用户加入")
		case client := <-r.leave:
			//退出
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("用户退出")
		case msg := <-r.forward:
			r.tracer.Trace("接收信息:", msg.Message)
			//实行广播
			for client := range r.clients {
				select {
				case client.send <- msg:
					//送信
					r.tracer.Trace(" == 用户信息已送达")
				default:
					//送信失败
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" == 信息发送失败,关闭客户端连接")
				}
			}
		}
	}
}

const (
	soketBufferSize   = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: soketBufferSize, WriteBufferSize: soketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
	}

	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatal("cookie获取失败:", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan *message, messageBufferSize),
		room:   r,
		userData: objx.MustFromBase64(authCookie.Value),
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
