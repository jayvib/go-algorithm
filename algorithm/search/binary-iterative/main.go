package main

func main() {

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
