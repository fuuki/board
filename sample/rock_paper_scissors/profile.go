package rock_paper_scissors

import (
	"fmt"

	"github.com/fuuki/board/player"
	"github.com/fuuki/board/resource"
)

type JankenBoardProfile struct {
	points map[player.Player]*resource.Point
}

func NewJankenBoardProfile(playerNum int) *JankenBoardProfile {
	p := &JankenBoardProfile{
		points: make(map[player.Player]*resource.Point),
	}
	for i := 0; i < playerNum; i++ {
		p.points[player.Player(i)] = resource.NewPoint()
	}
	return p
}

func (jp *JankenBoardProfile) Player(p player.Player) *resource.Point {
	return jp.points[p]
}

func (jp *JankenBoardProfile) PlayerNum() int {
	return len(jp.points)
}

// Show print all resources
func (jp *JankenBoardProfile) Show() string {
	s := ""
	for player, point := range jp.points {
		s += fmt.Sprintf("Player %d: %d pt(s)\n", player, point.Point())
	}
	return s
}

func (jp *JankenBoardProfile) Equal(jp2 *JankenBoardProfile) bool {
	for p, r := range jp.points {
		if r.Point() != jp2.points[p].Point() {
			return false
		}
	}
	return true
}
