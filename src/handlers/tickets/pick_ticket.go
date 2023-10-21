package tickets

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"sync"

	"github.com/UpRightSofia/lottolodge/src/models"
	"github.com/UpRightSofia/lottolodge/src/models/pool_store"
	"github.com/UpRightSofia/lottolodge/src/models/setting_store"
	"github.com/UpRightSofia/lottolodge/src/models/ticket_store"
	"github.com/UpRightSofia/lottolodge/src/models/user_store"
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

func (s *server) ticketPick() http.HandlerFunc {
	type MultiplierTickets struct {
		SmallMutiplier int `json:"small_multiplier"`
		BigMutiplier   int `json:"big_multiplier"`
	}
	type Ticket struct {
		UserUUID          string            `json:"user_uuid"`
		Numbers           []int             `json:"numbers"`
		MultiplierTickets MultiplierTickets `json:"multiplier_tickets"`
	}

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

		if len(ticket.Numbers) != 6 {
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

		details, err := json.Marshal(ticket)
		if err != nil {
			http.Error(w, "Failed to marshal ticket", http.StatusInternalServerError)
			return
		}

		dbTicket, err := s.db.Tickets.CreateTicket(ticket_store.CreateTicketRequest{UserID: ticket.UserUUID, Details: string(details), IsHandPicked: true, PoolID: pool_id})
		if err != nil {
			http.Error(w, "Failed to create ticket", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dbTicket)
		w.WriteHeader(http.StatusOK)
	}
}

func canUserPickTicket(db models.PostgresStore, user_id string) (string, error) {
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
		res, err := db.Pools.GetTodayPool()
		if err != nil {
			setError(errors.New("failed to get today pool"))
			return
		}
		pool = res
	}()

	go func() {
		defer wg.Done()
		res, err := db.Settings.GetLastSetting()
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
		return "", firstErr
	}

	pickedTicketsAlready, err := db.Tickets.GetUserTicketsCount(user.ID, pool.ID)
	if err != nil {
		return "", errors.New("failed to get user tickets count")
	}

	maxTickets := user.BalanceE5 / settings.TicketPrizeE5

	if pickedTicketsAlready >= int(maxTickets) {
		return "", &ClientError{Message: "Already picked today tickets"}
	}

	return pool.ID, nil
}
