package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"discovery/orm"
	"discovery/pg"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/orm", orm.ConnHandler).Methods(http.MethodGet)
	r.HandleFunc("/postgres", pg.ConnHandler).Methods(http.MethodGet)
	http.ListenAndServe("0.0.0.0:5555", r)
}
