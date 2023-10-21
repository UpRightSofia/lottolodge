package pool

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
)

func (pool *PoolService) finishPool() http.HandlerFunc {

	type PoolDetails struct {
		DrawnNumbers   []int `json:"drawn_numbers"`
		TwoMultiplier  int   `json:"two_multiplier"`
		FiveMultiplier int   `json:"five_multiplier"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		activePool, err := pool.poolStore.GetTodayPool()
		if err != nil {
			http.Error(w, "No active pool", http.StatusBadRequest)
			return
		}

		drawnNumbers, err := drawNumbers(8, 99)
		if err != nil {
			http.Error(w, "Failed to draw numbers", http.StatusInternalServerError)
			return
		}

		poolDetails := PoolDetails{
			DrawnNumbers:   drawnNumbers[:6],
			TwoMultiplier:  drawnNumbers[6],
			FiveMultiplier: drawnNumbers[7],
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
