package main

import (
	"fmt"
)

type colorList map[string]string

func main() {

	colors := colorList {
		"red":   "#FF0000",
		"green": "#4BF745",
		"blue":  "#233FF2",
		"white": "#FFFFFF",
	}

/* 	printMap(colors) */
	colors.print()
}

func printMap(c map[string]string) {
	for colorKey, hexValue := range c {
		fmt.Println("------------------------")
		fmt.Println("Colour: \t", colorKey)
		fmt.Println("Hex value: \t", hexValue)
	}
}

func (c colorList) print() {
	for colorKey, hexValue := range c {
		fmt.Println("------------------------")
		fmt.Println("Colour: \t", colorKey)
		fmt.Println("Hex value: \t", hexValue)
	}
}
