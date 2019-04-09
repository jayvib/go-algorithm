package main

import "fmt"

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func BubbleSort(arr []int, comp func(int, int)bool) {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		for j := 0; j < (size-i-1); j++ { // this is for the first element, then exclude the last item to hold
			if comp(arr[j], arr[j+1]) {
				// swap
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func bubbleSort(arr []string, more func(string, string)bool){
	size := len(arr)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if more(arr[j], arr[j+1]) {
				// swapping
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func main() {
	data := []int{9, 2, 3, 4, 5, 6}
	BubbleSort(data, func(a int, b int)bool{
		if a > b {
			return true
		}
		return false
	})

	fmt.Println(data)

	dataS := []string{"b", "d", "c", "f"}
	bubbleSort(dataS, func(a string, b string)bool{
		return a > b
	})
	fmt.Println(dataS)
}
