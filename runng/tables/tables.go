package tables

import (
	"minhajuddinkhan/runng/runng/cards"

	"github.com/google/uuid"
)

type TableWithColor interface {
	Recieve(tp *TablePlayer, c cards.Card) error
	Head() (TablePlayerPosition, error)
}

type Hand interface {
	Recieve(tp *TablePlayer, c cards.Card) error
	Head() (TablePlayerPosition, error)
	Complete() bool
}

type tableWithColor struct {
	id      uuid.UUID
	players map[TablePlayerPosition]*TablePlayer
	color   cards.CardKind
	head    TablePlayerPosition
	hands   []Hand
}

func NewTableWithColor(players map[TablePlayerPosition]*TablePlayer, color cards.CardKind, head TablePlayerPosition) TableWithColor {
	return &tableWithColor{
		id:      uuid.New(),
		players: players,
		color:   color,
		head:    head,
		hands:   []Hand{NewHand(head)},
	}
}

func (t *tableWithColor) Head() (TablePlayerPosition, error) {
	return t.head, nil
}

func (t *tableWithColor) Recieve(tp *TablePlayer, c cards.Card) error {
	latest := t.hands[len(t.hands)-1]
	if latest.Complete() {
		latest = NewHand(t.head)
		t.hands = append(t.hands, latest)
	}
	if !latest.Complete() {
		if err := latest.Recieve(tp, c); err != nil {
			return err
		}
	}

	head, err := latest.Head()
	if err != nil {
		return err
	}

	t.head = head
	t.hands[len(t.hands)-1] = latest

	return nil
}
