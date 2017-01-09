package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`      //`` 来补助反射
	Version     string   `xml:"version,attr"` // xml: ,attr 代表属性名
	Svs         []server `xml:"server"`       //也可以这样写 XXX>XXX>XXX
	Description string   `xml:",innerxml"`    //Unmarshal将会将此字段所对应的元素内所有内嵌的原始xml累加到此字段上

	//,any 子元素在不满足其他的规则的时候就会匹配到这个字段
	//,commentsXML元素包含一条或者多条注释，那么这些注释将被累加到第一个tag含有”,comments”的字段
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	//从第n行读入数据 至 结束
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("errorL %v", err)
		return
	}
	v := Recurlyservers{}
	//映射xml内容到指定对象中 v还可以为array或者slice
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}
