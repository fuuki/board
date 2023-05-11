package burst

import (
	"github.com/fuuki/board/logic"
	"github.com/fuuki/board/resource"
)

// burstBoardProfile is a profile of the board.
type burstBoardProfile struct {
	PlayerHands   []*resource.CardLine[*Card] `json:"player_hands"`
	Deck          *resource.CardLine[*Card]   `json:"deck"`
	PlayedHistory [][]*PlayedCard             `json:"played_history"`
	Count         int                         `json:"count"`
	Burster       logic.Player                `json:"burster"`
}

// burstBoardProfileDefinition is a definition of the board profile.
type burstBoardProfileDefinition struct {
	totalPlayer uint
}

var _ logic.BoardProfileDefinition[*burstBoardProfile] = (*burstBoardProfileDefinition)(nil)

// New returns a new board profile.
func (d *burstBoardProfileDefinition) New() *burstBoardProfile {
	p := &burstBoardProfile{
		PlayerHands: make([]*resource.CardLine[*Card], d.totalPlayer),
	}
	for i := uint(d.totalPlayer); i < d.totalPlayer; i++ {
		p.PlayerHands[logic.Player(i)] = resource.NewCardLine[*Card](nil)
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

func (jp *burstBoardProfile) Player(p logic.Player) *resource.CardLine[*Card] {
	return jp.PlayerHands[p]
}

type burstPlayerAction struct {
	// Action is an action of the player.
	Select resource.CardID
}

type burstConfig struct {
	totalPlayer uint
}

var _ logic.Config = (*burstConfig)(nil)

// NewConfig returns a new config.
func NewConfig(totalPlayer uint) *burstConfig {
	return &burstConfig{
		totalPlayer: totalPlayer,
	}
}

// TotalPlayer returns the total number of players.
func (c *burstConfig) TotalPlayer() uint {
	return c.totalPlayer
}
