package tables

import (
	"errors"
	"minhajuddinkhan/runng/runng/cards"

	"github.com/google/uuid"
)

var (
	ErrCardNotInHand  = errors.New("card not in hand")
	ErrNotPlayersTurn = errors.New("not players turn")
)

type TablePlayer struct {
	tableID  uuid.UUID
	position TablePlayerPosition
	*Player
	cards []cards.Card
}

func NewTablePlayer(tableID uuid.UUID, p *Player, pos TablePlayerPosition) *TablePlayer {
	return &TablePlayer{
		tableID:  tableID,
		position: pos,
		Player:   p,
		cards:    make([]cards.Card, 0),
	}
}

func (tp *TablePlayer) Take(c []cards.Card) {
	tp.cards = append(tp.cards, c...)
}

func (tp *TablePlayer) HasCard(card cards.Card) bool {
	for _, c := range tp.cards {
		if c.Kind() == card.Kind() && c.Number() == card.Number() {
			return true
		}
	}
	return false
}

func (tp *TablePlayer) GetPosition() TablePlayerPosition {
	return tp.position
}

func (tp *TablePlayer) Throw(h TableWithColor, c cards.Card) error {
	if !tp.HasCard(c) {
		return ErrCardNotInHand
	}

	if err := h.Recieve(tp, c); err != nil {
		return err
	}
	return nil
}
