package main

import "fmt"

func main() {
	var intSlice []int = newIntSlice(10)
	for _, i := range intSlice {
		even := ""
		if isEven(i) {
			even = "even"
		} else {
			even = "odd"
		}
		fmt.Printf("%v is %v\n", i, even)
	}

}

func newIntSlice(max int) []int {
	intSlice := []int{}
	for i := 0; i <= max; i++ {
		intSlice = append(intSlice, i)
	}
	return intSlice
}

func isEven(i int) bool {
	if i%2 != 0 {
		return false
	}
	return true
}
