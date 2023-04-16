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

var shortDeck = []*Card{
	{id: "s1", Rank: 1, Suit: Spade},
	{id: "s2", Rank: 2, Suit: Spade},
	{id: "s3", Rank: 3, Suit: Spade},
	{id: "s4", Rank: 4, Suit: Spade},
	{id: "s5", Rank: 5, Suit: Spade},
	{id: "h1", Rank: 1, Suit: Heart},
	{id: "h2", Rank: 2, Suit: Heart},
	{id: "h3", Rank: 3, Suit: Heart},
	{id: "h4", Rank: 4, Suit: Heart},
	{id: "h5", Rank: 5, Suit: Heart},
}

var deck = []*Card{
	{id: "s1", Rank: 1, Suit: Spade},
	{id: "s2", Rank: 2, Suit: Spade},
	{id: "s3", Rank: 3, Suit: Spade},
	{id: "s4", Rank: 4, Suit: Spade},
	{id: "s5", Rank: 5, Suit: Spade},
	{id: "s6", Rank: 6, Suit: Spade},
	{id: "s7", Rank: 7, Suit: Spade},
	{id: "s8", Rank: 8, Suit: Spade},
	{id: "s9", Rank: 9, Suit: Spade},
	{id: "s10", Rank: 10, Suit: Spade},
	{id: "s11", Rank: 11, Suit: Spade},
	{id: "s12", Rank: 12, Suit: Spade},
	{id: "s13", Rank: 13, Suit: Spade},
	{id: "h1", Rank: 1, Suit: Heart},
	{id: "h2", Rank: 2, Suit: Heart},
	{id: "h3", Rank: 3, Suit: Heart},
	{id: "h4", Rank: 4, Suit: Heart},
	{id: "h5", Rank: 5, Suit: Heart},
	{id: "h6", Rank: 6, Suit: Heart},
	{id: "h7", Rank: 7, Suit: Heart},
	{id: "h8", Rank: 8, Suit: Heart},
	{id: "h9", Rank: 9, Suit: Heart},
	{id: "h10", Rank: 10, Suit: Heart},
	{id: "h11", Rank: 11, Suit: Heart},
	{id: "h12", Rank: 12, Suit: Heart},
	{id: "h13", Rank: 13, Suit: Heart},
	{id: "d1", Rank: 1, Suit: Diamond},
	{id: "d2", Rank: 2, Suit: Diamond},
	{id: "d3", Rank: 3, Suit: Diamond},
	{id: "d4", Rank: 4, Suit: Diamond},
	{id: "d5", Rank: 5, Suit: Diamond},
	{id: "d6", Rank: 6, Suit: Diamond},
	{id: "d7", Rank: 7, Suit: Diamond},
	{id: "d8", Rank: 8, Suit: Diamond},
	{id: "d9", Rank: 9, Suit: Diamond},
	{id: "d10", Rank: 10, Suit: Diamond},
	{id: "d11", Rank: 11, Suit: Diamond},
	{id: "d12", Rank: 12, Suit: Diamond},
	{id: "d13", Rank: 13, Suit: Diamond},
	{id: "c1", Rank: 1, Suit: Club},
	{id: "c2", Rank: 2, Suit: Club},
	{id: "c3", Rank: 3, Suit: Club},
	{id: "c4", Rank: 4, Suit: Club},
	{id: "c5", Rank: 5, Suit: Club},
	{id: "c6", Rank: 6, Suit: Club},
	{id: "c7", Rank: 7, Suit: Club},
	{id: "c8", Rank: 8, Suit: Club},
	{id: "c9", Rank: 9, Suit: Club},
	{id: "c10", Rank: 10, Suit: Club},
	{id: "c11", Rank: 11, Suit: Club},
	{id: "c12", Rank: 12, Suit: Club},
	{id: "c13", Rank: 13, Suit: Club},
}
