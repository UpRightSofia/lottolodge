package winning_store

import (
	"database/sql"
	"sync"
	"testing"

	"github.com/UpRightSofia/lottolodge/src/models/pool_store"
	"github.com/UpRightSofia/lottolodge/src/models/ticket_store"
	"github.com/UpRightSofia/lottolodge/src/models/user_store"
	"github.com/UpRightSofia/lottolodge/src/models/utils"
	"github.com/google/uuid"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestWinningPsqlStore(t *testing.T) {

	utils.WithPostgres(t, func(db *sql.DB, wg *sync.WaitGroup) {
		store := WinningPostgresStore{db}
		userStore := user_store.NewUserPostgresStore(db)
		ticketStore := ticket_store.NewTicketPostgresStore(db)
		poolStore := pool_store.NewPoolPostgresStore(db)

		utils.WithParallel(wg, func() {
			t.Run("GetWinning returns Winning", func(t *testing.T) {
				uuId := uuid.New().String()

				userID := createUser(*userStore, t)
				ticketID := createTicket(*ticketStore, userID, t)
				poolID := createPool(*poolStore, t)
				prizeE5 := int64(100000)

				_, err := db.Exec(`insert into winnings (id, user_id, ticket_id, pool_id, prize_e5) values ($1, $2, $3, $4, $5);`, uuId, userID, ticketID, poolID, prizeE5)

				if err != nil {
					t.Error(err)
				}

				winning, getErr := store.GetWinning(uuId)
				if getErr != nil {
					t.Errorf("GetWinning failed: %s\n", getErr)
				}

				// The UpdatedAt field will be set to the current time when the record is inserted, so we don't set it here.
				expectedWinning := Winning{
					ID:       uuId,
					UserID:   userID,
					TicketID: ticketID,
					PoolID:   poolID,
					PrizeE5:  prizeE5,
					// UpdatedAt: This field will be set by the database, so we can skip it in this comparison or handle it 	separately.
				}

				compareWinnings(t, expectedWinning, winning)
			})
		})

		utils.WithParallel(wg, func() {
			t.Run("Create winning creates Winning", func(t *testing.T) {
				userID := createUser(*userStore, t)
				ticketID := createTicket(*ticketStore, userID, t)
				poolID := createPool(*poolStore, t)
				prizeE5 := int64(100000)

				request := CreateWinningRequest{
					UserID:   userID,
					TicketID: ticketID,
					PoolID:   poolID,
					PrizeE5:  prizeE5,
				}

				createdWinning, err := store.CreateWinning(request)
				if err != nil {
					t.Errorf("CreateWinning failed: %s\n", err)
				}

				winning, getErr := store.GetWinning(createdWinning.ID)
				if getErr != nil {
					t.Errorf("GetWinning failed: %s\n", err)
				}

				compareWinnings(t, createdWinning, winning)
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

func createTicket(store ticket_store.TicketPostgresStore, userID string, t *testing.T) string {
	ticketRequest := ticket_store.CreateTicketRequest{
		UserID:       userID,
		Details:      `{"event": "sample"}`, // Modify this as needed
		IsHandPicked: true,                  // Sample value
	}

	ticket, err := store.CreateTicket(ticketRequest)
	if err != nil {
		t.Errorf("CreateTicket failed: %s\n", err)
	}

	return ticket.ID
}

func createPool(store pool_store.PoolPostgresStore, t *testing.T) string {
	poolRequest := pool_store.CreatePoolRequest{
		Details: `{"description": "sample pool"}`,
	}

	pool, err := store.CreatePool(poolRequest)
	if err != nil {
		t.Errorf("CreatePool failed: %s\n", err)
	}

	return pool.ID
}

func compareWinnings(t *testing.T, expected, got Winning) {
	if expected.ID != got.ID ||
		expected.UserID != got.UserID ||
		expected.PoolID != got.PoolID ||
		expected.TicketID != got.TicketID ||
		expected.PrizeE5 != got.PrizeE5 {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
