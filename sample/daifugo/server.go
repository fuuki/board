package daifugo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fuuki/board/board"
)

type Server struct {
	g     *jGame
	chans []chan board.PhaseNo
}

func NewServer() *Server {
	g, ch := daifugoGame(2)
	s := &Server{
		g:     g,
		chans: []chan board.PhaseNo{},
	}
	go func() {
		for {
			n, ok := <-ch
			if !ok {
				log.Default().Printf("[ch] channel closed\n")
				break
			}
			for _, c := range s.chans {
				c <- n
			}
			log.Default().Printf("[ch] phase changed: %v\n", n)
		}
	}()

	g.Start()
	return s
}

func (s *Server) NewMux() *http.ServeMux {
	// ハンドラを登録
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.daifugoService)
	mux.HandleFunc("/notification", s.notificationService)
	return mux
}

func (s *Server) daifugoService(w http.ResponseWriter, r *http.Request) {
	// POST のみ受け付ける
	if r.Method != http.MethodPost {
		log.Default().Printf("method not allowed: %s\n", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// リクエストをパース
	req, err := s.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Default().Println(err)
		return
	}

	// アクションを実行
	if err := s.g.RegisterAction(board.Player(req.Player), req.Action); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Default().Println(err)
		return
	}
}

// PostActionRequest is a request to post an action.
type PostActionRequest struct {
	// PhaseNo is a phase number.
	PhaseNo int `json:"phaseNo"`
	// Player is a player who takes an action.
	Player int `json:"player"`
	// Action is an action.
	Action *daifugoPlayerAction `json:"action"`
}

// parseRequest parses a request.
func (s *Server) parseRequest(r *http.Request) (*PostActionRequest, error) {
	var req *PostActionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// notificationService is a service to notify a player of a game phase progress.
func (s *Server) notificationService(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("notificationService")
	flusher, _ := w.(http.Flusher)
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ch := make(chan board.PhaseNo)
	s.chans = append(s.chans, ch)

	for {
		n, ok := <-ch
		if !ok {
			log.Default().Printf("[ch] channel closed\n")
			break
		}
		_, err := fmt.Fprintf(w, "data: %d\n\n", n)
		if err != nil {
			log.Default().Printf("[ch] failed to write: %v\n", err)
			break
		}
		log.Default().Printf("[ch] phase changed: %v\n", n)
		flusher.Flush()
	}
	// ch を削除
	for i, c := range s.chans {
		if c == ch {
			s.chans = append(s.chans[:i], s.chans[i+1:]...)
			break
		}
	}
	<-r.Context().Done()
	log.Default().Printf("[ch] closed connection\n")
}
