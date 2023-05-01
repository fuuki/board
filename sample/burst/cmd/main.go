package main

import (
	"log"
	"net/http"

	"github.com/fuuki/board/sample/burst"
)

func main() {
	log.Default().Println("start server")
	s := burst.NewServer()
	mux := s.NewMux()
	http.ListenAndServe(":8080", mux)
}
