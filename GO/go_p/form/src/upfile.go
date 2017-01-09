package main
import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)
func postFile(filename string, targetUrl string) error {
	
	//缓存流
	bodyBuf := &bytes.Buffer{}
	//准备写入缓存流
	bodyWriter := multipart.NewWriter(bodyBuf)
	
	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	//iocopy 写入缓存流
	_, err = io.Copy(fileWriter, fh)
		if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	
	//发送请求
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	
	defer resp.Body.Close()
	
	//返回信息
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	
	return nil
}
// sample usage
func main() {
	target_url := "http://localhost:9090/upload"
	filename := "./up_test_file.txt"
	postFile(filename, target_url)
}