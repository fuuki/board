package action

import (
	"fmt"
	"reflect"

	"github.com/fuuki/board/board"
)

var ErrMustNotTakeAction = fmt.Errorf("player must not take action")
var ErrMustTakeAction = fmt.Errorf("player must take action")
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
				return ErrMustNotTakeAction
			}
		}
	}
	if !selected {
		return ErrMustTakeAction
	}
	return nil
}
