package board

import (
	"log"
	"time"

	"github.com/fuuki/board/internal/mech"
	"github.com/fuuki/board/internal/period"
	"github.com/fuuki/board/logic"
)

type Table[AD logic.PlayerActionDefinition, BP logic.BoardProfile, CF logic.Config] struct {
	// static information of games.
	mech   *mech.Mech[AD, BP, CF]
	config CF

	// dynamic information of games.
	periodHistory []*period.Period[AD, BP, CF]
	isOver        bool

	// channels
	eventChan *eventChan
}

// newTable returns a new table.
func newTable[AD logic.PlayerActionDefinition, BP logic.BoardProfile, CF logic.Config](
	mech *mech.Mech[AD, BP, CF], config CF,
) (*Table[AD, BP, CF], <-chan *Event) {
	rcvCh, ch := newEventChan()
	table := &Table[AD, BP, CF]{
		mech:      mech,
		config:    config,
		eventChan: ch,
	}
	return table, rcvCh
}

// initGame initializes the game.
// this method should called after channels are set.
func (t *Table[AD, BP, CF]) InitGame() {
	// initialize the first period
	t.createFirstPeriod()
}

// createFirstPeriod creates the first period.
func (t *Table[AD, BP, CF]) createFirstPeriod() {
	phase := t.mech.GetInitialPhase()
	bp := t.mech.GetBoardProfileDefinition().New()
	t.createPeriod(0, phase, bp)
}

// createContinuePeriod creates the continue period.
func (t *Table[AD, BP, CF]) createContinuePeriod(phaseName logic.PhaseName) {
	phase := t.mech.GetPhase(phaseName)
	count := t.currentPeriod().GetCount() + 1
	bp := t.mech.GetBoardProfileDefinition().Clone(t.currentPeriod().GetBoardProfile())
	t.createPeriod(count, phase, bp)
}

// createPeriod creates a new period.
func (t *Table[AD, BP, CF]) createPeriod(count int, phase logic.Phase[AD, BP, CF], bp BP) error {
	pr, result, err := period.NewPeriod(count, phase, bp, t.config)
	if err != nil {
		return err
	}
	t.registerPeriod(pr)
	if result.IsCompleted {
		t.afterPeriodCompleted(result.NextPhase)
	}
	return nil
}

// registerPeriod adds a new period to the game.
func (t *Table[AD, BP, CF]) registerPeriod(pr *period.Period[AD, BP, CF]) {
	t.eventChan.sendPhaseChange(pr.GetCount())
	t.periodHistory = append(t.periodHistory, pr)
}

// RegisterAction registers the action of players.
// It checks whether the action is valid.
func (t *Table[AD, BP, CF]) RegisterAction(p logic.Player, a AD) error {
	log.Default().Printf("register action: %d", p)
	result, err := t.currentPeriod().RegisterAction(p, a, time.Now())
	if err != nil {
		return err
	}

	if result.IsCompleted {
		t.afterPeriodCompleted(result.NextPhase)
	}
	return nil
}

// afterPeriodCompleted is called when a period is completed.
func (t *Table[AD, BP, CF]) afterPeriodCompleted(nextPhase logic.PhaseName) {
	if nextPhase == "" {
		t.eventChan.close()
		t.isOver = true
		return
	}
	t.createContinuePeriod(nextPhase)
}

// currentPeriod returns the current period.
func (t *Table[AD, BP, CF]) currentPeriod() *period.Period[AD, BP, CF] {
	if len(t.periodHistory) == 0 {
		return nil
	}
	return t.periodHistory[len(t.periodHistory)-1]
}

// IsOver returns true if the game is over.
func (t *Table[AD, BP, CF]) IsOver() bool {
	return t.isOver
}

// CurrentBoardProfile returns the current board profile.
func (t *Table[AD, BP, CF]) CurrentBoardProfile() logic.BoardProfile {
	return t.currentPeriod().GetBoardProfile()
}
