package daifugo

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fuuki/board/board"
)

type Server struct {
	g *jGame
}

func NewServer() *Server {
	g := daifugoGame(2)
	g.Start()
	return &Server{
		g: g,
	}
}

func (s *Server) NewMux() *http.ServeMux {
	// ハンドラを登録
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.daifugoService)
	return mux
}

func (s *Server) daifugoService(w http.ResponseWriter, r *http.Request) {
	// POST のみ受け付ける
	if r.Method != http.MethodPost {
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
