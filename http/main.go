package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.golang.org")
	if err != nil {
		fmt.Println("http response error : ", err)
		os.Exit(1)
	}
	/* io.Copy(os.Stdout, resp.Body) */

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:",len(bs))
	return len(bs), nil
}