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


func main() {
	arr := []int{1,2,3,4,5}
	fmt.Println(findMissingNumberEfficient(arr))
}
