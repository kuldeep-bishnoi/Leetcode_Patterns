package subsets

import "sort"

// SubsetsWithDuplicates generates all possible subsets of the given array that may contain duplicates
// Time Complexity: O(N * 2^N) where N is the number of elements in nums
// Space Complexity: O(2^N) for storing all subsets
func SubsetsWithDuplicates(nums []int) [][]int {
	// Sort the input to handle duplicates properly
	sort.Ints(nums)

	result := [][]int{{}} // Start with an empty set
	startIndex, endIndex := 0, 0

	for i := 0; i < len(nums); i++ {
		startIndex = 0

		// If current element is a duplicate, we only add it to the subsets created in the previous step
		if i > 0 && nums[i] == nums[i-1] {
			startIndex = endIndex + 1
		}

		endIndex = len(result) - 1

		// For each existing subset, create a new subset by adding the current number
		n := len(result)
		for j := startIndex; j < n; j++ {
			// Create a new subset by copying the existing subset and adding the current number
			newSet := make([]int, len(result[j]))
			copy(newSet, result[j])
			newSet = append(newSet, nums[i])
			result = append(result, newSet)
		}
	}

	return result
}

// SubsetsWithDuplicatesRecursive generates all possible subsets of the given array with duplicates using backtracking
// Time Complexity: O(N * 2^N) where N is the number of elements in nums
// Space Complexity: O(2^N) for storing all subsets
func SubsetsWithDuplicatesRecursive(nums []int) [][]int {
	// Sort the input to handle duplicates properly
	sort.Ints(nums)

	result := [][]int{}
	currentSet := []int{}
	generateSubsetsWithDuplicates(nums, 0, currentSet, &result)
	return result
}

// generateSubsetsWithDuplicates is a helper function that recursively generates all subsets
// with handling for duplicate elements
func generateSubsetsWithDuplicates(nums []int, index int, currentSet []int, result *[][]int) {
	// Add the current subset to the result
	subset := make([]int, len(currentSet))
	copy(subset, currentSet)
	*result = append(*result, subset)

	for i := index; i < len(nums); i++ {
		// Skip duplicates
		if i > index && nums[i] == nums[i-1] {
			continue
		}

		// Include the current element
		currentSet = append(currentSet, nums[i])

		// Recursively generate subsets with the current element included
		generateSubsetsWithDuplicates(nums, i+1, currentSet, result)

		// Backtrack by removing the current element
		currentSet = currentSet[:len(currentSet)-1]
	}
}

// Example usage:
//
// Input: nums = [1,2,2]
// Output: [[], [1], [2], [1,2], [2,2], [1,2,2]]
// Explanation: The set [2] appears only once even though 2 appears twice in the input.
//
// Input: nums = [1,1,2,2]
// Output: [[], [1], [1,1], [2], [1,2], [1,1,2], [2,2], [1,2,2], [1,1,2,2]]
// Explanation: We need to generate all unique subsets.
