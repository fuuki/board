package resource

import (
	"fmt"

	"github.com/fuuki/board/logic"
)

// Turn is a turn of the game.
type Turn struct {
	order   []logic.Player
	current int
}

// NewTurn returns a new Turn.
func NewTurn(order []logic.Player, startPlayer logic.Player) *Turn {
	c := 0
	for i, p := range order {
		if p == startPlayer {
			c = i
			break
		}
	}
	return &Turn{
		order:   order,
		current: c,
	}
}

// NewSimpleTurn returns a new Turn with simple order.
func NewSimpleTurn(totalPlayer uint) *Turn {
	order := make([]logic.Player, totalPlayer)
	for i := uint(0); i < totalPlayer; i++ {
		order[i] = logic.Player(i)
	}
	return NewTurn(order, 0)
}

// Next returns the next player.
func (t *Turn) Next() logic.Player {
	t.current++
	if t.current >= len(t.order) {
		t.current = 0
	}
	return t.order[t.current]
}

// Current returns the current player.
func (t *Turn) Current() logic.Player {
	return t.order[t.current]
}

// Pass pop the current player from the order.
func (t *Turn) Pass() {
	t.order = append(t.order[:t.current], t.order[t.current+1:]...)
	if t.current >= len(t.order) {
		t.current = 0
	}
}

// Order returns the order of the turn.
func (t *Turn) Order() []logic.Player {
	return t.order
}

// MarshalJSON implements json.Marshaler.
func (t *Turn) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", t.current)), nil
}
