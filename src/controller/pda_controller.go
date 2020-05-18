package controller

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

type PdaController struct {
	PdaManager usecase.PDAManager
}

func (pdaController *PdaController) ListAllPDA(writer http.ResponseWriter, request *http.Request) {
	names := pdaController.PdaManager.ListAllPDAs()
	json, _ := json.Marshal(names)
	writer.WriteHeader(200)
	writer.Write(json)

}

func (pdaController *PdaController) OpenPDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]

	all, err := ioutil.ReadAll(request.Body)
	pdaConf := string(all)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = pdaController.PdaManager.CreateNewPDA(pda_id, pdaConf)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

}
func (pdaController *PdaController) ResetPDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]

	//call manager for resetting the pda_id
	pdaController.PdaManager.Reset(pda_id)
}

func (pdaController *PdaController) PutsToken(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]

	pdaStatus := pdaController.getPDAStatus(request)

	all, _ := ioutil.ReadAll(request.Body)
	token := string(all)
	position, _ := strconv.Atoi(params["position"])

	//call manager to pass a token with position
	err := pdaController.PdaManager.Puts(pda_id, token, position, pdaStatus)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	cookieDetails := pdaController.PdaManager.GetCookieFor(pdaStatus.ReplicaId, pdaStatus.PdaId)
	val, _ := json.Marshal(cookieDetails)

	cookie := &http.Cookie{
		Name:  "pda",
		Value: strings.ReplaceAll(string(val), "\"", "'"),
	}
	http.SetCookie(writer, cookie)
}

func (pdaController *PdaController) getPDAStatus(request *http.Request) entity.PDAStatus {
	cookies := request.Cookies()
	var pdaStatus entity.PDAStatus
	for i := range cookies {
		if cookies[i].Name == "pda" {
			pdaStatus = getCookieValue(cookies[i].Value)
		}
	}
	return pdaStatus
}
func getCookieValue(value string) entity.PDAStatus {
	var pdaStatus entity.PDAStatus

	unescape, _ := url.PathUnescape(value)
	all := strings.ReplaceAll(unescape, "\"", "")
	all = strings.ReplaceAll(all, "'", "\"")
	bytes := []byte(all)
	json.Unmarshal(bytes, &pdaStatus)
	return pdaStatus
}

func (pdaController *PdaController) PutsEOS(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	position, _ := strconv.Atoi(params["position"])
	//call manager to call eos for pda_id, ignore position of eos
	pdaStatus := pdaController.getPDAStatus(request)

	err := pdaController.PdaManager.PutsEOS(pda_id, position, pdaStatus)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	cookieDetails := pdaController.PdaManager.GetCookieFor(pdaStatus.ReplicaId, pdaStatus.PdaId)
	val, _ := json.Marshal(cookieDetails)

	cookie := &http.Cookie{
		Name:  "pda",
		Value: strings.ReplaceAll(string(val), "\"", "'"),
	}
	http.SetCookie(writer, cookie)

}

func (pdaController *PdaController) IsPDAAccepted(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	//call manager to call is_accepted method for pda_id
	pdaStatus := pdaController.getPDAStatus(request)

	is_accepted, err := pdaController.PdaManager.Is_accepted(pda_id, pdaStatus)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	accepted, err := json.Marshal(is_accepted)
	writer.Write(accepted)
}

func (pdaController *PdaController) PeekStack(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	peek_k, _ := strconv.Atoi(params["k"])
	//call manager to call peek method for pda_id and position peek_k from top
	pdaStatus := pdaController.getPDAStatus(request)

	top_k, err := pdaController.PdaManager.Peek(pda_id, peek_k, pdaStatus)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	top_kjson, _ := json.Marshal(top_k)
	writer.Write(top_kjson)
}

func (pdaController *PdaController) StackSize(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	//call manager to return the value of stack.Size() for pda_id
	pdaStatus := pdaController.getPDAStatus(request)

	size, err := pdaController.PdaManager.Size(pda_id, pdaStatus)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	stacksize, _ := json.Marshal(size)
	writer.Write(stacksize)
}

