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
