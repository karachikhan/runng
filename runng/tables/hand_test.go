package tables_test

import (
	"minhajuddinkhan/runng/runng/cards"
	"minhajuddinkhan/runng/runng/tables"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHand(t *testing.T) {
	table := tables.NewTable()
	p1 := tables.NewPlayer()
	north := tables.NewTablePlayer(table.GetID(), &p1, tables.TablePlayerPositionNorth)
	p2 := tables.NewPlayer()
	south := tables.NewTablePlayer(table.GetID(), &p2, tables.TablePlayerPositionSouth)
	hand := tables.NewHand(tables.TablePlayerPositionSouth)
	p3 := tables.NewPlayer()
	west := tables.NewTablePlayer(table.GetID(), &p3, tables.TablePlayerPositionWest)
	p4 := tables.NewPlayer()
	east := tables.NewTablePlayer(table.GetID(), &p4, tables.TablePlayerPositionEast)

	// player ke pas woh patta hona chaiye jo wo phek raha hai.
	err := north.Throw(hand, cards.New(cards.CardAce, cards.CardKindSpade))
	assert.ErrorAs(t, err, &tables.ErrCardNotInHand)
	north.Take([]cards.Card{cards.New(cards.CardAce, cards.CardKindSpade)})

	err = north.Throw(hand, cards.New(cards.CardAce, cards.CardKindSpade))
	assert.Error(t, err, &tables.ErrNotPlayersTurn)
	// jo head hai wohi hand mein card daal sakta hai, agar hand me pehle se card hai toh wohi card daal sakta hai

	south.Take([]cards.Card{cards.New(cards.CardTwo, cards.CardKindSpade)})
	err = south.Throw(hand, cards.New(cards.CardTwo, cards.CardKindSpade))
	assert.NoError(t, err)

	// south ke baad east ki turn hai, to east wala card daal sakta hai
	west.Take([]cards.Card{cards.New(cards.CardThree, cards.CardKindSpade)})
	err = west.Throw(hand, cards.New(cards.CardThree, cards.CardKindSpade))
	assert.ErrorAs(t, err, &tables.ErrNotPlayersTurn)

	east.Take([]cards.Card{cards.New(cards.CardFour, cards.CardKindSpade)})
	east.Throw(hand, cards.New(cards.CardFour, cards.CardKindSpade))

	head, err := hand.Head()
	assert.NoError(t, err)
	assert.Equal(t, east.GetPosition(), head)

	err = north.Throw(hand, cards.New(cards.CardAce, cards.CardKindSpade))
	assert.Nil(t, err, &tables.ErrNotPlayersTurn)

	head, err = hand.Head()
	assert.NoError(t, err)
	assert.Equal(t, north.GetPosition(), head)

}

func TestHandScenario(t *testing.T) {

	table := tables.NewTable()

	p1 := tables.NewPlayer()
	north := tables.NewTablePlayer(table.GetID(), &p1, tables.TablePlayerPositionNorth)
	north.Take([]cards.Card{cards.New(cards.CardFive, cards.CardKindHeart)})

	p2 := tables.NewPlayer()
	south := tables.NewTablePlayer(table.GetID(), &p2, tables.TablePlayerPositionSouth)
	south.Take([]cards.Card{cards.New(cards.CardEight, cards.CardKindSpade)})

	p3 := tables.NewPlayer()
	west := tables.NewTablePlayer(table.GetID(), &p3, tables.TablePlayerPositionWest)
	west.Take([]cards.Card{cards.New(cards.CardThree, cards.CardKindSpade)})

	p4 := tables.NewPlayer()
	east := tables.NewTablePlayer(table.GetID(), &p4, tables.TablePlayerPositionEast)
	east.Take([]cards.Card{cards.New(cards.CardKing, cards.CardKindSpade)})

	hand := tables.NewHand(tables.TablePlayerPositionEast)

	assert.NoError(t, east.Throw(hand, cards.New(cards.CardKing, cards.CardKindSpade)))
	assert.NoError(t, north.Throw(hand, cards.New(cards.CardFive, cards.CardKindHeart)))
	assert.NoError(t, west.Throw(hand, cards.New(cards.CardThree, cards.CardKindSpade)))
	assert.NoError(t, south.Throw(hand, cards.New(cards.CardEight, cards.CardKindSpade)))

	head, err := hand.Head()
	assert.NoError(t, err)
	assert.Equal(t, tables.TablePlayerPositionEast, head)

}
