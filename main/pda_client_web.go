package main

import (
	"controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	ctrl := controller.PdaController{}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/pdas", ctrl.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id", ctrl.Gotopdas).Methods("PUT")
	router.HandleFunc("/pdas/id/", ctrl.Gotopdas).Methods("PUT")
	router.HandleFunc("/pdas/id/tokens/position", ctrl.Gotopdas).Methods("PUT")
	router.HandleFunc("/pdas/id/eos/position", ctrl.Gotopdas).Methods("PUT")
	router.HandleFunc("/pdas/id/is_accepted", ctrl.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id/stack/top/k", ctrl.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id/stack/len", ctrl.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id/state", ctrl.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id/tokens", ctrl.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id/snapshot/k", ctrl.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id/close", ctrl.Closepda).Methods("PUT")
	router.HandleFunc("/pdas/id/delete", ctrl.Gotopdas).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
