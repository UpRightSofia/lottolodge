package pool

import (
	"github.com/UpRightSofia/lottolodge/src/models"
	"github.com/UpRightSofia/lottolodge/src/models/pool_store"
	"github.com/gorilla/mux"
)

func NewServer(db models.PostgresStore, router *mux.Router) *server {
	s := &server{
		service: NewPoolService(db.PoolStore),
		router:  router,
	}

	s.routes()

	return s
}

type PoolService struct {
	poolStore pool_store.PoolStore
}

func NewPoolService(store pool_store.PoolStore) *PoolService {
	return &PoolService{poolStore: store}
}

type server struct {
	service *PoolService
	router  *mux.Router
}

func (s *server) routes() {
	s.router.HandleFunc("/pool/finish", s.service.finishPool())
}
