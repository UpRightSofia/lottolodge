package ticket_store

import (
	"database/sql"
	"sync"
	"testing"

	"github.com/UpRightSofia/lottolodge/src/models/user_store"
	"github.com/UpRightSofia/lottolodge/src/models/utils"
	"github.com/google/uuid"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestTicketPsqlStore(t *testing.T) {

	utils.WithPostgres(t, func(db *sql.DB, wg *sync.WaitGroup) {
		store := TicketPostgresStore{db}
		userStore := user_store.NewUserPostgresStore(db)

		utils.WithParallel(wg, func() {
			t.Run("GetTicket returns Ticket", func(t *testing.T) {
				uuId := uuid.New().String()

				// Sample data for insertion
				userID := createUser(*userStore, t)                                // You'll need to set this to an appropriate value, perhaps another UUID or an existing user ID.
				details := sql.NullString{String: `{"key": "value"}`, Valid: true} // Sample JSON for demonstration
				isHandPicked := true

				_, err := db.Exec(`insert into tickets (id, user_id, details, is_hand_picked) values ($1, $2, $3, $4);`, uuId, userID, details, isHandPicked)
				if err != nil {
					t.Error(err)
				}

				ticket, getErr := store.GetTicket(uuId)
				if getErr != nil {
					t.Errorf("GetTicket failed: %s\n", getErr)
				}

				// The UpdatedAt field will be set to the current time when the record is inserted, so we don't set it here.
				expectedTicket := Ticket{
					ID:           uuId,
					UserID:       userID,
					Details:      details,
					IsHandPicked: isHandPicked,
					// UpdatedAt: This field will be set by the database, so we can skip it in this comparison or handle it separately.
				}

				compareTickets(t, expectedTicket, ticket)
			})
		})

		utils.WithParallel(wg, func() {
			t.Run("GetTicketsInRange returns tickets", func(t *testing.T) {
				uuId1 := uuid.New().String()
				uuId2 := uuid.New().String()
				uuId3 := uuid.New().String()

				// Sample data for insertion
				userID := createUser(*userStore, t)                                // You'll need to set this to an appropriate value, perhaps another UUID or an existing user ID.
				details := sql.NullString{String: `{"key": "value"}`, Valid: true} // Sample JSON for demonstration
				isHandPicked := true

				// Insert one fresh ticket
				_, err := db.Exec(`insert into tickets (id, user_id, details, is_hand_picked) values ($1, $2, $3, $4);`, uuId1, userID, details, isHandPicked)
				if err != nil {
					t.Error(err)
				}

				// Insert one used
				_, err = db.Exec(`insert into tickets (id, user_id, details, is_hand_picked, is_used) values ($1, $2, $3, $4, $5);`, uuId2, userID, details, isHandPicked, true)
				if err != nil {
					t.Error(err)
				}

				// Insert second unused
				_, err = db.Exec(`insert into tickets (id, user_id, details, is_hand_picked) values ($1, $2, $3, $4);`, uuId3, userID, details, isHandPicked)
				if err != nil {
					t.Error(err)
				}

				tickets, getErr := store.GetUnusedTickets()
				if getErr != nil {
					t.Errorf("GetTicket failed: %s\n", getErr)
				}

				if len(tickets) != 2 {
					t.Errorf("Expected 2 tickets, but got %d", len(tickets))
				}
			})
		})

		utils.WithParallel(wg, func() {
			t.Run("Create ticket creates Ticket", func(t *testing.T) {
				userID := createUser(*userStore, t)
				request := CreateTicketRequest{
					UserID:       userID,
					Details:      `{"key": "value"}`,
					IsHandPicked: true,
				}

				createdTicket, err := store.CreateTicket(request)
				if err != nil {
					t.Errorf("CreateTicket failed: %s\n", err)
				}

				ticket, getErr := store.GetTicket(createdTicket.ID)
				if getErr != nil {
					t.Errorf("GetTicket failed: %s\n", err)
				}

				compareTickets(t, createdTicket, ticket)
			})
		})
	})
}

func createUser(store user_store.UserPostgresStore, t *testing.T) string {
	user, err := store.CreateUser(user_store.CreateUserRequest{ID: uuid.New().String(), BalanceE5: 100})
	if err != nil {
		t.Errorf("CreateUser failed: %s\n", err)
	}

	return user.ID
}

func compareTickets(t *testing.T, expected, got Ticket) {
	if expected.ID != got.ID ||
		expected.UserID != got.UserID ||
		expected.Details.String != got.Details.String ||
		expected.IsHandPicked != got.IsHandPicked {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
