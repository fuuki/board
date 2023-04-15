package resource

import "github.com/fuuki/board/board"

// Turn is a turn of the game.
type Turn struct {
	order   []board.Player
	current int
}

// NewTurn returns a new Turn.
func NewTurn(order []board.Player) *Turn {
	return &Turn{
		order:   order,
		current: 0,
	}
}

// NewSimpleTurn returns a new Turn with simple order.
func NewSimpleTurn(playerNum uint) *Turn {
	order := make([]board.Player, playerNum)
	for i := uint(0); i < playerNum; i++ {
		order[i] = board.Player(i)
	}
	return NewTurn(order)
}

// Next returns the next player.
func (t *Turn) Next() board.Player {
	t.current++
	if t.current >= len(t.order) {
		t.current = 0
	}
	return t.order[t.current]
}

// Current returns the current player.
func (t *Turn) Current() board.Player {
	return t.order[t.current]
}
