package controller

import (
	"net/http"
)

type PdaController struct {
}

func (pdaController *PdaController) Gotopdas(w http.ResponseWriter, r *http.Request) {

}

func (pdaController *PdaController) Closepda(w http.ResponseWriter, r *http.Request) {

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
}
