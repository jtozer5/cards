package main

import "fmt"

func main() {
	cards := []string{newCard(), "Ace of Diamonds"}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
