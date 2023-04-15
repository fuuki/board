package resource

import (
	"fmt"
	"strings"
)

// CardID is a card ID.
type CardID string

// Card is a playing card.
type Card interface {
	ID() CardID
}

// CardLine is a line of cards.
type CardLine[C Card] struct {
	line []C
}

// NewCardLine returns a new CardLine.
func NewCardLine[C Card](cards []C) *CardLine[C] {
	return &CardLine[C]{line: cards}
}

// Add adds a card to the line.
func (cl *CardLine[C]) Add(card C) {
	cl.line = append(cl.line, card)
}

// AddMulti adds multiple cards to the line.
func (cl *CardLine[C]) AddMulti(cards []C) {
	cl.line = append(cl.line, cards...)
}

// Cards returns all cards in the line.
func (cl *CardLine[C]) Cards() []C {
	return cl.line
}

// Pick picks a card from the line.
func (cl *CardLine[C]) Pick(id CardID) *C {
	for i, c := range cl.line {
		if c.ID() == id {
			cl.line = remove(cl.line, i)
			return &c
		}
	}
	return nil
}

// PickMulti picks multiple cards from the line.
func (cl *CardLine[C]) PickMulti(ids []CardID) []C {
	var cards []C
	for _, id := range ids {
		c := cl.Pick(id)
		if c != nil {
			cards = append(cards, *c)
		}
	}
	return cards
}

// Equals returns true if the line is equal to the other line.
func (cl *CardLine[C]) Equals(other *CardLine[C]) bool {
	if len(cl.line) != len(other.line) {
		return false
	}
	for i, c := range cl.line {
		if c.ID() != other.line[i].ID() {
			return false
		}
	}
	return true
}

// String returns a string representation of the line.
func (cl *CardLine[C]) String() string {
	var s []string
	for _, c := range cl.line {
		s = append(s, string(c.ID()))
	}
	str := strings.Join(s, ",")
	return fmt.Sprintf("[%s]", str)
}
