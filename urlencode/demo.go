package main

import (
	"fmt"
	"net/url"
)

func main() {
	urltest := "http://www.baidu.com/s?wd=连衣裙"
	fmt.Println(urltest)
	encodeurl := url.QueryEscape(urltest)
	fmt.Println(encodeurl)
	decodeurl, err := url.QueryUnescape(encodeurl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decodeurl)
}
