package controller

import (
	"encoding/json"
	"entity"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
	pda_id := mux.Vars(request)
	fmt.Println(pda_id)

	var conf entity.PDAConf
	err := json.NewDecoder(request.Body).Decode(&conf)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(writer, "GOT IT!!!!!: %+v", conf)
	//call manager and pass pda_id and conf

}
func (pdaController *PdaController) Callreset(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callreset")
	pda_id := mux.Vars(request)
	fmt.Println(pda_id)

	//call manager for resetting the pda_id
}

func (pdaController *PdaController) Calltoken(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("calltoken")
	params := mux.Vars(request)
	pda_id := params["id"]
	token := params["tokens"]
	position := params["position"]

	fmt.Println(pda_id, token, position)
	//call manager to pass a token with position
}

func (pdaController *PdaController) Calleos(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("calleos")
	params := mux.Vars(request)
	pda_id := params["id"]
	position := params["position"]
	fmt.Println(pda_id, position)
	//call manager to call eos for pda_id, ignore position of eos
}

func (pdaController *PdaController) Callis_accepted(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callis_accepted")
	pda_id := mux.Vars(request)
	fmt.Println(pda_id)
	//call manager to call is_accepted method for pda_id
}

func (pdaController *PdaController) Callpeek(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callpeek")
	params := mux.Vars(request)
	pda_id := params["id"]
	peek_k := params["k"]
	fmt.Println(pda_id, peek_k)
	//call manager to call peek method for pda_id and position peek_k from top
}

func (pdaController *PdaController) CallStackSize(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callstacksize")
	pda_id := mux.Vars(request)
	fmt.Println(pda_id)
	//call manager to return the value of stack.Size() for pda_id
}

func (pdaController *PdaController) Callcurrent_state(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callcurrent state")
	pda_id := mux.Vars(request)
	fmt.Println(pda_id)
	//call manager to return the value of pdaprocessor.current_state() for pda_id
}

func (pdaController *PdaController) Callqueued_token(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callqueued token")
	pda_id := mux.Vars(request)
	fmt.Println(pda_id)
	//call manager to return the value of queued tokens( that method id yet to be implemented) for pda_id
}
func (pdaController *PdaController) Callthree_methods(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callthree methods")
	params := mux.Vars(request)
	pda_id := params["id"]
	peek_k := params["k"]
	fmt.Println(pda_id, peek_k)
	//call manager to call three methods for pda_id: pdaprocessor.current_state(), queued tokens, peek(peek_k)
}
func (pdaController *PdaController) Callclose(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("callclose")
	pda_id := mux.Vars(request)
	fmt.Println(pda_id)
	//call manager to close the pda_id
}
func (pdaController *PdaController) Calldelete(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("calldelete")
	pda_id := mux.Vars(request)
	fmt.Println(pda_id)
	//call manager to delete the pda_id
}

/*func (pdaController *PdaController) Closepda(writer http.ResponseWriter, request *http.Request) {

	//TODO: Here "null" is going to create a problem, We are yet to handle it
	//
	//PDA1 := PDAConf{"HelloPDA", []string{"q1", "q2", "q3", "q4"}, []string{"0", "1"}, []string{"0", "1"},
	//	[]string{"q1", "q4"}, "q1", [][]string{{"q1", "null", "null", "q2", "$"}, {"q2", "0", "null", "q2", "0"},
	//		{"q2", "1", "0", "q3", "null"}, {"q3", "1", "0", "q3", "null"}, {"q3", "null", "$", "q4", "null"}}, "$"}
	//
	//PDA2 := PDAConf{
	//	"testPDA2",
	//	[]string{"q1", "q2", "q3", "q4"},
	//	[]string{"a", "b", "c"},
	//	[]string{"a", "c"},
	//	[]string{"q1", "q4"},
	//	"q1",
	//	[][]string{{"q1", "null", "null", "q2", "$"}, {"q2", "a", "null", "q2", "a"}, {"q2", "b", "null", "q3", "null"},
	//		{"q3", "c", "a", "q3", "null"}, {"q3", "null", "$", "q4", "null"}},
	//	"$"}
	//
	//PDA3 := PDAConf{"testPDA3",
	//	[]string{"q1", "q2", "q3", "q4", "q5", "q6"},
	//	[]string{"0", "1", "2"},
	//	[]string{"0"},
	//	[]string{"q1", "q6"},
	//	"q1", [][]string{
	//		{"q1", "null", "null", "q2", "$"},
	//		{"q2", "0", "null", "q3", "null"},
	//		{"q3", "1", "null", "q2", "1"},
	//		{"q2", "2", "null", "q4", "null"},
	//		{"q4", "0", "null", "q5", "null"},
	//		{"q4", "2", "null", "q4", "null"},
	//		{"q5", "1", "1", "q4", "null"},
	//		{"q4", "null", "$", "q6", "null"}}, "$"}
	//
	//PDA4 := PDAConf{"HelloPDA-EagerTest",
	//	[]string{"q1", "q2", "q3", "q4", "q5"},
	//	[]string{"0", "1"},
	//	[]string{"0", "1"},
	//	[]string{"q1", "q5"},
	//	"q1",
	//	[][]string{{"q1", "null", "null", "q2", "$"},
	//		{"q2", "null", "null", "q3", "null"},
	//		{"q3", "0", "null", "q3", "0"},
	//		{"q3", "1", "0", "q4", "null"},
	//		{"q4", "1", "0", "q4", "null"},
	//		{"q4", "null", "$", "q5", "null"}}, "$"}
	//
	//js1, err := json.Marshal(PDA1)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//js2, err := json.Marshal(PDA2)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//js3, err := json.Marshal(PDA3)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//js4, err := json.Marshal(PDA4)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	////w.Header().Set("Content-Type", "application/json")
	//fmt.Fprintf(w, "   ID:1   ")
	//w.Write(js1)
	//fmt.Fprintf(w, "\n\n")
	//fmt.Fprintf(w, "   ID:2   ")
	//w.Write(js2)
	//fmt.Fprintf(w, "\n\n")
	//fmt.Fprintf(w, "   ID:3   ")
	//w.Write(js3)
	//fmt.Fprintf(w, "\n\n")
	//fmt.Fprintf(w, "   ID:4   ")
	//w.Write(js4)
}*/
