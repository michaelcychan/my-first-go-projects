package server

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		// select will response to the first in channel
		select {

		// if data first
		case d := <-data:
			fmt.Fprint(w, d)

		// if cancelled first
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
