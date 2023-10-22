package winnings

import (
	"encoding/json"
	"net/http"

	"github.com/UpRightSofia/lottolodge/src/models/ticket_store"
	"github.com/gorilla/mux"
)

func (s *WinningService) GetWinningsForUserAndPool() http.HandlerFunc {

	type TicketDetailsWithPrize struct {
		DrawnNumbers    []int `json:"drawn_numbers"`
		SmallMultiplier int   `json:"small_multiplier"`
		BigMultiplier   int   `json:"big_multiplier"`
		WinnigsE5       int64 `json:"winnigsE5"`
	}

	type WinningsForUserResponse struct {
		Tickets []TicketDetailsWithPrize   `json:"tickets"`
		Pool    ticket_store.TicketDetails `json:"pool"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only Get method is allowed", http.StatusMethodNotAllowed)
			return
		}

		userId, found := mux.Vars(r)["user_id"]
		if !found {
			http.Error(w, "No user id supplied", http.StatusBadRequest)
			return
		}

		poolId, found := mux.Vars(r)["pool_id"]
		if !found {
			http.Error(w, "No pool id supplied", http.StatusBadRequest)
			return
		}

		_, err := s.userStore.GetUser(userId)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		requestedPool, err := s.poolStore.GetPool(poolId)
		if err != nil {
			http.Error(w, "Pool not found", http.StatusNotFound)
			return
		}

		var poolDetails ticket_store.TicketDetails
		if requestedPool.Details.Valid {
			detailsString := requestedPool.Details.String

			err := json.Unmarshal([]byte(detailsString), &poolDetails)
			if err != nil {
				http.Error(w, "Failed to unmarshal pool", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "This endpoint can only be used for finished pools", http.StatusBadRequest)
			return
		}

		winnings, err := s.winningStore.GetWinningsForUserAndPool(userId, poolId)
		if err != nil {
			http.Error(w, "Failed to get winnings", http.StatusInternalServerError)
			return
		}

		var ticketToPrize map[string]int64 = make(map[string]int64)
		for _, winning := range winnings {
			ticketToPrize[winning.TicketID] = winning.PrizeE5
		}

		tickets, err := s.ticketStore.GetTicketsForUser(userId, poolId)
		if err != nil {
			http.Error(w, "Failed to get tickets", http.StatusInternalServerError)
			return
		}

		ticketDetailsWithPrize := []TicketDetailsWithPrize{}
		for _, ticket := range tickets {

			if !ticket.Details.Valid {
				continue
			}

			detailsString := ticket.Details.String

			var ticketDetail ticket_store.TicketDetails
			err := json.Unmarshal([]byte(detailsString), &ticketDetail)
			if err != nil {
				http.Error(w, "Failed to unmarshal ticket", http.StatusInternalServerError)
				return
			}

			if prize, ok := ticketToPrize[ticket.ID]; ok {
				data := TicketDetailsWithPrize{
					DrawnNumbers:    ticketDetail.DrawnNumbers,
					SmallMultiplier: ticketDetail.SmallMultiplier,
					BigMultiplier:   ticketDetail.BigMultiplier,
					WinnigsE5:       prize,
				}
				ticketDetailsWithPrize = append(ticketDetailsWithPrize, data)
			} else {
				data := TicketDetailsWithPrize{
					DrawnNumbers:    ticketDetail.DrawnNumbers,
					SmallMultiplier: ticketDetail.SmallMultiplier,
					BigMultiplier:   ticketDetail.BigMultiplier,
					WinnigsE5:       0,
				}
				ticketDetailsWithPrize = append(ticketDetailsWithPrize, data)
			}
		}

		response := WinningsForUserResponse{
			Tickets: ticketDetailsWithPrize,
			Pool:    poolDetails,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusAccepted)
	}
}

func (s *WinningService) GetWinningsForUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only Get method is allowed", http.StatusMethodNotAllowed)
			return
		}

		userId, found := mux.Vars(r)["user_id"]
		if !found {
			http.Error(w, "No user id supplied", http.StatusBadRequest)
			return
		}

		winnings, err := s.winningStore.GetUserWinnings(userId)
		if err != nil {
			http.Error(w, "Failed to get winnings", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(winnings)
		w.WriteHeader(http.StatusAccepted)
	}
}
