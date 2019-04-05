package main

import (
	"fmt"
	"github.com/golang-collections/collections/set"
	"sort"
)

func printRepeating(data []int) {
	size := len(data)
	for i := 0; i < size; i++ {
		for j := i+1; j < size;j++ {
			if data[i] == data[j] {
				fmt.Println(" ", data[i])
			}
		}
	}
	fmt.Println()
}

func printRepeating2(data []int) {
	size := len(data)
	sort.Ints(data)
	for i :=  1; i < size; i++ {
		if data[i] == data[i-1] {
			fmt.Println(" ", data[i])
		}
	}
	fmt.Println()
}

func printRepeating3(data []int) {
	s := new(set.Set)
	size := len(data)
	fmt.Print("Repeating elements are: ")
	for i := 0; i < size; i++ {
		if s.Has(data[i]) {
			fmt.Println(" ", data[i])
		} else {
			s.Insert(data[i])
		}
	}
}
