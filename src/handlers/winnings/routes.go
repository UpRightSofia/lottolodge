package winnings

import (
	"github.com/UpRightSofia/lottolodge/src/models"
	"github.com/gorilla/mux"
)

func NewServer(db models.PostgresStore, router *mux.Router) *server {
	s := &server{
		db:     db,
		router: router,
	}

	s.routes()

	return s
}

type server struct {
	db     models.PostgresStore
	router *mux.Router
}

func (s *server) routes() {
	s.router.HandleFunc("/winnings/distribute", s.distribute())
}