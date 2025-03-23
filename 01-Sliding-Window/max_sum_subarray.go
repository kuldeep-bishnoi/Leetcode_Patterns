package slidingwindow

// MaxSumSubarrayOfSizeK finds the maximum sum of any contiguous subarray of size 'k'
// Time Complexity: O(n) where n is the length of the array
// Space Complexity: O(1)
func MaxSumSubarrayOfSizeK(arr []int, k int) int {
	n := len(arr)
	if n < k {
		return -1 // Invalid case
	}

	// Calculate sum of first window of size k
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	maxSum := windowSum

	// Slide the window and calculate the maximum sum
	for i := k; i < n; i++ {
		// Add the current element and subtract the element that's no longer in the window
		windowSum = windowSum + arr[i] - arr[i-k]

		// Update maximum sum if needed
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

// Example cases:
// Input: [2, 1, 5, 1, 3, 2], k=3
// Output: 9
// Explanation: Subarray with maximum sum is [5, 1, 3]

// Input: [2, 3, 4, 1, 5], k=2
// Output: 7
// Explanation: Subarray with maximum sum is [3, 4]
