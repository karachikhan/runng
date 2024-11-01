package decks_test

import (
	"minhajuddinkhan/runng/runng/cards"
	"minhajuddinkhan/runng/runng/decks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck := decks.New()
	assert.Len(t, deck, 52)

	// Test the first card
	firstCard := deck[0]
	assert.Equal(t, cards.CardTwo, firstCard.Number())
}
