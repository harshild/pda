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
	if pdaStatus.PdaId == pda_id {

		all, _ := ioutil.ReadAll(request.Body)
		token := string(all)
		position, _ := strconv.Atoi(params["position"])

		//call manager to pass a token with position
		err := pdaController.PdaManager.Puts(pda_id, token, position, pdaStatus)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		cookieDetails := pdaController.PdaManager.GetCookieFor(pdaStatus.ReplicaId, pdaStatus.PdaId)
		pdaController.WriteCookie(writer, cookieDetails)
	} else {
		http.Error(writer, "Connect to the PDA first!! then send token", 501)
	}

}

func (pdaController *PdaController) PutsEOS(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	position, _ := strconv.Atoi(params["position"])
	//call manager to call eos for pda_id, ignore position of eos
	pdaStatus := pdaController.getPDAStatus(request)

	if pdaStatus.PdaId == pda_id {
		err := pdaController.PdaManager.PutsEOS(pda_id, position, pdaStatus)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		cookieDetails := pdaController.PdaManager.GetCookieFor(pdaStatus.ReplicaId, pdaStatus.PdaId)

		pdaController.WriteCookie(writer, cookieDetails)
	} else {
		http.Error(writer, "Connect to the PDA first,then send EOS", 502)
	}
}

func (pdaController *PdaController) IsPDAAccepted(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	//call manager to call is_accepted method for pda_id
	pdaStatus := pdaController.getPDAStatus(request)

	if pdaStatus.PdaId == pda_id {
		is_accepted, err := pdaController.PdaManager.Is_accepted(pda_id, pdaStatus)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		accepted, err := json.Marshal(is_accepted)
		writer.Write(accepted)
	} else {
		http.Error(writer, "connect to the PDA first, then call is_accepted ", 503)
	}
}

func (pdaController *PdaController) PeekStack(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	peek_k, _ := strconv.Atoi(params["k"])
	//call manager to call peek method for pda_id and position peek_k from top
	pdaStatus := pdaController.getPDAStatus(request)

	if pdaStatus.PdaId == pda_id {
		top_k, err := pdaController.PdaManager.Peek(pda_id, peek_k, pdaStatus)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		top_kjson, _ := json.Marshal(top_k)
		writer.Write(top_kjson)
	} else {
		http.Error(writer, "First,Connect to the member of Replica group! ", 504)
	}
}

func (pdaController *PdaController) StackSize(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	//call manager to return the value of stack.Size() for pda_id
	pdaStatus := pdaController.getPDAStatus(request)

	if pdaStatus.PdaId == pda_id {
		size, err := pdaController.PdaManager.Size(pda_id, pdaStatus)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		stacksize, _ := json.Marshal(size)
		writer.Write(stacksize)
	} else {
		http.Error(writer, "First,Connect to the member of Replica group! ", 505)
	}
}

func (pdaController *PdaController) CurrentStatePDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	//call manager to return the value of pdaprocessor.current_state() for pda_id
	pdaStatus := pdaController.getPDAStatus(request)

	if pdaStatus.PdaId == pda_id {
		state, err := pdaController.PdaManager.Currentstate(pda_id, pdaStatus)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		currentstate, _ := json.Marshal(state)
		writer.Write(currentstate)
	} else {
		http.Error(writer, "First,Connect to the member of Replica group! ", 506)
	}
}

func (pdaController *PdaController) QueuedTokenPDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	//call manager to return the value of queued tokens( that method id yet to be implemented) for pda_id
	pdaStatus := pdaController.getPDAStatus(request)

	if pdaStatus.PdaId == pda_id {
		token, err := pdaController.PdaManager.Queued_token(pda_id, pdaStatus)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		q_token, _ := json.Marshal(token)
		writer.Write(q_token)
	} else {
		http.Error(writer, "First,Connect to the member of Replica group! ", 507)
	}

}
func (pdaController *PdaController) SnapshotPDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	peek_k, _ := strconv.Atoi(params["k"])
	//call manager to call three methods for pda_id: pdaprocessor.current_state(), queued tokens, peek(peek_k)
	pdaStatus := pdaController.getPDAStatus(request)

	if pdaStatus.PdaId == pda_id {
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
	} else {
		http.Error(writer, "First,Connect to the member of Replica group! ", 508)
	}

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
	groupMember := pdaController.PdaManager.GetRandomMemberAddress(replicaId)
	cookieDetails := pdaController.PdaManager.GetCookieFor(replicaId, groupMember)

	pdaController.WriteCookie(writer, cookieDetails)

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

	pdsStatus := pdaController.getPDAStatus(request)

	groupMember := pdaController.PdaManager.GetRandomMemberAddress(replicaId)
	cookieDetails := pdaController.PdaManager.GetCookieFor(replicaId, groupMember)

	if len(request.Cookies()) > 0 && pdsStatus.ReplicaId == replicaId {
		cookieDetails.LastConsumedIndex = pdsStatus.LastConsumedIndex
		cookieDetails.Stack = pdsStatus.Stack
		cookieDetails.State = pdsStatus.State
		cookieDetails.InputQueue = pdsStatus.InputQueue
		cookieDetails.Clock = pdsStatus.Clock
	}

	pdaController.WriteCookie(writer, cookieDetails)
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
	params := mux.Vars(request)
	replicaId, err := strconv.Atoi(params["gid"])

	if err != nil {
		http.Error(writer, "Invalid replica group ID provided", http.StatusBadRequest)
		return
	}
	pdaController.PdaManager.DeleteReplicaGrpAndMembers(replicaId)
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

func (pdaController *PdaController) GetC3State(writer http.ResponseWriter, request *http.Request) {
	pdsStatus := pdaController.getPDAStatus(request)
	cookieInfo := entity.CookieInfo{
		PdaId:     pdsStatus.PdaId,
		ReplicaId: pdsStatus.ReplicaId,
	}
	data, _ := json.Marshal(cookieInfo)
	writer.WriteHeader(200)
	writer.Write(data)
}

func (pdaController *PdaController) GetPDACode(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pdaId := params["id"]

	code := pdaController.PdaManager.GetPDACode(pdaId)

	data, _ := json.Marshal(code)
	writer.WriteHeader(200)
	writer.Write(data)
}

func (pdaController *PdaController) getPDAStatus(request *http.Request) entity.PDAStatus {
	cookies := request.Cookies()
	var pdaStatus entity.PDAStatus
	for i := range cookies {
		if cookies[i].Name == "pda" {
			cookieInfo := getCookieValue(cookies[i].Value)
			pdaStatus = pdaController.PdaManager.GetPDA(cookieInfo.ReplicaId, cookieInfo.PdaId)
		}
	}
	return pdaStatus
}
func getCookieValue(value string) entity.CookieInfo {
	var cookieInfo entity.CookieInfo

	unescape, _ := url.PathUnescape(value)
	all := strings.ReplaceAll(unescape, "\"", "")
	all = strings.ReplaceAll(all, "'", "\"")
	bytes := []byte(all)
	json.Unmarshal(bytes, &cookieInfo)
	return cookieInfo
}

func (pdaController *PdaController) WriteCookie(writer http.ResponseWriter, cookieDetails entity.PDAStatus) {
	cookieInfo := entity.CookieInfo{
		PdaId:     cookieDetails.PdaId,
		ReplicaId: cookieDetails.ReplicaId,
	}

	val, _ := json.Marshal(cookieInfo)

	cookie := &http.Cookie{
		Name:  "pda",
		Value: strings.ReplaceAll(string(val), "\"", "'"),
		Path:  "/",
	}
	http.SetCookie(writer, cookie)
}
