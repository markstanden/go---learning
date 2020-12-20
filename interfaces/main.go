package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hello my friends!"
}

func (spanishBot) getGreeting() string {
	return "Hola mon amigo!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/*
func printEbGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printSbGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
*/
