package stackmachine

import (
	"errors"
	"fmt"
	"strings"
)

func StackMachine(commands string) (int, error) {
	individualCommands := strings.Split(commands, " ")
	stack := NewStack()
	var err error
	var num int
	for _, cmd := range individualCommands {
		switch cmd {
		case "POP":
			_, err = stack.Pop()
			if err != nil {
				return ERROR_RETURN, err
			}
		case "DUP":
			stack.Duplicate()
		case "+":
			err = stack.Plus()
			if err != nil {
				return ERROR_RETURN, err
			}
		case "-":
			err = stack.Minus()
			if err != nil {
				return ERROR_RETURN, err
			}
		case "*":
			err = stack.Multiply()
			if err != nil {
				return ERROR_RETURN, err
			}
		case "CLEAR":
			stack.Clear()
		case "SUM":
			err = stack.Sum()
			if err != nil {
				return ERROR_RETURN, err
			}

		default:
			if isNumber(cmd, &num) {
				stack.Push(num)
			} else {
				// Non command given
				return 0, errors.New("Invalid command.")
			}
		}
	}

	fmt.Println(stack.StackNumbers)
	if stack.isEmpty() {
		return 0, errors.New("Empty stack, nothing to return.")
	} else {
		return stack.Peek(), nil
	}
}


