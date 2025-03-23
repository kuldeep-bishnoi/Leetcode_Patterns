package cyclicsort

// FindDuplicate finds the duplicate number in an array containing n+1 numbers where each number is between 1 and n
// Time Complexity: O(n) where n is the length of the array
// Space Complexity: O(1)
func FindDuplicate(nums []int) int {
	i := 0
	n := len(nums)

	// Cyclic sort
	for i < n {
		// For this problem, we know duplicate exists, so the correct position for nums[i] is (nums[i]-1)
		if nums[i] != i+1 {
			correctPos := nums[i] - 1

			// If the element at the correct position is already the correct number,
			// then nums[i] is the duplicate
			if nums[i] == nums[correctPos] {
				return nums[i]
			}

			// Otherwise, swap to move the current number to its correct position
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		} else {
			// Number is already at correct position
			i++
		}
	}

	// This line should never be reached if the problem statement is correct
	// (there's always exactly one duplicate)
	return -1
}

// Alternative approach using Floyd's Tortoise and Hare (Cycle Detection)
// This is more efficient as it doesn't modify the array and uses the Fast & Slow pointers pattern
func FindDuplicateFloyd(nums []int) int {
	// Find the intersection point of the slow and fast pointers
	slow, fast := nums[0], nums[0]

	for {
		slow = nums[slow]       // Move one step
		fast = nums[nums[fast]] // Move two steps
		if slow == fast {
			break
		}
	}

	// Find the "entrance" to the cycle (the duplicate number)
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

// Example cases:
// Input: [1, 3, 4, 2, 2]
// Output: 2
// Explanation: The array contains numbers 1 to 4, and the duplicate is 2

// Input: [3, 1, 3, 4, 2]
// Output: 3
// Explanation: The array contains numbers 1 to 4, and the duplicate is 3

// Visualization of the cyclic sort approach:
// Original: [1, 3, 4, 2, 2]
// Skip 1 (already at correct position): [1, 3, 4, 2, 2]
// Swap 3 with element at index 2: [1, 3, 2, 2, 4]
// Skip 3 (already at correct position): [1, 3, 2, 2, 4]
// Swap 2 with element at index 1: [1, 2, 3, 2, 4]
// Skip 2, 3, 4 (already at correct positions): [1, 2, 3, 2, 4]
// At this point, we find that the number 2 appears twice, at indices 1 and 3
