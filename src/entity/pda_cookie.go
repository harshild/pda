package entity

import "utility"

type Cookie struct {
	Stack             utility.Stack
	State             string
	Clock             int
	InputQueue        map[int]string
	LastConsumedIndex int
}
