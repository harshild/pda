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

func (pdaController *PdaController) Gotopdas(writer http.ResponseWriter, request *http.Request) {

	writer.WriteHeader(200)
	names := pdaController.PdaManager.ListAllPDAs()
	json, _ := json.Marshal(names)
	writer.Write(json)

}

func (pdaController *PdaController) Callopen(writer http.ResponseWriter, request *http.Request) {
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
func (pdaController *PdaController) Callreset(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callreset")
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)

	//call manager for resetting the pda_id
	pdaController.PdaManager.PdaProcessorcallsreset(pda_id)
}

func (pdaController *PdaController) Calltoken(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("calltoken")
	params := mux.Vars(request)
	pda_id := params["id"]

	all, _ := ioutil.ReadAll(request.Body)
	token := string(all)
	position, _ := strconv.Atoi(params["position"])

	fmt.Println(pda_id, token, position)
	//call manager to pass a token with position
	pdaController.PdaManager.PdaProcessorcallputs(pda_id, token, position)
}

func (pdaController *PdaController) Calleos(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("calleos")
	params := mux.Vars(request)
	pda_id := params["id"]
	eos := ""
	position, _ := strconv.Atoi(params["position"])
	fmt.Println(pda_id, position)
	//call manager to call eos for pda_id, ignore position of eos
	pdaController.PdaManager.PdaProcessorcallputs(pda_id, eos, position)
}

func (pdaController *PdaController) Callis_accepted(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callis_accepted")
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)
	//call manager to call is_accepted method for pda_id
	accepted, _ := json.Marshal(pdaController.PdaManager.PdaProcessorcallis_accepted(pda_id))
	writer.Write(accepted)
}

func (pdaController *PdaController) Callpeek(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callpeek")
	params := mux.Vars(request)
	pda_id := params["id"]
	peek_k, _ := strconv.Atoi(params["k"])
	fmt.Println(pda_id, peek_k)
	//call manager to call peek method for pda_id and position peek_k from top
	top_k, _ := json.Marshal(pdaController.PdaManager.Peek(pda_id, peek_k))
	writer.Write(top_k)
}

func (pdaController *PdaController) CallStackSize(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callstacksize")
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)
	//call manager to return the value of stack.Size() for pda_id
	stacksize, _ := json.Marshal(pdaController.PdaManager.Callsize(pda_id))
	writer.Write(stacksize)
}

func (pdaController *PdaController) Callcurrent_state(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callcurrent state")
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)
	//call manager to return the value of pdaprocessor.current_state() for pda_id
	currentstate, _ := json.Marshal(pdaController.PdaManager.Currentstate(pda_id))
	writer.Write(currentstate)
}

func (pdaController *PdaController) Callqueued_token(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callqueued token")
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)
	//call manager to return the value of queued tokens( that method id yet to be implemented) for pda_id
	q_token, _ := json.Marshal(pdaController.PdaManager.Q_token(pda_id))
	writer.Write(q_token)
}
func (pdaController *PdaController) Callthree_methods(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callthree methods")
	params := mux.Vars(request)
	pda_id := params["id"]
	peek_k, _ := strconv.Atoi(params["k"])
	fmt.Println(pda_id, peek_k)
	//call manager to call three methods for pda_id: pdaprocessor.current_state(), queued tokens, peek(peek_k)
	currentstate, _ := json.Marshal(pdaController.PdaManager.Currentstate(pda_id))
	q_token, _ := json.Marshal(pdaController.PdaManager.Q_token(pda_id))
	top_k, _ := json.Marshal(pdaController.PdaManager.Peek(pda_id, peek_k))

	writer.Write(currentstate)
	writer.Write(q_token)
	writer.Write(top_k)
}
func (pdaController *PdaController) Callclose(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callclose")
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)
	//call manager to close the pda_id
	pdaController.PdaManager.Callclose(pda_id)
}
func (pdaController *PdaController) Calldelete(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("calldelete")
	params := mux.Vars(request)
	pda_id := params["id"]
	fmt.Println(pda_id)
	//call manager to delete the pda_id
	pdaController.PdaManager.Deletepda(pda_id)
}
