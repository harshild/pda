# [Push-down Automata](./README.md)

This is an implementation of pushdown automata. For project definition refer the description file in root folder.

## Push-down Automata - Web

### Project Structure
#### Available REST APIs

|HTTP method | URL | Meaning|
|---|---|---|
|GET | base/pdas | List of names of PDAs available at the server
|PUT | base/pdas/id| Create at the server a PDA with the given id and the specification provided in the body of the request; calls open() method of PDA processor
|PUT | base/pdas/id/reset | Call reset() method
|PUT | base/pdas/id/tokens/position | Present a token at the given position
|PUT | base/pdas/id/eos/position |Call eos() with no tokens after (excluding) position
|GET | base/pdas/id/is_accepted |Call and return the value of is_accepted()
|GET | base/pdas/id/stack/top/k |Call and return the value of peek(k)
|GET | base/pdas/id/stack/len |Return the number of tokens currently in the stack
|GET | base/pdas/id/state |Call and return the value of current_state()
|GET | base/pdas/id/tokens |Call and return the value of queued_tokens()
|GET | base/pdas/id/snapshot/k |Return a JSON message (array) three components: the current_state(), queued_tokens(), and peek(k)
|PUT | base/pdas/id/close |Call close()
|DELETE | base/pdas/id/delete |Delete the PDA with name from the server

### Build 

To build , from the project root

`go build ./main/pda_client_web.go`

### Run

`./pda_client_web [port(optional)]`

The default port is **8080**