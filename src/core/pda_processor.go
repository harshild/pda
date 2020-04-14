package core

import (
	"encoding/json"
	"entity"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"utility"
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

func (pdaProcessor *PdaProcessor) Puts(position int, token string) {

	pdaProcessor.InputQueue[position] = token

	if pdaProcessor.LastConsumedIndex == position-1 {
		checkQueue(*pdaProcessor)
	}
}

func checkQueue(pdaProcessor PdaProcessor) {
	if _, ok := pdaProcessor.InputQueue[(pdaProcessor.LastConsumedIndex + 1)]; ok {
		pdaProcessor.Put(pdaProcessor.InputQueue[(pdaProcessor.LastConsumedIndex + 1)])
		checkQueue(pdaProcessor)
	}
}

func (pdaProcessor *PdaProcessor) Open(in []byte) bool {
	err := json.Unmarshal(in, &(pdaProcessor.PdaConf))
	if err != nil {
		return false
	}
	pdaProcessor.Stack = utility.Stack{}
	pdaProcessor.State = pdaProcessor.PdaConf.StartState
	pdaProcessor.LastConsumedIndex = -1
	pdaProcessor.InputQueue = make(map[int]string, 0)
	return true
}

func (pdaProcessor *PdaProcessor) Reset() {
	pdaProcessor.State = pdaProcessor.PdaConf.StartState
	pdaProcessor.Stack = utility.Stack{}
	pdaProcessor.Put(" ")
}

func (pdaProcessor *PdaProcessor) Put(token string) int {
	transitionCount := 0
	if utility.StringArrContains(pdaProcessor.PdaConf.InputAlphabet, token) || token == " " || token == pdaProcessor.PdaConf.Eos {
		transition := GetTransition(pdaProcessor.State, pdaProcessor.PdaConf.Transitions, token)

		if transition.currentAlphabet != "" && transition.currentState != "" && transition.nextState != "" {
			if transition.elementToBePopped != "" {
				if !pdaProcessor.Stack.IsEmpty() && pdaProcessor.Stack.TopElement() != transition.elementToBePopped {
					utility.Crash(&PDARuntimeError{message: "Element to be popped from Stack not found on top", forPDA: pdaProcessor})
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
			utility.Crash(&PDARuntimeError{message: "No transition found in configuration for STATE=" + pdaProcessor.Current_state(), forPDA: pdaProcessor})
		}
		pdaProcessor.LastConsumedIndex++
		return transitionCount
	} else {
		utility.Crash(&PDARuntimeError{message: "Invalid input sequence provided", forPDA: pdaProcessor})
		return 0
	}
}

func (pdaProcessor *PdaProcessor) Eos() {
	pdaProcessor.Put(" ")
}

func (pdaProcessor *PdaProcessor) Is_accepted() bool {
	return pdaProcessor.Stack.IsEmpty() && utility.StringArrContains(pdaProcessor.PdaConf.AcceptingStates, pdaProcessor.State)
}

func (pdaProcessor *PdaProcessor) Peek(len int) []string {
	return pdaProcessor.Stack.Peek(len)
}

func (pdaProcessor *PdaProcessor) Current_state() string {
	return pdaProcessor.State
}

func (pdaProcessor *PdaProcessor) Close() {
	pdaProcessor.PdaConf = entity.PDAConf{}
	pdaProcessor.Stack = utility.Stack{}
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
				utility.Crash(&PDARuntimeError{message: "Element to be popped from Stack not found on top", forPDA: pdaProcessor})
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

func (pdaProcessor *PdaProcessor) Queued_tokens() []string {
	keys := make([]int, 0, len(pdaProcessor.InputQueue))
	for k := range pdaProcessor.InputQueue {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	queuedTokens := make([]string, 0)

	for _, k := range keys {
		if k > pdaProcessor.LastConsumedIndex {
			queuedTokens = append(queuedTokens, pdaProcessor.InputQueue[k])
		}
	}

	return queuedTokens
}

func GetAllPDANames(pdas []PdaProcessor) []string {
	pdaNames := make([]string, 0)
	for _, pda := range pdas {
		pdaNames = append(pdaNames, pda.PdaConf.Name)
	}
	return pdaNames
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

func GetEagerTransition(currentState string, allTransitions [][]string, stack utility.Stack, eos string) PDATransition {
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
	Stack             utility.Stack
	PdaConf           entity.PDAConf
	State             string
	clock             int
	InputQueue        map[int]string
	LastConsumedIndex int
}

type PDATransition struct {
	currentState      string
	currentAlphabet   string
	elementToBePopped string
	nextState         string
	elementToBePushed string
}
