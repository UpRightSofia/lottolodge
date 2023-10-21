package winnings

import (
	"github.com/UpRightSofia/lottolodge/src/models"
	"github.com/UpRightSofia/lottolodge/src/models/winning_store"
	"github.com/gorilla/mux"
)

func NewServer(db models.PostgresStore, router *mux.Router) *server {
	s := &server{
		service: NewWinningsService(db.WinningStore),
		router:  router,
	}

	s.routes()

	return s
}

type WinningService struct {
	winningStore winning_store.WinningStore
}

func NewWinningsService(winningsStore winning_store.WinningStore) *WinningService {
	return &WinningService{winningStore: winningsStore}
}

type server struct {
	service *WinningService
	router  *mux.Router
}

func (s *server) routes() {
	s.router.HandleFunc("/user/{user_id}/pool/{pool_id}/winnings", s.service.GetWinningsForUser())
}
