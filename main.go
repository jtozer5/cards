package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("This is my hand: ")
	hand.print()
	fmt.Println("This is the remaining deck: ")
	remainingDeck.print()
}
