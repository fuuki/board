package daifugo

import (
	"github.com/fuuki/board/action"
	"github.com/fuuki/board/game"
	"github.com/fuuki/board/player"
	"github.com/fuuki/board/resource"
)

const (
	DealPhase game.PhaseName = "deal"
	PlayPhase game.PhaseName = "play"
)

func Play() {
	playerNum := 3
	g := daifugoGame(playerNum)
	g.Play(&action.InteractiveActionInputer{})
}

// daifugoGame returns a game of rock-paper-scissors.
func daifugoGame(playerNum int) *game.Game {
	rp := resourceProfile(playerNum)

	p1 := dealPhase(playerNum)
	g := game.NewGame(DealPhase, []*game.Phase{p1}, rp)
	return g
}

// resourceProfile returns a resource profile of rock-paper-scissors.
func resourceProfile(playerNum int) *resource.ResourceProfile {
	rp := resource.NewResourceProfile()
	for i := 0; i < playerNum; i++ {
		rp.AddResource(player.Player(i), resource.NewResource())
	}
	return rp
}

func dealPhase(playerNum int) *game.Phase {
	prepare := func(_ *game.Game) *action.ActionProfileDefinition {
		// Shuffle deck
		return nil
	}

	execute := func(g *game.Game, ap *action.ActionProfile) game.PhaseName {
		g.ResourceProfile().AddResource(0, resource.NewResource())
		return PlayPhase
	}

	p := game.NewPhase(DealPhase, prepare, execute)
	return p
}

func playPhase() *game.Phase {
	prepare := func(_ *game.Game) *action.ActionProfileDefinition {
		// Define action profile
		return nil
	}

	execute := func(g *game.Game, ap *action.ActionProfile) game.PhaseName {
		return PlayPhase
	}

	p := game.NewPhase(PlayPhase, prepare, execute)
	return p
}
