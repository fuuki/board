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
	phaseChangeChan chan<- int
}

// initGame initializes the game.
// this method should called after channels are set.
func (g *Table[AD, BP, CF]) InitGame() {
	// initialize the first period
	g.createFirstPeriod()
}

// createFirstPeriod creates the first period.
func (g *Table[AD, BP, CF]) createFirstPeriod() {
	initialPhase := g.mech.GetInitialPhase()
	pr, result := period.NewFirstPeriod(initialPhase, g.mech.GetBoardProfileDefinition(), g.config)
	g.registerPeriod(pr)

	if result.IsCompleted {
		g.afterPeriodCompleted(result.NextPhase)
	}
}

// createContinuePeriod creates a new period.
func (g *Table[AD, BP, CF]) createContinuePeriod(phaseName logic.PhaseName) {
	phase := g.mech.GetPhase(phaseName)
	count := g.currentPeriod().GetCount() + 1
	bp := g.mech.GetBoardProfileDefinition().Clone(g.currentPeriod().GetBoardProfile())

	pr, result := period.NewContinuePeriod(count, phase, bp, g.config)
	g.registerPeriod(pr)

	if result.IsCompleted {
		g.afterPeriodCompleted(result.NextPhase)
	}
}

// registerPeriod adds a new period to the game.
func (g *Table[AD, BP, CF]) registerPeriod(pr *period.Period[AD, BP, CF]) {
	if g.phaseChangeChan != nil {
		g.phaseChangeChan <- pr.GetCount()
	}
	g.periodHistory = append(g.periodHistory, pr)
}

// RegisterAction registers the action of players.
// It checks whether the action is valid.
func (g *Table[AD, BP, CF]) RegisterAction(p logic.Player, a AD) error {
	log.Default().Printf("register action: %d", p)
	result, err := g.currentPeriod().RegisterAction(p, a, time.Now())
	if err != nil {
		return err
	}

	if result.IsCompleted {
		g.afterPeriodCompleted(result.NextPhase)
	}
	return nil
}

// afterPeriodCompleted is called when a period is completed.
func (g *Table[AD, BP, CF]) afterPeriodCompleted(nextPhase logic.PhaseName) {
	if nextPhase == "" {
		if g.phaseChangeChan != nil {
			close(g.phaseChangeChan)
		}
		g.isOver = true
		return
	}
	g.createContinuePeriod(nextPhase)
}

// currentPeriod returns the current period.
func (g *Table[AD, BP, CF]) currentPeriod() *period.Period[AD, BP, CF] {
	if len(g.periodHistory) == 0 {
		return nil
	}
	return g.periodHistory[len(g.periodHistory)-1]
}

// IsOver returns true if the game is over.
func (g *Table[AD, BP, CF]) IsOver() bool {
	return g.isOver
}

// CurrentBoardProfile returns the current board profile.
func (g *Table[AD, BP, CF]) CurrentBoardProfile() logic.BoardProfile {
	return g.currentPeriod().GetBoardProfile()
}
