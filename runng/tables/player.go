package tables

import (
	"math/rand"
	"minhajuddinkhan/runng/runng/cards"
	"minhajuddinkhan/runng/runng/decks"

	"github.com/google/uuid"
)

type Shuffler interface {
	Shuffle(decks.Deck) decks.Deck
}

type Player struct {
	id uuid.UUID
}

func NewPlayer() Player {
	return Player{
		id: uuid.New(),
	}
}

func (p *Player) GetID() uuid.UUID {
	return p.id
}

func (p *Player) Join(t Table, pos TablePlayerPosition) (*TablePlayer, error) {
	return t.Join(p, pos)
}

func (p *Player) TakeFrom(start, end int, d decks.Deck) ([]cards.Card, decks.Deck) {
	cards := d[start:end]
	d = append(d[:start], d[end:]...)
	return cards, d
}

func (p *Player) PutOnDeck(c []cards.Card, d decks.Deck) decks.Deck {
	return append(c, d...)
}

func (p *Player) Shuffle(deck decks.Deck) decks.Deck {
	for i := 0; i < rand.Intn(6); i++ {
		start := 10 + rand.Intn(10)
		end := 20 + rand.Intn(10)
		cards, d := p.TakeFrom(start, end, deck)
		deck = p.PutOnDeck(cards, d)
	}
	return deck
}