func (pdaController *PdaController) CurrentStatePDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	//call manager to return the value of pdaprocessor.current_state() for pda_id
	pdaStatus := pdaController.getPDAStatus(request)

	state, err := pdaController.PdaManager.Currentstate(pda_id, pdaStatus)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	currentstate, _ := json.Marshal(state)
	writer.Write(currentstate)
}

func (pdaController *PdaController) QueuedTokenPDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	//call manager to return the value of queued tokens( that method id yet to be implemented) for pda_id
	pdaStatus := pdaController.getPDAStatus(request)

	token, err := pdaController.PdaManager.Queued_token(pda_id, pdaStatus)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	q_token, _ := json.Marshal(token)
	writer.Write(q_token)
}
func (pdaController *PdaController) SnapshotPDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	peek_k, _ := strconv.Atoi(params["k"])
	//call manager to call three methods for pda_id: pdaprocessor.current_state(), queued tokens, peek(peek_k)
	pdaStatus := pdaController.getPDAStatus(request)

	state, err := pdaController.PdaManager.Currentstate(pda_id, pdaStatus)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := pdaController.PdaManager.Queued_token(pda_id, pdaStatus)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	peek, err := pdaController.PdaManager.Peek(pda_id, peek_k, pdaStatus)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	out := make([]string, 0)

	out = append(out, state)
	out = append(out, strings.Join(token, ","))
	out = append(out, strings.Join(peek, " "))

	json_out, _ := json.Marshal(out)

	writer.Write(json_out)
}
func (pdaController *PdaController) ClosePDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	//call manager to close the pda_id
	err := pdaController.PdaManager.Close(pda_id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (pdaController *PdaController) DeletePDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	//call manager to delete the pda_id
	err := pdaController.PdaManager.Deletepda(pda_id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (pdaController *PdaController) GetAllReplicaIds(writer http.ResponseWriter, request *http.Request) {
	data, _ := json.Marshal(pdaController.PdaManager.GetAllReplicaIds())
	writer.WriteHeader(200)
	writer.Write(data)
}

func (pdaController *PdaController) CreateReplicaGroup(writer http.ResponseWriter, request *http.Request) {
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

	err = pdaController.PdaManager.CreateNewReplicaGroup(replicaId, conf)
}

func (pdaController *PdaController) ResetAllMembers(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	replicaId, err := strconv.Atoi(params["gid"])

	if err != nil {
		http.Error(writer, "Invalid replica group ID provided", http.StatusBadRequest)
		return
	}

	pdaController.PdaManager.ResetReplicaMembers(replicaId)
}

func (pdaController *PdaController) GetMembersAddress(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	replicaId, err := strconv.Atoi(params["gid"])

	if err != nil {
		http.Error(writer, "Invalid replica group ID provided", http.StatusBadRequest)
		return
	}
	groupMembers := pdaController.PdaManager.GetMemberAddress(replicaId)
	writer.WriteHeader(200)
	data, _ := json.Marshal(groupMembers)
	writer.Write(data)
}

func (pdaController *PdaController) ConnectToAMember(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	replicaId, err := strconv.Atoi(params["gid"])

	if err != nil {
		http.Error(writer, "Invalid replica group ID provided", http.StatusBadRequest)
		return
	}
	groupMember := pdaController.PdaManager.GetRandomMemberAddress(replicaId)
	cookieDetails := pdaController.PdaManager.GetCookieFor(replicaId, groupMember)
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

func (pdaController *PdaController) CloseReplicaGrp(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	replicaId, err := strconv.Atoi(params["gid"])

	if err != nil {
		http.Error(writer, "Invalid replica group ID provided", http.StatusBadRequest)
		return
	}
	pdaController.PdaManager.CloseReplicaGrpAndMembers(replicaId)

}

func (pdaController *PdaController) DeleteReplicaGrp(writer http.ResponseWriter, request *http.Request) {

}

func (pdaController *PdaController) Joinpda(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pdaId := params["id"]

	all, err := ioutil.ReadAll(request.Body)
	replicaId, _ := strconv.Atoi(string(all))

	if err != nil {
		http.Error(writer, "Please pass Replica ID in body of the request", http.StatusBadRequest)
		return
	}

	pdaController.PdaManager.JoinAReplicaGrp(pdaId, replicaId)
}

func (pdaController *PdaController) Pdacode(writer http.ResponseWriter, request *http.Request) {

}
