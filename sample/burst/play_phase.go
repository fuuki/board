package burst

import (
	"errors"
	"sort"

	"github.com/fuuki/board"
)

var ErrNoCardSelected = errors.New("no card selected")
var ErrInvalidCard = errors.New("invalid card")

func playPhase() *jPhase {
	p := board.NewPhase(PlayPhase, playPhasePrepare, playPhaseExecute)
	return p
}

func playPhasePrepare(st *board.Status, bp *burstBoardProfile) (*jActionReq, *burstBoardProfile) {
	// Define action profile
	apr := board.NewActionRequest[*burstPlayerAction](st.TotalPlayer())
	for p := 0; p < int(st.TotalPlayer()); p++ {
		p := p
		apr.RegisterValidator(board.Player(p), func(a *burstPlayerAction) error {
			if a.Select == "" {
				return ErrNoCardSelected
			}
			if !bp.PlayerHands[board.Player(p)].Has(a.Select) {
				return ErrInvalidCard
			}
			return nil
		})
	}
	return apr, bp
}

func playPhaseExecute(st *board.Status, bp *burstBoardProfile, ap *jAction) (board.PhaseName, *burstBoardProfile) {
	// 出したカードをプレイエリアに移動
	played := make([]*PlayedCard, st.TotalPlayer())
	for p := 0; p < int(st.TotalPlayer()); p++ {
		id := ap.Player(board.Player(p)).Select
		c := bp.PlayerHands[board.Player(p)].Pick(id)
		played[p] = &PlayedCard{
			Card:   c,
			Player: board.Player(p),
		}
	}
	// 数字が小さい順に並べる
	sortPlayedCards(played)
	bp.PlayedHistory = append(bp.PlayedHistory, played)

	// カウントを増やす
	for _, p := range played {
		bp.Count += p.Card.Number
		if bp.Count > 30 {
			bp.Burster = p.Player
			return "", bp
		}
	}

	// カードを補充
	for p := 0; p < int(st.TotalPlayer()); p++ {
		bp.PlayerHands[board.Player(p)].Add(bp.Deck.Draw())
	}
	return PlayPhase, bp
}

func sortPlayedCards(played []*PlayedCard) {
	// 数字が小さい順に並べる
	sort.Slice(played, func(i, j int) bool {
		return played[i].Card.Number < played[j].Card.Number
	})
}
