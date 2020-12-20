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

if serialCheckAll(urls) {
	fmt.Println("-- All OK --")
} else {
	fmt.Println("-- We have an issue --")
}
}

func serialCheckAll(urls []string) bool {
	defer timer(time.Now(), "Serial URL Checking Function")

	errLog := true
	for _, u := range urls {
		go e := checkLink(u)
		if e != nil {
			errLog = false
		}
	}
	if errLog {return true}
	return false
}

func checkLink(url string) error {
	_, e := http.Get(url)
	if e != nil {
		fmt.Printf("connection to %v failed, link is down\n", url)
		return e
	}
		fmt.Printf("connection to %v successful, link is up\n", url)
		return nil
}

func timer(st time.Time, n string) {
    et := time.Since(st)
    fmt.Printf("\n%s ran for %s seconds\n", n, et)
}