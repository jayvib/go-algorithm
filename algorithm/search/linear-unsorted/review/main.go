package review


// Time Complexity is O(n)
// Space Complexity is O(1)
//
// The downside of this is when the item that ur looking for
// is in last item.
func linearSearchUnsorted(data []int, value int) bool {
	size := len(data)
	for i :=0; i < size; i++ {
		if data[i] == value {
			return true
		}
	}
	return false
}


