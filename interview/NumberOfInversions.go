package interview

//inversions gets the count of all inversions in the values input
//It's based on merge-sort algorithm. While sorting the values,
//if two values need to be swap, then it's means a inversion.
func inversions(vals []int) int {
	if len(vals) < 2 {
		return 0
	}

	merged := make([]int, len(vals))
	mid := len(vals) / 2
	left := vals[:mid]
	right := vals[mid:]
	//divdie all the elements into two parts, and recursivly find inversions inside those two parts
	cnt := inversions(left)
	cnt += inversions(right)

	//compare two parts, and find inversions cross two parts
	var i, j, k int
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			merged[k] = right[j]
			k++

			//print all inversions
			// for icopy := i; icopy < len(left); icopy++ {
			// 	fmt.Printf("(%d,%d)\n", left[icopy], right[j])
			// }

			cnt += len(left) - i // elments from i in left are all bigger than right[j], as left is already sorted
			//forward right part
			j++
		} else {
			merged[k] = left[i]
			k++
			//forward left part
			i++
		}
	}
	//copy elements not merged in left or right part
	copy(merged[k:], right[j:])
	copy(merged[k:], left[i:])

	//save merge-sort result
	copy(vals, merged)

	return cnt
}
