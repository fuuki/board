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
	return logic.NewActionProfile(p.rawActionProfile)
}

// Period is a period of the game.
type Period[AD logic.PlayerActionDefinition, BP logic.BoardProfile, CF logic.Config] struct {
	config       CF
	count        int
	phase        logic.Phase[AD, BP, CF]
	boardProfile BP
	periodAction *periodAction[AD]
}

// NewFirstPeriod returns a new period.
func NewFirstPeriod[AD logic.PlayerActionDefinition, BP logic.BoardProfile, CF logic.Config](
	phase logic.Phase[AD, BP, CF],
	bpd logic.BoardProfileDefinition[BP],
	config CF,
) (*Period[AD, BP, CF], *periodExecuteResult) {
	bp := bpd.New()
	ar, bp := phase.Prepare(config, bp)
	p := &Period[AD, BP, CF]{
		config:       config,
		count:        0,
		phase:        phase,
		boardProfile: bp,
		periodAction: NewperiodAction(config.TotalPlayer(), ar),
	}
	r := p.executeChallenge()
	return p, r
}

// NewContinuePeriod returns a new period.
func NewContinuePeriod[AD logic.PlayerActionDefinition, BP logic.BoardProfile, CF logic.Config](
	count int,
	phase logic.Phase[AD, BP, CF],
	boardProfile BP,
	config CF,
) (*Period[AD, BP, CF], *periodExecuteResult) {
	ar, bp := phase.Prepare(config, boardProfile)
	p := &Period[AD, BP, CF]{
		config:       config,
		count:        count,
		phase:        phase,
		boardProfile: bp,
		periodAction: NewperiodAction(config.TotalPlayer(), ar),
	}
	r := p.executeChallenge()
	return p, r
}

// RegisterAction registers an action.
func (p *Period[AD, BP, CF]) RegisterAction(player logic.Player, action AD, now time.Time) (*periodExecuteResult, error) {
	// check whether the action is valid
	err := p.periodAction.insertAction(player, action, now)
	if err != nil {
		return nil, err
	}

	return p.executeChallenge(), nil
}

type periodExecuteResult struct {
	IsCompleted bool
	NextPhase   logic.PhaseName
}

// executeChallenge executes if the period action is completed.
func (p *Period[AD, BP, CF]) executeChallenge() *periodExecuteResult {
	// check whether the action is valid
	if ok := p.periodAction.isAllPlayerRegistered(); !ok {
		return &periodExecuteResult{
			IsCompleted: false,
			NextPhase:   "",
		}
	}

	// set the nature action
	// TODO: implement
	// p.actionProfile.natureActions = p.actionRequest.naturalPlayer.GetResults()

	// apply the action
	next, bp := p.phase.Execute(p.config, p.boardProfile, p.periodAction.GetActionProfile())
	p.boardProfile = bp
	return &periodExecuteResult{
		IsCompleted: true,
		NextPhase:   next,
	}
}

// GetBoardProfile returns the board profile.
func (p *Period[AD, BP, CF]) GetBoardProfile() BP {
	return p.boardProfile
}

// GetCount returns the count.
func (p *Period[AD, BP, CF]) GetCount() int {
	return p.count
}
