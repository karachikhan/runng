package cards_test

import (
	"minhajuddinkhan/runng/runng/cards"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	aceOfClub := cards.New(cards.CardAce, cards.CardKindClub)
	assert.Equal(t, aceOfClub.Number(), cards.CardAce)
}
