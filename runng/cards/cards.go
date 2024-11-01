package cards

type CardKind int

const (
	CardKindSpade CardKind = iota
	CardKindHeart
	CardKindDiamond
	CardKindClub
)

type CardNumber int

const (
	CardTwo CardNumber = iota
	CardThree
	CardFour
	CardFive
	CardSix
	CardSeven
	CardEight
	CardNine
	CardTen
	CardJack
	CardQueen
	CardKing
	CardAce
)

type Card interface {
	Number() CardNumber
	Kind() CardKind
}

func New(number CardNumber, kind CardKind) Card {
	return &card{
		number: number,
		kind:   kind,
	}
}

type card struct {
	number CardNumber
	kind   CardKind
}

func (c *card) Number() CardNumber {
	return c.number
}

func (c *card) Kind() CardKind {
	return c.kind
}

func IsJack(card Card) bool {
	return card.Number() == CardJack
}

// IsBigger
// returns true if c1 is bigger than c2
func IsBigger(c1, c2 Card) bool {
	return c1.Number() > c2.Number()
}
