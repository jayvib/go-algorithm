package main

func BinarySearchRecursive(data []int, low int, high int, value int) bool {
	if low > high {
		return false
	}

	mid := low + (high - low)/2

	if data[mid] == value {
		return true
	} else if data[mid] < value {
		return BinarySearchRecursive(data, mid+1, high, value)
	} else {
		return BinarySearchRecursive(data, low, mid-1, value)
	}
}
