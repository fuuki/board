package board

type PhaseName string

// Phase is a phase of the game.
type Phase[BP BoardProfile, AP PlayerActionDefinition] struct {
	name PhaseName
	// prepare returns the action profile definition.
	prepare func(*Game[BP, AP]) *ActionRequest[AP]
	// execute returns the next phase name.
	// if the next phase name is empty, the game is over.
	execute func(*Game[BP, AP], BP, *ActionProfile[AP]) (PhaseName, BP)
}

// NewPhase returns a new phase.
func NewPhase[BP BoardProfile, AP PlayerActionDefinition](
	name PhaseName,
	prepare func(*Game[BP, AP]) *ActionRequest[AP],
	execute func(*Game[BP, AP], BP, *ActionProfile[AP]) (PhaseName, BP),
) *Phase[BP, AP] {
	return &Phase[BP, AP]{
		name:    name,
		prepare: prepare,
		execute: execute,
	}
}
