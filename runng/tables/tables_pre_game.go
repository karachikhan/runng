package tables

import (
	"errors"
	"minhajuddinkhan/runng/runng/cards"
	"minhajuddinkhan/runng/runng/decks"

	"github.com/google/uuid"
)

type Table interface {
	GetID() uuid.UUID
	Join(p *Player, pos TablePlayerPosition) (*TablePlayer, error)
	WithDealer(pos TablePlayerPosition, deck decks.Deck) TableWithDealer
}

type TableWithDealer interface {
	Deal() (TableAwaitingColor, error)
}

type TableAwaitingColor interface {
	AwaitColor(ColorNominator) (TableWithColor, error)
}

type table struct {
	id      uuid.UUID
	dealer  Dealer
	players map[TablePlayerPosition]*TablePlayer
}

func NewTable() Table {
	tableID := uuid.New()
	return &table{
		id:      tableID,
		players: make(map[TablePlayerPosition]*TablePlayer),
	}
}

func (t *table) GetID() uuid.UUID {
	return t.id
}

func (t *table) Join(p *Player, pos TablePlayerPosition) (*TablePlayer, error) {
	if _, ok := t.players[pos]; ok {
		return nil, errors.New("player already at position")
	}
	if len(t.players) == 4 {
		return nil, errors.New("table full")
	}

	tp := &TablePlayer{
		tableID:  t.id,
		position: pos,
		Player:   p,
		cards:    make([]cards.Card, 0),
	}
	t.players[pos] = tp
	return tp, nil
}

func (t *table) Deal() (TableAwaitingColor, error) {
	if t.dealer == nil {
		return nil, errors.New("dealer not set")
	}

	t.dealer.Deal(t.players)
	return t, nil
}

func (t *table) WithDealer(pos TablePlayerPosition, deck decks.Deck) TableWithDealer {
	t.dealer = NewDealer(t.players[pos], deck)
	return t
}

func (t *table) AwaitColor(nominator ColorNominator) (TableWithColor, error) {
	color := nominator.AnnounceColor()
	t.dealer.Deal(t.players)
	t.dealer.Deal(t.players)
	return NewTableWithColor(t.players, color, nominator.Position()), nil
}
