package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	s := generateSliceInt(10, func() int {
		return rand.Intn(20)
	})
	sort.Ints(s) // the values must be sorted already
	fmt.Printf("%v\n", s)
	fmt.Println(len(s))

	fmt.Println(BinarySearch(s, 1))
}

func generateSliceInt(size int, fn func() int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = fn()
	}
	return slice
}

func BinarySearch(data []int, value int) bool {
	size := len(data)
	low := 0
	mid := 0
	high := size - 1

	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		} else if data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
