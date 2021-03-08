package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gitswarm.f5net.com/salton/reservations/pkg/reservation"

	"github.com/gorilla/mux"

	pg "orm/postgres"
)

func reservationsHandler(i *pg.Instance) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {

		log.Printf("List Request accepted : %v %v", req.Method, req.URL)
		rows, _ := i.AllReservations()

		log.Printf("rows: %v", rows)
		b, err := json.Marshal(rows)
		if err != nil {
			log.Printf("Error: unexpected decoding issue.\nreason: %v", err)
			wrt.WriteHeader(http.StatusBadRequest)
		}

		wrt.Write(b)

	})
}

func createReservationsHandler(i *pg.Instance) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {

		log.Printf("Create Request accepted : %v %v", req.Method, req.URL)
		var r reservation.Reservation
		body, err := ioutil.ReadAll(req.Body)

		if err = json.Unmarshal(body, &r); err != nil {
			log.Printf("Error: request body incorrect.\nreason: %v", err)
			wrt.WriteHeader(http.StatusBadRequest)
			return
		}
		defer req.Body.Close()

		i.CreateReservation(r)

	})
}

func updateHandler(i *pg.Instance) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {

		log.Printf("Update Request accepted: %v %v", req.Method, req.URL)
		var r reservation.Reservation
		body, err := ioutil.ReadAll(req.Body)

		if err = json.Unmarshal(body, &r); err != nil {
			log.Printf("Error: request body incorrect.\nreason: %v", err)
			wrt.WriteHeader(http.StatusBadRequest)
			return
		}
		defer req.Body.Close()

		i.UpdateReservation(r)

	})
}

func deleteHandler(i *pg.Instance) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {

		log.Printf("Delete Request accepted : %v %v", req.Method, req.URL)

		params := mux.Vars(req)
		id := params["id"]

		err := i.DeleteReservation(id)
		if err != nil {
			log.Printf("Error: failed to delete entry %v.\nreason: %v", id, err)
		}

	})
}

func main() {

	i, err := pg.New()
	if err != nil {
		log.Fatalf("could not establish a connection to the database, reason : %v", err)
	}
	defer i.Close()

	r := mux.NewRouter()

	r.HandleFunc("/reservations", reservationsHandler(i)).Methods(http.MethodGet)
	r.HandleFunc("/reservations", createReservationsHandler(i)).Methods(http.MethodPost)
	r.HandleFunc("/reservations/{id:[0-9]+}", deleteHandler(i)).Methods(http.MethodDelete)
	r.HandleFunc("/reservations/{id:[0-9]+}", updateHandler(i)).Methods(http.MethodPut)

	// Run the web server.
	log.Println("starting orm-service...")

	log.Fatal(http.ListenAndServe("0.0.0.0:5431", r))

}
