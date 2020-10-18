package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Kareem-Emad/simple-apm/auth"
	"github.com/Kareem-Emad/simple-apm/dal"
	"github.com/Kareem-Emad/simple-apm/producer"
	"github.com/gorilla/mux"
)

var tag = "[Server]"

// Start starts an http server with the api routes loaded as specified
func Start(pm producer.ProductionManager) {
	log.Println(fmt.Sprintf("%s Starting server ....", tag))
	r := mux.NewRouter()

	r.HandleFunc("/requests", func(w http.ResponseWriter, r *http.Request) {
		var payload dal.RequestStats

		if auth.Authenticate(r.Header) == false {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&payload)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// we are doing a lot of work here, so it's better to validate such stuff early on
		if dal.ValidateRequestStats(payload) != true {
			http.Error(w, "missing data fields in payload", http.StatusBadRequest)
			return
		}

		status := pm.PushJobInQueues(payload)
		if status == false {
			http.Error(w, "Failed to enqueue jobs in redis", http.StatusInternalServerError)
			return
		}
	}).Methods("POST")

	log.Println(fmt.Sprintf("%s  Listening to port %s", tag, serverPort))
	http.ListenAndServe(fmt.Sprintf(":%s", serverPort), r)
}
