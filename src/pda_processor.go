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
	if err != nil {
		return false
	}
	return true
}

func (pdaProcessor *PdaProcessor) Reset() {
	pdaProcessor.State = pdaProcessor.PdaConf.StartState
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
	pdaProcessor.PdaConf = PDAConf{}
	pdaProcessor.Stack = Stack{}
	pdaProcessor.State = ""
	pdaProcessor.transitions = 0

}

func (pdaProcessor *PdaProcessor) Put(token string) int {

	if StringArrContains(pdaProcessor.PdaConf.InputAlphabet, token) || token == " " || token == pdaProcessor.PdaConf.Eos {
		transition, err := GetTransition(pdaProcessor.State, pdaProcessor.PdaConf.Transitions, token)

		if err != nil {
			Crash(err)
		}

		if transition.elementToBePopped != "" {
			if !pdaProcessor.Stack.IsEmpty() && pdaProcessor.Stack.TopElement() != transition.elementToBePopped {
				Crash(&PDARuntimeError{"Element to be popped from Stack not found on top"})
			}
			pdaProcessor.Stack.Pop()
		}

		fmt.Printf("  %s => %s  ", pdaProcessor.State, transition.nextState)

		pdaProcessor.State = transition.nextState

		if transition.elementToBePushed != "" {
			pdaProcessor.Stack.Push(transition.elementToBePushed)
		}

		pdaProcessor.transitions++
		return pdaProcessor.transitions
	} else {
		Crash(&PDARuntimeError{"Invalid input sequence provided"})
		return -1
	}
}

func (pdaProcessor *PdaProcessor) Eos() {
	pdaProcessor.Put(" ")
}

func (pdaProcessor *PdaProcessor) GetPDAName() string {
	return pdaProcessor.PdaConf.Name
}

func GetTransition(currentState string, allTransitions [][]string, alphabet string) (PDATransition, error) {
	for _, transitions := range allTransitions {
		if transitions[0] == currentState && transitions[1] == strings.TrimSpace(alphabet) {
			return PDATransition{
				currentState:      currentState,
				currentAlphabet:   alphabet,
				elementToBePopped: transitions[2],
				nextState:         transitions[3],
				elementToBePushed: transitions[4],
			}, nil
		}
	}
	return PDATransition{}, &PDARuntimeError{"No transition found in configuration for current scenario"}
}

type PdaProcessor struct {
	Stack       Stack
	PdaConf     PDAConf
	State       string
	transitions int
}

type PDATransition struct {
	currentState      string
	currentAlphabet   string
	elementToBePopped string
	nextState         string
	elementToBePushed string
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
