package main

import (
	"fmt"
	"sort"
)

// Problem:
// Given a list of n elements. Find the majority element, which appears more than n/2
// times. Return 0 in case there is no majority element.

// First Approach:
func getMajority(data []int) (int, bool) {
	size := len(data)
	max := 0
	count := 0
	maxCount := 0
	for i := 0; i < size; i++ {
		for j := i+1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			max = data[i]
			maxCount = count
		}
	}
	if maxCount > size/2 {
		return max, true
	}
	return 0, false
}

// Second Approach
func getMajority2(data []int) (int, bool) {
	sort.Ints(data)
	size := len(data)
	majIndex := size/2
	candidate := data[majIndex]
	count := 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return candidate, true
	}
	return 0, false
}

// Third Approach: Moore Vooting Algorithm
func getMajority3(data []int) (int, bool) {
	majIndex := 0
	count := 1
	size := len(data)
	// --------MOORE VOTING ALGORITHM---------
	for i := 1; i < size; i++ {
		if data[majIndex] == data[i] {
			count++
		} else {
			count--
		}
		if count == 0 { // cancel the previous value.. a new major index
			majIndex = i
			count = 1
		}
	}
	// ---------------------------------------
	candidate := data[majIndex]
	count = 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return data[majIndex], true
	}
	return 0, false
}

func main() {
	arr := []int{3, 3, 4, 2, 4, 4, 2, 4, 4}
	r, _ := getMajority3(arr)
	fmt.Println(r)
}
