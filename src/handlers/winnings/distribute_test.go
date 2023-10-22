package winnings

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/UpRightSofia/lottolodge/src/models"
	"github.com/UpRightSofia/lottolodge/src/models/pool_store"
	"github.com/UpRightSofia/lottolodge/src/models/ticket_store"
	"github.com/UpRightSofia/lottolodge/src/models/winning_store"
)

func TestCalculateWinningsPerTicket(t *testing.T) {
	t.Run("Calculate 3/0/0", func(t *testing.T) {
		ticketDetails := ticket_store.TicketDetails{
			DrawnNumbers:    []int{1, 2, 3, 4, 5, 6},
			SmallMultiplier: 7,
			BigMultiplier:   8,
		}
		winningNumbers := map[int]struct{}{
			2: {},
			4: {},
			6: {},
		}
		expectedPrizeE5 := 150000
		actualPrizeE5 := calculateTicketPrize(&ticketDetails, winningNumbers, 13, 15)
		if expectedPrizeE5 != actualPrizeE5 {
			t.Errorf("Expected prize '%d', but got '%d", expectedPrizeE5, actualPrizeE5)
		}
	})
	t.Run("Calculate 3/1/0", func(t *testing.T) {
		ticketDetails := ticket_store.TicketDetails{
			DrawnNumbers:    []int{1, 2, 3, 4, 5, 6},
			SmallMultiplier: 7,
			BigMultiplier:   8,
		}
		winningNumbers := map[int]struct{}{
			2: {},
			4: {},
			6: {},
		}
		expectedPrizeE5 := 1500000
		actualPrizeE5 := calculateTicketPrize(&ticketDetails, winningNumbers, 7, 15)
		if expectedPrizeE5 != actualPrizeE5 {
			t.Errorf("Expected prize '%d', but got '%d", expectedPrizeE5, actualPrizeE5)
		}
	})
	t.Run("Calculate 3/0/1", func(t *testing.T) {
		ticketDetails := ticket_store.TicketDetails{
			DrawnNumbers:    []int{1, 2, 3, 4, 5, 6},
			SmallMultiplier: 7,
			BigMultiplier:   8,
		}
		winningNumbers := map[int]struct{}{
			2: {},
			4: {},
			6: {},
		}
		expectedPrizeE5 := 3000000
		actualPrizeE5 := calculateTicketPrize(&ticketDetails, winningNumbers, 13, 8)
		if expectedPrizeE5 != actualPrizeE5 {
			t.Errorf("Expected prize '%d', but got '%d", expectedPrizeE5, actualPrizeE5)
		}
	})
	t.Run("Calculate 3/1/1", func(t *testing.T) {
		ticketDetails := ticket_store.TicketDetails{
			DrawnNumbers:    []int{1, 2, 3, 4, 5, 6},
			SmallMultiplier: 7,
			BigMultiplier:   8,
		}
		winningNumbers := map[int]struct{}{
			2: {},
			4: {},
			6: {},
		}
		expectedPrizeE5 := 30000000
		actualPrizeE5 := calculateTicketPrize(&ticketDetails, winningNumbers, 7, 8)
		if expectedPrizeE5 != actualPrizeE5 {
			t.Errorf("Expected prize '%d', but got '%d", expectedPrizeE5, actualPrizeE5)
		}
	})
	t.Run("Calculate 4/0/0", func(t *testing.T) {
		ticketDetails := ticket_store.TicketDetails{
			DrawnNumbers:    []int{1, 2, 3, 4, 5, 6},
			SmallMultiplier: 7,
			BigMultiplier:   8,
		}
		winningNumbers := map[int]struct{}{
			1: {},
			2: {},
			3: {},
			4: {},
		}
		expectedPrizeE5 := 5000000
		actualPrizeE5 := calculateTicketPrize(&ticketDetails, winningNumbers, 13, 15)
		if expectedPrizeE5 != actualPrizeE5 {
			t.Errorf("Expected prize '%d', but got '%d", expectedPrizeE5, actualPrizeE5)
		}
	})
	t.Run("Calculate 5/0/0", func(t *testing.T) {
		ticketDetails := ticket_store.TicketDetails{
			DrawnNumbers:    []int{1, 2, 3, 4, 5, 6},
			SmallMultiplier: 7,
			BigMultiplier:   8,
		}
		winningNumbers := map[int]struct{}{
			1: {},
			2: {},
			3: {},
			4: {},
			5: {},
		}
		expectedPrizeE5 := 150000000
		actualPrizeE5 := calculateTicketPrize(&ticketDetails, winningNumbers, 13, 15)
		if expectedPrizeE5 != actualPrizeE5 {
			t.Errorf("Expected prize '%d', but got '%d", expectedPrizeE5, actualPrizeE5)
		}
	})
	t.Run("Calculate 6/0/0", func(t *testing.T) {
		ticketDetails := ticket_store.TicketDetails{
			DrawnNumbers:    []int{1, 2, 3, 4, 5, 6},
			SmallMultiplier: 7,
			BigMultiplier:   8,
		}
		winningNumbers := map[int]struct{}{
			1: {},
			2: {},
			3: {},
			4: {},
			5: {},
			6: {},
		}
		expectedPrizeE5 := 15000000000
		actualPrizeE5 := calculateTicketPrize(&ticketDetails, winningNumbers, 13, 15)
		if expectedPrizeE5 != actualPrizeE5 {
			t.Errorf("Expected prize '%d', but got '%d", expectedPrizeE5, actualPrizeE5)
		}
	})
}

