package setting_store

import "time"

type Setting struct {
	ID int `json:"id"`
	TicketPrizeE5 int64 `json:"ticket_prize_e5"`
	PayoutPercent int `json:"payout_percent"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSettingRequest struct {
	TicketPrizeE5 int64 `json:"ticket_prize_e5"`
	PayoutPercent int `json:"payout_percent"`
}
