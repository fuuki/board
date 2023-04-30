package board

import "log"

type PeriodCount int

type Game[BP BoardProfile, PD PlayerActionDefinition] struct {
	// definition of games.
	initialPhase           PhaseName
	boardProfileDefinition BoardProfileDefinition[BP]
	phaseMap               []*Phase[BP, PD]

	// status of games.
	status *Status

	// dynamic information of games.
	currentPeriod *Period[BP, PD]
	periodHistory []*Period[BP, PD]

	// options
	phaseChangeChan chan<- PeriodCount
}

type GameOption[BP BoardProfile, PD PlayerActionDefinition] interface {
	apply(*Game[BP, PD])
}

// phaseChangeChan はフェーズが次に進むときに、フェーズNoをチャネルに送信するオプション.
type phaseChangeChan[BP BoardProfile, PD PlayerActionDefinition] struct {
	ch chan<- PeriodCount
}

// apply はオプションを適用する.
func (o *phaseChangeChan[BP, PD]) apply(g *Game[BP, PD]) {
	g.phaseChangeChan = o.ch
}

// PhaseChangeChan はフェーズが次に進むときに、フェーズNoをチャネルに送信するオプション.
func PhaseChangeChan[BP BoardProfile, PD PlayerActionDefinition](ch chan<- PeriodCount) GameOption[BP, PD] {
	return &phaseChangeChan[BP, PD]{
		ch: ch,
	}
}

// NewGame returns a new game.
func NewGame[BP BoardProfile, PD PlayerActionDefinition](
	totalPlayer uint,
	initialPhase PhaseName,
	phases []*Phase[BP, PD],
	bpd BoardProfileDefinition[BP],
	options ...GameOption[BP, PD],
) *Game[BP, PD] {
	g := &Game[BP, PD]{
		status:                 NewStatus(totalPlayer),
		initialPhase:           initialPhase,
		phaseMap:               phases,
		boardProfileDefinition: bpd,
		periodHistory:          []*Period[BP, PD]{},
	}
	for _, o := range options {
		o.apply(g)
	}
	return g
}

// CurrentPeriod returns the current period.
func (g *Game[BP, PD]) CurrentPeriod() *Period[BP, PD] {
	if len(g.periodHistory) == 0 {
		return nil
	}
	return g.periodHistory[len(g.periodHistory)-1]
}

// preparePeriod adds a new period to the game.
func (g *Game[BP, PD]) preparePeriod(phaseName PhaseName) {
	log.Default().Printf("[Phase: %s]", phaseName)
	g.phaseChangeChan <- g.currentPeriod.count

	var bp BP
	var count PeriodCount
	if p := g.currentPeriod; p == nil {
		bp = g.boardProfileDefinition.New()
		count = 0
	} else {
		bp = g.boardProfileDefinition.Clone(p.boardProfile)
		count = p.count + 1
	}
	phase := g.getPhase(phaseName)
	g.currentPeriod = NewPeriod[BP, PD](count, phase, bp, g.status)
}

// Start returns the initial action profile definition.
func (g *Game[BP, PD]) Start() bool {
	g.phasePrepare(g.initialPhase)
	return g.travel()
}

// RegisterAction registers the action of players.
// It checks whether the action is valid.
func (g *Game[BP, PD]) RegisterAction(p Player, a PD) (isAllRegistered bool, gameContinues bool, err error) {
	cp := g.currentPeriod
	err = cp.actionRequest.IsValidPlayerAction(p, a)
	if err != nil {
		return
	}

	cp.actionProfile.SetPlayerAction(p, a)
	if err = cp.actionRequest.IsAllPlayerRegistered(cp.actionProfile); err == nil {
		isAllRegistered = true
		gameContinues = g.DirectRegisterAction(cp.actionProfile)
		return
	}
	return
}

// DirectRegisterAction registers the action of players.
// It does not check whether the action is valid.
// Use it for debugging and testing.
func (g *Game[BP, PD]) DirectRegisterAction(ap *ActionProfile[PD]) bool {
	cnt := g.incrementPeriod(ap)
	if !cnt {
		return false
	}
	return g.travel()
}

// travel travels the game to the next action input phase.
// it skips the phase that does not require action input.
// bool is true if the game continues.
func (g *Game[BP, PD]) travel() bool {
	for {
		ap := g.currentPeriod.actionProfile
		if err := g.currentPeriod.actionRequest.IsAllPlayerRegistered(ap); err != nil {
			return true
		}
		cnt := g.incrementPeriod(ap)
		if !cnt {
			return false
		}
	}
}

// incrementPeriod increments the phase.
// bool is true if the game continues.
func (g *Game[BP, PD]) incrementPeriod(ap *ActionProfile[PD]) {
	// execute the current phase
	next := g.phaseExecute(ap)

	// Prepare the next phase
	g.periodHistory = append(g.periodHistory, g.currentPeriod)
	if next == "" {
		g.currentPeriod = nil
		close(g.phaseChangeChan)
		return
	}
	g.preparePeriod(next)
	return
}

// phaseExecute executes the action profile definition.
// returns the next phase name.
func (g *Game[BP, PD]) phaseExecute(ap *ActionProfile[PD]) PhaseName {
	cp := g.currentPeriod
	next := cp.Execute(g.status)
	return next
}

func (g *Game[BP, PD]) getPhase(phaseName PhaseName) *Phase[BP, PD] {
	for _, p := range g.phaseMap {
		if p.name == phaseName {
			return p
		}
	}
	return nil
}

// IsOver returns true if the game is over.
func (g *Game[BP, PD]) IsOver() bool {
	return g.currentPeriod == nil
}

// Players returns the array of players.
func (g *Game[BP, PD]) Players() []Player {
	result := make([]Player, g.status.TotalPlayer())
	for i := uint(0); i < g.status.TotalPlayer(); i++ {
		result[i] = Player(i)
	}
	return result
}
