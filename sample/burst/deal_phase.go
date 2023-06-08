package burst

import (
	"github.com/fuuki/board/logic"
	"github.com/fuuki/board/resource"
)

type dealPhase struct{}

var _ bPhase = (*dealPhase)(nil)

// Name implement Phase.Name.
func (d *dealPhase) Name() logic.PhaseName {
	return DealPhase
}

// Prepare implement Phase.Prepare.
func (d *dealPhase) Prepare(config *burstConfig, bp *burstBoardProfile) (*bActionReq, *burstBoardProfile, error) {
	apr := logic.NewActionRequest[*burstPlayerAction](config.TotalPlayer())
	bp.Deck = resource.NewCardLine(newDeck())
	apr.AddShuffle("deck", bp.Deck.Len())
	return apr, bp, nil
}

// Execute implement Phase.Execute.
func (d *dealPhase) Execute(config *burstConfig, bp *burstBoardProfile, ap *bAction) (logic.PhaseName, *burstBoardProfile, error) {
	indexes := ap.NatureActionResult("deck")
	err := bp.Deck.ApplyShuffle(indexes)
	if err != nil {
		return "", bp, err
	}
	bp.dealCards(config.TotalPlayer())
	return PlayPhase, bp, nil
}

// dealCards prepares a new round.
func (bp *burstBoardProfile) dealCards(
	totalPlayer uint,
) {
	cards := bp.Deck.Cards()

	// Deal cards
	for p := 0; p < int(totalPlayer); p++ {
		bp.PlayerHands[p] = resource.NewCardLine(cards[p*3 : p*3+3])
	}
	bp.Deck = resource.NewCardLine(cards[totalPlayer*3:])
}
