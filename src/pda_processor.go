package src

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PDARuntimeError struct {
	message string
}

func (e *PDARuntimeError) Error() string {
	return e.message
}

func (pdaProcessor *PdaProcessor) Open(in []byte) bool {
	err := json.Unmarshal(in, &(pdaProcessor.PdaConf))
	pdaProcessor.State = pdaProcessor.PdaConf.StartState
	if err != nil {
		return false
	}
	return true
}

func (pdaProcessor *PdaProcessor) Reset() {
	pdaProcessor.Stack = Stack{}
	pdaProcessor.Put(" ")
}

func (pdaProcessor *PdaProcessor) Is_accepted() bool {
	return pdaProcessor.Stack.IsEmpty() && StringArrContains(pdaProcessor.PdaConf.AcceptingStates, pdaProcessor.State)
}

func (pdaProcessor *PdaProcessor) Current_state() string {
	return pdaProcessor.State
}

func (pdaProcessor *PdaProcessor) Peek(k int) []string {
	return pdaProcessor.Stack.Peek(k)
}

func (pdaProcessor *PdaProcessor) Close() {
	//TODO: garbage-collect/return any (re-usable) resources used by the PDA.
}

func (pdaProcessor *PdaProcessor) Put(token string) int {
	//numberOfTransitions := 0
	//tokenToBeProcessed := " " + token + " "
	//print("Start State ", pdaProcessor.State)
	//for _, alphabet := range tokenToBeProcessed {
	//if len(token) != 1 {
	//	return nilIntValue,&PDARuntimeError{"Invalid token length"}
	//}

	if StringArrContains(pdaProcessor.PdaConf.InputAlphabet, token) || token == " " || token == pdaProcessor.PdaConf.Eos {
		transition, err := GetTransition(pdaProcessor.State, pdaProcessor.PdaConf.Transitions, token)

		if err != nil {
			Crash(err)
		}

		if transition.ElementToBePopped != "" {
			if !pdaProcessor.Stack.IsEmpty() && pdaProcessor.Stack.TopElement() != transition.ElementToBePopped {
				Crash(&PDARuntimeError{"Element to be popped from Stack not found on top"})
			}
			pdaProcessor.Stack.Pop()
		}

		pdaProcessor.State = transition.NextState

		if transition.ElementToBePushed != "" {
			pdaProcessor.Stack.Push(transition.ElementToBePushed)
		}
		fmt.Printf("  %s =>  ", pdaProcessor.State)

		pdaProcessor.Transitions++
		return pdaProcessor.Transitions
	} else {
		Crash(&PDARuntimeError{"Invalid input sequence provided"})
		return -1
	}
}

func (pdaProcessor *PdaProcessor) Eos() {
	pdaProcessor.Put(" ")
}

func GetTransition(currentState string, allTransitions [][]string, alphabet string) (PDATransition, error) {
	for _, transitions := range allTransitions {
		if transitions[0] == currentState && transitions[1] == strings.TrimSpace(alphabet) {
			return PDATransition{
				CurrentState:      currentState,
				CurrentAlphabet:   alphabet,
				ElementToBePopped: transitions[2],
				NextState:         transitions[3],
				ElementToBePushed: transitions[4],
			}, nil
		}
	}
	return PDATransition{}, &PDARuntimeError{"No transaction found in configuration for current scenario"}
}

type PdaProcessor struct {
	Stack       Stack
	PdaConf     PDAConf
	State       string
	Transitions int
}

type PDATransition struct {
	CurrentState      string
	CurrentAlphabet   string
	ElementToBePopped string
	NextState         string
	ElementToBePushed string
}

type PDAConf struct {
	Name            string     `json:"name"`
	States          []string   `json:"states"`
	InputAlphabet   []string   `json:"input_alphabet"`
	StackAlphabet   []string   `json:"stack_alphabet"`
	AcceptingStates []string   `json:"accepting_states"`
	StartState      string     `json:"start_state"`
	Transitions     [][]string `json:"transitions"`
	Eos             string     `json:"eos"`
}
