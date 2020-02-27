package main

import "testing"

func TestOpen(t *testing.T) {
	t.Run("Negative", func(t *testing.T) {
		str := "{}"
		got := open([]byte(str))
		if got != false {
			t.Errorf("output for %s is \n %t; want false", str, got)
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
		got := open([]byte(str))
		if got != true {
			t.Errorf("output for %s is \n %t; want true", str, got)
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
		got := open([]byte(str))
		if got != false {
			t.Errorf("output for %s is \n %t; want false", str, got)
		}
	})

	t.Run("Negative - wrong fields", func(t *testing.T) {
		str := `{"name1": "HelloPDA"}`
		got := open([]byte(str))
		if got != false {
			t.Errorf("output for %s is \n %t; want false", str, got)
		}
	})
}
