package setting_store

import (
	"fmt"
	"database/sql"
)

type SettingPostgresStore struct{ db *sql.DB }

func NewSettingPostgresStore(db *sql.DB) *SettingPostgresStore {
	return &SettingPostgresStore{db: db}
}

func (s *SettingPostgresStore) GetLastSetting() (Setting, error) {
	var setting Setting

	err := s.db.QueryRow(`select * from settings order by updated_at desc limit 1`).
		Scan(&setting.ID, &setting.TicketPrizeE5, &setting.PayoutPercent, &setting.UpdatedAt)

	if err != nil {
		return setting, fmt.Errorf("unable to fetch setting: %w", err)
	}

	return setting, nil
}

func (s *SettingPostgresStore) CreateSetting(request CreateSettingRequest) (Setting, error) {
	var setting Setting

	err := s.db.QueryRow(`insert into settings (ticket_prize_e5, payout_percent) values ($1, $2) returning *`, request.TicketPrizeE5, request.PayoutPercent). 
		Scan(&setting.ID, &setting.TicketPrizeE5, &setting.PayoutPercent, &setting.UpdatedAt)

	if err != nil {
		return setting, fmt.Errorf("unable to create setting: %w", err)
	}

	return setting, nil
}
