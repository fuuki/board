package daifugo

import (
	"fmt"

	"github.com/fuuki/board/board"
	"github.com/fuuki/board/resource"
)

// daifugoBoardProfile is a profile of the board.
// ゲームはターン、シーケンス、ラウンド、ゲームの4つのレベルで構成される。
type daifugoBoardProfile struct {
	Turn        *resource.Turn              `json:"turn"`
	PlayerHands []*resource.CardLine[*Card] `json:"player_hands"`
	PlayArea    *resource.CardLine[*Card]   `json:"play_area"`
}

func NewDaifugoBoardProfile(totalPlayer uint) *daifugoBoardProfile {
	p := &daifugoBoardProfile{
		Turn:        resource.NewSimpleTurn(totalPlayer),
		PlayerHands: make([]*resource.CardLine[*Card], totalPlayer),
	}
	for i := uint(totalPlayer); i < totalPlayer; i++ {
		p.PlayerHands[board.Player(i)] = resource.NewCardLine[*Card](nil)
	}
	return p
}

// PrepareNewRound prepares a new round.
func (bp *daifugoBoardProfile) PrepareNewRound(
	players []board.Player,
	startPlayer board.Player,
) {
	// dealCards deals cards to players.
	// Shuffle cards
	cards := shortDeck

	// Deal cards
	c := len(cards) / len(players)
	for i, p := range players {
		bp.PlayerHands[p] = resource.NewCardLine(cards[i*c : (i+1)*c])
	}

	bp.PrepareNewSequence(players, startPlayer)
}

// PrepareNewSequence prepares a new sequence.
func (bp *daifugoBoardProfile) PrepareNewSequence(
	players []board.Player,
	startPlayer board.Player,
) {
	bp.PlayArea = resource.NewCardLine([]*Card{})
	bp.Turn = resource.NewTurn(players, startPlayer)
}

func (jp *daifugoBoardProfile) Player(p board.Player) *resource.CardLine[*Card] {
	return jp.PlayerHands[p]
}

// Show print all resources
func (jp *daifugoBoardProfile) Show() string {
	s := ""
	for player, hand := range jp.PlayerHands {
		s += fmt.Sprintf("Player %d: %v\n", player, hand)
	}
	s += fmt.Sprintf("PlayArea: %v\n", jp.PlayArea)
	s += fmt.Sprintf("Turn: %v\n", jp.Turn)
	return s
}

type daifugoPlayerAction struct {
	// Action is an action of the player.
	Select []resource.CardID
	// Pass is true if the player pass.
	Pass bool
}
