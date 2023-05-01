package burst

import (
	"github.com/fuuki/board"
	"github.com/fuuki/board/resource"
)

// burstBoardProfile is a profile of the board.
// ゲームはターン、シーケンス、ラウンド、ゲームの4つのレベルで構成される。
type burstBoardProfile struct {
	PlayerHands   []*resource.CardLine[*Card] `json:"player_hands"`
	Deck          *resource.CardLine[*Card]   `json:"deck"`
	PlayedHistory [][]*PlayedCard             `json:"played_history"`
	Count         int                         `json:"count"`
	Burster       board.Player                `json:"burster"`
}

// burstBoardProfileDefinition is a definition of the board profile.
type burstBoardProfileDefinition struct {
	totalPlayer uint
}

var _ board.BoardProfileDefinition[*burstBoardProfile] = (*burstBoardProfileDefinition)(nil)

// New returns a new board profile.
func (d *burstBoardProfileDefinition) New() *burstBoardProfile {
	p := &burstBoardProfile{
		PlayerHands: make([]*resource.CardLine[*Card], d.totalPlayer),
	}
	for i := uint(d.totalPlayer); i < d.totalPlayer; i++ {
		p.PlayerHands[board.Player(i)] = resource.NewCardLine[*Card](nil)
	}
	return p
}

// Clone returns a cloned board profile.
func (d *burstBoardProfileDefinition) Clone(bp *burstBoardProfile) *burstBoardProfile {
	result := &burstBoardProfile{
		PlayerHands:   make([]*resource.CardLine[*Card], len(bp.PlayerHands)),
		Deck:          bp.Deck.Clone(),
		PlayedHistory: make([][]*PlayedCard, len(bp.PlayedHistory)),
		Count:         bp.Count,
		Burster:       bp.Burster,
	}
	for i, h := range bp.PlayerHands {
		result.PlayerHands[i] = h.Clone()
	}
	for i, h := range bp.PlayedHistory {
		hc := make([]*PlayedCard, len(h))
		copy(hc, h)
		result.PlayedHistory[i] = hc
	}
	return result
}

func (jp *burstBoardProfile) Player(p board.Player) *resource.CardLine[*Card] {
	return jp.PlayerHands[p]
}

type burstPlayerAction struct {
	// Action is an action of the player.
	Select resource.CardID
}
