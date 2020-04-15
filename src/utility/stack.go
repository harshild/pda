package utility

import "strconv"

type Stack []string

type StackRuntimeError struct {
	message string
}

func (e *StackRuntimeError) Error() string {
	return e.message
}

func (stack *Stack) Size() int {
	return len(*stack)
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Push(str string) {
	*stack = append(*stack, str)
}

func (stack *Stack) Pop() string {
	if stack.Size() == 0 {
		return ""
	}
	topElement := stack.TopElement()
	elementIndex := stack.Size() - 1
	*stack = (*stack)[:elementIndex]
	return topElement

}

func (stack *Stack) TopElement() string {
	elementIndex := stack.Size() - 1
	return (*stack)[elementIndex]
}

func (stack *Stack) Peek(len int) ([]string, error) {
	if stack.Size() < len {
		return nil, &StackRuntimeError{"Peek length is " + strconv.Itoa(len) + ", But the size of stack is " + strconv.Itoa(stack.Size())}
	}
	topKValues := stack.Size() - len - 1
	return (*stack)[topKValues:(stack.Size() - 1)], nil
}
