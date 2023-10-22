package pool

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/UpRightSofia/lottolodge/src/models/ticket_store"
)

func (pool *PoolService) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		_, err := pool.poolStore.GetActivePool()
		if err == nil {
			http.Error(w, "There's already an active pool", http.StatusBadRequest)
			return
		}

		pool, err := pool.poolStore.CreatePool()
		if err != nil {
			http.Error(w, "Could not create pool", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pool)
		w.WriteHeader(http.StatusAccepted)
	}
}

func (pool *PoolService) last() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
			return
		}

		lastCompleted, err := pool.poolStore.GetLastCompleted()
		if err != nil {
			http.Error(w, "Could not obtain completed pool", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lastCompleted)
		w.WriteHeader(http.StatusAccepted)
	}
}

func (pool *PoolService) finishPool() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		activePool, err := pool.poolStore.GetActivePool()
		if err != nil {
			http.Error(w, "No active pool", http.StatusBadRequest)
			return
		}

		poolDetails, err := DrawNumbers()
		if err != nil {
			http.Error(w, "Failed to draw numbers", http.StatusInternalServerError)
			return
		}

		jsonBytes, err := json.Marshal(poolDetails)
		if err != nil {
			fmt.Println("Error marshaling to JSON:", err)
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}

		details := string(jsonBytes)

		pool.poolStore.MarkPoolAsDone(activePool.ID, details)

		w.WriteHeader(http.StatusAccepted)
	}
}

func drawNumbers(n, total int) ([]int, error) {
	balls := make([]int, 0, n)

	for len(balls) < n {
		max := big.NewInt(int64(total))
		randNum, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}
		num := int(randNum.Int64()) + 1 // Convert to 1-based index

		if !contains(balls, num) {
			balls = append(balls, num)
		}
	}

	return balls, nil
}

func contains(slice []int, item int) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func DrawNumbers() (ticket_store.TicketDetails, error) {
	const regualarNumbers = 6
	const twoMultiplier = 1
	const fiveMultiplier = 1
	const totalNumbersToDraw = regualarNumbers + twoMultiplier + fiveMultiplier

	const twoMultiplierIndex = totalNumbersToDraw - fiveMultiplier - 1
	const fiveMultiplierIndex = totalNumbersToDraw - 1

	// Draw from 1 to 99
	const maxNumber = 99
	drawnNumbers, err := drawNumbers(totalNumbersToDraw, maxNumber)
	if err != nil {
		return ticket_store.TicketDetails{}, err
	}

	return ticket_store.TicketDetails{
		DrawnNumbers:    drawnNumbers[:regualarNumbers],
		SmallMultiplier: drawnNumbers[twoMultiplierIndex],
		BigMultiplier:   drawnNumbers[fiveMultiplierIndex],
	}, nil
}
