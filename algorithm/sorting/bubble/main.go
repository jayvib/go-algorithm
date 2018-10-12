package main

import (
	"math/rand"
	"time"
)

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func bubbleSort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n--

	}
}

func bubbleSort2(slice []int) {
	var (
		size   = len(slice)
		sorted = false
	)

	for !sorted {
		swapped := false
		for i := 0; i < size-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i+1], slice[i] = slice[i], slice[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		size--
	}
}
