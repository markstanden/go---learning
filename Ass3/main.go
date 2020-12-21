package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/* 	Get the filename from the commandline argument provided,
	assign to fn assume filename is the last provided argument. */
	fn := os.Args[len(os.Args)-1]

	/* 
	Open the file, check for errors in loading file
	 */
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("--- Issues loading file ---")
	} else {
		/* If the file loads OK, copy contents to the OS's stdout */
		fmt.Printf("--- Reading file : %s ---\n\n", fn)
		io.Copy(os.Stdout, f)
		fmt.Println("\n\n--- File read complete! ---")
	}
}
