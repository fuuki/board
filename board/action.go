package board

type PlayerActionDefinition interface {
}

type ActionRequest[AP PlayerActionDefinition] interface {
	IsValid(ActionProfile[AP]) bool
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

// NewActionProfileWithAction returns a new action profile with the action.
func NewActionProfileWithAction[AP PlayerActionDefinition](actions []AP) *ActionProfile[AP] {
	ap := &ActionProfile[AP]{
		playerActions: actions,
	}
	return ap
}

// Player returns the player's action.
func (ap *ActionProfile[AP]) Player(p Player) AP {
	return ap.playerActions[p]
}

// PlayerActions returns all player's actions.
func (ap *ActionProfile[AP]) PlayerActions() []AP {
	return ap.playerActions
}

// SetPlayerAction sets the player's action.
func (ap *ActionProfile[AP]) SetPlayerAction(p Player, a AP) {
	ap.playerActions[p] = a
}
