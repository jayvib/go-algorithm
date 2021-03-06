package main

import (
	"fmt"
	"sort"
)

// First Approach:
func getMax(data []int) int {
	size := len(data)
	max := data[0]
	count := 1
	maxCount := 1
	for i := 0; i < size; i++ {
		count = 1 // this is needed so that the next iteration of index i, this value will come back to 1.
		for j := i+1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			max = data[i]
		}
	}
	return max
}

// Brute Force(Iteration)
// Time Complexity: O(n)
// Space Complexity: O(1)
func getMax1Review(data []int) int {
	size := len(data)
	max := data[0]
	count := 1
	maxCount := 1

	// Find the items that has highest
	// frequency count.
	for i := 0; i < size; i++ {
		count = 1
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

	return max
}

func getMax2(data []int) int {
	sort.Ints(data)
	size := len(data)
	max := data[0] // ???
	maxCount := 1
	curr := data[0]
	currCount := 1
	//fmt.Println(curr, max)
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] { //
			currCount++
		} else {
			currCount = 1
			curr = data[i]
		}
		if currCount > maxCount {
			maxCount = currCount
			max = curr
		}
	}
	return max
}

func getMax2Review(data []int) int {
	sort.Ints(data)
	size := len(data)
	max := data[0] // initial max
	maxCount := 1
	currentSelected := data[0]
	currentCount := 1

	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			currentCount++
		} else {
			// New item to be compared.
			currentCount = 1
			currentSelected = data[i]
		}

		if currentCount > maxCount {
			maxCount = currentCount
			max = currentSelected
		}
	}
	return max
}

func getMax3(data []int, dataRange int) int {
	maxCount := 1
	max := data[0]
	size := len(data)
	count := make([]int, dataRange)
	for i := 0; i < size; i++ {
		count[data[i]]++
		if count[data[i]] > maxCount {
			maxCount = count[data[i]]
			max = data[i]
		}
	}
	return max
}

func getMax3V2(data []int, dataRange int) int {
	maxCount := 1
	max := data[0]
	size := len(data)
	count := make([]int, dataRange)
	for i := 0; i < size; i++ {
		count[data[i]]++
		if count[data[i]] > maxCount {
			maxCount = count[data[i]]
			max = data[i]
		}
	}
	return max
}

func main() {
	arr := []int{7, 2, 1, 29, 29,29,3, 2,5,2,3,3,3, 3, 3, 3, 3,7,9,29,19, 29, 29, 3, 3, 3, 1,1,1,1,1,1}
	max := getMax1Review(arr)
	fmt.Println(max)
}
