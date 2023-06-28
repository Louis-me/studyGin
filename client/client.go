package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 1、测试post请求的结构体
func TestPostLogin() {
	reqMap := map[string]interface{}{
		"name":     "test1",
		"password": "123456",
	}

	body, err := json.Marshal(reqMap)
	if err != nil {
		fmt.Println("TestPostReq json.Marshal err:", err)
		return
	}

	url := "http://127.0.0.1:8000/login"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("TestPostReq http.NewRequest err:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second} // 设置请求超时时长5s
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("TestPostReq http.DefaultClient.Do() err: ", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("TestPostReq ioutil.ReadAll() err: ", err)
		return
	}
	fmt.Println("respBody: ", string(respBody))

	rsp := make(map[string]interface{})
	err = json.Unmarshal(respBody, &rsp)
	if err != nil {
		fmt.Println("TestPostReq json.Unmarshal() err: ", err)
		return
	}
	fmt.Printf("rsp: %+v", rsp)

	// 最后经过字段筛选后，再序列化成json格式即可
	// result, err := json.Marshal(rsp)
	// if err != nil {
	// 	fmt.Println("TestPostReq json.Marrshal() err2: ", err)
	// 	return
	// }
	// fmt.Println(string(result))
}
func TestGetUserList() {
	resp, err := http.Get("http://127.0.0.1:8000/GetUserList")
	if err != nil {
		fmt.Println("http get error:", err)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024) // 创建一个切片，用例接受服务器返回的数据
	for {
		m, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("resp.Body.Read error:", err)
			return
		}
		res := string(buf[:m])
		fmt.Println("get server content,", res)
		break

	}
}

func main() {
	TestPostLogin()
	TestGetUserList()
}
