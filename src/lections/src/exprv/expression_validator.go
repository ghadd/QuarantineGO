package main

import (
	"bufio"
	"fmt"
	"lections/src/exprv/stack"
	"math/rand"
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	colors = []color.Color{
		*color.New(color.FgRed),
		*color.New(color.FgGreen),
		*color.New(color.FgBlue),
		*color.New(color.FgYellow),
		*color.New(color.FgMagenta),
	}
	brackets = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	filePath string
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

func isValid(str string) (valid bool, errpos int) {
	var charStack, colorStack stack.Stack
	charStack.Init()
	colorStack.Init()

	var firstErrFound bool
	defer fmt.Println()

	for pos, char := range str {
		switch {
		case containsValue(brackets, char):
			charStack.Push(char)
			var choice int
			for {
				choice = rand.Intn(len(colors))
				if colorStack.Size == 0 {
					break
				}
				if choice != colorStack.Peek().(int) {
					break
				}
			}
			colorStack.Push(choice)
			colors[colorStack.Peek().(int)].Printf("%c", char)
		case containsKey(brackets, char):
			if charStack.Peek() == brackets[char] {
				charStack.Pop()
				colors[colorStack.Peek().(int)].Printf("%c", char)
				colorStack.Pop()
			} else {
				fmt.Printf("%c", char)
				if !firstErrFound {
					valid, errpos = false, pos
				}
			}
		default:
			fmt.Printf("%c", char)
		}
	}

	if firstErrFound {
		return valid, errpos
	}

	if charStack.Size != 0 {
		return false, charStack.Size - 1
	}
	return true, -1
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func pseudoCompile(path string) {
	data, err := readLines(path)

	if err != nil {
		panic(err)
	}

	code := strings.Join(data, "\n")
	res, er := isValid(code)
	if res {
		fmt.Println("Compilation done ok.")
	} else {
		fmt.Println("Compilation error. Unresolved charachter at", er)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filePath = scanner.Text()

	pseudoCompile(filePath)
}
