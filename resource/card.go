package resource

import (
	"encoding/json"
	"fmt"
	"reflect"
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
func (cl *CardLine[C]) Pick(id CardID) C {
	for i, c := range cl.line {
		if c.ID() == id {
			cl.line = remove(cl.line, i)
			return c
		}
	}
	return *new(C)
}

// Draw draws a card from the line.
func (cl *CardLine[C]) Draw() C {
	if len(cl.line) == 0 {
		return *new(C)
	}
	c := cl.line[0]
	cl.line = cl.line[1:]
	return c
}

// Has returns true if the line has the card.
func (cl *CardLine[C]) Has(id CardID) bool {
	for _, c := range cl.line {
		if c.ID() == id {
			return true
		}
	}
	return false
}

// PickMulti picks multiple cards from the line.
func (cl *CardLine[C]) PickMulti(ids []CardID) []C {
	var cards []C
	for _, id := range ids {
		c := cl.Pick(id)
		if reflect.ValueOf(c).IsZero() {
			cards = append(cards, c)
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

// MarshalJSON implements json.Marshaler.
func (cl *CardLine[C]) MarshalJSON() ([]byte, error) {
	return json.Marshal(cl.line)
}

// Clone returns a cloned CardLine.
func (cl *CardLine[C]) Clone() *CardLine[C] {
	line := make([]C, len(cl.line))
	copy(line, cl.line)
	return NewCardLine(line)
}
