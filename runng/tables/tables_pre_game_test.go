package tables_test

import (
	"minhajuddinkhan/runng/runng/cards"
	"minhajuddinkhan/runng/runng/decks"
	"minhajuddinkhan/runng/runng/tables"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableWithDealer(t *testing.T) {
	p1 := tables.NewPlayer()
	p2 := tables.NewPlayer()
	p3 := tables.NewPlayer()
	p4 := tables.NewPlayer()
	table := tables.NewTable()
	south, err := p1.Join(table, tables.TablePlayerPositionSouth)
	assert.NoError(t, err)
	east, err := p2.Join(table, tables.TablePlayerPositionEast)
	assert.NoError(t, err)
	north, err := p3.Join(table, tables.TablePlayerPositionNorth)
	assert.NoError(t, err)
	west, err := p4.Join(table, tables.TablePlayerPositionWest)
	assert.NoError(t, err)

	deck := decks.New()
	nominator := tables.NewFirstJackDealNominator()
	dealerPosition, err := nominator.Nominate(tables.TablePlayerPositionEast, deck)
	assert.NoError(t, err)
	assert.Equal(t, tables.TablePlayerPositionNorth, dealerPosition)

	tableWithDealer := table.WithDealer(dealerPosition, deck)
	tableAwaitingColor, err := tableWithDealer.Deal()
	assert.NoError(t, err)

	colorNominator := tables.NewColorNominator(west)
	assert.Error(t, colorNominator.Nominate(cards.CardKindHeart))
	go colorNominator.Nominate(cards.CardKindSpade)
	tableWithColor, err := tableAwaitingColor.AwaitColor(colorNominator)
	assert.NoError(t, err)

	assert.True(t, west.HasCard(cards.New(cards.CardTwo, cards.CardKindSpade)))
	assert.True(t, west.HasCard(cards.New(cards.CardThree, cards.CardKindSpade)))
	assert.True(t, west.HasCard(cards.New(cards.CardFour, cards.CardKindSpade)))
	assert.True(t, west.HasCard(cards.New(cards.CardFive, cards.CardKindSpade)))
	assert.True(t, west.HasCard(cards.New(cards.CardSix, cards.CardKindSpade)))

	assert.True(t, south.HasCard(cards.New(cards.CardSeven, cards.CardKindSpade)))
	assert.True(t, south.HasCard(cards.New(cards.CardEight, cards.CardKindSpade)))
	assert.True(t, south.HasCard(cards.New(cards.CardNine, cards.CardKindSpade)))
	assert.True(t, south.HasCard(cards.New(cards.CardTen, cards.CardKindSpade)))
	assert.True(t, south.HasCard(cards.New(cards.CardJack, cards.CardKindSpade)))

	assert.True(t, east.HasCard(cards.New(cards.CardQueen, cards.CardKindSpade)))
	assert.True(t, east.HasCard(cards.New(cards.CardKing, cards.CardKindSpade)))
	assert.True(t, east.HasCard(cards.New(cards.CardAce, cards.CardKindSpade)))
	assert.True(t, east.HasCard(cards.New(cards.CardTwo, cards.CardKindHeart)))
	assert.True(t, east.HasCard(cards.New(cards.CardThree, cards.CardKindHeart)))

	assert.True(t, north.HasCard(cards.New(cards.CardFour, cards.CardKindHeart)))
	assert.True(t, north.HasCard(cards.New(cards.CardFive, cards.CardKindHeart)))
	assert.True(t, north.HasCard(cards.New(cards.CardSix, cards.CardKindHeart)))
	assert.True(t, north.HasCard(cards.New(cards.CardSeven, cards.CardKindHeart)))
	assert.True(t, north.HasCard(cards.New(cards.CardEight, cards.CardKindHeart)))

	assert.True(t, west.HasCard(cards.New(cards.CardNine, cards.CardKindHeart)))
	assert.True(t, west.HasCard(cards.New(cards.CardTen, cards.CardKindHeart)))
	assert.True(t, west.HasCard(cards.New(cards.CardJack, cards.CardKindHeart)))
	assert.True(t, west.HasCard(cards.New(cards.CardQueen, cards.CardKindHeart)))

	assert.True(t, south.HasCard(cards.New(cards.CardKing, cards.CardKindHeart)))
	assert.True(t, south.HasCard(cards.New(cards.CardAce, cards.CardKindHeart)))
	assert.True(t, south.HasCard(cards.New(cards.CardTwo, cards.CardKindDiamond)))
	assert.True(t, south.HasCard(cards.New(cards.CardThree, cards.CardKindDiamond)))

	assert.True(t, east.HasCard(cards.New(cards.CardFour, cards.CardKindDiamond)))
	assert.True(t, east.HasCard(cards.New(cards.CardFive, cards.CardKindDiamond)))
	assert.True(t, east.HasCard(cards.New(cards.CardSix, cards.CardKindDiamond)))
	assert.True(t, east.HasCard(cards.New(cards.CardSeven, cards.CardKindDiamond)))

	assert.True(t, north.HasCard(cards.New(cards.CardEight, cards.CardKindDiamond)))
	assert.True(t, north.HasCard(cards.New(cards.CardNine, cards.CardKindDiamond)))
	assert.True(t, north.HasCard(cards.New(cards.CardTen, cards.CardKindDiamond)))
	assert.True(t, north.HasCard(cards.New(cards.CardJack, cards.CardKindDiamond)))

	assert.True(t, west.HasCard(cards.New(cards.CardQueen, cards.CardKindDiamond)))
	assert.True(t, west.HasCard(cards.New(cards.CardKing, cards.CardKindDiamond)))
	assert.True(t, west.HasCard(cards.New(cards.CardAce, cards.CardKindDiamond)))
	assert.True(t, west.HasCard(cards.New(cards.CardTwo, cards.CardKindClub)))

	assert.True(t, south.HasCard(cards.New(cards.CardThree, cards.CardKindClub)))
	assert.True(t, south.HasCard(cards.New(cards.CardFour, cards.CardKindClub)))
	assert.True(t, south.HasCard(cards.New(cards.CardFive, cards.CardKindClub)))
	assert.True(t, south.HasCard(cards.New(cards.CardSix, cards.CardKindClub)))

	assert.True(t, east.HasCard(cards.New(cards.CardSeven, cards.CardKindClub)))
	assert.True(t, east.HasCard(cards.New(cards.CardEight, cards.CardKindClub)))
	assert.True(t, east.HasCard(cards.New(cards.CardNine, cards.CardKindClub)))
	assert.True(t, east.HasCard(cards.New(cards.CardTen, cards.CardKindClub)))

	assert.True(t, north.HasCard(cards.New(cards.CardJack, cards.CardKindClub)))
	assert.True(t, north.HasCard(cards.New(cards.CardQueen, cards.CardKindClub)))
	assert.True(t, north.HasCard(cards.New(cards.CardKing, cards.CardKindClub)))
	assert.True(t, north.HasCard(cards.New(cards.CardAce, cards.CardKindClub)))

	head, err := tableWithColor.Head()
	assert.NoError(t, err)
	assert.Equal(t, dealerPosition.Next(), head)
}
