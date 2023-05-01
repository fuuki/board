package board

// Period is a phase of the game.
type Period[BP BoardProfile, PD PlayerActionDefinition] struct {
	count         PeriodCount
	phase         *Phase[BP, PD]
	boardProfile  BP
	actionProfile *ActionProfile[PD]
	actionRequest *ActionRequest[PD]
	status        *Status
}

// NewFirstPeriod returns a new period.
func NewFirstPeriod[BP BoardProfile, PD PlayerActionDefinition](
	phase *Phase[BP, PD],
	bpd BoardProfileDefinition[BP],
	status *Status,
) (*Period[BP, PD], *PeriodExecuteResult) {
	bp := bpd.New()
	ar, bp := phase.prepare(status, bp)
	p := &Period[BP, PD]{
		count:         0,
		phase:         phase,
		boardProfile:  bp,
		actionProfile: NewActionProfile[PD](status.TotalPlayer()),
		actionRequest: ar,
		status:        status,
	}
	r := p.executeChallenge()
	return p, r
}

// NewContinuePeriod returns a new period.
func NewContinuePeriod[BP BoardProfile, PD PlayerActionDefinition](
	count PeriodCount,
	phase *Phase[BP, PD],
	boardProfile BP,
	status *Status,
) (*Period[BP, PD], *PeriodExecuteResult) {
	ar, bp := phase.prepare(status, boardProfile)

	p := &Period[BP, PD]{
		count:         count,
		phase:         phase,
		boardProfile:  bp,
		actionProfile: NewActionProfile[PD](status.TotalPlayer()),
		actionRequest: ar,
		status:        status,
	}
	r := p.executeChallenge()
	return p, r
}

// registerAction registers an action.
func (p *Period[BP, PD]) registerAction(player Player, action PD) (*PeriodExecuteResult, error) {
	// check whether the action is valid
	err := p.actionRequest.IsValidPlayerAction(player, action)
	if err != nil {
		return nil, err
	}

	// register the action
	p.actionProfile.SetPlayerAction(player, action)
	return p.executeChallenge(), nil
}

type PeriodExecuteResult struct {
	IsCompleted bool
	NextPhase   PhaseName
}

// executeChallenge applies the action.
func (p *Period[BP, PD]) executeChallenge() *PeriodExecuteResult {
	// check whether the action is valid
	if ok := p.actionRequest.IsAllPlayerRegistered(p.actionProfile); !ok {
		return &PeriodExecuteResult{
			IsCompleted: false,
			NextPhase:   "",
		}
	}

	// apply the action
	next, bp := p.phase.execute(p.status, p.boardProfile, p.actionProfile)
	p.boardProfile = bp
	return &PeriodExecuteResult{
		IsCompleted: true,
		NextPhase:   next,
	}
}
