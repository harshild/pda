package main

import (
	"controller"
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

	ctrl := controller.PdaController{
		usecase.PDAManager{
			db.InMemoryStore{},
		},
	}

	ctrl.PdaManager.PdaStore.InitStore()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/pdas", ctrl.ListAllPDA).Methods("GET")
	router.HandleFunc("/pdas/{id}", ctrl.OpenPDA).Methods("PUT")
	router.HandleFunc("/pdas/{id}/reset", ctrl.ResetPDA).Methods("PUT")
	router.HandleFunc("/pdas/{id}/tokens/{position}", ctrl.PutsToken).Methods("PUT")
	router.HandleFunc("/pdas/{id}/eos/{position}", ctrl.PutsEOS).Methods("PUT")
	router.HandleFunc("/pdas/{id}/is_accepted", ctrl.IsPDAAccepted).Methods("GET")
	router.HandleFunc("/pdas/{id}/stack/top/{k}", ctrl.PeekStack).Methods("GET")
	router.HandleFunc("/pdas/{id}/stack/len", ctrl.StackSize).Methods("GET")
	router.HandleFunc("/pdas/{id}/state", ctrl.CurrentStatePDA).Methods("GET")
	router.HandleFunc("/pdas/{id}/tokens", ctrl.QueuedTokenPDA).Methods("GET")
	router.HandleFunc("/pdas/{id}/snapshot/{k}", ctrl.SnapshotPDA).Methods("GET")
	router.HandleFunc("/pdas/{id}/close", ctrl.ClosePDA).Methods("PUT")
	router.HandleFunc("/pdas/{id}/delete", ctrl.DeletePDA).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
