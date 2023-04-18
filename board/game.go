package board

import (
	"log"
)

type Game[BP BoardProfile, AP PlayerActionDefinition] struct {
	// definition of games.
	initialPhase PhaseName
	phaseMap     []*Phase[BP, AP]
	resultFn     func(*Game[BP, AP]) *Result

	// dynamic information of games.
	gameState GameState[BP, AP]
}

type GameState[BP BoardProfile, AP PlayerActionDefinition] struct {
	BoardProfile  BP
	CurrentPhase  PhaseName
	ActionProfile *ActionProfile[AP]
	ActionRequest *ActionRequest[AP]
}

func NewGame[BP BoardProfile, AP PlayerActionDefinition](
	initialPhase PhaseName,
	phases []*Phase[BP, AP],
	boardProfile BP,
	resultFn func(*Game[BP, AP]) *Result,
) *Game[BP, AP] {
	state := GameState[BP, AP]{
		BoardProfile: boardProfile,
	}
	return &Game[BP, AP]{
		initialPhase: initialPhase,
		phaseMap:     phases,
		resultFn:     resultFn,
		gameState:    state,
	}
}

// Start returns the initial action profile definition.
func (g *Game[BP, AP]) Start() bool {
	g.gameState.CurrentPhase = g.initialPhase
	g.phasePrepare()
	return g.travel()
}

// RegisterAction registers the action of players.
func (g *Game[BP, AP]) RegisterAction(p Player, a AP) error {
	err := g.gameState.ActionRequest.IsValidPlayerAction(p, a)
	if err != nil {
		return err
	}

	g.gameState.ActionProfile.SetPlayerAction(p, a)
	if err := g.gameState.ActionRequest.IsAllPlayerRegistered(g.gameState.ActionProfile); err == nil {
		g.Next(g.gameState.ActionProfile)
	}
	return nil
}

// Next returns the next action profile definition.
// bool is true if the game continues.
func (g *Game[BP, AP]) Next(ap *ActionProfile[AP]) bool {
	cnt := g.incrementPhase(ap)
	if !cnt {
		return false
	}
	return g.travel()
}

// travel travels the game to the next action input phase.
// bool is true if the game continues.
func (g *Game[BP, AP]) travel() bool {
	for {
		ap := g.gameState.ActionProfile
		if err := g.gameState.ActionRequest.IsAllPlayerRegistered(ap); err != nil {
			return true
		}
		cnt := g.incrementPhase(ap)
		if !cnt {
			return false
		}
	}
}

// incrementPhase increments the phase.
func (g *Game[BP, AP]) incrementPhase(ap *ActionProfile[AP]) bool {
	g.phaseExecute(ap)
	log.Default().Printf("== BoardProfile ==\n%s\n", g.BoardProfile().Show())

	next := g.gameState.CurrentPhase
	if next == "" {
		return false
	}
	g.phasePrepare()
	return true
}

// phasePrepare prepares the action profile definition.
func (g *Game[BP, AP]) phasePrepare() {
	name := g.gameState.CurrentPhase
	log.Default().Printf("[Phase: %s]", name)
	phase := g.getPhase(name)
	ar := phase.prepare(g)
	g.gameState.ActionRequest = ar
	g.gameState.ActionProfile = NewActionProfile[AP](2) // FIXME: 2 is a number of players.
}

// phaseExecute executes the action profile definition.
func (g *Game[BP, AP]) phaseExecute(ap *ActionProfile[AP]) {
	phase := g.getPhase(g.gameState.CurrentPhase)
	bp := g.gameState.BoardProfile
	next, bp := phase.execute(g, bp, ap)
	g.gameState.BoardProfile = bp
	g.gameState.CurrentPhase = next
}

func (g *Game[BP, AP]) getPhase(phaseName PhaseName) *Phase[BP, AP] {
	for _, p := range g.phaseMap {
		if p.name == phaseName {
			return p
		}
	}
	return nil
}

func (g *Game[BP, AP]) BoardProfile() BP {
	return g.gameState.BoardProfile
}

// IsOver returns true if the game is over.
func (g *Game[BP, AP]) IsOver() bool {
	return g.gameState.CurrentPhase == ""
}
