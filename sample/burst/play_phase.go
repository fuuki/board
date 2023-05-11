package burst

import (
	"errors"
	"sort"

	"github.com/fuuki/board/logic"
)

var ErrNoCardSelected = errors.New("no card selected")
var ErrInvalidCard = errors.New("invalid card")

type playPhase struct{}

var _ jPhase = (*playPhase)(nil)

// Name returns a phase name.
func (p *playPhase) Name() logic.PhaseName {
	return PlayPhase
}

// Prepare implements Phase.Prepare.
func (p *playPhase) Prepare(config *burstConfig, bp *burstBoardProfile) (*logic.ActionRequest[*burstPlayerAction], *burstBoardProfile) {
	// Define action profile
	apr := logic.NewActionRequest[*burstPlayerAction](config.TotalPlayer())
	for p := 0; p < int(config.TotalPlayer()); p++ {
		p := p
		apr.RegisterValidator(logic.Player(p), func(a *burstPlayerAction) error {
			if a == nil {
				return ErrNoCardSelected
			}
			if a.Select == "" {
				return ErrNoCardSelected
			}
			if !bp.PlayerHands[logic.Player(p)].Has(a.Select) {
				return ErrInvalidCard
			}
			return nil
		})
	}
	return apr, bp
}

func (p *playPhase) Execute(config *burstConfig, bp *burstBoardProfile, ap *logic.ActionProfile[*burstPlayerAction]) (logic.PhaseName, *burstBoardProfile) {
	// 出したカードをプレイエリアに移動
	played := make([]*PlayedCard, config.TotalPlayer())
	for p := 0; p < int(config.TotalPlayer()); p++ {
		id := ap.Player(logic.Player(p)).Select
		c := bp.PlayerHands[logic.Player(p)].Pick(id)
		played[p] = &PlayedCard{
			Card:   c,
			Player: logic.Player(p),
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
	for p := 0; p < int(config.TotalPlayer()); p++ {
		bp.PlayerHands[logic.Player(p)].Add(bp.Deck.Draw())
	}
	return PlayPhase, bp
}

func sortPlayedCards(played []*PlayedCard) {
	// 数字が小さい順に並べる
	sort.Slice(played, func(i, j int) bool {
		return played[i].Card.Number < played[j].Card.Number
	})
}
