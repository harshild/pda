package entity

import "utility"

type PDAStatus struct {
	Stack             utility.Stack  `json:"Stack"`
	State             string         `json:"State"`
	Clock             int            `json:"Clock"`
	InputQueue        map[int]string `json:"InputQueue"`
	LastConsumedIndex int            `json:"LastConsumedIndex"`
	PdaId             string         `json:"PdaId"`
	ReplicaId         int            `json:"ReplicaId"`
}

type CookieInfo struct {
	PdaId     string `json:"PdaId"`
	ReplicaId int    `json:"ReplicaId"`
}