type FakePoolStore struct{}

func (f *FakePoolStore) GetPool(id string) (pool_store.Pool, error) {
	return pool_store.Pool{
		ID: "pool_1",
		Details: sql.NullString{
			Valid:  true,
			String: `{"drawn_numbers":[1,2,3,4,5,6],"small_multiplier":7,"big_multiplier":8}`,
		},
	}, nil
}

func (f *FakePoolStore) CreatePool() (pool_store.Pool, error) {
	return pool_store.Pool{}, errors.New("not implemented")
}

func (f *FakePoolStore) GetActivePool() (pool_store.Pool, error) {
	return pool_store.Pool{}, errors.New("not implemented")
}

func (f *FakePoolStore) GetLastCompleted() (pool_store.Pool, error) {
	return pool_store.Pool{}, errors.New("not implemented")
}

func (f *FakePoolStore) MarkPoolAsDone(id string, details string) (pool_store.Pool, error) {
	return pool_store.Pool{}, errors.New("not implemented")
}

type FakeTicketStore struct{}

func (s *FakeTicketStore) GetTicketsForUser(user_id string, pool_id string) ([]ticket_store.Ticket, error) {
	return []ticket_store.Ticket{}, errors.New("not implemented")
}

func (s *FakeTicketStore) GetTicket(id string) (ticket_store.Ticket, error) {
	return ticket_store.Ticket{}, errors.New("not implemented")
}

func (s *FakeTicketStore) GetUnusedTickets(poolID string) ([]ticket_store.Ticket, error) {
	if poolID != "pool_1" {
		return []ticket_store.Ticket{}, errors.New("not supported for that pool")
	}

	return []ticket_store.Ticket{
		{
			ID:     "ticket_1",
			UserID: "user_1",
			PoolID: "pool_1",
			Details: sql.NullString{
				Valid:  true,
				String: `{"drawn_numbers":[1,2,3,4,5,6],"small_multiplier":7,"big_multiplier":8}`,
			},
		},
	}, nil
}

func (s *FakeTicketStore) CreateTicket(request ticket_store.CreateTicketRequest) (ticket_store.Ticket, error) {
	return ticket_store.Ticket{}, errors.New("not implemented")
}
func (s *FakeTicketStore) GetUserTicketsCount(user_id string, pool_id string) (int, error) {
	return 0, errors.New("not implemented")
}

func (s *FakeTicketStore) BatchInsertTicket(requests []ticket_store.CreateTicketRequest) error {
	return errors.New("not implemented")
}

type FakeWinningStore struct{}

func (f *FakeWinningStore) GetWinning(id string) (winning_store.Winning, error) {
	return winning_store.Winning{}, errors.New("not implemented")
}

func (f *FakeWinningStore) CreateWinning(request winning_store.CreateWinningRequest) (winning_store.Winning, error) {
	expectedRequest := winning_store.CreateWinningRequest{
		UserID:   "user_1",
		TicketID: "ticket_1",
		PoolID:   "pool_1",
		PrizeE5:  int64(10 * 20 * 15000000000),
	}
	if request == expectedRequest {
		return winning_store.Winning{}, nil
	}
	return winning_store.Winning{}, errors.New("not implemented")
}

func (f *FakeWinningStore) GetWinningsForUserAndPool(user_id string, pool_id string) ([]winning_store.Winning, error) {
	return []winning_store.Winning{}, errors.New("not implemented")
}

func (f *FakeWinningStore) GetUserWinnings(user_id string) ([]winning_store.UserWinning, error) {
	return []winning_store.UserWinning{}, errors.New("not implemented")
}

func newFakeServer(db models.PostgresStore) *server {
	s := &server{
		service: NewWinningsService(db),
	}

	return s
}

func Test_canUserPickTicket(t *testing.T) {
	db := models.PostgresStore{
		PoolStore:    &FakePoolStore{},
		TicketStore:  &FakeTicketStore{},
		WinningStore: &FakeWinningStore{},
	}
	server := newFakeServer(db)

	t.Run("Distribute write jackpot", func(t *testing.T) {
		err := server.distributePoolWinnings("pool_1")
		if err != nil {
			t.Errorf("Got an error: %s", err)
		}
	})
}
