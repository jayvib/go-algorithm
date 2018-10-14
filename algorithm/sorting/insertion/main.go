package main

import "fmt"

func InsertionSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	var temp, i, j int
	for i = 1; i < size; i++ { // starting from the second element
		temp = arr[i] // hold the 2nd element value
		for j = i; j > 0 && comp(arr[j-1], temp); j-- { // get the previous item and compare... if the previous item is greater then swapped
			arr[j] = arr[j-1] // swapped
		}
		arr[j] = temp
	}
}

func more(a, b int) bool {
	return a > b
}

func main() {
	data := []int{9,1,8,9,3,7,3,6,4,5}
	InsertionSort(data, more)
	fmt.Println(data)
}
