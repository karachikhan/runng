package decks

import (
	"math/rand"
	"minhajuddinkhan/runng/runng/cards"
)

type Deck []cards.Card

func New() Deck {
	deck := make(Deck, 0)
	for kind := cards.CardKindSpade; kind <= cards.CardKindClub; kind++ {
		for number := cards.CardTwo; number <= cards.CardAce; number++ {
			deck = append(deck, cards.New(number, kind))
		}
	}
	return deck
}

func (d Deck) Random() int {
	return rand.Intn(51)
}
