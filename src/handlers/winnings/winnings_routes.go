package winnings

import (
	"github.com/UpRightSofia/lottolodge/src/models"
	"github.com/UpRightSofia/lottolodge/src/models/pool_store"
	"github.com/UpRightSofia/lottolodge/src/models/ticket_store"
	"github.com/UpRightSofia/lottolodge/src/models/user_store"
	"github.com/UpRightSofia/lottolodge/src/models/winning_store"
	"github.com/gorilla/mux"
)

func NewServer(db models.PostgresStore, router *mux.Router) *server {
	s := &server{
		service: NewWinningsService(db),
		router:  router,
	}

	s.routes()

	return s
}

type WinningService struct {
	winningStore winning_store.WinningStore
	userStore    user_store.UserStore
	poolStore    pool_store.PoolStore
	ticketStore  ticket_store.TicketStore
}

func NewWinningsService(db models.PostgresStore) *WinningService {
	return &WinningService{
		winningStore: db.WinningStore,
		userStore:    db.UserStore,
		poolStore:    db.PoolStore,
		ticketStore:  db.TicketStore,
	}
}

type server struct {
	service *WinningService
	router  *mux.Router
}

func (s *server) routes() {
	s.router.HandleFunc("/user/{user_id}/pool/{pool_id}/winnings", s.service.GetWinningsForUserAndPool())
	s.router.HandleFunc("/user/{user_id}/winnings", s.service.GetWinningsForUser())
	s.router.HandleFunc("/winnings/distribute", s.distribute())
}
