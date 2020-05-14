package controller

import (
	"net/http"
	"usecase"
)

type ReplicaController struct {
	ReplicaManager usecase.ReplicaManager
}

func (replicaController *ReplicaController) GetAllReplicaIds(writer http.ResponseWriter, request *http.Request) {

}

func (replicaController *ReplicaController) CreateReplicagrp(writer http.ResponseWriter, request *http.Request) {

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
