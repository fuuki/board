package board

import (
	"fmt"
	"log"
)

type CurrentPhase struct {
	phaseName PhaseName
}

type Game[BP BoardProfile, AP PlayerActionDefinition] struct {
	// definition of games.
	initialPhase PhaseName
	phaseMap     []*Phase[BP, AP]
	resultFn     func(*Game[BP, AP]) *Result
	totalPlayers uint

	// dynamic information of games.
	gameState GameState[BP, AP]
}

type GameState[BP BoardProfile, AP PlayerActionDefinition] struct {
	BoardProfile BP
	CurrentPhase *CurrentPhase
}

func NewGame[BP BoardProfile, AP PlayerActionDefinition](
	totalPlayers uint,
	initialPhase PhaseName,
	phases []*Phase[BP, AP],
	boardProfile BP,
	resultFn func(*Game[BP, AP]) *Result,
) *Game[BP, AP] {
	state := GameState[BP, AP]{
		BoardProfile: boardProfile,
	}
	return &Game[BP, AP]{
		totalPlayers: totalPlayers,
		initialPhase: initialPhase,
		phaseMap:     phases,
		resultFn:     resultFn,
		gameState:    state,
	}
}

func (g *Game[BP, AP]) Play(inputer ActionInputer[AP]) {
	var ap *ActionProfile[AP]
	for {
		cnt, apr := g.Next(ap)
		if !cnt {
			break
		}
		ap = inputer.Input(apr)
	}
	result := g.resultFn(g)
	fmt.Printf("%+v", result)
}

// Start returns the initial action profile definition.
// bool is true if the game continues.
func (g *Game[BP, AP]) Start() (bool, ActionRequest[AP]) {
	return g.Next(nil)
}

// Next returns the next action profile definition.
// bool is true if the game continues.
func (g *Game[BP, AP]) Next(ap *ActionProfile[AP]) (bool, ActionRequest[AP]) {
	var next PhaseName
	if g.gameState.CurrentPhase == nil {
		next = g.initialPhase
	} else {
		phase := g.getPhase(g.gameState.CurrentPhase.phaseName)
		bp := g.gameState.BoardProfile
		next, bp = phase.execute(g, bp, ap)
		g.gameState.BoardProfile = bp
	}
	log.Default().Printf("== BoardProfile ==\n%s\n", g.BoardProfile().Show())

	if next == "" {
		return false, nil
	}

	log.Default().Printf("[Phase: %s]", next)
	nextPhase := g.getPhase(next)
	apr := nextPhase.prepare(g)
	g.gameState.CurrentPhase = &CurrentPhase{
		phaseName: next,
	}
	return true, apr
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

// TotalPlayers returns the total number of players.
func (g *Game[BP, AP]) TotalPlayers() uint {
	return g.totalPlayers
}

// Players returns the players.
func (g *Game[BP, AP]) Players() []Player {
	players := make([]Player, g.totalPlayers)
	for i := uint(0); i < g.totalPlayers; i++ {
		players[i] = Player(i)
	}
	return players
}
