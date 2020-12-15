package main

/*
Import format package
*/
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/** seperator Constant details */
const seperator string = ","

// Create a new type 'deck'
//which is a slice of strings/** SEPERATOR Constant details */
type deck []string

/*
Creates a new deck
*/
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

/*
Deals a hand from the deck, returns the new hand, and the remainder of the deck
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*
Displays the content of the deck
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

/*
turns the deck into a single string, items separated by the separator constant
*/
func (d deck) toString() string {
	return strings.Join([]string(d), seperator)
}

func fromString(s string) deck {
	return deck(strings.Split(s, seperator))
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	bs, err := (ioutil.ReadFile(filename))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return fromString(string(bs))
}

func (d deck) shuffle(swaps int) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := 0; i <= swaps; i++ {
		/*
			'easy' way from go library
			rand.Shuffle(len(d), func(j int, k int) {
				d[j], d[k] = d[k], d[j]

			})
		*/

		/* Hard Way ;-) */
		for j := range d {
			/* rand.Intn(n int) returns a random int from [0,n) */
			k := r.Intn(len(d))
			d[j], d[k] = d[k], d[j]
		}

	}
}
