package tables

import (
	"errors"
	"minhajuddinkhan/runng/runng/cards"
	"minhajuddinkhan/runng/runng/decks"
)

type FirstJackDealNominator struct {
}

func NewFirstJackDealNominator() *FirstJackDealNominator {
	return &FirstJackDealNominator{}
}

func (n *FirstJackDealNominator) Nominate(startingWith TablePlayerPosition, d decks.Deck) (TablePlayerPosition, error) {
	pos := startingWith
	for _, card := range d {
		if cards.IsJack(card) {
			return pos, nil
		}
		pos = pos.Next()
	}
	return -1, errors.New("no jack found")
}
