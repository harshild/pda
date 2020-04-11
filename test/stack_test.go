package test

import (
	"utility"
)

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Push new element should result in size increase", func(t *testing.T) {
		stack := utility.Stack{}
		size := stack.Size()
		if size != 0 {
			t.Errorf("Expected initial size to be 0 but got %d", size)
		}

		stack.Push("a")
		size = stack.Size()
		if size != 1 {
			t.Errorf("Expected the size to be 1 but got %d", size)
		}
	})

	t.Run("Size of stack should increment with number of variables", func(t *testing.T) {
		stack := utility.Stack{}
		size := stack.Size()
		if size != 0 {
			t.Errorf("Expected initial size to be 0 but got %d", size)
		}

		stack.Push("a")
		size = stack.Size()
		if size != 1 {
			t.Errorf("Expected the size to be 1 but got %d", size)
		}

		stack.Push("b")
		size = stack.Size()
		if size != 2 {
			t.Errorf("Expected the size to be 2 but got %d", size)
		}

		stack.Push("a")
		stack.Push("a")
		stack.Push("a")
		size = stack.Size()
		if size != 5 {
			t.Errorf("Expected the size to be 5 but got %d", size)
		}

	})

	t.Run("isEmpty Positive", func(t *testing.T) {
		stack := utility.Stack{}
		empty := stack.IsEmpty()
		if !empty {
			t.Errorf("Expected isEmpty to return false got %t", empty)
		}
	})

	t.Run("isEmpty Negative", func(t *testing.T) {
		stack := utility.Stack{}
		stack.Push("a")
		empty := stack.IsEmpty()
		if empty {
			t.Errorf("Expected isEmpty to return true got %t", empty)
		}
	})

	t.Run("We should be able to get top element on the stack", func(t *testing.T) {
		stack := utility.Stack{}
		stack.Push("a")
		stack.Push("b")
		stack.Push("c")
		element := stack.TopElement()
		if element != "c" {
			t.Errorf("Expected topmost element to be c got %s", element)
		}
	})

	t.Run("Pop element should pop last added element", func(t *testing.T) {
		stack := utility.Stack{}
		stack.Push("a")
		stack.Push("b")
		stack.Push("c")
		element := stack.Pop()
		if element != "c" {
			t.Errorf("Expected poped element to return c got %s", element)
		}
	})

	t.Run("Pop should return null when stack is empty", func(t *testing.T) {
		stack := utility.Stack{}
		element := stack.Pop()
		if element != "" {
			t.Errorf("Empty stack pop:Expected null got %s", element)
		}

	})

	t.Run("Peek should return top 3 elements", func(t *testing.T) {
		stack := utility.Stack{}
		stack.Push("a")
		stack.Push("b")
		stack.Push("c")
		stack.Push("d")
		stack.Push("e")
		stack.Push("f")
		stack.Push("g")
		got := stack.Peek(3)
		if len(got) != 3 && utility.StringArrContains(got, "e") && utility.StringArrContains(got, "f") && utility.StringArrContains(got, "g") {
			t.Errorf("Peek should return top 3 elements got %d and elements as %s", len(got), got)
		}

	})
}
