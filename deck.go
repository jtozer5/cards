package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Function that creates and returns a ceck of playing cards
func newDeck() deck {
	// create a new deck list named cards
	cards := deck{}

	// Creates two slices of strings that include all card suits and values
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// Loops that append the card value and suit for each card
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	// return the newly created deck
	return cards

}

// create a method 'print' for any deck
// the method prints each card within the range of the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal 3 cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Save representation of the deck to a file on the harddrive
func saveToFile() {

}
