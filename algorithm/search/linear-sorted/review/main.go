package review


// Time Complexity: O(n)
// Space Complexity: O(1)
//
// Worst Case: When the item that your looking for is in the
// last position
func linearSearchSorted(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if data[i] == value {
			return true
		} else if data[i] > value {
			return false
		}
	}
	return false
}
