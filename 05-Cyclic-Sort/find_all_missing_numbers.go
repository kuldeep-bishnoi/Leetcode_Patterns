package cyclicsort

// FindAllMissingNumbers finds all missing numbers in an array containing numbers from 1 to n
// Time Complexity: O(n) where n is the length of the array
// Space Complexity: O(1) excluding the output array
func FindAllMissingNumbers(nums []int) []int {
	i := 0
	n := len(nums)

	// Cyclic sort (place each number at index (number-1))
	for i < n {
		// Correct position for the current number
		correctPos := nums[i] - 1

		// If the number is within range and not at its correct position, swap
		if nums[i] > 0 && nums[i] <= n && nums[i] != nums[correctPos] {
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		} else {
			// If the number is out of range or already at the correct position
			i++
		}
	}

	// Find all missing numbers
	missingNumbers := []int{}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			missingNumbers = append(missingNumbers, i+1)
		}
	}

	return missingNumbers
}

// Example cases:
// Input: [2, 3, 1, 8, 2, 3, 5, 1]
// Output: [4, 6, 7]
// Explanation: After cyclic sort, the array becomes [1, 2, 3, ?, 5, ?, ?, 8], and the missing numbers are 4, 6, and 7

// Input: [2, 4, 1, 2]
// Output: [3]
// Explanation: After cyclic sort, the array becomes [1, 2, ?, 4], and the missing number is 3

// Input: [2, 3, 2, 1]
// Output: [4]
// Explanation: After cyclic sort, the array becomes [1, 2, 3, ?], and the missing number is 4

// Visualization:
// Original: [2, 3, 1, 8, 2, 3, 5, 1]
// After cyclic sort: [1, 2, 3, 8, 5, 3, 2, 1]
// At index 3, we expect 4 but found 8, so 4 is missing
// At index 5, we expect 6 but found 3, so 6 is missing
// At index 6, we expect 7 but found 2, so 7 is missing
