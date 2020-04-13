package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	arr = make([]int, 10)
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	InsertionSort(&arr)
	fmt.Println(arr)
}

// InsertionSort sort in array in ascending order
func InsertionSort(arr *[]int) {
	var arrCopy []int
	Copy(arr, &arrCopy)

	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = MinInt(arrCopy)
		ind := Index(arrCopy, (*arr)[i])
		Remove(&arrCopy, ind)
	}
}

// MinInt returns min element of a slice
func MinInt(arr []int) (minEl int) {
	minEl = arr[0]
	for i := 0; i < len(arr); i++ {
		if minEl > arr[i] {
			minEl = arr[i]
		}
	}
	return
}

// Index return index of 'el' in arr[]
func Index(arr []int, el int) (i int) {
	for i = 0; i < len(arr); i++ {
		if arr[i] == el {
			return
		}
	}
	i = -1
	return
}

// Copy makes a copy of 'from' as 'to'
func Copy(from *[]int, to *[]int) {
	*to = make([]int, len(*from))
	for i := 0; i < len(*from); i++ {
		(*to)[i] = (*from)[i]
	}
}

// Remove gets rid of element in pos' index
func Remove(arr *[]int, pos int) {
	(*arr)[pos] = (*arr)[len(*arr)-1]
	(*arr) = (*arr)[:len(*arr)-1]
}
