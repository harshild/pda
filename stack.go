package main

type Stack []string

func (stack *Stack) size() int {
	return len(*stack)
}

func (stack *Stack) isEmpty() bool {
	return stack.size() == 0
}

func (stack *Stack) push(str string) {
	*stack = append(*stack, str)
}

func (stack *Stack) pop() string {
	if stack.size() == 0 {
		return ""
	}
	topElement := stack.topElement()
	elementIndex := stack.size() - 1
	*stack = (*stack)[:elementIndex]
	return topElement

}

func (stack *Stack) topElement() string {
	elementIndex := stack.size() - 1
	return (*stack)[elementIndex]
}
