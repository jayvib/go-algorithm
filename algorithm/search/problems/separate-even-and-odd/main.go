package main

import "fmt"

// Problem: Given a list of even and odd numbers, write a program to separte even numbers from the
// odd numbers.

func separateEvenAndOdd(data []int) {
	size := len(data)
	left := 0
	right := size-1
	for left < right {
		switch {
		case data[left]%2 == 0: // find the odd
			left++
		case data[right]%2 == 1: // find the even
			right--
		default:
			data[left], data[right] = data[right], data[left] // swap the odd and even position
			left++
			right--
		}
	}
}

func separteEvenAndOdd(data []int) {
	size := len(data)
	left := 0
	right := size-1
	for left < right {
		switch {
		case data[left]%2 == 0: // find odd
			left++
		case data[right]%2 == 1: // find even
			right--
		default:
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
	}
}

func main() {
	data := []int{99, 12, 34, 45, 9, 8, 90, 3, 5, 4, 2, 1, 67,38}
	separateEvenAndOdd(data)
	fmt.Println(data)
}
