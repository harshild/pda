package src

import "encoding/json"

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

func (pda *PdaController) close() {

}

type PdaController struct {
	Stack   Stack
	PdaConf PDAConf
	State   string
}

type PDATransitions struct {
	CurrentState      string `json:"currentState"`
	CurrentAlphabet   string `json:"currentAlphabet"`
	ElementToBePopped string `json:"elementToBePopped"`
	NextState         string `json:"nextState"`
	ElementToBePushed string `json:"elementToBePushed"`
}

type PDAConf struct {
	Name            string   `json:"name"`
	States          []string `json:"states"`
	InputAlphabet   []string `json:"input_alphabet"`
	StackAlphabet   []string `json:"stack_alphabet"`
	AcceptingStates []string `json:"accepting_states"`
	StartState      string   `json:"start_state"`
	Transitions     []string `json:"transitions"`
	Eos             string   `json:"eos"`
}
