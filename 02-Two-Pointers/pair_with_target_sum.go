package twopointers

// PairWithTargetSum finds a pair of elements in a sorted array that sum to a target value
// Returns the indices of the pair
// Time Complexity: O(n) where n is the length of the array
// Space Complexity: O(1)
func PairWithTargetSum(arr []int, targetSum int) []int {
	left, right := 0, len(arr)-1

	for left < right {
		currentSum := arr[left] + arr[right]

		if currentSum == targetSum {
			return []int{left, right} // Found the pair
		}

		if currentSum < targetSum {
			left++ // Need a pair with a larger sum
		} else {
			right-- // Need a pair with a smaller sum
		}
	}

	return []int{-1, -1} // No pair found
}

// AlternativePairWithTargetSum finds a pair of elements that sum to a target using a hash map
// This approach is more efficient for unsorted arrays
// Time Complexity: O(n)
// Space Complexity: O(n) for the hash map
func AlternativePairWithTargetSum(arr []int, targetSum int) []int {
	// Create a hash map to store numbers and their indices
	nums := make(map[int]int)

	// Iterate through the array
	for i, num := range arr {
		// Check if we've seen the complement before
		if j, exists := nums[targetSum-num]; exists {
			return []int{j, i}
		}
		// Store the current number and its index
		nums[num] = i
	}

	return []int{-1, -1} // No pair found
}

// Example cases:
// Input: [1, 2, 3, 4, 6], targetSum=6
// Output: [1, 3]
// Explanation: The numbers at index 1 (2) and 3 (4) add up to 6

// Input: [2, 5, 9, 11], targetSum=11
// Output: [0, 2]
// Explanation: The numbers at index 0 (2) and 2 (9) add up to 11
