package src

import (
	"encoding/json"
)

func (pda *PdaController) Open(in []byte) bool {
	err := json.Unmarshal(in, &(pda.PdaConf))
	if err != nil {
		return false
	}
	return true
}

func (pda *PdaController) Reset() {
	pda.Stack = Stack{}
}

func (pda *PdaController) Is_accepted() bool {
	return pda.Stack.IsEmpty() && StringArrContains(pda.PdaConf.AcceptingStates, pda.State)
}

func (pda *PdaController) Current_state() string {
	return pda.State
}

func (pda *PdaController) Peek(k int) []string {
	return pda.Stack.Peek(k)
}

func (pda *PdaController) Close() {
	//TODO: garbage-collect/return any (re-usable) resources used by the PDA.
}

func (pda *PdaController) Put(token string) int {
	numberOfTransitions := 0
	tokenToBeProcessed := " " + token + pda.PdaConf.Eos
	for _, alphabet := range tokenToBeProcessed {
		transition := GetTransition(pda.State, pda.PdaConf.Transitions, string(alphabet))

		if transition.ElementToBePopped != "" {
			if pda.Stack.Pop() != transition.ElementToBePopped {
				return -1
			}
		}

		if transition.NextState != "" {
			pda.State = transition.NextState
		}

		if transition.ElementToBePushed != "" {
			pda.Stack.Push(transition.ElementToBePushed)
		}

		numberOfTransitions++

	}
	return numberOfTransitions
}

func GetTransition(currentState string, allTransitions [][]string, alphabet string) PDATransition {
	for _, transitions := range allTransitions {
		if transitions[0] == currentState && transitions[1] == alphabet {
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

type PdaController struct {
	Stack   Stack
	PdaConf PDAConf
	State   string
}

type PDATransition struct {
	CurrentState      string `json:"currentState"`
	CurrentAlphabet   string `json:"currentAlphabet"`
	ElementToBePopped string `json:"elementToBePopped"`
	NextState         string `json:"nextState"`
	ElementToBePushed string `json:"elementToBePushed"`
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
