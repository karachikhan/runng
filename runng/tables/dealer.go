package tables

import (
	"minhajuddinkhan/runng/runng/cards"
	"minhajuddinkhan/runng/runng/decks"
)

type Dealer interface {
	Deal(players map[TablePlayerPosition]*TablePlayer)
}

func getPlayerAt(pos TablePlayerPosition, players []*TablePlayer) *TablePlayer {
	for _, p := range players {
		if p.position == pos {
			return p
		}
	}
	return nil
}

type dealer struct {
	*TablePlayer
	nextToDeal TablePlayerPosition
	dealRound  dealRound
}

type dealRound int

func (d dealRound) drawCount() int {
	switch d {
	case 1:
		return 5
	default:
		return 4
	}
}

func NewDealer(p *TablePlayer, deck decks.Deck) Dealer {
	p.Take(deck)
	return &dealer{
		TablePlayer: p,
		nextToDeal:  p.position.Next(),
		dealRound:   1,
	}
}

func (d *dealer) Deal(players map[TablePlayerPosition]*TablePlayer) {
	for i := 0; i < len(players); i++ {
		cards := d.dealHand(d.dealRound.drawCount())
		next := players[d.nextToDeal]
		next.Take(cards)
		d.nextToDeal = next.position.Next()
	}
	d.dealRound++
}

func (d *dealer) dealHand(n int) []cards.Card {
	cards := d.TablePlayer.cards[0:n]
	d.TablePlayer.cards = d.TablePlayer.cards[n:]
	return cards
}
