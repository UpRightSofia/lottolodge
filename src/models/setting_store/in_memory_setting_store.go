package setting_store

import (
	"errors"
	"time"
)

type InMemorySettingStore struct {
	lastSetting *Setting
}

func NewInMemorySettingStore(lastSetting *Setting) *InMemorySettingStore {
	return &InMemorySettingStore{lastSetting: lastSetting}
}

func (s *InMemorySettingStore) GetLastSetting() (Setting, error) {
	if s.lastSetting == nil {
		return Setting{}, errors.New("not found")
	}

	return *s.lastSetting, nil
}

func (s *InMemorySettingStore) CreateSetting(request CreateSettingRequest) (Setting, error) {
	id := 0
	if s.lastSetting != nil {
		id = s.lastSetting.ID + 1
	}

	s.lastSetting = &Setting{
		ID: id,
		TicketPrizeE5: request.TicketPrizeE5,
		PayoutPercent: request.PayoutPercent,
		UpdatedAt: time.Now(),
	}

	return *s.lastSetting, nil
}
