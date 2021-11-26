package main

import (
	"fmt"
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
		}()
		time.Sleep(time.Millisecond)
	}

}
