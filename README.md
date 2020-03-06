# Push-down Automata

This is an implementation of pushdown automata. For project definition refer the description file in root folder.

### Assumptions

For the purposes of tasks of this project, you may make the following assumptions:
1. PDAs are deterministic (eg, at most one transition from a current state, input, top of stack token
triplet). In general, PDAs are non-deterministic devices.
2. PDAs are eager and instantaneously take all possible transitions from a current state that do not
need the consumption of any input tokens.
3. The transition that pushes the $ (bottom-of-the-stack marker) onto the stack from the initial state is
taken lazily at (right before) the first put() or eos(). The transition that pops $ from the stack is
taken lazily at the eos().
4. All (non-special) tokens in the input and stack alphabets are sequences of ASCII (printable) non-
whitespace characters.
5. A token-stream is marshalled by joining the sequence of tokens with a whitespace into an ASCII
string
6. A token-stream is unmarshalled from an ASCII string by tokenizing (splitting) it on contiguous
whitespaces (ie. consecutive whitespaces are effectively suppressed into one whitespace).
7. PDA JSON specifications follow the structure of the example PDA above.
8. PDA JSON specifications are well-formed and valid.

### Build 

To build , from the project root

`go build ./main/pda_client.go`

### Run

`./pda_client [CONFIG FILE PATH] [INPUT TOKEN or INPUT FILE PATH]`

### Run Script

`/bin/bash  ./sample_run.sh > standard-output 2>standard-error`

or 

`/bin/bash  ./sample_run.sh`

## Samples

### Sample Test Files

You may look up `sample` folder in project root and sample_run.sh for some examples

### Sample Config JSON

```
{
  "name": "PDA Name",
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
    ["q3", null, "$", "q4", null]
  ],
  "eos": "$"
}
```