package board

type PlayerActionDefinition interface {
}

type ActionRequest[AP PlayerActionDefinition] interface {
	IsCompleted(ActionProfile[AP]) bool
}

type ActionProfile[AP PlayerActionDefinition] struct {
	playerActions []AP
}

// NewActionProfile returns a new action profile.
func NewActionProfile[AP PlayerActionDefinition](playerNum uint) *ActionProfile[AP] {
	ap := &ActionProfile[AP]{
		playerActions: make([]AP, playerNum),
	}
	return ap
}

// Player returns the player's action.
func (ap *ActionProfile[AP]) Player(p Player) AP {
	return ap.playerActions[p]
}

// SetPlayerAction sets the player's action.
func (ap *ActionProfile[AP]) SetPlayerAction(p Player, a AP) {
	ap.playerActions[p] = a
}
