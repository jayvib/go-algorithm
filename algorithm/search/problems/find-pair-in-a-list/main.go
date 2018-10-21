package main

import "fmt"

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
