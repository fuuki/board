package game

import (
	"github.com/fuuki/board/action"
)

type PhaseName string

// Phase is a phase of the game.
type Phase struct {
	name PhaseName
	// prepare returns the action profile definition.
	prepare func(*Game) *action.ActionProfileDefinition
	// execute returns the next phase name.
	// if the next phase name is empty, the game is over.
	execute func(*Game, *action.ActionProfile) PhaseName
}

// NewPhase returns a new phase.
func NewPhase(
	name PhaseName,
	prepare func(*Game) *action.ActionProfileDefinition,
	execute func(*Game, *action.ActionProfile) PhaseName,
) *Phase {
	return &Phase{
		name:    name,
		prepare: prepare,
		execute: execute,
	}
}
