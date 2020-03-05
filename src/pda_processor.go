package src

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PDARuntimeError struct {
	message string
	forPDA  *PdaProcessor
}

func (e *PDARuntimeError) Error() string {
	pdaInfo := ""
	if e.forPDA != nil && e.forPDA.PdaConf.Name != "" {
		pdaInfo = "Error occurred for PDA=" + e.forPDA.PdaConf.Name + " At clock=" + string(e.forPDA.clock) + " At state=" + e.forPDA.State
	}
	return e.message + "\n" + pdaInfo
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

func (pdaProcessor *PdaProcessor) Put(token string) int {

	transitions :=0
	if StringArrContains(pdaProcessor.PdaConf.InputAlphabet, token) || token == " " || token == pdaProcessor.PdaConf.Eos {
		transition:= GetTransition(pdaProcessor.State ,pdaProcessor.PdaConf.Transitions, token)

			if transition.elementToBePopped != "" {
				if !pdaProcessor.Stack.IsEmpty() && pdaProcessor.Stack.TopElement() != transition.elementToBePopped {
					Crash(&PDARuntimeError{message: "Element to be popped from Stack not found on top", forPDA: pdaProcessor})
				}
				pdaProcessor.Stack.Pop()
			}

			fmt.Printf("  %s => %s  ", pdaProcessor.State, transition.nextState)

			pdaProcessor.State = transition.nextState

			if transition.elementToBePushed != "" {
				pdaProcessor.Stack.Push(transition.elementToBePushed)
			}

			pdaProcessor.clock++

			transitions = 1 + pdaProcessor.takeEagerSteps()
		}
		if transitions < 1 {
			Crash(&PDARuntimeError{message: "No transition found in configuration for STATE="+pdaProcessor.Current_state()+" Token="+token})
		}
		return transitions
	} else {
		Crash(&PDARuntimeError{message: "Invalid input sequence provided", forPDA: pdaProcessor})
		return 0
	}
}

func (pdaProcessor *PdaProcessor) Eos() {
	pdaProcessor.Put(" ")
}

func (pdaProcessor *PdaProcessor) Is_accepted() bool {
	return pdaProcessor.Stack.IsEmpty() && StringArrContains(pdaProcessor.PdaConf.AcceptingStates, pdaProcessor.State)
}

func (pdaProcessor *PdaProcessor) Peek(len int) []string {
	return pdaProcessor.Stack.Peek(len)
}

func (pdaProcessor *PdaProcessor) Current_state() string {
	return pdaProcessor.State
}

func (pdaProcessor *PdaProcessor) Close() {
	pdaProcessor.PdaConf = PDAConf{}
	pdaProcessor.Stack = Stack{}
	pdaProcessor.State = ""
	pdaProcessor.clock = 0

}

func (pdaProcessor *PdaProcessor) GetPDAName() string {
	return pdaProcessor.PdaConf.Name
}

func (pdaProcessor *PdaProcessor) GetClock() int {
	return pdaProcessor.clock
}

func (pdaProcessor *PdaProcessor) takeEagerSteps() {

}

func GetTransition(currentState string, allTransitions [][]string,alphabet string) PDATransition {
	for _, transitions := range allTransitions {
		if transitions[0] == currentState && transitions[1] == strings.TrimSpace(alphabet) {
			return PDATransition{
				currentState:      currentState,
				currentAlphabet:   alphabet,
				elementToBePopped: transitions[2],
				nextState:         transitions[3],
				elementToBePushed: transitions[4],
			}
		}
	}
	return PDATransition{}
}

type PdaProcessor struct {
	Stack   Stack
	PdaConf PDAConf
	State   string
	clock   int
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
