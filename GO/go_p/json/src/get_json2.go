package main

import (
	"encoding/json"
	"os"
)

type Server struct {

	//JSON对象只支持string作为key，所以要编码一个map，那么必须是map[string]T这种类型(T是Go语言中任意的类型) !!!!

	// ID 不会导出到JSON中
	ID int `json:"-"`

	// ServerName2 的值会进行二次JSON编码
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"` //表示为纯文本输出

	// 如果 ServerIP 为空，则不输出到JSON串中
	ServerIP string `json:"serverIP,omitempty"`
}

func main() {

	s := Server{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:    ``,
	}

	//输出JSON格式
	b, _ := json.Marshal(s)
	os.Stdout.Write(b)
}
