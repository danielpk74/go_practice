package main

import (
	"os"
	"testing"
)

var d = newDeck()

func TestNewDeck(t *testing.T) {
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but gut %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d.saveToFile("_desktesting")
	loaded := newDeckFromFile("_desktesting")
	if len(loaded) != 16 {
		t.Errorf("Expected a deck with a length of 16, but got %v", len(loaded))
	}
	os.Remove("_desktesting")
}
