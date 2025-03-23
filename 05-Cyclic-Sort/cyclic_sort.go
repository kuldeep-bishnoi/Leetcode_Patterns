package cyclicsort

// CyclicSort sorts an array containing numbers from 1 to n
// Time Complexity: O(n) where n is the length of the array
// Space Complexity: O(1) as we sort in-place
func CyclicSort(nums []int) {
	i := 0
	for i < len(nums) {
		// The correct position for the current number is (nums[i] - 1)
		// For example, number 1 should be at index 0, number 2 at index 1, and so on
		correctPos := nums[i] - 1

		// If the number is not at its correct position, swap
		if nums[i] != nums[correctPos] {
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		} else {
			// If the number is already at the correct position, move to the next element
			i++
		}
	}
}

// Example cases:
// Input: [3, 1, 5, 4, 2]
// Output: [1, 2, 3, 4, 5]
// Explanation: After cyclic sort, each number i is placed at index (i-1)

// Input: [2, 6, 4, 3, 1, 5]
// Output: [1, 2, 3, 4, 5, 6]
// Explanation: Each number i is placed at index (i-1)

// Visualization:
// Original: [3, 1, 5, 4, 2]
// Swap 3 with element at index 2: [5, 1, 3, 4, 2]
// Swap 5 with element at index 4: [2, 1, 3, 4, 5]
// Swap 2 with element at index 1: [1, 2, 3, 4, 5]
// Now all elements are at their correct positions
