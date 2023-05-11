package logic

type PhaseName string

type Config interface {
	// TotalPlayer returns the total number of players.
	TotalPlayer() uint
}

// Phase is a phase of the game.
type Phase[AD PlayerActionDefinition, BP BoardProfile, CF Config] interface {
	Name() PhaseName
	// prepare returns the action profile definition.
	Prepare(CF, BP) (*ActionRequest[AD], BP)
	// execute returns the next phase name.
	// if the next phase name is empty, the game is over.
	Execute(CF, BP, *ActionProfile[AD]) (PhaseName, BP)
}
