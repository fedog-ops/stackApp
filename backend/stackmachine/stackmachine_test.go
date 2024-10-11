package stackmachine

import (
	// "errors"
	"fmt"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	stack := NewStackFromNumbers([]int{1, 2, 3})

	poppedValue, err := stack.Pop()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if poppedValue != 3 {
		t.Errorf("expected popped value to be 3, but got %d", poppedValue)
	}

	if stack.len() != 2 {
		t.Errorf("expected stack size to be 2 after pop, but got %d", len(stack.StackNumbers))
	}

	emptyStack := NewStack()
	_, err = emptyStack.Pop()
	if err == nil {
		t.Error("expected an error when popping from an empty stack, but got none")
	}
}

func TestStack_Push(t *testing.T) {
	stack := NewStack()

	stack.Push(5)

	if stack.len() != 1 {
		t.Errorf("expected stack size to be 1 after pushing, but got %d", len(stack.StackNumbers))
	}

	top := stack.Peek()
	if top != 5 {
		t.Errorf("expected top element to be 5 after pushing, but got %d", stack.StackNumbers[0])
	}

	stack.Push(10)
	stack.Push(15)

	if stack.len() != 3 {
		t.Errorf("expected stack size to be 3 after pushing multiple elements, but got %d", len(stack.StackNumbers))
	}

	top = stack.Peek()
	if top != 15 {
		t.Errorf("expected top element to be 15 after pushing, but got %d", stack.StackNumbers[2])
	}

	stackWithNumbers := NewStackFromNumbers([]int{1, 2, 3})
	stackWithNumbers.Push(4)

	if stackWithNumbers.len() != 4 {
		t.Errorf("expected stack size to be 4 after pushing, but got %d", len(stackWithNumbers.StackNumbers))
	}

	if stackWithNumbers.StackNumbers[3] != 4 {
		t.Errorf("expected top element to be 4 after pushing, but got %d", stackWithNumbers.StackNumbers[3])
	}

	stackWithBigNumbers := NewStackFromNumbers([]int{20001, 30000})
	err := stackWithBigNumbers.Plus()
	if err == nil {
		t.Errorf("expected an error when summing 20001 and 30000 but did not receive one.")
	}
}

func TestStack_Duplicate(t *testing.T) {
	stack := NewStackFromNumbers([]int{1, 2, 3})

	stack.Duplicate()

	if stack.len() != 4 {
		t.Errorf("expected stack size to be 4 after DUP, but got %d", len(stack.StackNumbers))
	}

	top := stack.Peek()
	if top != 3 {
		t.Errorf("expected top element to be 3 after DUP, but got %d", stack.StackNumbers[3])
	}
	emptyStack := NewStack()

	if emptyStack.len() != 0 {
		t.Errorf("expected stack size to remain 0 after DUP on empty stack, but got %d", emptyStack.len())
	}
}

func TestStack_Plus(t *testing.T) {
	stack := NewStackFromNumbers([]int{10, 20})

	err := stack.Plus()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if stack.len() != 1 {
		t.Errorf("expected stack size to be 1 after addition, but got %d", len(stack.StackNumbers))
	}

	top := stack.Peek()
	if top != 30 {
		t.Errorf("expected top element to be 30 after addition, but got %d", stack.StackNumbers[0])
	}

	stackWithOverflow := NewStackFromNumbers([]int{25000, 26000})

	err = stackWithOverflow.Plus()
	if err == nil {
		t.Fatal("expected an overflow error, but got none")
	}

	stackWithOneElement := NewStackFromNumbers([]int{10})

	err = stackWithOneElement.Plus()
	if err == nil {
		t.Fatal("expected an error due to insufficient elements, but got none")
	}

}

func TestStack_Minus(t *testing.T) {
	stack := NewStackFromNumbers([]int{10, 20})

	err := stack.Minus()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if stack.len() != 1 {
		t.Errorf("expected stack size to be 1 after subtraction, but got %d", len(stack.StackNumbers))
	}

	top := stack.Peek()
	if top != 10 {
		t.Errorf("expected top element to be 10 after subtraction, but got %d", stack.StackNumbers[0])
	}

	stackWithNegativeResult := NewStackFromNumbers([]int{20, 10})

	err = stackWithNegativeResult.Minus()
	if err == nil {
		t.Fatal("expected a result below zero error, but got none")
	}

	stackWithOneElement := NewStackFromNumbers([]int{10})

	err = stackWithOneElement.Minus()
	if err == nil {
		t.Fatal("expected an error due to insufficient elements, but got none")
	}
}

