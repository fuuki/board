package rock_paper_scissors

import (
	"github.com/fuuki/board/action"
	"github.com/fuuki/board/board"
	"github.com/fuuki/board/game"
	"github.com/fuuki/board/player"
	"github.com/fuuki/board/result"
)

const (
	PLAY_PHASE game.PhaseName = "play"
)

type jGame = game.Game[*JankenBoardProfile, *JankenActionProfile]
type jPhase = game.Phase[*JankenBoardProfile, *JankenActionProfile]
type jAction = board.ActionProfile[*JankenActionProfile]
type jActionReq = board.ActionRequest[*JankenActionProfile]

func Play() {
	g := rockPaperScissorsGame()
	inputer := &action.InteractiveActionInputer[*JankenActionProfile]{}
	g.Play(inputer)
}

// rockPaperScissorsGame returns a game of rock-paper-scissors.
func rockPaperScissorsGame() *jGame {
	rp := resourceProfile()

	p1 := playPhase()
	g := game.NewGame(PLAY_PHASE, []*jPhase{p1}, rp, resultFn)
	return g
}

// resourceProfile returns a resource profile of rock-paper-scissors.
func resourceProfile() *JankenBoardProfile {
	rp := NewJankenBoardProfile(2)
	return rp
}

// playPhase returns a phase of rock-paper-scissors.
func playPhase() *jPhase {
	prepare := func(_ *jGame) *jActionReq {
		// Define action profile
		apr := profileDef()
		return apr
	}

	execute := func(g *jGame, ap *jAction) game.PhaseName {
		getReward(ap, g.BoardProfile())
		if isFinished(g.BoardProfile()) {
			return ""
		}
		return PLAY_PHASE
	}

	p := game.NewPhase(PLAY_PHASE, prepare, execute)
	return p
}

func profileDef() *jActionReq {
	// TODO: implement
	r := &jActionReq{}
	return r
}

// getReward returns the result of the game.
func getReward(ap *jAction, rp *JankenBoardProfile) {
	a0 := (*ap.Player(player.Player(0))).Hand
	a1 := (*ap.Player(player.Player(1))).Hand

	switch (a0 - a1 + 3) % 3 {
	case 0:
		return
	case 1:
		// Player 0 win
		rp.Player(0).AddPoint(1)
		return
	case 2:
		// Player 1 win
		rp.Player(1).AddPoint(1)
		return
	}
}

func isFinished(jp *JankenBoardProfile) bool {
	for i := 0; i < jp.PlayerNum(); i++ {
		if jp.Player(player.Player(i)).Point() >= 3 {
			return true
		}
	}
	return false
}

func resultFn(g *jGame) *result.Result {
	r := result.NewResult()
	rank := func(point int) uint {
		if point == 3 {
			return 1
		}
		return 2
	}
	for i := 0; i < g.BoardProfile().PlayerNum(); i++ {
		score := g.BoardProfile().Player(player.Player(i)).Point()
		r.AddPlayerResult(
			result.PlayerResult{
				Player: player.Player(i),
				Score:  score,
				Rank:   rank(score),
			})
	}
	return r
}
