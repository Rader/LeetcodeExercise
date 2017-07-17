package LeetcodeExercise

import (
	"sort"
)

//   Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func findRightInterval(intervals []Interval) []int {
	idxMap := make(map[int]int)
	//keep the original position of each interval
	for i, itv := range intervals {
		idxMap[itv.Start] = i
	}
	//sort by start point
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	result := make([]int, len(intervals))
	for _, itv := range intervals {
		//index of an interval before sort
		oldIdx := idxMap[itv.Start]
		found := binarySearch(intervals, itv.End)
		if found != -1 {
			result[oldIdx] = idxMap[intervals[found].Start]
		} else {
			result[oldIdx] = -1
		}
	}
	return result
}

//find the min index of an interval whose start point >= end
func binarySearch(intervals []Interval, point int) int {
	idx := -1 //default, not found
	begin := 0
	end := len(intervals) - 1
	for begin <= end {
		if begin == end { //last try, should end loop
			if intervals[begin].Start >= point {
				idx = begin
			}
			break
		}

		mid := (begin + end) / 2
		if intervals[mid].Start < point {
			//search in right
			begin = mid + 1
		} else { //search in left
			end = mid
		}

	}
	return idx
}
