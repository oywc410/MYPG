package main

import (
	"encoding/binary"
	"bytes"
	"fmt"
)

func main() {
	var x int32
	x = 106
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	fmt.Println(bytesBuffer.Bytes())
}
