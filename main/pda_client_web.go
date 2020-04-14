package main

import (
	"controller"
	"core"
	"db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"usecase"
)

func main() {

	//store := db.SqliteStore{
	//	"sqlite3",
	//	"./test.db",
	//	nil,
	//	nil,
	//}

	//store := db.FileStore{
	//	"./test.db",
	//}

	store := db.InMemoryStore{}
	ctrl := controller.PdaController{
		usecase.PDAManager{
			core.PdaProcessor{},
			store,
		},
	}
	store.InitStore()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/pdas", ctrl.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/{id}", ctrl.Callopen).Methods("PUT")
	router.HandleFunc("/pdas/{id}/reset", ctrl.Callreset).Methods("PUT")
	router.HandleFunc("/pdas/{id}/{tokens}/{position}", ctrl.Calltoken).Methods("PUT")
	router.HandleFunc("/pdas/{id}/eos/{position}", ctrl.Calleos).Methods("PUT")
	router.HandleFunc("/pdas/{id}/is_accepted", ctrl.Callis_accepted).Methods("GET")
	router.HandleFunc("/pdas/{id}/stack/top/{k}", ctrl.Callpeek).Methods("GET")
	router.HandleFunc("/pdas/{id}/stack/len", ctrl.CallStackSize).Methods("GET")
	router.HandleFunc("/pdas/{id}/state", ctrl.Callcurrent_state).Methods("GET")
	router.HandleFunc("/pdas/{id}/tokens", ctrl.Callqueued_token).Methods("GET")
	router.HandleFunc("/pdas/{id}/snapshot/{k}", ctrl.Callthree_methods).Methods("GET")
	router.HandleFunc("/pdas/{id}/close", ctrl.Callclose).Methods("PUT")
	router.HandleFunc("/pdas/{id}/delete", ctrl.Calldelete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
