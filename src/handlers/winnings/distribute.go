package winnings

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/UpRightSofia/lottolodge/src/models/ticket_store"
	"github.com/UpRightSofia/lottolodge/src/models/winning_store"
)

const (
	smallMultiplier = 2
	bigMultiplier   = 5
)

var winningMappingE5 = map[int]int{
	3: 40000,
	4: 1000000,
	5: 50000000,
	6: 10000000000,
}

func (s *server) distribute() http.HandlerFunc {
	type DistrtibuteRequest struct {
		PoolUUID string `json:"pool_uuid"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var request DistrtibuteRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}

		if request.PoolUUID == "" {
			http.Error(w, "Expecting pool UUID", http.StatusBadRequest)
			return
		}

		err = s.distributePoolWinnings(request.PoolUUID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to distribute pool winnings: %s", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Winnings distributet successfully!")
		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) distributePoolWinnings(poolID string) error {
	pool, err := s.service.poolStore.GetPool(poolID)
	if err != nil {
		return fmt.Errorf("could not obtain pool: %s", err)
	}

	tickets, err := s.service.ticketStore.GetUnusedTickets(pool.ID)
	if err != nil {
		return fmt.Errorf("could not obtain tickets for pool: %s", err)
	}

	if len(tickets) == 0 {
		return nil
	}

	if !pool.Details.Valid {
		return errors.New("pool details are empty")
	}

	var poolDetails ticket_store.TicketDetails
	err = json.Unmarshal([]byte(pool.Details.String), &poolDetails)
	if err != nil {
		return fmt.Errorf("failed to parse pool details: %s", err)
	}

	winningNumbers := make(map[int]struct{})
	for _, n := range poolDetails.DrawnNumbers {
		winningNumbers[n] = struct{}{}
	}
	for _, ticket := range tickets {
		// TODO: fan out
		_, err := s.distributeTicketWinnings(&ticket, winningNumbers, poolDetails.SmallMultiplier, poolDetails.BigMultiplier)
		if err != nil {
			return fmt.Errorf("failed to distribute winnings for ticket: %s", err)
		}
	}

	return nil
}

func (s *server) distributeTicketWinnings(
	ticket *ticket_store.Ticket,
	winningNumbers map[int]struct{},
	winningSmallMultiplier, winningBigMultiplier int) (*winning_store.Winning, error) {
	if !ticket.Details.Valid {
		return nil, errors.New("empty ticket details")
	}

	var ticketDetails ticket_store.TicketDetails
	err := json.Unmarshal([]byte(ticket.Details.String), &ticketDetails)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ticket details: %s", err)
	}

	prize := calculateTicketPrize(&ticketDetails, winningNumbers, winningSmallMultiplier, winningBigMultiplier)
	if prize > 0 {
		winningRequest := winning_store.CreateWinningRequest{
			UserID:   ticket.UserID,
			TicketID: ticket.ID,
			PoolID:   ticket.PoolID,
			PrizeE5:  int64(prize),
		}
		winning, err := s.service.winningStore.CreateWinning(winningRequest)
		if err != nil {
			return nil, fmt.Errorf("failed to write a winning record: %s", err)
		}

		return &winning, nil
	}

	return nil, nil
}

func calculateTicketPrize(ticketDetails *ticket_store.TicketDetails, winningNumbers map[int]struct{}, winningSmallMultiplier, winningBigMultiplier int) int {
	matchCount := 0
	multiplier := 1
	for _, n := range ticketDetails.DrawnNumbers {
		if _, ok := winningNumbers[n]; ok {
			matchCount += 1
		}
	}
	if ticketDetails.SmallMultiplier == winningSmallMultiplier {
		multiplier *= smallMultiplier
	}
	if ticketDetails.BigMultiplier == winningBigMultiplier {
		multiplier *= bigMultiplier
	}

	prize, _ := winningMappingE5[matchCount]
	return prize * multiplier
}
