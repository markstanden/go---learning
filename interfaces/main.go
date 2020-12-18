package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}

func main() {
eb := englishBot{}
sb := spanishBot{}
printEbGreeting(eb)
printSbGreeting(sb)
}

func (_ englishBot) getGreeting() string {
	return "Hello my friends!"
}

func (_ spanishBot) getGreeting() string {
	return "Hola mon amigo!"
}

func printEbGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printSbGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
