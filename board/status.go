package board

// Status is a status of the game.
type Status struct {
	totalPlayer uint
}

// NewStatus returns a new status.
func NewStatus(totalPlayer uint) *Status {
	return &Status{
		totalPlayer: totalPlayer,
	}
}

// TotalPlayer returns the total number of players.
func (s *Status) TotalPlayer() uint {
	return s.totalPlayer
}
