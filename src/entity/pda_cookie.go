package entity

import "utility"

type PDAStatus struct {
	Stack             utility.Stack
	State             string
	Clock             int
	InputQueue        map[int]string
	LastConsumedIndex int
	PdaName           string
	ReplicaName       string
}
