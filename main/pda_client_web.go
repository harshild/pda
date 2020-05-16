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

	pdaController := controller.PdaController{
		usecase.PDAManager{
			db.InMemoryStore{},
		},
	}

	pdaController.PdaManager.PdaStore.InitStore()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/pdas", pdaController.ListAllPDA).Methods("GET")
	router.HandleFunc("/pdas/{id}", pdaController.OpenPDA).Methods("PUT")
	router.HandleFunc("/pdas/{id}/reset", pdaController.ResetPDA).Methods("PUT")
	router.HandleFunc("/pdas/{id}/tokens/{position}", pdaController.PutsToken).Methods("PUT")
	router.HandleFunc("/pdas/{id}/eos/{position}", pdaController.PutsEOS).Methods("PUT")
	router.HandleFunc("/pdas/{id}/is_accepted", pdaController.IsPDAAccepted).Methods("GET")
	router.HandleFunc("/pdas/{id}/stack/top/{k}", pdaController.PeekStack).Methods("GET")
	router.HandleFunc("/pdas/{id}/stack/len", pdaController.StackSize).Methods("GET")
	router.HandleFunc("/pdas/{id}/state", pdaController.CurrentStatePDA).Methods("GET")
	router.HandleFunc("/pdas/{id}/tokens", pdaController.QueuedTokenPDA).Methods("GET")
	router.HandleFunc("/pdas/{id}/snapshot/{k}", pdaController.SnapshotPDA).Methods("GET")
	router.HandleFunc("/pdas/{id}/close", pdaController.ClosePDA).Methods("PUT")
	router.HandleFunc("/pdas/{id}/delete", pdaController.DeletePDA).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
