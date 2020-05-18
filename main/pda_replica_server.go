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

	pdaController := controller.PdaController{
		usecase.PDAManager{
			db.InMemoryStore{},
		},
	}

	pdaController.PdaManager.PdaStore.InitStore()

	replicaController := controller.ReplicaController{
		ReplicaManager: usecase.ReplicaManager{
			ReplicaStore: db.ReplicaInMemoryStore{},
		},
	}

	replicaController.ReplicaManager.ReplicaStore.InitStore()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/pdas", replicaController.ListAllPDA).Methods("GET")
	//router.HandleFunc("/pdas/{id}", replicaController.OpenPDA).Methods("PUT")
	//router.HandleFunc("/pdas/{id}/reset", replicaController.ResetPDA).Methods("PUT")
	//router.HandleFunc("/pdas/{id}/tokens/{position}", replicaController.PutsToken).Methods("PUT")
	//router.HandleFunc("/pdas/{id}/eos/{position}", replicaController.PutsEOS).Methods("PUT")
	//router.HandleFunc("/pdas/{id}/is_accepted", replicaController.IsPDAAccepted).Methods("GET")
	//router.HandleFunc("/pdas/{id}/stack/top/{k}", replicaController.PeekStack).Methods("GET")
	//router.HandleFunc("/pdas/{id}/stack/len", replicaController.StackSize).Methods("GET")
	//router.HandleFunc("/pdas/{id}/state", replicaController.CurrentStatePDA).Methods("GET")
	//router.HandleFunc("/pdas/{id}/tokens", replicaController.QueuedTokenPDA).Methods("GET")
	//router.HandleFunc("/pdas/{id}/snapshot/{k}", replicaController.SnapshotPDA).Methods("GET")
	//router.HandleFunc("/pdas/{id}/close", replicaController.ClosePDA).Methods("PUT")
	//router.HandleFunc("/pdas/{id}/delete", replicaController.DeletePDA).Methods("DELETE")

	router.HandleFunc("/replica_pdas", replicaController.GetAllReplicaIds).Methods("GET")
	router.HandleFunc("/replica_pdas/{gid}", replicaController.CreateReplicaGroup).Methods("PUT") //takes pda specification and group member ids
	router.HandleFunc("/replica_pdas/{gid}/reset", replicaController.ResetAllMembers).Methods("PUT")
	router.HandleFunc("/replica_pdas/{gid}/members", replicaController.GetMembersAddress).Methods("GET")
	router.HandleFunc("/replica_pdas/{gid}/connect", replicaController.ConnectToAMember).Methods("GET") //Return the address of a random member that a client could connect to
	router.HandleFunc("/replica_pdas/{gid}/close", replicaController.CloseReplicaGrp).Methods("PUT")    //close all members
	router.HandleFunc("/replica_pdas/{gid}/delete", replicaController.DeleteReplicaGrp).Methods("DELETE")
	router.HandleFunc("/pdas/{id}/join", replicaController.Joinpda).Methods("PUT")
	router.HandleFunc("/pdas/{id}/code", replicaController.Pdacode).Methods("GET")
	router.HandleFunc("/pdas/{id}/c3state", replicaController.Pdacode).Methods("GET") //state information

	log.Fatal(http.ListenAndServe(":8080", router))
}
