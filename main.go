package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	httpMux := http.NewServeMux()

	httpMux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	log.Fatal(http.ListenAndServe(":80", httpMux))
}
