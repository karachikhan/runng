package tables

import (
	"errors"
	"minhajuddinkhan/runng/runng/cards"
)

var (
	ErrPlayerDoesNotHaveColor = errors.New("player does not have the color")
)

type ColorNominator interface {
	Nominate(color cards.CardKind) error
	AnnounceColor() cards.CardKind
	Position() TablePlayerPosition
}

type colorNominator struct {
	*TablePlayer
	nomination chan cards.CardKind
}

func NewColorNominator(tp *TablePlayer) ColorNominator {
	return &colorNominator{
		nomination:  make(chan cards.CardKind),
		TablePlayer: tp,
	}
}

func (n *colorNominator) Nominate(color cards.CardKind) error {
	found := false
	for _, c := range n.cards {
		if c.Kind() == color {
			found = true
		}
	}
	if !found {
		return ErrPlayerDoesNotHaveColor
	}
	n.nomination <- color
	return nil
}

func (n *colorNominator) AnnounceColor() cards.CardKind {
	return <-n.nomination
}

func (n *colorNominator) Position() TablePlayerPosition {
	return n.position
}
