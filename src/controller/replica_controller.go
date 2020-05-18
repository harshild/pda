package controller

import "C"
import (
	"encoding/json"
	"entity"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
	params := mux.Vars(request)
	replica_id := params["gid"]

	replicaController.ReplicaManager.ResetReplicaMembers(replica_id)
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
		Value: strings.ReplaceAll(string(val), "\"", "'"),
	}
	http.SetCookie(writer, cookie)
	data, _ := json.Marshal(groupMember)
	writer.WriteHeader(200)
	writer.Write(data)
}

func (replicaController *ReplicaController) CloseReplicaGrp(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	replicaId, err := strconv.Atoi(params["gid"])

	if err != nil {
		http.Error(writer, "Invalid replica group ID provided", http.StatusBadRequest)
		return
	}
	replicaController.ReplicaManager.CloseReplicaGrpAndMembers(replicaId)

}

func (replicaController *ReplicaController) DeleteReplicaGrp(writer http.ResponseWriter, request *http.Request) {

}

func (replicaController *ReplicaController) Joinpda(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pdaId := params["id"]

	all, err := ioutil.ReadAll(request.Body)
	replicaId, _ := strconv.Atoi(string(all))

	if err != nil {
		http.Error(writer, "Please pass Replica ID in request body", http.StatusBadRequest)
		return
	}

	replicaController.ReplicaManager.JoinAReplicaGrp(pdaId, replicaId)
}

func (replicaController *ReplicaController) Pdacode(writer http.ResponseWriter, request *http.Request) {

}

// PDA APIs with respect to Replica

func (replicaController *ReplicaController) ListAllPDA(writer http.ResponseWriter, request *http.Request) {
	cookies := request.Cookies()
	var pdaStatus entity.PDAStatus
	for i := range cookies {
		if cookies[i].Name == "pda" {
			pdaStatus = getCookieValue(cookies[i].Value)
		}
	}

	names := replicaController.ReplicaManager.ListAllPDAs(pdaStatus)
	json, _ := json.Marshal(names)
	writer.WriteHeader(200)
	writer.Write(json)

}

func getCookieValue(value string) entity.PDAStatus {
	var pdaStatus entity.PDAStatus

	unescape, _ := url.PathUnescape(value)
	all := strings.ReplaceAll(unescape, "\"", "")
	all = strings.ReplaceAll(all, "'", "\"")
	print(all)
	bytes := []byte(all)
	json.Unmarshal(bytes, &pdaStatus)
	return pdaStatus
}

//func (replicaController *ReplicaController) OpenPDA(writer http.ResponseWriter, request *http.Request) {
//	params := mux.Vars(request)
//	pda_id := params["id"]
//
//	all, err := ioutil.ReadAll(request.Body)
//	pdaConf := string(all)
//
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusBadRequest)
//		return
//	}
//	err = replicaController.ReplicaManager.CreateNewPDA(pda_id, pdaConf)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusBadRequest)
//	}
//
//}
//func (replicaController *ReplicaController) ResetPDA(writer http.ResponseWriter, request *http.Request) {
//	params := mux.Vars(request)
//	pda_id := params["id"]
//
//	//call manager for resetting the pda_id
//	replicaController.ReplicaManager.Reset(pda_id)
//}
//
//func (replicaController *ReplicaController) PutsToken(writer http.ResponseWriter, request *http.Request) {
//	params := mux.Vars(request)
//	pda_id := params["id"]
//
//	all, _ := ioutil.ReadAll(request.Body)
//	token := string(all)
//	position, _ := strconv.Atoi(params["position"])
//
//	//call manager to pass a token with position
//	err := replicaController.ReplicaManager.Puts(pda_id, token, position)
//
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//func (replicaController *ReplicaController) PutsEOS(writer http.ResponseWriter, request *http.Request) {
//	params := mux.Vars(request)
//	pda_id := params["id"]
//	position, _ := strconv.Atoi(params["position"])
//	//call manager to call eos for pda_id, ignore position of eos
//	err := replicaController.ReplicaManager.PutsEOS(pda_id, position)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}
//
//func (replicaController *ReplicaController) IsPDAAccepted(writer http.ResponseWriter, request *http.Request) {
//	params := mux.Vars(request)
//	pda_id := params["id"]
//	//call manager to call is_accepted method for pda_id
//	is_accepted, err := replicaController.ReplicaManager.Is_accepted(pda_id)
//
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	accepted, err := json.Marshal(is_accepted)
//	writer.Write(accepted)
//}
//
//func (replicaController *ReplicaController) PeekStack(writer http.ResponseWriter, request *http.Request) {
//	params := mux.Vars(request)
//	pda_id := params["id"]
//	peek_k, _ := strconv.Atoi(params["k"])
//	//call manager to call peek method for pda_id and position peek_k from top
//	top_k, err := replicaController.ReplicaManager.Peek(pda_id, peek_k)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	top_kjson, _ := json.Marshal(top_k)
//	writer.Write(top_kjson)
//}
//
//func (replicaController *ReplicaController) StackSize(writer http.ResponseWriter, request *http.Request) {
//	params := mux.Vars(request)
//	pda_id := params["id"]
//	//call manager to return the value of stack.Size() for pda_id
//	size, err := replicaController.ReplicaManager.Size(pda_id)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	stacksize, _ := json.Marshal(size)
//	writer.Write(stacksize)
//}
//
//func (replicaController *ReplicaController) CurrentStatePDA(writer http.ResponseWriter, request *http.Request) {
//	params := mux.Vars(request)
//	pda_id := params["id"]
//	//call manager to return the value of pdaprocessor.current_state() for pda_id
//	state, err := replicaController.ReplicaManager.Currentstate(pda_id)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	currentstate, _ := json.Marshal(state)
//	writer.Write(currentstate)
//}
//
//func (replicaController *ReplicaController) QueuedTokenPDA(writer http.ResponseWriter, request *http.Request) {
//	params := mux.Vars(request)
//	pda_id := params["id"]
//	//call manager to return the value of queued tokens( that method id yet to be implemented) for pda_id
//	token, err := replicaController.ReplicaManager.Queued_token(pda_id)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	q_token, _ := json.Marshal(token)
//	writer.Write(q_token)
//}
//func (replicaController *ReplicaController) SnapshotPDA(writer http.ResponseWriter, request *http.Request) {
//	params := mux.Vars(request)
//	pda_id := params["id"]
//	peek_k, _ := strconv.Atoi(params["k"])
//	//call manager to call three methods for pda_id: pdaprocessor.current_state(), queued tokens, peek(peek_k)
//	state, err := replicaController.ReplicaManager.Currentstate(pda_id)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	token, err := replicaController.ReplicaManager.Queued_token(pda_id)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	peek, err := replicaController.ReplicaManager.Peek(pda_id, peek_k)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	out := make([]string, 0)
//
//	out = append(out, state)
//	out = append(out, strings.Join(token, ","))
//	out = append(out, strings.Join(peek, " "))
//
//	json_out, _ := json.Marshal(out)
//
//	writer.Write(json_out)
//}
//func (replicaController *ReplicaController) ClosePDA(writer http.ResponseWriter, request *http.Request) {
//	params := mux.Vars(request)
//	pda_id := params["id"]
//	//call manager to close the pda_id
//	err := replicaController.ReplicaManager.Close(pda_id)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}
//func (replicaController *ReplicaController) DeletePDA(writer http.ResponseWriter, request *http.Request) {
//	params := mux.Vars(request)
//	pda_id := params["id"]
//	//call manager to delete the pda_id
//	err := replicaController.ReplicaManager.Deletepda(pda_id)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}
