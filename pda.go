package main

import "encoding/json"

func (pda *PDA) open(in []byte) bool {
	err := json.Unmarshal(in, &pda.pdaConf)
	if err == nil {
		return false
	}
	return true
}

func (pda *PDA) reset() {
	pda.stack = Stack{}
}

func (pda *PDA) is_accepted() bool {
	return pda.stack.isEmpty() && stringArrContains(pda.pdaConf.acceptingStates, pda.state)
}

func (pda *PDA) current_state() string {
	return pda.state
}

func (pda *PDA) close() {

}

type PDA struct {
	stack   Stack
	pdaConf PDAConf
	state   string
}

type PDATransitions struct {
	currentState      string
	currentAlphabet   string
	elementToBePopped string
	nextState         string
	elementToBePushed string
}

type PDAConf struct {
	name            string
	states          []string
	inputAlphabet   []string
	stackAlphabet   []string
	acceptingStates []string
	startState      string
	transitions     []PDATransitions
	eos             string
}
