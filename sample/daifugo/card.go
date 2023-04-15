package daifugo

import (
	"github.com/fuuki/board/resource"
)

// Suit is a suit of a playing card.
type Suit string

const (
	Spade   Suit = "Spade"
	Heart   Suit = "Heart"
	Diamond Suit = "Diamond"
	Club    Suit = "Club"
)

// Card is a playing card.
type Card struct {
	id   string
	Rank int
	Suit Suit
}

var _ resource.Card = (*Card)(nil)

// ID returns a card ID.
func (c *Card) ID() resource.CardID {
	return resource.CardID(c.id)
}
