package card

import "errors"

type Suit int

const (
	SuitSpade Suit = iota
	SuitHeart
	SuitDiamond
	SuitClub
)

type Formation int

const (
	FormationInvalid Formation = iota
	FormationSingle
	FormationSameNum
	FormationStairs
)

type Card struct {
	Suit Suit
	Num  int
}

func Deck() []Card {
	var deck []Card
	for i := 1; i <= 10; i++ {
		deck = append(deck, Card{SuitSpade, i})
		deck = append(deck, Card{SuitHeart, i})
		deck = append(deck, Card{SuitDiamond, i})
		deck = append(deck, Card{SuitClub, i})
	}
	return deck
}

type CardLine struct {
	Cards []Card
}

var ErrFormationNotMatch = errors.New("formation not match")
var ErrBeated = errors.New("beated")

func Beat(active, challenge *CardLine) error {
	// 同じ陣形でないと勝負できない
	if active.Formation() != challenge.Formation() {
		return ErrFormationNotMatch
	}
	// 枚数が違うと勝負できない
	if len(active.Cards) != len(challenge.Cards) {
		return ErrFormationNotMatch
	}
	// カードの強さを比較する
	if active.Cards[0].Num > challenge.Cards[0].Num {
		return ErrBeated
	}
	return nil
}

// Formation returns the type of the card line.
func (c *CardLine) Formation() Formation {
	if c.IsFormationSingle() {
		return FormationSingle
	}
	if c.IsFormationSameNum() {
		return FormationSameNum
	}
	if c.IsFormationStairs() {
		return FormationStairs
	}
	return FormationInvalid
}

func (c *CardLine) IsFormationSingle() bool {
	return len(c.Cards) == 1
}

func (c *CardLine) IsFormationSameNum() bool {
	for _, card := range c.Cards[1:] {
		if card.Num != c.Cards[0].Num {
			return false
		}
	}
	return true
}

func (c *CardLine) IsFormationStairs() bool {
	if len(c.Cards) < 2 {
		return false
	}
	for _, card := range c.Cards[1:] {
		if card.Suit != c.Cards[0].Suit {
			return false
		}
	}

	return true
}
