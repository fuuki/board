package period

import (
	"time"

	"github.com/fuuki/board/logic"
)

// Period is a period of the game.
type Period[AD logic.PlayerActionDefinition, BP logic.BoardProfile, CF logic.Config] struct {
	config       CF
	count        int
	phase        logic.Phase[AD, BP, CF]
	boardProfile BP
	periodAction *periodAction[AD]
}

// NewPeriod returns a new period.
func NewPeriod[AD logic.PlayerActionDefinition, BP logic.BoardProfile, CF logic.Config](
	count int,
	phase logic.Phase[AD, BP, CF],
	boardProfile BP,
	config CF,
) (*Period[AD, BP, CF], *periodExecuteResult, error) {
	ar, bp, err := phase.Prepare(config, boardProfile)
	if err != nil {
		return nil, nil, err
	}

	p := &Period[AD, BP, CF]{
		config:       config,
		count:        count,
		phase:        phase,
		boardProfile: bp,
		periodAction: NewperiodAction(config.TotalPlayer(), ar),
	}
	r, err := p.executeChallenge()
	if err != nil {
		return nil, nil, err
	}
	return p, r, err
}

// RegisterAction registers an action.
func (p *Period[AD, BP, CF]) RegisterAction(player logic.Player, action AD, now time.Time) (*periodExecuteResult, error) {
	// check whether the action is valid
	err := p.periodAction.insertAction(player, action, now)
	if err != nil {
		return nil, err
	}

	result, err := p.executeChallenge()
	return result, err
}

type periodExecuteResult struct {
	IsCompleted bool
	NextPhase   logic.PhaseName
}

// executeChallenge executes if the period action is completed.
func (p *Period[AD, BP, CF]) executeChallenge() (*periodExecuteResult, error) {
	// check whether the action is valid
	if ok := p.periodAction.isAllPlayerRegistered(); !ok {
		return &periodExecuteResult{
			IsCompleted: false,
			NextPhase:   "",
		}, nil
	}

	// apply the action
	next, bp, err := p.phase.Execute(p.config, p.boardProfile, p.periodAction.GetActionProfile())
	if err != nil {
		return nil, err
	}
	p.boardProfile = bp
	return &periodExecuteResult{
		IsCompleted: true,
		NextPhase:   next,
	}, nil
}

// GetBoardProfile returns the board profile.
func (p *Period[AD, BP, CF]) GetBoardProfile() BP {
	return p.boardProfile
}

// GetCount returns the count.
func (p *Period[AD, BP, CF]) GetCount() int {
	return p.count
}
