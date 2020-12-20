package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.co.uk",
		"http://microsoft.com",
		"http://scratch.mit.edu",
		"http://open.ac.uk",
		"http://amazon.com",
		"http://github.com",
	}
	serialCheckAll(urls)
}

func serialCheckAll(urls []string) {
	c := make(chan string)

	defer timer(time.Now(), "Serial URL Checking Function")

	for _, u := range urls {
		go checkLink(u, c)
	}
	for link := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(link)

	}

}

func checkLink(url string, c chan string) {
	_, e := http.Get(url)
	if e != nil {
		fmt.Printf("connection to %v failed, link is down\n", url)
		c <- url
		return
	}
	fmt.Printf("connection to %v successful, link is up\n", url)
	c <- url
	return
}

func timer(st time.Time, n string) {
	et := time.Since(st)
	fmt.Printf("\n%s ran for %s seconds\n", n, et)
}
