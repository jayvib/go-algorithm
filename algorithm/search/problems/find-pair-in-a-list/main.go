package main

import (
	"fmt"
	"sort"
)

// PROBLEM: Given a list of n numbers, find two elements sch that their sum is equal to "value"


// First Approach: Exhaustive search or Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func findPair(data []int, value int) bool {
	size := len(data)
	found := false
	for i := 0; i < size; i++ {
		for j := i+1; j < size; j++ {
			if (data[i]+data[j]) == value {
				fmt.Println("The pair is:", data[i], ",", data[j])
				found = true
			}
		}
	}
	return found
}

// Second Approach: Sorting
// Time Complexity: O(n.logn)
// Space Complexity: O(1)
func findPair2(data []int, value int) bool {
	size := len(data)
	firstIndex := 0
	secondIndex := size-1
	found := false
	sort.Ints(data)
	for firstIndex < secondIndex {
		curr := data[firstIndex] + data[secondIndex]
		if curr == value {
			fmt.Println("The pair is", data[firstIndex], data[secondIndex])
			found = true
		}
		if curr < value {
			firstIndex++
		} else {
			secondIndex--
		}
	}
	return found
}

func findPair2v2(data []int, value int) bool {
	size := len(data)
	firstIndex := 0
	secondIndex := size-1
	isFound := false
	sort.Ints(data)
	for firstIndex < secondIndex {
		current := data[firstIndex] + data[secondIndex]
		if current == value {
			fmt.Println("The pair is", data[firstIndex], data[secondIndex])
			isFound = true
		}
		if current < value {
			firstIndex++
		} else {
			secondIndex--
		}
	}
	return isFound
}

type Set map[int]struct{}

func (s Set) Add(i int) {
	s[i] = struct{}{}
}

func (s Set) Find(value int) bool {
	_, ok := s[value]
	return ok
}


// Third Approach: Hash Table
func findPair3(data []int, value int) bool {
	s := make(Set)
	size := len(data)
	isFound := false
	for i := 0; i < size; i++ {
		if s.Find(value-data[i]) {
			fmt.Println(i, "The pair is:", data[i], ",", value-data[i])
			isFound = true
		} else {
			s.Add(data[i])
		}
	}
	return isFound
}
