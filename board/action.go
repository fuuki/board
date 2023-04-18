package board

import (
	"fmt"
	"reflect"
)

var ErrInvalidPlayer = fmt.Errorf("invalid player")
var ErrNotActionablePlayer = fmt.Errorf("not actionable")

type PlayerActionDefinition interface{}

type actionRequestDefinition[PD PlayerActionDefinition] struct {
	actionable bool
	validator  func(PD) error
}

type ActionRequest[PD PlayerActionDefinition] struct {
	defs []actionRequestDefinition[PD]
}

// NewActionRequest returns a new action request.
func NewActionRequest[PD PlayerActionDefinition](totalPlayer uint) *ActionRequest[PD] {
	defs := make([]actionRequestDefinition[PD], totalPlayer)
	for i := uint(0); i < totalPlayer; i++ {
		defs[i] = actionRequestDefinition[PD]{
			actionable: false,
			validator:  func(PD) error { return nil },
		}
	}
	return &ActionRequest[PD]{
		defs: defs,
	}
}

// RegisterValidator registers the validator for the player.
func (ar ActionRequest[PD]) RegisterValidator(p Player, fn func(PD) error) {
	ar.defs[p] = actionRequestDefinition[PD]{
		actionable: true,
		validator:  fn,
	}
}

// IsValidPlayerAction returns nil if the action is valid.
func (ar ActionRequest[PD]) IsValidPlayerAction(p Player, a PD) error {
	if int(p) >= len(ar.defs) {
		return fmt.Errorf("player %d is %w", p, ErrInvalidPlayer)
	}
	def := ar.defs[p]
	if !def.actionable {
		return fmt.Errorf("player %d is %w", p, ErrNotActionablePlayer)
	}
	return def.validator(a)
}

// IsAllPlayerRegistered returns nil if all players are registered.
func (ar ActionRequest[PD]) IsAllPlayerRegistered(ap *ActionProfile[PD]) error {
	for p, def := range ar.defs {
		if !def.actionable {
			continue
		}
		a := ap.Player(Player(p))
		if reflect.ValueOf(a).IsZero() {
			return fmt.Errorf("player %d is not registered", p)
		}
	}
	return nil
}

type ActionProfile[PD PlayerActionDefinition] struct {
	playerActions []PD
}

// NewActionProfile returns a new action profile.
func NewActionProfile[PD PlayerActionDefinition](totalPlayer uint) *ActionProfile[PD] {
	ap := &ActionProfile[PD]{
		playerActions: make([]PD, totalPlayer),
	}
	return ap
}

// NewActionProfileWithAction returns a new action profile with the action.
func NewActionProfileWithAction[PD PlayerActionDefinition](actions []PD) *ActionProfile[PD] {
	ap := &ActionProfile[PD]{
		playerActions: actions,
	}
	return ap
}

// Player returns the player's action.
func (ap *ActionProfile[PD]) Player(p Player) PD {
	return ap.playerActions[p]
}

// SetPlayerAction sets the player's action.
func (ap *ActionProfile[PD]) SetPlayerAction(p Player, a PD) {
	ap.playerActions[p] = a
}
