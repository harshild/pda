package main

import "encoding/json"

var pdaConf PDAConf

func open(in []byte) bool {
	err := json.Unmarshal(in, &pdaConf)
	if err == nil {
		return false
	}
	return true
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
