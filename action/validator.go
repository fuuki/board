package action

import (
	"fmt"
	"reflect"

	"github.com/fuuki/board/board"
)

var ErrMustNotTakeAction = fmt.Errorf("must not take action")
var ErrMustTakeAction = fmt.Errorf("must take action")
var ErrInvalidAction = fmt.Errorf("invalid action")

func TurnValid[AP board.PlayerActionDefinition](ap board.ActionProfile[AP], currentPlayer board.Player) error {
	selected := false
	for i, a := range ap.PlayerActions() {
		p := board.Player(i)
		if p == currentPlayer {
			// current player should be not nil
			if !reflect.DeepEqual(a, *new(AP)) {
				selected = true
			}
		} else {
			// not current player should be nil
			if !reflect.DeepEqual(a, *new(AP)) {
				return fmt.Errorf("player %d %w", p, ErrMustNotTakeAction)
			}
		}
	}
	if !selected {
		return fmt.Errorf("player %d %w", currentPlayer, ErrMustTakeAction)
	}
	return nil
}
