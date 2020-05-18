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
	Message string
	ForPDA  *PdaProcessor
}

func (e *PDARuntimeError) Error() string {
	pdaInfo := "No PDA Info available for \t"
	if e.ForPDA != nil && e.ForPDA.PdaConf.Name != "" {
		pdaInfo = "Error occurred for PDA=" + e.ForPDA.PdaConf.Name + " At Clock=" + strconv.Itoa(e.ForPDA.Clock) + " At state=" + e.ForPDA.State
	}
	return e.Message + " || " + pdaInfo
}

func (pdaProcessor *PdaProcessor) Puts(position int, token string) error {
	if pdaProcessor.Is_accepted() {
		return &PDARuntimeError{Message: "You have to reset the PDA first. Token DISCARDED"}
	}

	pdaProcessor.InputQueue[position] = token

	if pdaProcessor.LastConsumedIndex == position-1 {
		pdaProcessor.checkQueue()
	}
	return nil
}

func (pdaProcessor *PdaProcessor) checkQueue() {
	if _, ok := pdaProcessor.InputQueue[(pdaProcessor.LastConsumedIndex + 1)]; ok {
		pdaProcessor.Put(pdaProcessor.InputQueue[(pdaProcessor.LastConsumedIndex + 1)])
		pdaProcessor.LastConsumedIndex++
		pdaProcessor.checkQueue()
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

func (pdaProcessor *PdaProcessor) Reset() error {
	pdaProcessor.State = pdaProcessor.PdaConf.StartState
	pdaProcessor.Stack = utility.Stack{}
	pdaProcessor.Clock = 0
	pdaProcessor.InputQueue = make(map[int]string, 0)
	pdaProcessor.LastConsumedIndex = -1

	_, err := pdaProcessor.Put(" ")
	return err
}

func (pdaProcessor *PdaProcessor) Put(token string) (int, error) {
	transitionCount := 0
	if utility.StringArrContains(pdaProcessor.PdaConf.InputAlphabet, token) || token == " " || token == pdaProcessor.PdaConf.Eos {
		transition := GetTransition(pdaProcessor.State, pdaProcessor.PdaConf.Transitions, token)

		if transition.currentAlphabet != "" && transition.currentState != "" && transition.nextState != "" {
			if transition.elementToBePopped != "" {
				if !pdaProcessor.Stack.IsEmpty() && pdaProcessor.Stack.TopElement() != transition.elementToBePopped {
					return 0, &PDARuntimeError{Message: "Element to be popped from Stack not found on top", ForPDA: pdaProcessor}
				}
				pdaProcessor.Stack.Pop()
			}

			fmt.Printf("  %s => %s  ", pdaProcessor.State, transition.nextState)

			pdaProcessor.State = transition.nextState

			if transition.elementToBePushed != "" {
				pdaProcessor.Stack.Push(transition.elementToBePushed)
			}

			pdaProcessor.Clock++

			transitionCount = 1 + pdaProcessor.takeEagerSteps()
		}

		if transitionCount < 1 {
			return 0, &PDARuntimeError{Message: "No transition found in configuration for STATE=" + pdaProcessor.Current_state(), ForPDA: pdaProcessor}
		}

		fmt.Printf("PDA Name=%s \tToken=%s \t Transitions Took=%d\tClock Ticks=%d \n", pdaProcessor.GetPDAName(), token, transitionCount, pdaProcessor.GetClock())
		return transitionCount, nil
	} else {
		return 0, &PDARuntimeError{Message: "Put method called: No transition  Invalid input sequence provided Input:" + token, ForPDA: pdaProcessor}
	}
}

func (pdaProcessor *PdaProcessor) Eos() {
	pdaProcessor.Put(" ")
}

func (pdaProcessor *PdaProcessor) Is_accepted() bool {
	return pdaProcessor.Stack.IsEmpty() && utility.StringArrContains(pdaProcessor.PdaConf.AcceptingStates, pdaProcessor.State)
}

func (pdaProcessor *PdaProcessor) Peek(len int) ([]string, error) {
	return pdaProcessor.Stack.Peek(len)
}

func (pdaProcessor *PdaProcessor) Current_state() string {
	return pdaProcessor.State
}

func (pdaProcessor *PdaProcessor) Close() {
	pdaProcessor.Stack = utility.Stack{}
	pdaProcessor.State = ""
	pdaProcessor.Clock = 0
	pdaProcessor.State = pdaProcessor.PdaConf.StartState
}

func (pdaProcessor *PdaProcessor) GetPDAName() string {
	return pdaProcessor.PdaConf.Name
}

func (pdaProcessor *PdaProcessor) GetClock() int {
	return pdaProcessor.Clock
}

func (pdaProcessor *PdaProcessor) takeEagerSteps() int {
	transition := GetEagerTransition(pdaProcessor.State, pdaProcessor.PdaConf.Transitions, pdaProcessor.Stack, pdaProcessor.PdaConf.Eos)

	if transition.currentState != "" && transition.nextState != "" {

		if transition.elementToBePopped != "" {
			if !pdaProcessor.Stack.IsEmpty() && pdaProcessor.Stack.TopElement() != transition.elementToBePopped {
				utility.Crash(&PDARuntimeError{Message: "Element to be popped from Stack not found on top", ForPDA: pdaProcessor})
			}
			pdaProcessor.Stack.Pop()
		}
		fmt.Print("\n")
		fmt.Printf("  %s => %s  ", pdaProcessor.State, transition.nextState)

		pdaProcessor.State = transition.nextState

		if transition.elementToBePushed != "" {
			pdaProcessor.Stack.Push(transition.elementToBePushed)
		}

		pdaProcessor.Clock++

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

func (pdaProcessor *PdaProcessor) UpdateStatus(status entity.PDAStatus) {
	pdaProcessor.LastConsumedIndex = status.LastConsumedIndex
	pdaProcessor.State = status.State
	pdaProcessor.Stack = status.Stack
	pdaProcessor.InputQueue = status.InputQueue
	pdaProcessor.Clock = status.Clock
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
	Clock             int
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
