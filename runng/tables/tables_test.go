package tables_test

import (
	"minhajuddinkhan/runng/runng/cards"
	"minhajuddinkhan/runng/runng/decks"
	"minhajuddinkhan/runng/runng/tables"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
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
	theTable, err := tableAwaitingColor.AwaitColor(colorNominator)
	assert.NoError(t, err)

	assert.NoError(t, west.Throw(theTable, cards.New(cards.CardTwo, cards.CardKindSpade)))
	assert.NoError(t, south.Throw(theTable, cards.New(cards.CardSeven, cards.CardKindSpade)))
	assert.NoError(t, east.Throw(theTable, cards.New(cards.CardQueen, cards.CardKindSpade)))
	assert.NoError(t, north.Throw(theTable, cards.New(cards.CardFour, cards.CardKindHeart)))

	assert.NoError(t, east.Throw(theTable, cards.New(cards.CardKing, cards.CardKindSpade)))
	assert.NoError(t, north.Throw(theTable, cards.New(cards.CardFive, cards.CardKindHeart)))
	assert.NoError(t, west.Throw(theTable, cards.New(cards.CardThree, cards.CardKindSpade)))
	assert.NoError(t, south.Throw(theTable, cards.New(cards.CardEight, cards.CardKindSpade)))

	head, err := theTable.Head()
	assert.NoError(t, err)
	assert.Equal(t, tables.TablePlayerPositionEast, head)
}

// TODO:: kaatney ka scene
// TODO:: jab card same house ka available ho aur phir bhi player us house ka card nahi daal raha toh error aana chaiye
// TODO:: panchwey hath ke baad  player agar consecutive head rahe to wo saare hath player jeet jaega
