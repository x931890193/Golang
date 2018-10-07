/*
Create on: 2018-10-05 下午6:12
author: sato
mail: ysudqfs@163.com
life is short, you need go go go
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

var ch chan int = make(chan int)

//func init() { // init 初始化
//	if len(os.Args) != 2 {
//		fmt.Println("Need args")
//		os.Exit(-1)
//	}
//}

// 函数入口
func main() {
	// 获得响应r, err := http.Get(os.Args[1])
	//	if err != nil{
	//		fmt.Println(err)
	//		return
	//	}
	//	// 从body中赋值的stdout 标准输出
	//	io.Copy(os.Stdout, r.Body)
	//	if err := r.Body.Close(); err != nil{
	//		fmt.Println(err)
	//	}
	count := 10
	for i := 0; i < count; i++ {
		//go fork()
		go postFile("/home/kali/Desktop/IMG_0001.JPG", "http://69.171.75.147:3306/baidu_if")
	}
	for i := 0; i < count; i++ {
		<-ch
	}
}

func fork() {
	r, err := http.PostForm(os.Args[1], url.Values{"key": {"Value"}, "id": {"123"}})
	//http.POSt
	if err != nil {
		fmt.Println(666, err)

		return
	}
	// 从body中赋值的stdout 标准输出
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
	ch <- 0
}

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("img", filename)
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
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	ch <- 0
	return nil

}
