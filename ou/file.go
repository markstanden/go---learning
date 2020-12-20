package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/* Reads data from file, returns the read data as a byte slice. */
func bsFromFile(filename string) []byte {
	bs, err := (ioutil.ReadFile(filename))
	if err != nil {

		fmt.Println("File read error:", err)
		os.Exit(1)
	}
	return bs
}
