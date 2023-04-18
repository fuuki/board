package board

import (
	"log"
)

type Game[BP BoardProfile, PD PlayerActionDefinition] struct {
	// definition of games.
	totalPlayer  uint
	initialPhase PhaseName
	phaseMap     []*Phase[BP, PD]
	resultFn     func(*Game[BP, PD]) *Result

	// dynamic information of games.
	gameState GameState[BP, PD]
}

type GameState[BP BoardProfile, PD PlayerActionDefinition] struct {
	BoardProfile  BP
	CurrentPhase  PhaseName
	ActionProfile *ActionProfile[PD]
	ActionRequest *ActionRequest[PD]
}

func NewGame[BP BoardProfile, PD PlayerActionDefinition](
	totalPlayer uint,
	initialPhase PhaseName,
	phases []*Phase[BP, PD],
	boardProfile BP,
	resultFn func(*Game[BP, PD]) *Result,
) *Game[BP, PD] {
	state := GameState[BP, PD]{
		BoardProfile: boardProfile,
	}
	return &Game[BP, PD]{
		totalPlayer:  totalPlayer,
		initialPhase: initialPhase,
		phaseMap:     phases,
		resultFn:     resultFn,
		gameState:    state,
	}
}

// Start returns the initial action profile definition.
func (g *Game[BP, PD]) Start() bool {
	g.gameState.CurrentPhase = g.initialPhase
	g.phasePrepare()
	return g.travel()
}

// RegisterAction registers the action of players.
func (g *Game[BP, PD]) RegisterAction(p Player, a PD) error {
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
func (g *Game[BP, PD]) Next(ap *ActionProfile[PD]) bool {
	cnt := g.incrementPhase(ap)
	if !cnt {
		return false
	}
	return g.travel()
}

// travel travels the game to the next action input phase.
// bool is true if the game continues.
func (g *Game[BP, PD]) travel() bool {
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
func (g *Game[BP, PD]) incrementPhase(ap *ActionProfile[PD]) bool {
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
func (g *Game[BP, PD]) phasePrepare() {
	name := g.gameState.CurrentPhase
	log.Default().Printf("[Phase: %s]", name)
	phase := g.getPhase(name)
	ar := phase.prepare(g)
	g.gameState.ActionRequest = ar
	g.gameState.ActionProfile = NewActionProfile[PD](g.TotalPlayer())
}

// phaseExecute executes the action profile definition.
func (g *Game[BP, PD]) phaseExecute(ap *ActionProfile[PD]) {
	phase := g.getPhase(g.gameState.CurrentPhase)
	bp := g.gameState.BoardProfile
	next, bp := phase.execute(g, bp, ap)
	g.gameState.BoardProfile = bp
	g.gameState.CurrentPhase = next
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
	return g.gameState.BoardProfile
}

// IsOver returns true if the game is over.
func (g *Game[BP, PD]) IsOver() bool {
	return g.gameState.CurrentPhase == ""
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
