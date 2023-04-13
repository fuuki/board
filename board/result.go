package board

type Result struct {
	Players []PlayerResult
}

func NewResult() *Result {
	return &Result{}
}

type PlayerResult struct {
	Player Player
	Score  int
	Rank   uint
}

func (r *Result) AddPlayerResult(pr PlayerResult) {
	r.Players = append(r.Players, pr)
}
