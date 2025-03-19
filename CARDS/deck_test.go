package main

import "testing"

// TestNewDeck checks that the length of the new deck is 16.
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
