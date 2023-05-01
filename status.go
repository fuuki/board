package board

// Status is a status of the game.
type Status struct {
	totalPlayer uint
}

// NewStatus returns a new status.
func NewStatus[BP BoardProfile, PD PlayerActionDefinition](g *Game[BP, PD]) *Status {
	return &Status{
		totalPlayer: g.totalPlayer,
	}
}

// TotalPlayer returns the total number of players.
func (s *Status) TotalPlayer() uint {
	return s.totalPlayer
}
