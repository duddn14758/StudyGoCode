package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	// 초기화 되지 않은 map에는 인자를 추가시킬 수 없다.
	//var results = map[string]string{}
	//var results = make(map[string]string)

	c := make(chan string)

	urls := []string{
		"https://www.naver.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.amazon.com",
	}

	fmt.Println("Waiting for checking...")
	for _, url := range urls {
		//result := "OK"
		go hitUrl(url, c)
		/*
			err := hitUrl(url)
			if err != nil {
				result = "FAILED"
			}
			results[url] = result
		*/
	}
	//time.Sleep(time.Second * 10)

	for range urls {
		fmt.Println(<-c)
		//fmt.Println(url, result)
	}
}

func hitUrl(url string, c chan string) error {
	fmt.Println("CHECKING :", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- "FAILED"
		return errRequestFailed
	}
	c <- url + " OK"
	return nil
}
