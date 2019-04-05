package review

func BinarySearch(data []int, value int) bool {
	size := len(data)
	low := 0
	mid := 0
	high := size - 1

	for low <= high {
		mid = low + (high-low)/2 // get the middle index
		if  data[mid] == value { // then check if matched
			return true
		} else if data[mid] < value { // if the value if greater than the current middle index then search in the upper bound
			low = mid + 1
		} else { // else search in the lower bound
			high = mid - 1
		}
	}
	return false
}