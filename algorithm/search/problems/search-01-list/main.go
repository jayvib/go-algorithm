package main

// Problem: Given a list of 0's ad 1's. All the 0's come before 1's. Write an algorithm to find
// the index of the first 1.

// Approach: Binary Search
func BinarySearch01(data []int) int {
	size := len(data)
	if size == 1 && data[0] == 1 {
		return -1
	} else if size == 0 {
		return -1
	}
	return BinarySearch01Util(data, 0, size-1)
}

func BinarySearch01Util(data []int, start, end int) int {
	if end < start {
		return -1
	}
	mid := (start+end)/2
	if data[mid] == 1 && data[mid-1] == 0 {
		return mid
	}
	if data[mid] == 0 {
		// search to the upper portion
		return BinarySearch01Util(data, mid+1, end)
	}
	return BinarySearch01Util(data, start, mid-1)
}
