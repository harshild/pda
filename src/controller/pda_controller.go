package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"usecase"
)

type PdaController struct {
	PdaManager usecase.PDAManager
}

func (pdaController *PdaController) ListAllPDA(writer http.ResponseWriter, request *http.Request) {

	writer.WriteHeader(200)
	names := pdaController.PdaManager.ListAllPDAs()
	json, _ := json.Marshal(names)
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
	pdaController.PdaManager.CreateNewPDA(pda_id, pdaConf)

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

	all, _ := ioutil.ReadAll(request.Body)
	token := string(all)
	position, _ := strconv.Atoi(params["position"])

	fmt.Println(pda_id, token, position)
	//call manager to pass a token with position
	pdaController.PdaManager.Puts(pda_id, token, position)
}

func (pdaController *PdaController) PutsEOS(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	eos := " "
	position, _ := strconv.Atoi(params["position"])
	fmt.Println(pda_id, position)
	//call manager to call eos for pda_id, ignore position of eos
	pdaController.PdaManager.Puts(pda_id, eos, position)
}

func (pdaController *PdaController) IsPDAAccepted(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)
	//call manager to call is_accepted method for pda_id
	accepted, _ := json.Marshal(pdaController.PdaManager.Is_accepted(pda_id))
	writer.Write(accepted)
}

func (pdaController *PdaController) PeekStack(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	peek_k, _ := strconv.Atoi(params["k"])
	fmt.Println(pda_id, peek_k)
	//call manager to call peek method for pda_id and position peek_k from top
	top_k, _ := json.Marshal(pdaController.PdaManager.Peek(pda_id, peek_k))
	writer.Write(top_k)
}

func (pdaController *PdaController) StackSize(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)
	//call manager to return the value of stack.Size() for pda_id
	stacksize, _ := json.Marshal(pdaController.PdaManager.Size(pda_id))
	writer.Write(stacksize)
}

func (pdaController *PdaController) CurrentStatePDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)
	//call manager to return the value of pdaprocessor.current_state() for pda_id
	currentstate, _ := json.Marshal(pdaController.PdaManager.Currentstate(pda_id))
	writer.Write(currentstate)
}

func (pdaController *PdaController) QueuedTokenPDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)
	//call manager to return the value of queued tokens( that method id yet to be implemented) for pda_id
	q_token, _ := json.Marshal(pdaController.PdaManager.Queued_token(pda_id))
	writer.Write(q_token)
}
func (pdaController *PdaController) SnapshotPDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	peek_k, _ := strconv.Atoi(params["k"])
	fmt.Println(pda_id, peek_k)
	//call manager to call three methods for pda_id: pdaprocessor.current_state(), queued tokens, peek(peek_k)
	currentstate, _ := json.Marshal(pdaController.PdaManager.Currentstate(pda_id))
	q_token, _ := json.Marshal(pdaController.PdaManager.Queued_token(pda_id))
	top_k, _ := json.Marshal(pdaController.PdaManager.Peek(pda_id, peek_k))

	writer.Write(currentstate)
	writer.Write(q_token)
	writer.Write(top_k)
}
func (pdaController *PdaController) ClosePDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)
	//call manager to close the pda_id
	pdaController.PdaManager.Close(pda_id)
}
func (pdaController *PdaController) DeletePDA(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)
	//call manager to delete the pda_id
	pdaController.PdaManager.Deletepda(pda_id)
}
