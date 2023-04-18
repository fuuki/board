package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/fuuki/board/sample/daifugo"
	"github.com/fuuki/board/sample/rock_paper_scissors"
)

func main() {
	var (
		g = flag.String("g", "rock_paper_scissors", "select game")
		s = flag.Bool("s", false, "start server")
	)
	flag.Parse()

	switch *g {
	case "rock_paper_scissors":
		log.Default().Println("start rock_paper_scissors terminal game")
		rock_paper_scissors.Play()
	case "daifugo":
		if *s {
			log.Default().Println("start daifugo server")
			daifugoServer()
			return
		}
		log.Default().Println("start daifugo terminal game")
		daifugo.Play()
	default:
		panic("unknown game")
	}
}

func daifugoServer() {
	// サーバを起動
	log.Default().Println("start server")
	s := daifugo.NewServer()
	mux := s.NewMux()
	http.ListenAndServe(":8080", mux)
}
