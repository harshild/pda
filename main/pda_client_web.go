package main

import (
	. "../src/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	controller:= PdaController{}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/pdas",controller.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id",controller.Gotopdas).Methods("PUT")
	router.HandleFunc("/pdas/id/",controller.Gotopdas).Methods("PUT")
	router.HandleFunc("/pdas/id/tokens/position",controller.Gotopdas).Methods("PUT")
	router.HandleFunc("/pdas/id/eos/position",controller.Gotopdas).Methods("PUT")
	router.HandleFunc("/pdas/id/is_accepted",controller.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id/stack/top/k",controller.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id/stack/len",controller.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id/state",controller.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id/tokens",controller.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id/snapshot/k",controller.Gotopdas).Methods("GET")
	router.HandleFunc("/pdas/id/close",controller.Gotopdas).Methods("PUT")
	router.HandleFunc("/pdas/id/delete",controller.Gotopdas).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

