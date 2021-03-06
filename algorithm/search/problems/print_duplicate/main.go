package main

import (
	"fmt"
	"github.com/golang-collections/collections/set"
	"sort"
)

// Problem:
// given a list of n numbers, print the duplicate elements in the list.

// First Approach:
// Exhaustive search or Brute force, for each element in list find if there is
// some other element with the same value. This is done using two for loop,
// first loop to select the element and second loop to find its duplicate entry.

func printRepeating(data []int) {

	size := len(data)
	fmt.Println("Repeating elements are:")
	for i := 0; i < size; i++ { // iterate each element...
		for j := i+1; j < size; j ++ { // holding the value from index i.. iterate succedding element.
			if data[i] == data[j] { // compare the value from index i to index j.
				fmt.Print(" ", data[i])
				break
			}
		}
	}
	fmt.Println()
}

func printRepeating2(data []int) {
	size := len(data)
	sort.Ints(data)

	fmt.Println("Repeating elements are:")
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			fmt.Print(" ", data[i])
		}
	}
	fmt.Println()
}

// Third Approach: using the hash-table
func printRepeating3(data []int) {
	s := set.New() // create a storage to store the unique item.
	size := len(data)
	printedElements := set.New() // store unique element that has duplicate
	fmt.Println("Repeating elements are:")
	for i := 0; i < size; i++ { // iterate all the list.. O(n)
		if s.Has(data[i]) && !printedElements.Has(data[i]) { // if duplicate and the element wasn't printed yet.
			fmt.Print(" ", data[i])
			printedElements.Insert(data[i])
		} else {
			s.Insert(data[i])
		}
	}
	fmt.Println()
}

// Forth approach:
// intrange is the maximum value of the input.
func printRepeating4(data []int, intrange int) {
	size := len(data)
	count := make([]int, intrange)

	fmt.Println("Repeating elements are:")
	for i := 0; i < size; i++ {
		if count[data[i]] == 1 {
			fmt.Print(" ", data[i])
		} else {
			count[data[i]]++
		}
	}
	fmt.Println()
}
func main() {
	data := []int{2, 5, 3, 6, 8, 8, 2, 5, 10, 3, 2, 3, 1, 6, 7}
	printRepeating(data)
	printRepeating2(data)
	printRepeating3(data)

	sort.Ints(data)
	printRepeating4(data, len(data))
}



