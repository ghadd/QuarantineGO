package main

import (
	"fmt"
	"math/rand"
)

func main() {
	height := 25
	for i := 0; i < height; i++ {
		for j := 0; j < height*3; j++ {
			if j < height-i || j > height+i {
				fmt.Print(" ")
			} else if (j == height-1 || j == height+1) && i > 3 {
				fmt.Print("|")
			} else {
				if rand.Intn(10) == 1 {
					fmt.Print("0")
				} else {
					fmt.Print("*")
				}
			}
		}
		fmt.Println()
	}

}
