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
	"os"
	"strings"
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
	count := 100000
	for i := 0; i < count; i++ {
		//go fork()
		go httpDo()
		//go postFile("/home/kali/Desktop/新建文本.txt", "127.0.0.1:8080/api/student/event/edit")
	}
	for i := 0; i < count; i++ {
		<-ch
	}
}

func fork() {
	r, err := http.Get(os.Args[1])
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

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://perpetua.hexin.im/api/user/own", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "role=super; auth=; exclude_test=false; exclude_zero=true; passport=5bb9e75cb7e8759d13d0c19f3dc1b93683aa7b61")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	ch <- 0

	fmt.Println(string(body))
}
