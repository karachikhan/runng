package tables_test

import (
	"minhajuddinkhan/runng/runng/decks"
	"minhajuddinkhan/runng/runng/tables"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerTakeFromDeckAndPutOnTop(t *testing.T) {
	p := tables.NewPlayer()
	deck := decks.New()
	cards, deck := p.TakeFrom(13, 26, deck)
	assert.Len(t, cards, 13)
	assert.Len(t, deck, 39)
	deck = p.PutOnDeck(cards, deck)
	assert.Len(t, deck, 52)
	assert.Equal(t, cards[0], deck[0])
}

func TestPlayerAsShuffler(t *testing.T) {
	p := tables.NewPlayer()
	deck := decks.New()
	firstDraw := deck[0]
	assert.Equal(t, firstDraw, deck[0])
	deck = p.Shuffle(deck)
	firstDrawAfterShuffle := deck[0]
	assert.NotEqual(t, firstDraw, firstDrawAfterShuffle)
}
