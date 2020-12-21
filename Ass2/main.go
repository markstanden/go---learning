package main

import (
	"fmt"
)

type triangle struct {
	height float64
	base  float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tri triangle) getArea() float64 {
	return 0.5 * tri.base * tri.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	s := square{sideLength: 10}
	t := triangle{height: 10, base: 20}

	printArea(s)
	printArea(t)
}
