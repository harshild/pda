package test

import "testing"
import . "../src"

func TestOpen(t *testing.T) {
	t.Run("Negative", func(t *testing.T) {
		pda := PdaController{}
		str := "{a = a}"
		isParsingSuccess := pda.Open([]byte(str))
		if isParsingSuccess != false {
			t.Errorf("output for %s is \n %t; want false", str, isParsingSuccess)
		}
	})

	t.Run("Positive", func(t *testing.T) {
		pda := PdaController{}
		str := `{"name":"hello"}`
		isParsingSuccess := pda.Open([]byte(str))
		if isParsingSuccess != true {
			t.Errorf("output for %s is \n %t; want true", str, isParsingSuccess)
		}

		if pda.PdaConf.Name != "hello" {
			t.Errorf("Parsing went wrong, start state %s is parsed", pda.PdaConf.StartState)
		}
	})

	/*	t.Run("Negative - missing fields", func(t *testing.T) {
		pda := PdaController{}
			str := `{"name": "HelloPDA", "states": ["q1", "q2", "q3", "q4"], ` +
				`"input_alphabet": ["0", "1"], ` +
				`“stack_alphabet” : [“0”, “1”], ` +
				`"accepting_states": ["q1", "q4"], ` +
				`"start_state": "q1", ` +
				`"transitions": [ ` +
				`["q1", null, null, "q2", "$"], ` +
				`["q2", "0", null, "q2", "0"], ` +
				`["q2", "1", "0", "q3", null], ` +
				`["q3", "1", "0", "q3", null], ` +
				`["q3", null, "$", "q4", null]]}`
			isParsingSuccess := open([]byte(str))
			if isParsingSuccess != false {
				t.Errorf("output for %s is \n %t; want false", str, isParsingSuccess)
			}
		})*/

	/*	t.Run("Negative - wrong fields", func(t *testing.T) {
		pda := PdaController{}
		str := `{"name1": "HelloPDA"}`
		json := pda.Open([]byte(str))
		if json != false {
			t.Errorf("output for %s is \n %t; want false", str, json)
		}
	})*/
}

func TestReset(t *testing.T) {
	t.Run("should reset the pda stack to empty", func(t *testing.T) {
		pda := PdaController{}
		pda.Stack.Push("a")
		if pda.Stack.IsEmpty() {
			t.Errorf("initial stack is empty")
		}

		pda.Reset()

		if !pda.Stack.IsEmpty() {
			t.Errorf("stack is not reset")
		}
	})
}

func TestIsAccepted(t *testing.T) {
	t.Run("return True if PdaController is currently at an accepting state with empty stack", func(t *testing.T) {
		pda := PdaController{}
		pda.PdaConf.AcceptingStates = append(pda.PdaConf.AcceptingStates, "q1", "q2")

		pda.State = "q1"
		accepted := pda.Is_accepted()

		if !accepted {
			t.Errorf("expecting the state to be accepting and stack to be empty but failed")
		}
	})
}

func TestCurrentState(t *testing.T) {
	t.Run("check current pda state", func(t *testing.T) {
		pda := PdaController{}
		state := "q1"
		pda.State = state
		got := pda.Current_state()

		if got != state {
			t.Errorf("expecting the state to be q1 got %s", got)
		}
	})
}
