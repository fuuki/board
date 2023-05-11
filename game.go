package board

import (
	"log"
)

type PeriodCount int

type Game[BP BoardProfile, PD PlayerActionDefinition] struct {
	// definition of games.
	initialPhase           PhaseName
	boardProfileDefinition BoardProfileDefinition[BP]
	phaseMap               []*Phase[BP, PD]
	totalPlayer            uint

	// dynamic information of games.
	periodHistory []*period[BP, PD]
	isOver        bool

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

var _ GameOption[BoardProfile, PlayerActionDefinition] = (*phaseChangeChan[BoardProfile, PlayerActionDefinition])(nil)

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
		periodHistory:          []*period[BP, PD]{},
	}
	for _, o := range options {
		o.apply(g)
	}
	return g
}

// initGame initializes the game.
// this method should called after channels are set.
func (g *Game[BP, PD]) InitGame() {
	// initialize the first period
	g.createFirstPeriod()
}

// createFirstPeriod creates the first period.
func (g *Game[BP, PD]) createFirstPeriod() {
	phase := g.getPhase(g.initialPhase)
	pr, result := newFirstPeriod(phase, g.boardProfileDefinition, g.Status())
	g.registerPeriod(pr)

	if result.IsCompleted {
		g.afterPeriodCompleted(result.NextPhase)
	}
}

// createContinuePeriod creates a new period.
func (g *Game[BP, PD]) createContinuePeriod(phaseName PhaseName) {
	phase := g.getPhase(phaseName)
	count := g.currentPeriod().count + 1
	bp := g.boardProfileDefinition.Clone(g.currentPeriod().boardProfile)

	pr, result := newContinuePeriod(count, phase, bp, g.Status())
	g.registerPeriod(pr)

	if result.IsCompleted {
		g.afterPeriodCompleted(result.NextPhase)
	}
}

// registerPeriod adds a new period to the game.
func (g *Game[BP, PD]) registerPeriod(pr *period[BP, PD]) {
	if g.phaseChangeChan != nil {
		g.phaseChangeChan <- pr.count
	}
	g.periodHistory = append(g.periodHistory, pr)
}

// RegisterAction registers the action of players.
// It checks whether the action is valid.
func (g *Game[BP, PD]) RegisterAction(p Player, a PD) error {
	log.Default().Printf("register action: %d", p)
	result, err := g.currentPeriod().registerAction(p, a)
	if err != nil {
		return err
	}

	if result.IsCompleted {
		g.afterPeriodCompleted(result.NextPhase)
	}
	return nil
}

// afterPeriodCompleted is called when a period is completed.
func (g *Game[BP, PD]) afterPeriodCompleted(nextPhase PhaseName) {
	if nextPhase == "" {
		if g.phaseChangeChan != nil {
			close(g.phaseChangeChan)
		}
		g.isOver = true
		return
	}
	g.createContinuePeriod(nextPhase)
}

func (g *Game[BP, PD]) getPhase(phaseName PhaseName) *Phase[BP, PD] {
	for _, p := range g.phaseMap {
		if p.name == phaseName {
			return p
		}
	}
	return nil
}

// currentPeriod returns the current period.
func (g *Game[BP, PD]) currentPeriod() *period[BP, PD] {
	if len(g.periodHistory) == 0 {
		return nil
	}
	return g.periodHistory[len(g.periodHistory)-1]
}

// IsOver returns true if the game is over.
func (g *Game[BP, PD]) IsOver() bool {
	return g.isOver
}

// Status returns the current status.
func (g *Game[BP, PD]) Status() *Status {
	return NewStatus(g)
}

// CurrentBoardProfile returns the current board profile.
func (g *Game[BP, PD]) CurrentBoardProfile() BoardProfile {
	return g.currentPeriod().boardProfile
}
