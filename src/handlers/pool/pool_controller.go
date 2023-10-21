package pool

import (
	"net/http"
)

func (pool *PoolServer) finishPool() http.HandlerFunc {

	type PoolDetails struct {
		//drawn number
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		pool, err := pool.poolStore.GetPool("active")
		if err != nil {
			http.Error(w, "No active pool", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
