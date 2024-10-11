package stackmachine

import (
	"errors"
	"strconv"
)

const ERROR_RETURN = -1

type Stack struct {
	StackNumbers     []int
	lastElementIndex int
}

func NewStack() *Stack {
	var emptyStackNumbers []int = []int{}
	return &Stack{
		emptyStackNumbers,
		-1,
	}
}

func NewStackFromNumbers(stackNumbers []int) *Stack {
	return &Stack{
		stackNumbers,
		len(stackNumbers) - 1,
	}
}

func (stack *Stack) Peek() int {
	return stack.StackNumbers[stack.lastElementIndex]
}

func (stack *Stack) Push(value int) error {
	if isNumberAllowedOnStack(value) {
		stack.StackNumbers = append(stack.StackNumbers, value)
		stack.lastElementIndex += 1
		return nil
	} else {
		return errors.New("This value is not allowed on the stack.")
	}
}

func (stack *Stack) Duplicate() {
	// Non empty stack
	if stack.lastElementIndex >= 0 {
		stack.Push(stack.Peek())
	}
}

func isNumberAllowedOnStack(number int) bool {
	return number <= 50000 && number >= 0
}

func (stack *Stack) len() int {
	return len(stack.StackNumbers)
}
func (stack *Stack) hasTwoNumbers() bool {
	return stack.len() >= 2
}

func (stack *Stack) PopTwoMostRecentNumbers(err *error) (int, int) {
	if stack.hasTwoNumbers() {
		firstNumber, _ := stack.Pop()
		secondNumber, _ := stack.Pop()
		*err = nil
		return firstNumber, secondNumber
	} else {
		returnedError := errors.New("Failed to pop two numbers from the stack.")
		*err = returnedError
		return 0, 0
	}
}

func (stack *Stack) Plus() error {
	var err error
	firstNumber, secondNumber := stack.PopTwoMostRecentNumbers(&err)
	if err == nil {
		sum := firstNumber + secondNumber
		err = stack.Push(sum)
	}
	return err
}

func (stack *Stack) Pop() (int, error) {
	if !stack.isEmpty() {
		lastElement := stack.Peek()
		stack.StackNumbers = stack.StackNumbers[:stack.lastElementIndex]
		stack.lastElementIndex -= 1
		return lastElement, nil
	} else {
		return ERROR_RETURN, errors.New("Cannot remove from empty stack.")
	}

}

func (stack *Stack) Minus() error {
	var err error
	firstNumber, secondNumber := stack.PopTwoMostRecentNumbers(&err)
	if err == nil {
		err = stack.Push(firstNumber - secondNumber)
		return err
	}
	return err
}

func (stack *Stack) Multiply() error {
	var err error
	firstNumber, secondNumber := stack.PopTwoMostRecentNumbers(&err)
	if err == nil {
		err = stack.Push(firstNumber * secondNumber)
		return err
	}
	return err
}

func (stack *Stack) Clear() {
	stack.StackNumbers = []int{}
	stack.lastElementIndex = -1
}

func (stack *Stack) isEmpty() bool {
	return stack.lastElementIndex == -1
}
func (stack *Stack) Sum() error {
	if stack.isEmpty() {
		return errors.New("Cannot sum over an empty stack.")
	}
	var err error
	for {
		if err != nil {
			break
		}
		err = stack.Plus()
	}

	return nil
}

func isNumber(cmd string, num *int) bool {
	number, err := strconv.Atoi(cmd)
	if err == nil {
		*num = number
		return true
	}

	return false

}
