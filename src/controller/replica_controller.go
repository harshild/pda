package controller

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"usecase"
)

type ReplicaController struct {
	ReplicaManager usecase.ReplicaManager
}

func (replicaController *ReplicaController) GetAllReplicaIds(writer http.ResponseWriter, request *http.Request) {

}

func (replicaController *ReplicaController) CreateReplicagrp(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	Replica_id := params["gid"]

	all, err := ioutil.ReadAll(request.Body)
	conf := string(all)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = replicaController.ReplicaManager.CreateNewReplicagrp(Replica_id, conf)

}

func (replicaController *ReplicaController) ResetAllMembers(writer http.ResponseWriter, request *http.Request) {

}

func (replicaController *ReplicaController) GetMembersAddress(writer http.ResponseWriter, request *http.Request) {

}

func (replicaController *ReplicaController) ConnectToAMember(writer http.ResponseWriter, request *http.Request) {

}

func (replicaController *ReplicaController) CloseReplicaGrp(writer http.ResponseWriter, request *http.Request) {

}

func (replicaController *ReplicaController) DeleteReplicaGrp(writer http.ResponseWriter, request *http.Request) {

}

func (replicaController *ReplicaController) Joinpda(writer http.ResponseWriter, request *http.Request) {

}

func (replicaController *ReplicaController) Pdacode(writer http.ResponseWriter, request *http.Request) {

}
