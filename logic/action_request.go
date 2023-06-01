package logic

import (
	"fmt"

	"github.com/fuuki/board/logic/internal/action"
)

var ErrInvalidPlayer = fmt.Errorf("invalid player")
var ErrNotActionablePlayer = fmt.Errorf("not actionable")

type PlayerActionDefinition interface{}

type playerActionRequest[AD PlayerActionDefinition] struct {
	actionable bool
	validator  func(AD) error
}

type ActionRequest[AD PlayerActionDefinition] struct {
	players []playerActionRequest[AD]
	natural *action.NaturalPlayer
}

// NewActionRequest returns a new action request.
func NewActionRequest[AD PlayerActionDefinition](totalPlayer uint) *ActionRequest[AD] {
	defs := make([]playerActionRequest[AD], totalPlayer)
	for i := uint(0); i < totalPlayer; i++ {
		defs[i] = playerActionRequest[AD]{
			actionable: false,
			validator:  func(AD) error { return nil },
		}
	}
	naturalPlayer := action.NewNaturalPlayer()
	return &ActionRequest[AD]{
		players: defs,
		natural: naturalPlayer,
	}
}

// RegisterValidator registers the validator for the player.
func (ar ActionRequest[AD]) RegisterValidator(p Player, fn func(AD) error) {
	ar.players[p] = playerActionRequest[AD]{
		actionable: true,
		validator:  fn,
	}
}

// AddDice adds dice to the nature's action.
func (ar ActionRequest[AD]) AddDice(name string, numDice int, numFace int) {
	ar.natural.AddDice(name, numDice, numFace)
}

// AddShuffle adds shuffle to the nature's action.
func (ar ActionRequest[AD]) AddShuffle(name string, numCard int) {
	ar.natural.AddShuffle(name, numCard)
}

// GetNaturalActionResults returns the nature's actions.
func (ar ActionRequest[AD]) GetNaturalActionResults() map[string][]int {
	return ar.natural.GetResults()
}

// ValidateAction validates the action.
// It returns nil if the action is valid.
// It returns an error if the action is invalid.
func (ar ActionRequest[AD]) ValidateAction(p Player, a AD) error {
	if int(p) >= len(ar.players) {
		return fmt.Errorf("player %d is %w", p, ErrInvalidPlayer)
	}
	def := ar.players[p]

	if !def.actionable {
		return fmt.Errorf("player %d is %w", p, ErrNotActionablePlayer)
	}
	return def.validator(a)
}

// IsActionable returns true if the player is actionable.
func (ar ActionRequest[AD]) IsActionable(p Player) bool {
	if int(p) >= len(ar.players) {
		return false
	}
	return ar.players[p].actionable
}
