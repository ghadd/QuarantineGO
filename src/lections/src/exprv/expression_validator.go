package main

import (
	"bufio"
	"fmt"
	"lections/src/exprv/stack"
	"os"
)

var (
	brackets = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
		'>': '<',
	}
	inputStatement string
)

func containsKey(m map[rune]rune, item rune) bool {
	for k := range m {
		if k == item {
			return true
		}
	}
	return false
}

func containsValue(m map[rune]rune, item rune) bool {
	for _, v := range m {
		if v == item {
			return true
		}
	}
	return false
}

func isValid(str string) (bool, int) {
	var charStack stack.Stack
	charStack.Init()

	for pos, char := range str {
		switch {
		case containsValue(brackets, char):
			charStack.Push(char)
		case containsKey(brackets, char):
			if charStack.Peek() == brackets[char] {
				charStack.Pop()
			} else {
				return false, pos
			}
		}
	}
	if charStack.Size != 0 {
		return false, charStack.Size - 1
	}
	return true, -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStatement = scanner.Text()

	res, err := isValid(inputStatement)
	switch res {
	case true:
		fmt.Println("Expression is ok!")
	case false:
		fmt.Println("Unexpected token at pos ", err)
	}
}
