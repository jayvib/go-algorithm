package main

import "fmt"

func main() {
	arrays := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(incSort(arrays))
	fmt.Println(decSort(arrays))
}

func incSort(a []int) []int {
	o := a
	for j := 1; j < len(o); j++ {
		key := o[j]
		i := j - 1
		for i >= 0 && o[i] > key {
			o[i + 1] = o[i]
			i--
		}
		o[i + 1] = key
	}
	return o
}

func decSort(a []int) []int {
	o := a
	for j := 1; j < len(o); j++ {
		key := o[j]
		i := j - 1
		for i >= 0 && o[i] < key {
			o[i + 1] = o[i]
			i--
		}
		o[i + 1] = key
	}
	 
	return o
}