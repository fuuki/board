package period

import (
	"time"

	"github.com/fuuki/board/logic"
)

// playerAction is a player's action.
type playerAction[AD logic.PlayerActionDefinition] struct {
	player       logic.Player
	action       AD
	registeredAt time.Time
}

// periodAction is a period action profile.
type periodAction[AD logic.PlayerActionDefinition] struct {
	actionRequest *logic.ActionRequest[AD]
	playerActions []*playerAction[AD]

	// rawActionProfile is a raw action profile. It is derivable from playerActions.
	rawActionProfile []AD
}

// NewperiodAction returns a new period action profile.
func NewperiodAction[AD logic.PlayerActionDefinition](totalPlayer uint, request *logic.ActionRequest[AD]) *periodAction[AD] {
	return &periodAction[AD]{
		actionRequest:    request,
		playerActions:    []*playerAction[AD]{},
		rawActionProfile: make([]AD, totalPlayer),
	}
}

// insertAction inserts an action.
func (p *periodAction[AD]) insertAction(player logic.Player, action AD, now time.Time) error {
	// check whether the player is valid
	if player >= logic.Player(len(p.rawActionProfile)) {
		return logic.ErrInvalidPlayer
	}
	// check whether the player is actionable
	if !p.actionRequest.IsActionable(player) {
		return logic.ErrNotActionablePlayer
	}
	// check whether the action is valid
	if err := p.actionRequest.ValidateAction(player, action); err != nil {
		return err
	}

	p.playerActions = append(p.playerActions, &playerAction[AD]{
		player:       player,
		action:       action,
		registeredAt: now,
	})
	p.rawActionProfile[player] = action
	return nil
}

// isAllPlayerRegistered returns nil if all players are registered.
func (p *periodAction[AD]) isAllPlayerRegistered() bool {
	// check whether all players are registered
	for i, action := range p.rawActionProfile {
		// check whether the action is valid
		if !p.actionRequest.IsActionable(logic.Player(i)) {
			continue
		}
		if p.actionRequest.ValidateAction(logic.Player(i), action) != nil {
			return false
		}
	}
	return true
}

// GetActionProfile returns the action profile.
func (p *periodAction[AD]) GetActionProfile() *logic.ActionProfile[AD] {
	nat := p.actionRequest.GetNaturalActionResults()
	return logic.NewActionProfile(p.rawActionProfile, nat)
}
