package board

import (
	"log"
)

type PeriodCount int

type Game[BP BoardProfile, PD PlayerActionDefinition] struct {
	// definition of games.
	totalPlayer            uint
	initialPhase           PhaseName
	boardProfileDefinition BoardProfileDefinition[BP]
	phaseMap               []*Phase[BP, PD]

	// dynamic information of games.
	periodHistory []Period[BP, PD]

	// options
	phaseChangeChan chan<- PeriodCount
}

// Period is a phase of the game.
type Period[BP BoardProfile, PD PlayerActionDefinition] struct {
	count         PeriodCount
	phase         PhaseName
	boardProfile  BP
	actionProfile *ActionProfile[PD]
	actionRequest *ActionRequest[PD]
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
		totalPlayer:            totalPlayer,
		initialPhase:           initialPhase,
		phaseMap:               phases,
		boardProfileDefinition: bpd,
		periodHistory:          []Period[BP, PD]{},
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
	return &g.periodHistory[len(g.periodHistory)-1]
}

// addPeriod adds a period to the game.
func (g *Game[BP, PD]) addPeriod(phase PhaseName) {
	var bp BP
	var count PeriodCount
	if p := g.CurrentPeriod(); p == nil {
		bp = g.boardProfileDefinition.New()
		count = 0
	} else {
		bp = g.boardProfileDefinition.Clone(p.boardProfile)
		count = p.count + 1
	}

	g.periodHistory = append(g.periodHistory, Period[BP, PD]{
		count:         count,
		phase:         phase,
		boardProfile:  bp,
		actionProfile: NewActionProfile[PD](g.totalPlayer),
	})
}

// Start returns the initial action profile definition.
func (g *Game[BP, PD]) Start() bool {
	g.addPeriod(g.initialPhase)
	g.phasePrepare()
	return g.travel()
}

// RegisterAction registers the action of players.
// It checks whether the action is valid.
func (g *Game[BP, PD]) RegisterAction(p Player, a PD) (isAllRegistered bool, gameContinues bool, err error) {
	cp := g.CurrentPeriod()
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
		ap := g.CurrentPeriod().actionProfile
		if err := g.CurrentPeriod().actionRequest.IsAllPlayerRegistered(ap); err != nil {
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
func (g *Game[BP, PD]) incrementPeriod(ap *ActionProfile[PD]) bool {
	// execute the current phase
	next := g.phaseExecute(ap)
	// log.Default().Printf("== BoardProfile ==\n%s\n", g.BoardProfile().Show())
	if next == "" {
		close(g.phaseChangeChan)
		return false
	}

	// go to next period
	g.addPeriod(next)
	g.phaseChangeChan <- g.CurrentPeriod().count

	// Prepare the next phase
	g.phasePrepare()
	return true
}

// phasePrepare prepares the action profile definition.
func (g *Game[BP, PD]) phasePrepare() {
	cp := g.CurrentPeriod()
	log.Default().Printf("[Phase: %s]", cp.phase)
	phase := g.getPhase(cp.phase)
	ar := phase.prepare(g)
	cp.actionRequest = ar
	cp.actionProfile = NewActionProfile[PD](g.TotalPlayer())
}

// phaseExecute executes the action profile definition.
// returns the next phase name.
func (g *Game[BP, PD]) phaseExecute(ap *ActionProfile[PD]) PhaseName {
	cp := g.CurrentPeriod()
	phase := g.getPhase(cp.phase)
	bp := cp.boardProfile
	next, bp := phase.execute(g, bp, ap)
	cp.boardProfile = bp
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

func (g *Game[BP, PD]) BoardProfile() BP {
	return g.CurrentPeriod().boardProfile
}

// IsOver returns true if the game is over.
func (g *Game[BP, PD]) IsOver() bool {
	return g.CurrentPeriod().phase == ""
}

// Players returns the array of players.
func (g *Game[BP, PD]) Players() []Player {
	result := make([]Player, g.totalPlayer)
	for i := uint(0); i < g.totalPlayer; i++ {
		result[i] = Player(i)
	}
	return result
}

// TotalPlayer returns the number of players.
func (g *Game[BP, PD]) TotalPlayer() uint {
	return g.totalPlayer
}
