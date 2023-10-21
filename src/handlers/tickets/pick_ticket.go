package tickets

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"sync"

	"github.com/UpRightSofia/lottolodge/src/handlers/pool"
	"github.com/UpRightSofia/lottolodge/src/models"
	"github.com/UpRightSofia/lottolodge/src/models/pool_store"
	"github.com/UpRightSofia/lottolodge/src/models/setting_store"
	"github.com/UpRightSofia/lottolodge/src/models/ticket_store"
	"github.com/UpRightSofia/lottolodge/src/models/user_store"
	"github.com/gorilla/mux"
)

type ClientError struct {
	Message string `json:"message"`
}

func (e *ClientError) Error() string {
	return e.Message
}

func isClientError(err error) bool {
	_, ok := err.(*ClientError)
	return ok
}

type Ticket struct {
	UserUUID      string                     `json:"user_id"`
	TicketDetails ticket_store.TicketDetails `json:"ticket_details"`
}

func (s *server) ticketPick() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var ticket Ticket

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&ticket)
		if err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}

		if len(ticket.TicketDetails.DrawnNumbers) != 6 {
			http.Error(w, "Expecting 6 numbers", http.StatusBadRequest)
			return
		}

		pool_id, err := canUserPickTicket(s.db, ticket.UserUUID)
		if isClientError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		details, err := json.Marshal(ticket.TicketDetails)
		if err != nil {
			http.Error(w, "Failed to marshal ticket", http.StatusInternalServerError)
			return
		}

		dbTicket, err := s.db.TicketStore.CreateTicket(ticket_store.CreateTicketRequest{UserID: ticket.UserUUID, Details: string(details), IsHandPicked: true, PoolID: pool_id})
		if err != nil {
			http.Error(w, "Failed to create ticket", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dbTicket)
		w.WriteHeader(http.StatusOK)
	}
}

func RemainingTicketCount(db models.PostgresStore, user_id string) (int, string, error) {
	var wg sync.WaitGroup
	var once sync.Once
	var firstErr error
	var pool pool_store.Pool
	var settings setting_store.Setting
	var user user_store.User

	setError := func(err error) {
		once.Do(func() {
			firstErr = err
		})
	}

	wg.Add(3)

	go func() {
		defer wg.Done()
		res, err := db.PoolStore.GetTodayPool()
		if err != nil {
			setError(errors.New("failed to get today pool"))
			return
		}
		pool = res
	}()

	go func() {
		defer wg.Done()
		res, err := db.SettingsStore.GetLastSetting()
		if err != nil {
			setError(errors.New("failed to get ticket price"))
			return
		}
		settings = res
	}()

	go func() {
		defer wg.Done()
		res, err := db.UserStore.GetUser(user_id)
		if err == sql.ErrNoRows {
			setError(&ClientError{Message: "No such user"})
			return
		}
		if err != nil {
			setError(errors.New("failed to get user"))
			return
		}
		user = res
	}()

	wg.Wait()

	if firstErr != nil {
		return 0, "", firstErr
	}

	pickedTicketsAlready, err := db.TicketStore.GetUserTicketsCount(user.ID, pool.ID)
	if err != nil {
		return 0, "", errors.New("failed to get user tickets count")
	}

	maxTickets := user.BalanceE5 / settings.TicketPrizeE5

	return int(maxTickets) - pickedTicketsAlready, pool.ID, nil
}

func canUserPickTicket(db models.PostgresStore, user_id string) (string, error) {
	remaingTicket, poolId, err := RemainingTicketCount(db, user_id)
	if err != nil {
		return "", err
	}

	if remaingTicket <= 0 {
		return "", &ClientError{Message: "Already picked today tickets"}
	}

	return poolId, nil
}

func (s *server) pickBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var user_id struct {
			UserID string `json:"user_id"`
		}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user_id)
		if err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}

		request, err := s.pickTickets(user_id.UserID)
		if isClientError(err) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = s.db.TicketStore.BatchInsertTicket(request)
		if err != nil {
			http.Error(w, "Failed to create tickets", http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) pickTickets(user_id string) ([]ticket_store.CreateTicketRequest, error) {
	remaingTicket, poolId, err := RemainingTicketCount(s.db, user_id)
	if err != nil {
		return nil, err
	}

	batchTickets := make([]Ticket, 0, remaingTicket)
	for batchTicketsCount := 0; batchTicketsCount < remaingTicket; batchTicketsCount++ {
		ticketDetails, err := pool.DrawNumbers()
		if err != nil {
			return nil, err
		}

		batchTickets = append(batchTickets, Ticket{UserUUID: user_id, TicketDetails: ticketDetails})
	}

	request := make([]ticket_store.CreateTicketRequest, 0, remaingTicket)
	for _, ticket := range batchTickets {
		details, err := json.Marshal(ticket.TicketDetails)
		if err != nil {
			return nil, err
		}

		request = append(request, ticket_store.CreateTicketRequest{UserID: ticket.UserUUID, Details: string(details), IsHandPicked: false, PoolID: poolId})
	}

	return request, nil
}

func (s *server) returnPickedTickets() http.HandlerFunc {

	type PickedTicketsResponse struct {
		Tickets        []ticket_store.TicketDetails `json:"tickets"`
		MaximumTickets int64                        `json:"maximum_tickets"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only Get method is allowed", http.StatusMethodNotAllowed)
			return
		}

		vars := mux.Vars(r)
		userId, found := vars["user_id"]

		if !found {
			http.Error(w, "User id not supplied in request", http.StatusBadRequest)
			return
		}

		user, err := s.db.UserStore.GetUser(userId)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		activePool, err := s.db.PoolStore.GetTodayPool()
		if err != nil {
			http.Error(w, "No active pool", http.StatusBadRequest)
			return
		}

		settings, err := s.db.SettingsStore.GetLastSetting()
		if err != nil {
			http.Error(w, "Failed to get settings", http.StatusInternalServerError)
			return
		}

		maximumTickets := user.BalanceE5 / settings.TicketPrizeE5

		tickets, err := s.db.TicketStore.GetTicketsForUser(userId, activePool.ID)
		if err != nil {
			http.Error(w, "Failed to return ticket", http.StatusInternalServerError)
			return
		}

		var ticketDetails []ticket_store.TicketDetails = make([]ticket_store.TicketDetails, 0, len(tickets))
		for _, ticket := range tickets {
			if ticket.Details.Valid {
				detailsString := ticket.Details.String

				var ticketDetail ticket_store.TicketDetails
				err := json.Unmarshal([]byte(detailsString), &ticketDetail)
				if err != nil {
					http.Error(w, "Failed to unmarshal ticket", http.StatusInternalServerError)
					return
				}

				ticketDetails = append(ticketDetails, ticketDetail)
			}
		}

		response := PickedTicketsResponse{
			Tickets:        ticketDetails,
			MaximumTickets: maximumTickets,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusOK)
	}
}
