package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	data, lseq string
)

func longestSortedSequence(data string) (seq string, maxIndex int) {
	var maxLen, length = 1, 1
	maxIndex = 0

	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			length++
		} else {
			if length > maxLen {
				maxLen = length
				maxIndex = i - maxLen
			}
			length = 1
		}
	}

	if length > maxLen {
		maxLen = length
		maxIndex = len(data) - maxLen
	}

	seq = ""
	for i := maxIndex; i < maxIndex+maxLen; i++ {
		seq += string(data[i])
	}

	return
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	data = sc.Text()

	lseq, ind := longestSortedSequence(data)
	fmt.Println("Longest sequence in alphabetical order: ", lseq, "\nLength = ", len(lseq))

	red := color.New(color.FgRed)
	for pos, char := range data {
		if pos >= ind && pos < ind+len(lseq) {
			red.Printf("%c", char)
		} else {
			fmt.Printf("%c", char)
		}
	}
	fmt.Println()
}
