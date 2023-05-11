package logic

import (
	"fmt"

	"github.com/fuuki/board/logic/internal/action"
)

var ErrInvalidPlayer = fmt.Errorf("invalid player")
var ErrNotActionablePlayer = fmt.Errorf("not actionable")

type PlayerActionDefinition interface{}

type actionRequestDefinition[AD PlayerActionDefinition] struct {
	actionable bool
	validator  func(AD) error
}

type ActionRequest[AD PlayerActionDefinition] struct {
	defs          []actionRequestDefinition[AD]
	naturalPlayer *action.NaturalPlayer
}

// NewActionRequest returns a new action request.
func NewActionRequest[AD PlayerActionDefinition](totalPlayer uint) *ActionRequest[AD] {
	defs := make([]actionRequestDefinition[AD], totalPlayer)
	for i := uint(0); i < totalPlayer; i++ {
		defs[i] = actionRequestDefinition[AD]{
			actionable: false,
			validator:  func(AD) error { return nil },
		}
	}
	naturalPlayer := action.NewNaturalPlayer()
	return &ActionRequest[AD]{
		defs:          defs,
		naturalPlayer: naturalPlayer,
	}
}

// RegisterValidator registers the validator for the player.
func (ar ActionRequest[AD]) RegisterValidator(p Player, fn func(AD) error) {
	ar.defs[p] = actionRequestDefinition[AD]{
		actionable: true,
		validator:  fn,
	}
}

// AddDice adds dice to the nature's action.
func (ar ActionRequest[AD]) AddDice(name string, numDice int, numFace int) {
	ar.naturalPlayer.AddDice(name, numDice, numFace)
}

// AddShuffle adds shuffle to the nature's action.
func (ar ActionRequest[AD]) AddShuffle(name string, numCard int) {
	ar.naturalPlayer.AddShuffle(name, numCard)
}

// ValidateAction validates the action.
// It returns nil if the action is valid.
// It returns an error if the action is invalid.
func (ar ActionRequest[AD]) ValidateAction(p Player, a AD) error {
	if int(p) >= len(ar.defs) {
		return fmt.Errorf("player %d is %w", p, ErrInvalidPlayer)
	}
	def := ar.defs[p]

	if !def.actionable {
		return fmt.Errorf("player %d is %w", p, ErrNotActionablePlayer)
	}
	return def.validator(a)
}

// IsActionable returns true if the player is actionable.
func (ar ActionRequest[AD]) IsActionable(p Player) bool {
	if int(p) >= len(ar.defs) {
		return false
	}
	return ar.defs[p].actionable
}

type ActionProfile[AD PlayerActionDefinition] struct {
	playerActions []AD
	natureActions map[string][]int
}

// NewActionProfile returns a new action profile.
func NewActionProfile[AD PlayerActionDefinition](playerActions []AD) *ActionProfile[AD] {
	ap := &ActionProfile[AD]{
		playerActions: playerActions,
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
