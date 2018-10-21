package main

import (
	"fmt"
)


func sum(data []int) int {
	total := 0
	for _, d := range data {
		total += d
	}
	return total
}

func findMissingNumberEfficient(data []int) (missing int, found bool) {
	size := len(data)
	total := (size+1)*(size+2)/2
	sumOfData := sum(data)
	missing = total - sumOfData
	if missing == size+1 {
		return 0, false
	}
	return missing, true
}

func findMissingNumberXOR(data []int, dataRange int) int {
	size := len(data)
	xorSum := 0
	for i := 1; i <= dataRange; i++ {
		xorSum ^= i
	}
	for i := 0; i < size; i++ {
		xorSum ^= data[i]
	}
	return xorSum
}

func findMissingNumberXOR2(data []int, datarange int) int {
	size := len(data)
	xorSum := 0
	for i := 1; i <= datarange; i++ {
		xorSum ^= i
	}
	for i := 0; i < size; i++ {
		xorSum ^= data[i]
	}
	return xorSum
}

func main() {
	arr := []int{1,2,3,6,8,4,5}
	fmt.Println(findMissingNumberEfficient(arr))
	fmt.Println(findMissingNumberXOR(arr, 8))
}
