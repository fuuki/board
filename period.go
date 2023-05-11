package board

// period is a phase of the game.
type period[BP BoardProfile, PD PlayerActionDefinition] struct {
	count         PeriodCount
	phase         *Phase[BP, PD]
	boardProfile  BP
	actionProfile *ActionProfile[PD]
	actionRequest *ActionRequest[PD]
	status        *Status
}

// newFirstPeriod returns a new period.
func newFirstPeriod[BP BoardProfile, PD PlayerActionDefinition](
	phase *Phase[BP, PD],
	bpd BoardProfileDefinition[BP],
	status *Status,
) (*period[BP, PD], *periodExecuteResult) {
	bp := bpd.New()
	ar, bp := phase.prepare(status, bp)
	p := &period[BP, PD]{
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

// newContinuePeriod returns a new period.
func newContinuePeriod[BP BoardProfile, PD PlayerActionDefinition](
	count PeriodCount,
	phase *Phase[BP, PD],
	boardProfile BP,
	status *Status,
) (*period[BP, PD], *periodExecuteResult) {
	ar, bp := phase.prepare(status, boardProfile)

	p := &period[BP, PD]{
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
func (p *period[BP, PD]) registerAction(player Player, action PD) (*periodExecuteResult, error) {
	// check whether the action is valid
	err := p.actionRequest.isValidPlayerAction(player, action)
	if err != nil {
		return nil, err
	}

	// register the action
	p.actionProfile.SetPlayerAction(player, action)
	return p.executeChallenge(), nil
}

type periodExecuteResult struct {
	IsCompleted bool
	NextPhase   PhaseName
}

// executeChallenge executes if the period action is completed.
func (p *period[BP, PD]) executeChallenge() *periodExecuteResult {
	// check whether the action is valid
	if ok := p.actionRequest.isAllPlayerRegistered(p.actionProfile); !ok {
		return &periodExecuteResult{
			IsCompleted: false,
			NextPhase:   "",
		}
	}

	// set the nature action
	p.actionProfile.natureActions = p.actionRequest.naturalPlayer.GetResults()

	// apply the action
	next, bp := p.phase.execute(p.status, p.boardProfile, p.actionProfile)
	p.boardProfile = bp
	return &periodExecuteResult{
		IsCompleted: true,
		NextPhase:   next,
	}
}
