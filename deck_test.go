package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, instead got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Something has gone wrong: First card is not 'Ace of Spades', it is %v", d[0])
	}

	if d[51] != "King of Clubs" {
		t.Errorf("Something has gone wrong: Last card is not 'King of Clubs', it is %v", d[51])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