func TestStack_Multiply(t *testing.T) {
	stack := NewStackFromNumbers([]int{10, 5})

	err := stack.Multiply()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if stack.len() != 1 {
		t.Errorf("expected stack size to be 1 after multiplication, but got %d", len(stack.StackNumbers))
	}
	top := stack.Peek()

	if top != 50 {
		t.Errorf("expected top element to be 50 after multiplication, but got %d", stack.StackNumbers[0])
	}

	stackWithOneElement := NewStackFromNumbers([]int{10})

	err = stackWithOneElement.Multiply()
	if err == nil {
		t.Fatal("expected an error due to insufficient elements, but got none")
	}

	stackWithLargeNumbers := NewStackFromNumbers([]int{1000, 2000})

	err = stackWithLargeNumbers.Multiply()
	if err == nil {
		t.Error("expected error when multiplying two large numbers but did not receieve one.")
	}

}

func TestStack_Sum(t *testing.T) {
	stack := NewStackFromNumbers([]int{10, 20, 30})

	err := stack.Sum()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if stack.len() != 1 {
		t.Errorf("expected stack size to be 1 after SUM, but got %d", len(stack.StackNumbers))
	}

	top := stack.Peek()
	if top != 60 {
		t.Errorf("expected top element to be 60 after SUM, but got %d", stack.StackNumbers[0])
	}

	emptyStack := NewStack()

	err = emptyStack.Sum()
	if err == nil {
		t.Fatal("expected an error due to empty stack, but got none")
	}

	stackWithOneElement := NewStackFromNumbers([]int{42})

	err = stackWithOneElement.Sum()
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if stack.len() != 1 {
		t.Errorf("expected stack size to be 1 after SUM, but got %d", len(stackWithOneElement.StackNumbers))
	}

	top = stackWithOneElement.Peek()
	if top != 42 {
		t.Errorf("expected top element to be 42 after SUM, but got %d", top)
	}
}

/*
All these tests must pass for completion
*/

func TestAcceptanceTests(t *testing.T) {
	tests := []struct {
		name        string
		commands    string
		expected    int
		shouldError bool
	}{
		{name: "empty error", commands: "", expected: 0, shouldError: true},
		{name: "add overflow", commands: "50000 DUP +", expected: 0, shouldError: true},
		{name: "too few add", commands: "99 +", expected: 0, shouldError: true},
		{name: "too few minus", commands: "99 -", expected: 0, shouldError: true},
		{name: "too few multiply", commands: "99 *", expected: 0, shouldError: true},
		{name: "empty stack", commands: "99 CLEAR", expected: 0, shouldError: true},
		{name: "sum single value", commands: "99 SUM", expected: 99, shouldError: false},
		{name: "sum empty", commands: "SUM", expected: 0, shouldError: true},
		{name: "normal +*", commands: "5 6 + 2 *", expected: 22, shouldError: false},
		{name: "clear too few", commands: "1 2 3 4 + CLEAR 12 +", expected: 0, shouldError: true},
		{name: "normal after clear", commands: "1 CLEAR 2 3 +", expected: 5, shouldError: false},
		{name: "single integer", commands: "9876", expected: 9876, shouldError: false},
		{name: "invalid command", commands: "DOGBANANA", expected: 0, shouldError: true},
		{name: "normal +-*", commands: "5 9 DUP + + 43 - 3 *", expected: 60, shouldError: false},
		{name: "minus", commands: "2 5 -", expected: 3, shouldError: false},
		{name: "underflow minus", commands: "5 2 -", expected: 0, shouldError: true},
		{name: "at overflow limit", commands: "25000 DUP +", expected: 50000, shouldError: false},
		{name: "at overflow limit single value", commands: "50000 0 +", expected: 50000, shouldError: false},
		{name: "overflow plus", commands: "50000 1 +", expected: 0, shouldError: true},
		{name: "overflow single value", commands: "50001", expected: 0, shouldError: true},
		{name: "times zero at overflow limit", commands: "50000 0 *", expected: 0, shouldError: false},
		{name: "too few at first", commands: "1 2 3 4 5 + + + + * 999", expected: 0, shouldError: true},
		{name: "normal simple", commands: "1 2 - 99 +", expected: 100, shouldError: false},
		{name: "at overflow minus to zero", commands: "50000 50000 -", expected: 0, shouldError: false},
		{name: "clear empties stack", commands: "CLEAR", expected: 0, shouldError: true},
		{name: "normal sum", commands: "3 4 3 5 5 1 1 1 SUM", expected: 23, shouldError: false},
		{name: "sum after clear stack", commands: "3 4 3 5 CLEAR 5 1 1 1 SUM", expected: 8, shouldError: false},
		{name: "sum then too few", commands: "3 4 3 5 5 1 1 1 SUM -", expected: 0, shouldError: true},
		{name: "fibonacci", commands: "1 2 3 4 5 * * * *", expected: 120, shouldError: false},
	}

	for _, test := range tests {

		fmt.Println(test.name)
		got, err := StackMachine(test.commands)
		fmt.Println(got)

		if test.shouldError {
			if err == nil {
				t.Errorf("%s (%s) Expected error, but received nil", test.name, test.commands)
			}
		} else if got != test.expected {
			t.Errorf("%s (%s) got %v, want %v", test.name, test.commands, got, test.expected)
		}
	}
}
