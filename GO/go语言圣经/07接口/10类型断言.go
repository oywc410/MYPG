package main
import (
	"io"
	"os"
	"bytes"
	"fmt"
)

func main() {
	var w io.Writer
	w = os.Stdout

	//进行类型断言
	osFile := w.(*os.File)

	bytesBuffer := w.(*bytes.Buffer)//此处抛出错误

	//不抛出异常方法
	bytesBuffer, ok := w.(*bytes.Buffer)

	_ = osFile
	_ = bytesBuffer
	_ = ok

	switch w.(type) {
	case nil:
	case int, uint:
	case bool:
	case string:
	default:
	}
}

func sqlQuote(x interface{}) string {
	//重用变量名
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	}
}
