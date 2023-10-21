package winnings

import (
	"net/http"
)

type TicketDetailsWithPrize struct {
	DrawnNumbers    []int `json:"drawn_numbers"`
	SmallMultiplier int   `json:"small_multiplier"`
	BigMultiplier   int   `json:"big_multiplier"`
	WinnigsE5       int64 `json:"winnigsE5"`
}

type WinningsForUserResponse struct {
}

func (pool *WinningService) GetWinningsForUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only Get method is allowed", http.StatusMethodNotAllowed)
			return
		}

	}
}
