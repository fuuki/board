package daifugo

import (
	"fmt"

	"github.com/fuuki/board/board"
	"github.com/fuuki/board/resource"
)

type daifugoBoardProfile struct {
	turn        *resource.Turn
	playerHands map[board.Player]*resource.CardLine[*Card]
	PlayArea    *resource.CardLine[*Card]
}

func NewDaifugoBoardProfile(playerNum uint) *daifugoBoardProfile {
	p := &daifugoBoardProfile{
		turn:        resource.NewSimpleTurn(playerNum),
		playerHands: make(map[board.Player]*resource.CardLine[*Card]),
	}
	for i := uint(playerNum); i < playerNum; i++ {
		p.playerHands[board.Player(i)] = resource.NewCardLine[*Card](nil)
	}
	return p
}

func (jp *daifugoBoardProfile) Player(p board.Player) *resource.CardLine[*Card] {
	return jp.playerHands[p]
}

// Show print all resources
func (jp *daifugoBoardProfile) Show() string {
	s := ""
	for player, hand := range jp.playerHands {
		s += fmt.Sprintf("Player %d: %v", player, hand)
	}
	return s
}

type daifugoActionProfile struct {
	// Action is an action of the player.
	Select []resource.CardID
}
