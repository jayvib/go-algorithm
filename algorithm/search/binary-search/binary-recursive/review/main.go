package review

func BinarySearchRecursive(data []int, high int, low int, value int) bool {
	if low > high {
		return false
	}
	mid := low + (high-low)/2
	if data[mid] == value {
		return true
	} else if data[mid] < value {
		// find in the upper band
		return BinarySearchRecursive(data, high, mid+1, value)
	} else if data[mid] > value {
		// search in the lower bound
		return BinarySearchRecursive(data, mid-1, low, value)
	}
}
