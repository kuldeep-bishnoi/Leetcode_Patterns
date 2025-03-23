package slidingwindow

import "math"

// SmallestSubarrayWithGivenSum finds the smallest subarray with a sum greater than or equal to 's'
// Time Complexity: O(n) where n is the length of the array
// Space Complexity: O(1)
func SmallestSubarrayWithGivenSum(arr []int, s int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	windowSum := 0
	minLength := math.MaxInt32
	windowStart := 0

	// Try to extend the range [windowStart, windowEnd]
	for windowEnd := 0; windowEnd < n; windowEnd++ {
		// Add the next element to the window
		windowSum += arr[windowEnd]

		// Shrink the window as small as possible until the windowSum is less than s
		for windowSum >= s {
			// Update the minimum length
			currentLength := windowEnd - windowStart + 1
			minLength = int(math.Min(float64(minLength), float64(currentLength)))

			// Remove the leftmost element from the window
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	// If we didn't find any subarray
	if minLength == math.MaxInt32 {
		return 0
	}

	return minLength
}

// Example cases:
// Input: [2, 1, 5, 2, 3, 2], s=7
// Output: 2
// Explanation: The smallest subarray with a sum greater than or equal to '7' is [5, 2]

// Input: [2, 1, 5, 2, 8], s=7
// Output: 1
// Explanation: The smallest subarray with a sum greater than or equal to '7' is [8]

// Input: [3, 4, 1, 1, 6], s=8
// Output: 3
// Explanation: Smallest subarrays with a sum greater than or equal to '8' are [3, 4, 1] or [1, 1, 6]
