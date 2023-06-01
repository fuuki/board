package logic

type ActionProfile[AD PlayerActionDefinition] struct {
	playerActions []AD
	natureActions map[string][]int
}

// NewActionProfile returns a new action profile.
func NewActionProfile[AD PlayerActionDefinition](playerActions []AD, naturalActions map[string][]int) *ActionProfile[AD] {
	ap := &ActionProfile[AD]{
		playerActions: playerActions,
		natureActions: naturalActions,
	}
	return ap
}

// Player returns the player's action.
func (ap *ActionProfile[AD]) Player(p Player) AD {
	return ap.playerActions[p]
}

// NatureActionResult returns the nature's action result.
func (ap *ActionProfile[AD]) NatureActionResult(name string) []int {
	return ap.natureActions[name]
}
