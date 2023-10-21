package tickets

import (
	"errors"
	"testing"

	"github.com/UpRightSofia/lottolodge/src/models"
	"github.com/UpRightSofia/lottolodge/src/models/pool_store"
	"github.com/UpRightSofia/lottolodge/src/models/setting_store"
	"github.com/UpRightSofia/lottolodge/src/models/ticket_store"
	"github.com/UpRightSofia/lottolodge/src/models/user_store"
)

type UserStoreMock struct{}

func (s *UserStoreMock) GetUser(user_id string) (user_store.User, error) {
	if user_id == "61cee350-89e8-4990-911f-138d944d119b" {
		return user_store.User{
			ID:        "61cee350-89e8-4990-911f-138d944d119b",
			BalanceE5: 750000000,
		}, nil
	}
	return user_store.User{}, errors.New("text")
}

func (s *UserStoreMock) CreateUser(request user_store.CreateUserRequest) (user_store.User, error) {
	return user_store.User{}, nil
}

type SettingStoreMock struct{}

func (s *SettingStoreMock) GetLastSetting() (setting_store.Setting, error) {
	return setting_store.Setting{
		TicketPrizeE5: 2500000,
	}, nil

}
func (s *SettingStoreMock) CreateSetting(request setting_store.CreateSettingRequest) (setting_store.Setting, error) {
	return setting_store.Setting{}, nil
}

type PoolStoreMock struct{}

func (s *PoolStoreMock) GetPool(id string) (pool_store.Pool, error) {
	return pool_store.Pool{
		ID: "1",
	}, nil
}
func (s *PoolStoreMock) CreatePool(request pool_store.CreatePoolRequest) (pool_store.Pool, error) {
	return pool_store.Pool{}, nil
}
func (s *PoolStoreMock) GetTodayPool() (pool_store.Pool, error) {
	return pool_store.Pool{
		ID: "1",
	}, nil
}

func (s *PoolStoreMock) MarkPoolAsDone(id string, details string) (pool_store.Pool, error) {
	return pool_store.Pool{}, nil
}

type TicketStoreMock struct{}

func (s *TicketStoreMock) GetTicket(id string) (ticket_store.Ticket, error) {
	return ticket_store.Ticket{}, nil
}

func (s *TicketStoreMock) GetUnusedTickets(poolID string) ([]ticket_store.Ticket, error) {
	return []ticket_store.Ticket{}, nil
}

func (s *TicketStoreMock) CreateTicket(request ticket_store.CreateTicketRequest) (ticket_store.Ticket, error) {
	return ticket_store.Ticket{
		ID: "345",
	}, nil
}
func (s *TicketStoreMock) GetUserTicketsCount(user_id string, pool_id string) (int, error) {
	if user_id == "61cee350-89e8-4990-911f-138d944d119b" && pool_id == "1" {
		return 42, nil
	}
	return 0, errors.New("text")
}
func (s *TicketStoreMock) BatchInsertTicket(requests []ticket_store.CreateTicketRequest) error {
	return nil
}

type TicketStoreMock300 struct{}

func (s *TicketStoreMock300) GetTicket(id string) (ticket_store.Ticket, error) {
	return ticket_store.Ticket{}, nil
}

func (s *TicketStoreMock300) GetUnusedTickets(poolID string) ([]ticket_store.Ticket, error) {
	return []ticket_store.Ticket{}, nil
}

func (s *TicketStoreMock300) CreateTicket(request ticket_store.CreateTicketRequest) (ticket_store.Ticket, error) {
	return ticket_store.Ticket{
		ID: "345",
	}, nil
}

func (s *TicketStoreMock300) GetUserTicketsCount(user_id string, pool_id string) (int, error) {
	if user_id == "61cee350-89e8-4990-911f-138d944d119b" && pool_id == "1" {
		return 300, nil
	}
	return 0, errors.New("text")
}

func (s *TicketStoreMock300) BatchInsertTicket(requests []ticket_store.CreateTicketRequest) error {
	return nil
}

func Test_canUserPickTicket(t *testing.T) {
	tests := []struct {
		name    string
		db      models.PostgresStore
		user_id string
		want    string
		wantErr string
	}{
		{
			name: "can pick ticket",
			db: models.PostgresStore{
				UserStore:     &UserStoreMock{},
				SettingsStore: &SettingStoreMock{},
				PoolStore:     &PoolStoreMock{},
				TicketStore:   &TicketStoreMock{},
			},
			user_id: "61cee350-89e8-4990-911f-138d944d119b",
			want:    "1",
		},
		{
			name: "error",
			db: models.PostgresStore{
				UserStore:     &UserStoreMock{},
				SettingsStore: &SettingStoreMock{},
				PoolStore:     &PoolStoreMock{},
				TicketStore:   &TicketStoreMock{},
			},
			user_id: "61cee350-89e8-4990-911f-138d944d119a",
			wantErr: "failed to get user",
		},
		{
			name: "error_no_more_tickets",
			db: models.PostgresStore{
				UserStore:     &UserStoreMock{},
				SettingsStore: &SettingStoreMock{},
				PoolStore:     &PoolStoreMock{},
				TicketStore:   &TicketStoreMock300{},
			},
			user_id: "61cee350-89e8-4990-911f-138d944d119b",
			wantErr: "Already picked today tickets",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := canUserPickTicket(tt.db, tt.user_id)
			if tt.wantErr != "" {
				if err == nil {
					t.Errorf("canUserPickTicket() error = %v, wantErr %v", err, tt.wantErr)
				}
				if err.Error() != tt.wantErr {
					t.Errorf("canUserPickTicket() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("canUserPickTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}
