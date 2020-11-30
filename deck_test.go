package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected Length of 16 but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of  but got: %v", d[0])
	}

	if d[len(d)-1] != "Four of Diamonds" {
		t.Errorf("Expected last card of  but got: %v", d[len(d)-1])
	}

}

func TestSaveToFileAndSaveToFile(t *testing.T) {
	os.Remove("_testingDeck")
	d := newDeck()
	d.saveToFile("_testingDeck")
	loadedDeck := newDeckFromFile("_testingDeck")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected Length of 16 but got: %v", len(loadedDeck))
	}
	os.Remove("_testingDeck")
}
