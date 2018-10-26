package main

import "fmt"

// Problem: Given a sorted list array find the number of occurrences of a number

func findKeyCount2(data []int, key int) int {
	size := len(data)
	firstIndex := findFirstIndex(data, 0, size-1, key)
	lastIndex := findLastIndex(data, 0, size-1, key)
	return lastIndex-firstIndex+1
}

// Using the Binary Search
func findFirstIndex(data []int, start, end, key int) int {
	if end < start {
		return -1
	}
	mid := (start + end)/2 // get the middle
	if key == data[mid] && (mid == start || data[mid-1] != key) {
		return mid
	}
	if key < data[mid] {
		return findFirstIndex(data, start, mid-1, key)
	}
	return findFirstIndex(data, mid+1, end, key)
}

func findLastIndex(data []int, start, end, key int) int {
	if end < start {
		return -1
	}
	mid := (start+end)/2
	if key == data[mid] && (mid == end || data[mid+1] != key) {
		return mid
	}
	if key < data[mid] {
		return findLastIndex(data, start, mid-1, key)
	}
	return findLastIndex(data, mid+1, end, key)
}

func main() {
	arr := []int{1, 2, 3, 3, 4, 5, 2}
	occur := findKeyCount2(arr, 3)
	fmt.Println(occur)
}