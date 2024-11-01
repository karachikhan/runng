package tables_test

import (
	"minhajuddinkhan/runng/runng/cards"
	"minhajuddinkhan/runng/runng/decks"
	"minhajuddinkhan/runng/runng/tables"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJackDealNominator(t *testing.T) {
	mockDeck := decks.Deck{
		cards.New(cards.CardAce, cards.CardKindClub),
		cards.New(cards.CardJack, cards.CardKindClub),
	}

	nominator := tables.NewFirstJackDealNominator()

	nominatedPosition, err := nominator.Nominate(tables.TablePlayerPositionSouth, mockDeck)
	assert.NoError(t, err)
	assert.Equal(t, tables.TablePlayerPositionEast, nominatedPosition)
}
