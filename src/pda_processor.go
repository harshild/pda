package src

import (
	"encoding/json"
	"strings"
)

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
	numberOfTransitions := 0
	tokenToBeProcessed := " " + token + " "
	print("Start State ", pdaProcessor.State)
	for _, alphabet := range tokenToBeProcessed {
		transition := GetTransition(pdaProcessor.State, pdaProcessor.PdaConf.Transitions, string(alphabet))

		if transition.ElementToBePopped != "" {
			if pdaProcessor.Stack.Pop() != transition.ElementToBePopped {
				return -1
			}
		}

		if transition.NextState != "" {
			pdaProcessor.State = transition.NextState
			print("  =>  ", pdaProcessor.State)
		}

		if transition.ElementToBePushed != "" {
			pdaProcessor.Stack.Push(transition.ElementToBePushed)
		}

		numberOfTransitions++

	}

	return numberOfTransitions
}

func GetTransition(currentState string, allTransitions [][]string, alphabet string) PDATransition {
	for _, transitions := range allTransitions {
		if transitions[0] == currentState && transitions[1] == strings.TrimSpace(alphabet) {
			return PDATransition{
				CurrentState:      currentState,
				CurrentAlphabet:   alphabet,
				ElementToBePopped: transitions[2],
				NextState:         transitions[3],
				ElementToBePushed: transitions[4],
			}
		}
	}
	return PDATransition{}
}

type PdaProcessor struct {
	Stack   Stack
	PdaConf PDAConf
	State   string
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
