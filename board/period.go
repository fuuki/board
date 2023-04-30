package board

// Period is a phase of the game.
type Period[BP BoardProfile, PD PlayerActionDefinition] struct {
	count         PeriodCount
	phase         *Phase[BP, PD]
	boardProfile  BP
	actionProfile *ActionProfile[PD]
	actionRequest *ActionRequest[PD]
}

// NewPeriod returns a new period.
func NewPeriod[BP BoardProfile, PD PlayerActionDefinition](
	count PeriodCount,
	phase *Phase[BP, PD],
	boardProfile BP,
	status *Status,
) *Period[BP, PD] {
	ar, bp := phase.prepare(status, boardProfile)

	return &Period[BP, PD]{
		count:         count,
		phase:         phase,
		boardProfile:  bp,
		actionProfile: NewActionProfile[PD](status.TotalPlayer()),
		actionRequest: ar,
	}
}

// Execute is called when the period is ended.
func (p *Period[BP, PD]) Execute(status *Status) PhaseName {
	next, bp := p.phase.execute(status, p.boardProfile, p.actionProfile)
	p.boardProfile = bp
	return next
}
