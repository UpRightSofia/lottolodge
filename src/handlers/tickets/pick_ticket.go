package tickets

import (
	"encoding/json"
	"net/http"
)

func (s *server) ticketPick() http.HandlerFunc {
	type Ticket struct {
		UserUUID  string `json:"user_uuid"`
		Numbers   []int  `json:"numbers"`
		RedNumber int    `json:"red_number"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var ticket Ticket

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&ticket)
		if err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}

		if len(ticket.Numbers) != 6 {
			http.Error(w, "Expecting 6 numbers", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Ticket received successfully!")
		w.WriteHeader(http.StatusOK)
	}
}
