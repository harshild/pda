package src

type Stack []string

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

func (stack *Stack) Peek(k int) []string {
	topKValues := stack.Size() - k - 1
	return (*stack)[topKValues:(stack.Size() - 1)]
}
