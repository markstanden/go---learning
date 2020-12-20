package main

import (
	"fmt"
)

type triangle struct {
	height float64
	width float64
}

type square struct {
	height float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.height * s.height
}

func (t triangle) getArea() float64 {
	return 0.5*t.width*t.height
}

func (shape) printArea() {
	fmt.Println(s.getArea())
}

func main() {
	fmt.Println("Hello")
}