package src

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type PDARuntimeError struct {
	message string
	forPDA  *PdaProcessor
}

func (e *PDARuntimeError) Error() string {
	pdaInfo := "No PDA Info available for \t"
	if e.forPDA != nil && e.forPDA.PdaConf.Name != "" {
		pdaInfo = "Error occurred for PDA=" + e.forPDA.PdaConf.Name + " At clock=" + strconv.Itoa(e.forPDA.clock) + " At state=" + e.forPDA.State
	}
	return e.message + " || " + pdaInfo
}

func (pdaProcessor *PdaProcessor) Open(in []byte) bool {
	err := json.Unmarshal(in, &(pdaProcessor.PdaConf))
	if err != nil {
		return false
	}
	pdaProcessor.Stack = Stack{}
	pdaProcessor.State = pdaProcessor.PdaConf.StartState
	return true
}

func (pdaProcessor *PdaProcessor) Reset() {
	pdaProcessor.State = pdaProcessor.PdaConf.StartState
	pdaProcessor.Stack = Stack{}
	pdaProcessor.Put(" ")
}

func (pdaProcessor *PdaProcessor) Put(token string) int {
	transitionCount := 0
	if StringArrContains(pdaProcessor.PdaConf.InputAlphabet, token) || token == " " || token == pdaProcessor.PdaConf.Eos {
		transition := GetTransition(pdaProcessor.State, pdaProcessor.PdaConf.Transitions, token)

		if transition.currentAlphabet != "" && transition.currentState != "" && transition.nextState != "" {
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

			transitionCount = 1 + pdaProcessor.takeEagerSteps()
		}

		if transitionCount < 1 {
			Crash(&PDARuntimeError{message: "No transition found in configuration for STATE=" + pdaProcessor.Current_state(), forPDA: pdaProcessor})
		}
		return transitionCount
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
	pdaProcessor.State = pdaProcessor.PdaConf.StartState
}

func (pdaProcessor *PdaProcessor) GetPDAName() string {
	return pdaProcessor.PdaConf.Name
}

func (pdaProcessor *PdaProcessor) GetClock() int {
	return pdaProcessor.clock
}

func (pdaProcessor *PdaProcessor) takeEagerSteps() int {
	transition := GetEagerTransition(pdaProcessor.State, pdaProcessor.PdaConf.Transitions, pdaProcessor.Stack, pdaProcessor.PdaConf.Eos)

	if transition.currentState != "" && transition.nextState != "" {

		if transition.elementToBePopped != "" {
			if !pdaProcessor.Stack.IsEmpty() && pdaProcessor.Stack.TopElement() != transition.elementToBePopped {
				Crash(&PDARuntimeError{message: "Element to be popped from Stack not found on top", forPDA: pdaProcessor})
			}
			pdaProcessor.Stack.Pop()
		}
		fmt.Print("\n")
		fmt.Printf("  %s => %s  ", pdaProcessor.State, transition.nextState)

		pdaProcessor.State = transition.nextState

		if transition.elementToBePushed != "" {
			pdaProcessor.Stack.Push(transition.elementToBePushed)
		}

		pdaProcessor.clock++

		return 1 + pdaProcessor.takeEagerSteps()
	}
	return 0

}

func GetTransition(currentState string, allTransitions [][]string, alphabet string) PDATransition {
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

func GetEagerTransition(currentState string, allTransitions [][]string, stack Stack, eos string) PDATransition {
	for _, transitions := range allTransitions {
		if transitions[0] == currentState && transitions[1] == "" && (transitions[2] == "" || transitions[2] == stack.TopElement()) &&
			transitions[2] != eos && transitions[4] != eos { //whitespace is used to indicate null I/P
			return PDATransition{
				currentState:      currentState,
				currentAlphabet:   transitions[1],
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
