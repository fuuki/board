package main

import (
	"flag"

	"github.com/fuuki/board/sample/daifugo"
	"github.com/fuuki/board/sample/rock_paper_scissors"
)

func main() {
	var (
		g = flag.String("g", "rock_paper_scissors", "select game")
	)
	//ここで解析されます
	flag.Parse()
	switch *g {
	case "rock_paper_scissors":
		rock_paper_scissors.Play()
	case "daifugo":
		daifugo.Play()
	default:
		panic("unknown game")
	}
}
