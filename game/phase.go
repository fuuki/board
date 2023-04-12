package game

import (
	"github.com/fuuki/board/action"
	"github.com/fuuki/board/board"
)

type PhaseName string

// Phase is a phase of the game.
type Phase[BP board.BoardProfile] struct {
	name PhaseName
	// prepare returns the action profile definition.
	prepare func(*Game[BP]) *action.ActionProfileDefinition
	// execute returns the next phase name.
	// if the next phase name is empty, the game is over.
	execute func(*Game[BP], *action.ActionProfile) PhaseName
}

// NewPhase returns a new phase.
func NewPhase[BP board.BoardProfile](
	name PhaseName,
	prepare func(*Game[BP]) *action.ActionProfileDefinition,
	execute func(*Game[BP], *action.ActionProfile) PhaseName,
) *Phase[BP] {
	return &Phase[BP]{
		name:    name,
		prepare: prepare,
		execute: execute,
	}
}
