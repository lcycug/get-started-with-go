package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://baidu.com",
		"https://douban.com",
		"https://golang.cn",
		"https://cn.bing.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {

		go func(l string) {
			time.Sleep(time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be crashed!")
		c <- link
		return
	}

	fmt.Println(link, "goes well.")
	c <- link
}
