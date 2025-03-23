package cyclicsort

// FindMissingNumber finds the missing number in an array containing n distinct numbers in the range [0, n]
// Time Complexity: O(n) where n is the length of the array
// Space Complexity: O(1)
func FindMissingNumber(nums []int) int {
	i, n := 0, len(nums)

	// Cyclic sort (note that array contains 0 to n, with one number missing)
	for i < n {
		// Skip if the number is n or already at the correct position
		if nums[i] < n && nums[i] != i {
			// Swap with the correct position
			correctPos := nums[i]
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		} else {
			// If the number is out of range (equal to n) or already at the correct position
			i++
		}
	}

	// Find the first missing number
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}

	// If all numbers from 0 to n-1 are present, the missing number is n
	return n
}

// Example cases:
// Input: [4, 0, 3, 1]
// Output: 2
// Explanation: After cyclic sort, the array becomes [0, 1, ?, 3, 4], and the missing number is 2

// Input: [8, 3, 5, 2, 4, 6, 0, 1]
// Output: 7
// Explanation: After cyclic sort, the array becomes [0, 1, 2, 3, 4, 5, 6, ?], and the missing number is 7

// Visualization:
// Original: [4, 0, 3, 1]
// Swap 4 with element at index 4 (out of bounds, skip): [4, 0, 3, 1]
// Skip 0 (already at correct position): [4, 0, 3, 1]
// Skip 3 (already at correct position): [4, 0, 3, 1]
// Swap 1 with element at index 1: [4, 1, 3, 0]
// Swap 4 with element at index 4 (out of bounds, skip): [4, 1, 3, 0]
// Skip 1 (already at correct position): [4, 1, 3, 0]
// Skip 3 (already at correct position): [4, 1, 3, 0]
// Skip 0 (already at correct position): [4, 1, 3, 0]
// After sorting, the array is [0, 1, ?, 3, 4], and the missing number is 2
