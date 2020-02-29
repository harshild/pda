package main

import "encoding/json"

var pda PDA

func open(in []byte) bool {
	err := json.Unmarshal(in, &pda.pdaConf)
	if err == nil {
		return false
	}
	return true
}

func reset() {
	pda.stack = Stack{}
}

func is_accepted() bool {
	return pda.stack.isEmpty() && stringArrContains(pda.pdaConf.acceptingStates, pda.state)
}

func stringArrContains(str_arr []string, lookup_item string) bool {
	for _, elem := range str_arr {
		if elem == lookup_item {
			return true
		}
	}
	return false

}

func current_state() string {
	return pda.state
}

type PDA struct {
	stack   Stack
	pdaConf PDAConf
	state   string
}

type PDAConf struct {
	name            string
	states          []string
	inputAlphabet   []string
	stackAlphabet   []string
	acceptingStates []string
	startState      string
	transitions     [][]string
	eos             string
}
