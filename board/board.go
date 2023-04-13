package board

import "github.com/fuuki/board/player"

type BoardProfile interface {
	Show() string
}

type ActionProfile[AP PlayerActionDefinition] struct {
	playerActions []*AP
}

// NewActionProfile returns a new action profile.
func NewActionProfile[AP PlayerActionDefinition](playerNum uint) *ActionProfile[AP] {
	ap := &ActionProfile[AP]{
		playerActions: make([]*AP, playerNum),
	}
	return ap
}

// Player returns the player's action.
func (ap *ActionProfile[AP]) Player(p player.Player) *AP {
	return ap.playerActions[p]
}

// SetPlayerAction sets the player's action.
func (ap *ActionProfile[AP]) SetPlayerAction(p player.Player, a *AP) {
	ap.playerActions[p] = a
}

type PlayerActionDefinition interface {
}

type ActionRequest[AP PlayerActionDefinition] struct {
}

func (ar *ActionRequest[AP]) IsValid(ap *ActionProfile[AP]) bool {
	for _, a := range ap.playerActions {
		if a == nil {
			return false
		}
	}
	return true
}
