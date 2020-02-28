package main

import "testing"

func TestOpen(t *testing.T) {
	t.Run("Negative", func(t *testing.T) {
		str := "{}"
		json := open([]byte(str))
		if json != false {
			t.Errorf("output for %s is \n %t; want false", str, json)
		}
	})

	t.Run("Positive", func(t *testing.T) {
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
			`["q3", null, "$", "q4", null]], ` +
			`“ eos”: “$”}`
		json := open([]byte(str))
		if json != true {
			t.Errorf("output for %s is \n %t; want true", str, json)
		}
	})

	t.Run("Negative - missing fields", func(t *testing.T) {
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
		json := open([]byte(str))
		if json != false {
			t.Errorf("output for %s is \n %t; want false", str, json)
		}
	})

	t.Run("Negative - wrong fields", func(t *testing.T) {
		str := `{"name1": "HelloPDA"}`
		json := open([]byte(str))
		if json != false {
			t.Errorf("output for %s is \n %t; want false", str, json)
		}
	})
}

func TestReset(t *testing.T) {
	t.Run("should reset the pda stack to empty", func(t *testing.T) {
		pda.stack.push("a")
		if pda.stack.isEmpty() {
			t.Errorf("initial stack is empty")
		}

		reset()

		if !pda.stack.isEmpty() {
			t.Errorf("stack is not reset")
		}
	})
}
