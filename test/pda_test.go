package test

import "testing"
import . "../src"

func Test_Open(t *testing.T) {
	t.Run("open should fail", func(t *testing.T) {
		pda := PdaProcessor{}
		str := "{a = a}"
		isParsingSuccess := pda.Open([]byte(str))
		if isParsingSuccess != false {
			t.Errorf("output for %s is \n %t; want false", str, isParsingSuccess)
		}
	})

	t.Run("open should work as expected", func(t *testing.T) {
		pda := PdaProcessor{}
		str := `{"name": "HelloPDA",
  				"states": ["q1", "q2", "q3", "q4"],
  				"input_alphabet": ["0", "1"],
  				"stack_alphabet" : ["0", "1"],
  				"accepting_states": ["q1", "q4"],
  				"start_state": "q1",
  				"transitions": [
  				  ["q1", null, null, "q2", "$"],
  				  ["q2", "0", null, "q2", "0"],
  				  ["q2", "1", "0", "q3", null],
  				  ["q3", "1", "0", "q3", null],
  				  ["q3", null, "$", "q4", null]],
  				"eos": "$"}`
		isParsingSuccess := pda.Open([]byte(str))
		if isParsingSuccess != true {
			t.Errorf("output for %s is \n %t; want true", str, isParsingSuccess)
		}

		if pda.GetPDAName() != "HelloPDA" {
			t.Errorf("Parsing went wrong, PDA name is different")
		}
	})

}

func Test_Reset(t *testing.T) {
	t.Run("should reset the pda ", func(t *testing.T) {
		pda := PdaProcessor{
			PdaConf: PDAConf{
				Name:            "Test PDA",
				States:          []string{"q1", "q2", "q3", "q4"},
				InputAlphabet:   []string{"0", "1"},
				StackAlphabet:   []string{"0", "1"},
				AcceptingStates: []string{"q1", "q4"},
				StartState:      "q1",
				Transitions:     [][]string{{"q1", "", "", "q2", ""}},
				Eos:             "$",
			},
			State: "q1",
		}
		pda.Stack.Push("a")
		if pda.Stack.IsEmpty() {
			t.Errorf("initial Stack is empty")
		}

		pda.Reset()

		if !pda.Stack.IsEmpty() {
			t.Errorf("stack is not reset")
		}
	})
}

func Test_Is_Accepted(t *testing.T) {
	t.Run("return True if PdaProcessor is currently at an accepting state with empty stack", func(t *testing.T) {
		pda := PdaProcessor{}
		pda.PdaConf.AcceptingStates = append(pda.PdaConf.AcceptingStates, "q1", "q2")

		pda.State = "q1"
		accepted := pda.Is_accepted()

		if !accepted {
			t.Errorf("expecting the state to be accepting and stack to be empty but failed")
		}
	})
}

func Test_Current_State(t *testing.T) {
	t.Run("check current pda state", func(t *testing.T) {
		pda := PdaProcessor{}
		state := "q1"
		pda.State = state
		got := pda.Current_state()

		if got != state {
			t.Errorf("expecting the state to be q1 got %s", got)
		}
	})
}

func Test_Put(t *testing.T) {
	t.Run("Put token should return transitions taken", func(t *testing.T) {
		pda := PdaProcessor{
			PdaConf: PDAConf{
				Name:            "Test PDA",
				States:          []string{"q1", "q2", "q3", "q4"},
				InputAlphabet:   []string{"0", "1"},
				StackAlphabet:   []string{"0", "1"},
				AcceptingStates: []string{"q1", "q4"},
				StartState:      "q1",
				Transitions: [][]string{{"q1", "", "", "q2", "$"},
					{"q2", "0", "", "q2", "0"},
					{"q2", "1", "0", "q3", ""},
					{"q3", "1", "0", "q3", ""},
					{"q3", "", "$", "q4", ""}},
				Eos: "$",
			},
			State: "q1",
		}
		transitionCount := pda.Put(" ")

		if transitionCount != 1 {
			t.Errorf("Expected transition count to be 1 got %d", transitionCount)
		}

		transitionCount = pda.Put("0")

		if transitionCount != 2 {
			t.Errorf("Expected transition count to be 1 got %d", transitionCount)
		}

		transitionCount = pda.Put("0")

		if transitionCount != 3 {
			t.Errorf("Expected transition count to be 2 got %d", transitionCount)
		}
	})
}

//func TestProcessAlphabet(t *testing.T) {
//	t.Run("Process alphabet should return transition required for current scenario", func(t *testing.T) {
//		pda := PdaProcessor{
//			pdaConf: PDAConf{
//				Name:            "Test PDA",
//				States:          []string{"q1", "q2", "q3", "q4"},
//				InputAlphabet:   []string{"0", "1"},
//				StackAlphabet:   []string{"0", "1"},
//				AcceptingStates: []string{"q1", "q4"},
//				StartState:      "q1",
//				transitions: [][]string{{"q1", "", "", "q2", "$"},
//					{"q2", "0", "", "q2", "0"},
//					{"q2", "1", "0", "q3", ""},
//					{"q3", "1", "0", "q3", ""},
//					{"q3", "", "$", "q4", ""}},
//				Eos: "$",
//			},
//			state: "q1",
//		}
//
//		result,err := pda.Put("0011")
//
//		if err != nil{
//			t.Errorf("%s",err)
//		}
//		print(result)
//	})
//}
