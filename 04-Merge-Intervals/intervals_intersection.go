package mergeintervals

// IntervalsIntersection finds the intersection of two lists of intervals
// Time Complexity: O(m + n) where m and n are the lengths of the interval lists
// Space Complexity: O(min(m, n)) in the worst case for storing the result
func IntervalsIntersection(intervals1 []Interval, intervals2 []Interval) []Interval {
	result := []Interval{}
	i, j := 0, 0 // Pointers for each interval list

	// Compare intervals from both lists
	for i < len(intervals1) && j < len(intervals2) {
		// Check if intervals overlap
		// For overlap: start of one interval <= end of another interval && vice versa
		a := intervals1[i]
		b := intervals2[j]

		// Find the overlapping interval
		aOverlapsB := a.Start <= b.End && b.Start <= a.End

		if aOverlapsB {
			// Add the intersection to the result
			overlapStart := max(a.Start, b.Start)
			overlapEnd := min(a.End, b.End)
			result = append(result, Interval{Start: overlapStart, End: overlapEnd})
		}

		// Move the pointer of the interval that ends earlier
		if a.End < b.End {
			i++
		} else {
			j++
		}
	}

	return result
}

// Example cases:
// Input:
// A = [[1, 3], [5, 6], [7, 9]]
// B = [[2, 3], [5, 7]]
// Output: [[2, 3], [5, 6], [7, 7]]
// Explanation:
// [1, 3] intersects with [2, 3] at [2, 3]
// [5, 6] intersects with [5, 7] at [5, 6]
// [7, 9] intersects with [5, 7] at [7, 7]

// Input:
// A = [[1, 3], [5, 7], [9, 12]]
// B = [[5, 10]]
// Output: [[5, 7], [9, 10]]
// Explanation:
// [5, 7] and [9, 12] from A intersect with [5, 10] from B

// Visualization:
// A: |---|   |---|   |---|
// B:    |---|    |------|
// R:    |-|   |--|   |-|
