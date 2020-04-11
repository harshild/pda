package main

import (
	. "../src"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	pdaProcessor := PdaProcessor{}

	if len(os.Args) == 3 {
		if os.Args[1] == "" {
			println("Give PDA specifications ")
		} else {

			transition := 1
			specJSON, _ := ioutil.ReadFile(os.Args[1])
			if pdaProcessor.Open(specJSON) {
				if os.Args[2] != "" {
					inputString := ""
					inputBytes, err := ioutil.ReadFile(os.Args[2])
					if err != nil {
						inputString = os.Args[2]
					} else {
						inputString = string(inputBytes)
					}
					fmt.Printf("PDA Name=%s \tMethod=Is_Accepted =%t \n", pdaProcessor.GetPDAName(), pdaProcessor.Is_accepted())

					pdaProcessor.Reset()
					fmt.Printf("PDA Name=%s \tToken=START \t Transitions Took=%d\tClock Ticks=%d \n", pdaProcessor.GetPDAName(), pdaProcessor.GetClock(), pdaProcessor.GetClock())

					for _, alphabet := range inputString {
						transition = pdaProcessor.Put(string(alphabet))
						fmt.Printf("PDA Name=%s \tToken=%s \t Transitions Took=%d\tClock Ticks=%d \n", pdaProcessor.GetPDAName(), string(alphabet), transition, pdaProcessor.GetClock())
					}
					prevClock := pdaProcessor.GetClock()
					pdaProcessor.Eos()
					fmt.Printf("PDA Name=%s \tToken=EOS \t Transitions Took=%d\tClock Ticks=%d \n", pdaProcessor.GetPDAName(), pdaProcessor.GetClock()-prevClock, pdaProcessor.GetClock())

					fmt.Print("For PDA=" + pdaProcessor.GetPDAName() + "\t Current State=" + pdaProcessor.Current_state() + "\n")
					pdaProcessor.Peek(5)

					pdaProcessor.Close()

				} else {
					println("PDA no input stream specified")
				}

			} else {
				println("PDA open (spec) API failed!!")
			}
		}
	} else {
		println("Wrong number of inputs provided, Input length", len(os.Args))
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", gotopdas)
	router.HandleFunc("/pdas", pdalist)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func gotopdas(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go to PDA list: http://localhost:8080/pdas")
}

func pdalist(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Select one PDA and Enter the ID of it\n\n ")

	//TODO: Here "null" is going to create a problem, We are yet to handle it

	PDA1 := PDAConf{"HelloPDA", []string{"q1", "q2", "q3", "q4"}, []string{"0", "1"}, []string{"0", "1"},
		[]string{"q1", "q4"}, "q1", [][]string{{"q1", "null", "null", "q2", "$"}, {"q2", "0", "null", "q2", "0"},
			{"q2", "1", "0", "q3", "null"}, {"q3", "1", "0", "q3", "null"}, {"q3", "null", "$", "q4", "null"}}, "$"}

	PDA2 := PDAConf{
		"testPDA2",
		[]string{"q1", "q2", "q3", "q4"},
		[]string{"a", "b", "c"},
		[]string{"a", "c"},
		[]string{"q1", "q4"},
		"q1",
		[][]string{{"q1", "null", "null", "q2", "$"}, {"q2", "a", "null", "q2", "a"}, {"q2", "b", "null", "q3", "null"},
			{"q3", "c", "a", "q3", "null"}, {"q3", "null", "$", "q4", "null"}},
		"$"}

	PDA3 := PDAConf{"testPDA3",
		[]string{"q1", "q2", "q3", "q4", "q5", "q6"},
		[]string{"0", "1", "2"},
		[]string{"0"},
		[]string{"q1", "q6"},
		"q1", [][]string{
			{"q1", "null", "null", "q2", "$"},
			{"q2", "0", "null", "q3", "null"},
			{"q3", "1", "null", "q2", "1"},
			{"q2", "2", "null", "q4", "null"},
			{"q4", "0", "null", "q5", "null"},
			{"q4", "2", "null", "q4", "null"},
			{"q5", "1", "1", "q4", "null"},
			{"q4", "null", "$", "q6", "null"}}, "$"}

	PDA4 := PDAConf{"HelloPDA-EagerTest",
		[]string{"q1", "q2", "q3", "q4", "q5"},
		[]string{"0", "1"},
		[]string{"0", "1"},
		[]string{"q1", "q5"},
		"q1",
		[][]string{{"q1", "null", "null", "q2", "$"},
			{"q2", "null", "null", "q3", "null"},
			{"q3", "0", "null", "q3", "0"},
			{"q3", "1", "0", "q4", "null"},
			{"q4", "1", "0", "q4", "null"},
			{"q4", "null", "$", "q5", "null"}}, "$"}

	js1, err := json.Marshal(PDA1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js2, err := json.Marshal(PDA2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js3, err := json.Marshal(PDA3)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js4, err := json.Marshal(PDA4)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "   ID:1   ")
	w.Write(js1)
	fmt.Fprintf(w, "\n\n")
	fmt.Fprintf(w, "   ID:2   ")
	w.Write(js2)
	fmt.Fprintf(w, "\n\n")
	fmt.Fprintf(w, "   ID:3   ")
	w.Write(js3)
	fmt.Fprintf(w, "\n\n")
	fmt.Fprintf(w, "   ID:4   ")
	w.Write(js4)
}
