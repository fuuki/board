package result

import "github.com/fuuki/board/player"

type Result struct {
	Players []PlayerResult
}

func NewResult() *Result {
	return &Result{}
}

type PlayerResult struct {
	Player player.Player
	Score  int
	Rank   uint
}

func (r *Result) AddPlayerResult(pr PlayerResult) {
	r.Players = append(r.Players, pr)
}
