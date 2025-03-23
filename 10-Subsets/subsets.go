package subsets

// Subsets generates all possible subsets of the given array using an iterative approach (BFS)
// Time Complexity: O(N * 2^N) where N is the number of elements in nums
// Space Complexity: O(2^N) for storing all subsets
func Subsets(nums []int) [][]int {
	result := [][]int{{}} // Start with an empty set

	for _, num := range nums {
		// For each existing subset, create a new subset by adding the current number
		n := len(result)
		for i := 0; i < n; i++ {
			// Create a new subset by copying the existing subset and adding the current number
			newSet := make([]int, len(result[i]))
			copy(newSet, result[i])
			newSet = append(newSet, num)
			result = append(result, newSet)
		}
	}

	return result
}

// SubsetsRecursive generates all possible subsets of the given array using a recursive approach (DFS/Backtracking)
// Time Complexity: O(N * 2^N) where N is the number of elements in nums
// Space Complexity: O(2^N) for storing all subsets
func SubsetsRecursive(nums []int) [][]int {
	result := [][]int{}
	currentSet := []int{}
	generateSubsets(nums, 0, currentSet, &result)
	return result
}

// generateSubsets is a helper function that recursively generates all subsets
func generateSubsets(nums []int, index int, currentSet []int, result *[][]int) {
	// Add the current subset to the result
	subset := make([]int, len(currentSet))
	copy(subset, currentSet)
	*result = append(*result, subset)

	// Generate subsets by including elements from the current index onwards
	for i := index; i < len(nums); i++ {
		// Include the current element
		currentSet = append(currentSet, nums[i])

		// Recursively generate subsets with the current element included
		generateSubsets(nums, i+1, currentSet, result)

		// Backtrack by removing the current element
		currentSet = currentSet[:len(currentSet)-1]
	}
}

// Example usage:
//
// Input: nums = [1,2,3]
// Output: [[], [1], [2], [1,2], [3], [1,3], [2,3], [1,2,3]]
// Explanation: This is all possible subsets of the array [1,2,3]
//
// Input: nums = [0]
// Output: [[], [0]]
// Explanation: This is all possible subsets of the array [0]
