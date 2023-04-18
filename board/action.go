package board

import (
	"fmt"
	"reflect"
)

var ErrInvalidPlayer = fmt.Errorf("invalid player")
var ErrNotActionablePlayer = fmt.Errorf("not actionable")

type PlayerActionDefinition interface{}

type actionRequestDefinition[AP PlayerActionDefinition] struct {
	actionable bool
	validator  func(AP) error
}

type ActionRequest[AP PlayerActionDefinition] struct {
	defs map[Player]actionRequestDefinition[AP]
}

// NewActionRequest returns a new action request.
func NewActionRequest[AP PlayerActionDefinition](totalPlayer uint) *ActionRequest[AP] {
	defs := make(map[Player]actionRequestDefinition[AP], totalPlayer)
	for i := uint(0); i < totalPlayer; i++ {
		defs[Player(i)] = actionRequestDefinition[AP]{
			actionable: false,
			validator:  func(AP) error { return nil },
		}
	}
	return &ActionRequest[AP]{
		defs: defs,
	}
}

// RegisterValidator registers the validator for the player.
func (ar ActionRequest[AP]) RegisterValidator(p Player, fn func(AP) error) {
	ar.defs[p] = actionRequestDefinition[AP]{
		actionable: true,
		validator:  fn,
	}
}

// IsValidPlayerAction returns nil if the action is valid.
func (ar ActionRequest[AP]) IsValidPlayerAction(p Player, a AP) error {
	def, ok := ar.defs[p]
	if !ok {
		return fmt.Errorf("player %d is %w", p, ErrInvalidPlayer)
	}
	if !def.actionable {
		return fmt.Errorf("player %d is %w", p, ErrNotActionablePlayer)
	}
	return def.validator(a)
}

// IsAllPlayerRegistered returns nil if all players are registered.
func (ar ActionRequest[AP]) IsAllPlayerRegistered(ap *ActionProfile[AP]) error {
	for p, def := range ar.defs {
		if !def.actionable {
			continue
		}
		a := ap.Player(p)
		if reflect.ValueOf(a).IsZero() {
			return fmt.Errorf("player %d is not registered", p)
		}
	}
	return nil
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
