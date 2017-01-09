package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

//文件头信息容器
type BitmapInfoHeader struct {
	Size           uint32
	Width          int32
	Height         int32
	Places         uint16
	BitCount       uint16
	Compression    uint32
	SizeImage      uint32
	XperlsPerMeter int32
	YperlsPerMeter int32
	ClsrUsed       uint32
	ClrImportant   uint32
}

func main() {
	file, err := os.Open("image.bmp")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	//读取方法一

	//文件类型
	var headA, headB byte

	//二进制包
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)

	//文件大小
	var size uint32
	binary.Read(file, binary.LittleEndian, &size)

	var reservedA, reservedB byte
	binary.Read(file, binary.LittleEndian, &reservedA)
	binary.Read(file, binary.LittleEndian, &reservedB)

	//文件内容起始位置(图片数据)
	var offbits uint32
	binary.Read(file, binary.LittleEndian, &offbits)

	fmt.Println(headA, headB, size, reservedA, reservedB, offbits)

	//读取方法二
	infoHeader := new(BitmapInfoHeader)
	//结构体发射所有数据
	binary.Read(file, binary.LittleEndian, infoHeader)
	fmt.Println(infoHeader)
}
