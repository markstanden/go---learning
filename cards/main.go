package main
import "fmt"
func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)


	fmt.Println("Hand: ")
	hand.print()
	fmt.Println("Remainder of Deck: ")
	remainingDeck.print()
	fmt.Println("Single string: ", hand.toString())

}
