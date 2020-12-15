package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle(100)
	hand, remainingDeck := deal(cards, 5)

	fmt.Println("Hand: ")
	hand.print()
	fmt.Println("Remainder of Deck: ")
	remainingDeck.print()
	fmt.Println("Single string: ", hand.toString())

	fmt.Println("Saving...")
	cards.saveToFile("myCards")

	fmt.Println("loading...")
	readFromFile("myCards").print()

}
