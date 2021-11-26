package main

import (
	"fmt"
	"net/http"
	"time"
)

//以每秒300次的速度向http://www.pkuzj.com/ 发送get请求
func main() {

	for {
		//获取当前时间
		now := time.Now()
		url := "http://www.baidu.com/" + now.String()
		//输出url到控制台
		//一直循环
		go func() {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
		}()
		time.Sleep(time.Millisecond)
	}

}
