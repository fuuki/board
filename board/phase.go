package board

type PhaseName string

// Phase is a phase of the game.
type Phase[BP BoardProfile, PD PlayerActionDefinition] struct {
	name PhaseName
	// prepare returns the action profile definition.
	prepare func(*Game[BP, PD]) *ActionRequest[PD]
	// execute returns the next phase name.
	// if the next phase name is empty, the game is over.
	execute func(*Game[BP, PD], BP, *ActionProfile[PD]) (PhaseName, BP)
}

// NewPhase returns a new phase.
func NewPhase[BP BoardProfile, PD PlayerActionDefinition](
	name PhaseName,
	prepare func(*Game[BP, PD]) *ActionRequest[PD],
	execute func(*Game[BP, PD], BP, *ActionProfile[PD]) (PhaseName, BP),
) *Phase[BP, PD] {
	return &Phase[BP, PD]{
		name:    name,
		prepare: prepare,
		execute: execute,
	}
}
