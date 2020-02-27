package main

import "encoding/json"

func main() {
	print("Hello")
}

func open(in []byte) bool {
	var d PDAConf
	err := json.Unmarshal(in, &d)
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
