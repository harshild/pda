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
	replicaController.ReplicaManager.GetAllReplicaIds()
}

func (replicaController *ReplicaController) CreateReplicaGroup(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	replicaId := params["gid"]

	all, err := ioutil.ReadAll(request.Body)
	conf := string(all)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = replicaController.ReplicaManager.CreateNewReplicaGroup(replicaId, conf)

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
