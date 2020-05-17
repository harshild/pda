package controller

import "C"
import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"usecase"
)

type ReplicaController struct {
	ReplicaManager usecase.ReplicaManager
}

func (replicaController *ReplicaController) GetAllReplicaIds(writer http.ResponseWriter, request *http.Request) {
	data, _ := json.Marshal(replicaController.ReplicaManager.GetAllReplicaIds())
	writer.WriteHeader(200)
	writer.Write(data)
}

func (replicaController *ReplicaController) CreateReplicaGroup(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	replicaId, err := strconv.Atoi(params["gid"])

	if err != nil {
		http.Error(writer, "Invalid replica group ID provided", http.StatusBadRequest)
		return
	}

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
	params := mux.Vars(request)
	replicaId, err := strconv.Atoi(params["gid"])

	if err != nil {
		http.Error(writer, "Invalid replica group ID provided", http.StatusBadRequest)
		return
	}
	groupMembers := replicaController.ReplicaManager.GetMemberAddress(replicaId)
	writer.WriteHeader(200)
	data, _ := json.Marshal(groupMembers)
	writer.Write(data)
}

func (replicaController *ReplicaController) ConnectToAMember(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	replicaId, err := strconv.Atoi(params["gid"])

	if err != nil {
		http.Error(writer, "Invalid replica group ID provided", http.StatusBadRequest)
		return
	}
	groupMember := replicaController.ReplicaManager.GetRandomMemberAddress(replicaId)
	cookieDetails := replicaController.ReplicaManager.GetCookieFor(replicaId, groupMember)
	val, _ := json.Marshal(cookieDetails)

	cookie := &http.Cookie{
		Name:  "pda",
		Value: string(val),
	}
	http.SetCookie(writer, cookie)
	data, _ := json.Marshal(groupMember)
	writer.WriteHeader(200)
	writer.Write(data)
}

func (replicaController *ReplicaController) CloseReplicaGrp(writer http.ResponseWriter, request *http.Request) {

}

func (replicaController *ReplicaController) DeleteReplicaGrp(writer http.ResponseWriter, request *http.Request) {

}

func (replicaController *ReplicaController) Joinpda(writer http.ResponseWriter, request *http.Request) {

}

func (replicaController *ReplicaController) Pdacode(writer http.ResponseWriter, request *http.Request) {

}
