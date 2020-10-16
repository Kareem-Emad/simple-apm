package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Kareem-Emad/new-new-relic/auth"
	"github.com/Kareem-Emad/new-new-relic/producer"
	"github.com/gorilla/mux"
)

var tag = "[NEW_NEW_RELIC_Server]"

// Start starts an http server with the api routes loaded as specified
func Start(pm producer.ProductionManager) {
	log.Println(fmt.Sprintf("%s Starting server ....", tag))
	r := mux.NewRouter()

	r.HandleFunc("/requests", func(w http.ResponseWriter, r *http.Request) {
		var message Message

		if auth.Authenticate(r.Header) == false {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		status := pm.PushJobInQueues(message.Method, message.URL, message.TimeInMilliseconds)
		if status == false {
			http.Error(w, "Failed to enqueue jobs in redis", http.StatusInternalServerError)
			return
		}
	}).Methods("POST")

	log.Println(fmt.Sprintf("%s  Listening to port %s", tag, serverPort))
	http.ListenAndServe(fmt.Sprintf(":%s", serverPort), r)
}
