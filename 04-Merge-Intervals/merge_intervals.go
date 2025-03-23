package mergeintervals

import "sort"

// Interval represents a time interval with start and end times
type Interval struct {
	Start int
	End   int
}

// MergeIntervals merges all overlapping intervals
// Time Complexity: O(n log n) where n is the number of intervals
// Space Complexity: O(n) for storing the result
func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort intervals based on start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	merged := []Interval{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := &merged[len(merged)-1]

		// If current interval overlaps with the last merged interval, merge them
		if current.Start <= lastMerged.End {
			// Update end of the last merged interval if necessary
			if current.End > lastMerged.End {
				lastMerged.End = current.End
			}
		} else {
			// No overlap, add current interval to result
			merged = append(merged, current)
		}
	}

	return merged
}

// Example cases:
// Input: [[1,3], [2,6], [8,10], [15,18]]
// Output: [[1,6], [8,10], [15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, they are merged into [1,6]

// Input: [[1,4], [4,5]]
// Output: [[1,5]]
// Explanation: Since intervals [1,4] and [4,5] have a common point at 4, they are merged into [1,5]

// Visualization:
//
// Interval 1: |--------|
// Interval 2:     |---------|
// Merged:     |------------|
//
// Interval 3:                 |-----|
// Interval 4:                           |-----|
// Final:      |------------|  |-----|  |-----|
