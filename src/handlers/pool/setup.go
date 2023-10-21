package pool

import (
	"log"
	"net/http"
	"strings"

	"github.com/UpRightSofia/lottolodge/src/models/pool_store"
	"github.com/UpRightSofia/lottolodge/src/models/ticket_store"
	"github.com/UpRightSofia/lottolodge/src/models/user_store"
)

func (h *PoolServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/hud/")

	appError := handleLogic(h, id, w, r)
	if appError != nil {
		log.Printf("%d - %s - %s", appError.Code, r.URL.Path, appError.Message)
		http.Error(w, appError.Message, appError.Code)
	}
}

func handleLogic(p *PoolServer, id string, w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return p.GetPool(id, w)
	case http.MethodPost:
		return h.CreateHud(id, w, r)
	}

	return &core.AppErrorResponse{Message: "Method not allowed", Code: http.StatusMethodNotAllowed}
}

type PoolServer struct {
	poolStore   pool_store.PoolStore
	userStore   user_store.UserStore
	ticketStore ticket_store.TicketStore
}

func NewPoolServer(store pool_store.PoolStore) *PoolServer {
	return &PoolServer{poolStore: store}
}
