package board

type PhaseName string

// Phase is a phase of the game.
type Phase[BP BoardProfile, PD PlayerActionDefinition] struct {
	name PhaseName
	// prepare returns the action profile definition.
	prepare func(*Status, BP) (*ActionRequest[PD], BP)
	// execute returns the next phase name.
	// if the next phase name is empty, the game is over.
	execute func(*Status, BP, *ActionProfile[PD]) (PhaseName, BP)
}

// NewPhase returns a new phase.
func NewPhase[BP BoardProfile, PD PlayerActionDefinition](
	name PhaseName,
	prepare func(*Status, BP) (*ActionRequest[PD], BP),
	execute func(*Status, BP, *ActionProfile[PD]) (PhaseName, BP),
) *Phase[BP, PD] {
	return &Phase[BP, PD]{
		name:    name,
		prepare: prepare,
		execute: execute,
	}
}
