package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	for {

		now := time.Now()
		url := "http://www.pkuzj.com/" + now.String()
		go func() {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(body))
		}()
		time.Sleep(time.Millisecond)
	}

}
