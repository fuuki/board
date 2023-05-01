package burst

import (
	"fmt"

	"github.com/fuuki/board"
	"github.com/fuuki/board/resource"
)

// Card is a playing card.
type Card struct {
	Number int `json:"number"`
}

var _ resource.Card = (*Card)(nil)

// ID returns a card ID.
func (c *Card) ID() resource.CardID {
	id := fmt.Sprintf("%d", c.Number)
	return resource.CardID(id)
}

// PlayedCard is a played card.
type PlayedCard struct {
	Card   *Card        `json:"card"`
	Player board.Player `json:"player"`
}

func newDeck() []*Card {
	return []*Card{
		{Number: 100},
		{Number: -5},
		{Number: -4},
		{Number: -3},
		{Number: -2},
		{Number: -1},
		{Number: 0},
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
		{Number: 5},
		{Number: 6},
		{Number: 7},
		{Number: 8},
		{Number: 9},
		{Number: 10},
		{Number: 15},
		{Number: 20},
	}
}
