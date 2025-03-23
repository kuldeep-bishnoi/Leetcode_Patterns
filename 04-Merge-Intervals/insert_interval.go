package mergeintervals

// InsertInterval inserts a new interval into a list of non-overlapping intervals and merges if necessary
// Time Complexity: O(n) where n is the number of intervals
// Space Complexity: O(n) for storing the result
func InsertInterval(intervals []Interval, newInterval Interval) []Interval {
	result := []Interval{}
	i := 0
	n := len(intervals)

	// Add all intervals that come before newInterval
	for i < n && intervals[i].End < newInterval.Start {
		result = append(result, intervals[i])
		i++
	}

	// Merge all overlapping intervals
	for i < n && intervals[i].Start <= newInterval.End {
		newInterval.Start = min(newInterval.Start, intervals[i].Start)
		newInterval.End = max(newInterval.End, intervals[i].End)
		i++
	}

	// Add the merged interval
	result = append(result, newInterval)

	// Add all intervals that come after newInterval
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

// Helper functions to find min and max
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example cases:
// Input: Intervals = [[1,3], [6,9]], newInterval = [2,5]
// Output: [[1,5], [6,9]]
// Explanation: The new interval [2,5] overlaps with [1,3], so they are merged into [1,5]

// Input: Intervals = [[1,2], [3,5], [6,7], [8,10], [12,16]], newInterval = [4,8]
// Output: [[1,2], [3,10], [12,16]]
// Explanation: Intervals [3,5], [6,7], [8,10] overlap with the new interval [4,8], so they are merged into [3,10]

// Visualization:
//
// Intervals:   |--|  |---|  |--|  |---|  |-----|
// NewInterval:          |-----|
// Result:      |--|  |-------|  |-----|
