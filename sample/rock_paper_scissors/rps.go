package rock_paper_scissors

import (
	"github.com/fuuki/board/board"
)

const (
	PLAY_PHASE board.PhaseName = "play"
)

type jGame = board.Game[*JankenBoardProfile, *JankenActionProfile]
type jPhase = board.Phase[*JankenBoardProfile, *JankenActionProfile]
type jAction = board.ActionProfile[*JankenActionProfile]
type jActionReq = board.ActionRequest[*JankenActionProfile]

func Play() {
	g := rockPaperScissorsGame()
	inputer := &board.InteractiveActionInputer[*JankenActionProfile]{}
	g.Play(inputer)
}

// rockPaperScissorsGame returns a game of rock-paper-scissors.
func rockPaperScissorsGame() *jGame {
	rp := resourceProfile()

	p1 := playPhase()
	g := board.NewGame(2, PLAY_PHASE, []*jPhase{p1}, rp, resultFn)
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

	execute := func(g *jGame, bp *JankenBoardProfile, ap *jAction) (board.PhaseName, *JankenBoardProfile) {
		getReward(ap, g.BoardProfile())
		if isFinished(g.BoardProfile()) {
			return "", bp
		}
		return PLAY_PHASE, bp
	}

	p := board.NewPhase(PLAY_PHASE, prepare, execute)
	return p
}

func profileDef() *jActionReq {
	// TODO: implement
	r := &jActionReq{}
	return r
}

// getReward returns the result of the board.
func getReward(ap *jAction, rp *JankenBoardProfile) {
	a0 := (*ap.Player(board.Player(0))).Hand
	a1 := (*ap.Player(board.Player(1))).Hand

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
		if jp.Player(board.Player(i)).Point() >= 3 {
			return true
		}
	}
	return false
}

func resultFn(g *jGame) *board.Result {
	r := board.NewResult()
	rank := func(point int) uint {
		if point == 3 {
			return 1
		}
		return 2
	}
	for i := 0; i < g.BoardProfile().PlayerNum(); i++ {
		score := g.BoardProfile().Player(board.Player(i)).Point()
		r.AddPlayerResult(
			board.PlayerResult{
				Player: board.Player(i),
				Score:  score,
				Rank:   rank(score),
			})
	}
	return r
}
