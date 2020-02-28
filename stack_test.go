package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Push new element should result in size increase", func(t *testing.T) {
		stack := Stack{}
		size := stack.size()
		if size != 0 {
			t.Errorf("Expected initial size to be 0 but got %d", size)
		}

		stack.push("a")
		size = stack.size()
		if size != 1 {
			t.Errorf("Expected the size to be 1 but got %d", size)
		}
	})

	t.Run("Size of stack should increment with number of variables", func(t *testing.T) {
		stack := Stack{}
		size := stack.size()
		if size != 0 {
			t.Errorf("Expected initial size to be 0 but got %d", size)
		}

		stack.push("a")
		size = stack.size()
		if size != 1 {
			t.Errorf("Expected the size to be 1 but got %d", size)
		}

		stack.push("b")
		size = stack.size()
		if size != 2 {
			t.Errorf("Expected the size to be 2 but got %d", size)
		}

		stack.push("a")
		stack.push("a")
		stack.push("a")
		size = stack.size()
		if size != 5 {
			t.Errorf("Expected the size to be 5 but got %d", size)
		}

	})

	t.Run("isEmpty Positive", func(t *testing.T) {
		stack := Stack{}
		empty := stack.isEmpty()
		if !empty {
			t.Errorf("Expected isEmpty to return false got %t", empty)
		}
	})

	t.Run("isEmpty Negative", func(t *testing.T) {
		stack := Stack{}
		stack.push("a")
		empty := stack.isEmpty()
		if empty {
			t.Errorf("Expected isEmpty to return true got %t", empty)
		}
	})

	t.Run("We should be able to get top element on the stack", func(t *testing.T) {
		stack := Stack{}
		stack.push("a")
		stack.push("b")
		stack.push("c")
		element := stack.topElement()
		if element != "c" {
			t.Errorf("Expected topmost element to be c got %s", element)
		}
	})

	t.Run("Pop element should pop last added element", func(t *testing.T) {
		stack := Stack{}
		stack.push("a")
		stack.push("b")
		stack.push("c")
		element := stack.pop()
		if element != "c" {
			t.Errorf("Expected poped element to return c got %s", element)
		}
	})

	t.Run("Pop should return null when stack is empty", func(t *testing.T) {
		stack := Stack{}
		element := stack.pop()
		if element != "" {
			t.Errorf("Empty stack pop:Expected null got #{element}")
		}

	})
}
