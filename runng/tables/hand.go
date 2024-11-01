package tables

import (
	"minhajuddinkhan/runng/runng/cards"

	"github.com/google/uuid"
)

type hand struct {
	id       uuid.UUID
	head     TablePlayerPosition
	headCard cards.Card
	cards    map[TablePlayerPosition]cards.Card
	next     TablePlayerPosition
}

func NewHand(head TablePlayerPosition) Hand {
	return &hand{
		id:    uuid.New(),
		head:  head,
		cards: make(map[TablePlayerPosition]cards.Card),
		next:  head,
	}

}

func (h *hand) Recieve(tp *TablePlayer, c cards.Card) error {
	// only the player whose turn it is can play
	if tp.position != h.next {
		return ErrNotPlayersTurn
	}

	// if the hand is empty, set the kind of the hand
	if h.isEmpty() {
		h.setHead(tp.position, c)
	}

	h.cards[tp.position] = c
	h.next = tp.position.Next()

	if cards.IsBigger(c, h.headCard) {
		h.setHead(tp.position, c)
	}

	if h.Complete() {
		h.next = h.head
	}

	return nil

}

func (h *hand) Head() (TablePlayerPosition, error) {
	return h.head, nil
}

func (h *hand) Complete() bool {
	return len(h.cards) == 4
}

func (h *hand) isEmpty() bool {
	return len(h.cards) == 0
}

func (h *hand) setHead(player TablePlayerPosition, c cards.Card) {
	h.head = player
	h.headCard = c
}
