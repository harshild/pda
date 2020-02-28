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

type PDA struct {
	stack   Stack
	pdaConf PDAConf
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
