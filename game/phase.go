package game

import (
	"github.com/fuuki/board/board"
)

type PhaseName string

// Phase is a phase of the game.
type Phase[BP board.BoardProfile, AP board.PlayerActionDefinition] struct {
	name PhaseName
	// prepare returns the action profile definition.
	prepare func(*Game[BP, AP]) *board.ActionRequest[AP]
	// execute returns the next phase name.
	// if the next phase name is empty, the game is over.
	execute func(*Game[BP, AP], *board.ActionProfile[AP]) PhaseName
}

// NewPhase returns a new phase.
func NewPhase[BP board.BoardProfile, AP board.PlayerActionDefinition](
	name PhaseName,
	prepare func(*Game[BP, AP]) *board.ActionRequest[AP],
	execute func(*Game[BP, AP], *board.ActionProfile[AP]) PhaseName,
) *Phase[BP, AP] {
	return &Phase[BP, AP]{
		name:    name,
		prepare: prepare,
		execute: execute,
	}
}
