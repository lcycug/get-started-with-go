package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected 16 but got %v", len(d))
	}
}

func TestSave2FileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.save2File("_decktesting")

	newDeck := newDeckFromFile("_decktesting")
	if len(newDeck) != 16 {
		t.Errorf("Expected 16 but got %v", len(deck))
	}
	os.Remove("_decktesting")
}
