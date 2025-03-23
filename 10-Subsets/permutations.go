package subsets

// Permute generates all possible permutations of the given array
// Time Complexity: O(N!), where N is the number of elements in nums
// Space Complexity: O(N!) for storing all permutations
func Permute(nums []int) [][]int {
	result := [][]int{}
	generatePermutations(nums, 0, &result)
	return result
}

// generatePermutations recursively generates all permutations by swapping elements
func generatePermutations(nums []int, index int, result *[][]int) {
	// If we've reached the end of the array, add the current permutation to the result
	if index == len(nums) {
		permutation := make([]int, len(nums))
		copy(permutation, nums)
		*result = append(*result, permutation)
		return
	}
	
	// Try placing every element at the current position
	for i := index; i < len(nums); i++ {
		// Swap the current element with the element at the index position
		nums[i], nums[index] = nums[index], nums[i]
		
		// Recursively generate permutations for the next position
		generatePermutations(nums, index+1, result)
		
		// Backtrack by swapping back to the original order
		nums[i], nums[index] = nums[index], nums[i]
	}
}

// PermuteIterative generates all possible permutations of the given array using an iterative approach
// Time Complexity: O(N!), where N is the number of elements in nums
// Space Complexity: O(N!) for storing all permutations
func PermuteIterative(nums []int) [][]int {
	result := [][]int{}
	
	// Start with a single empty permutation
	result = append(result, []int{})
	
	for _, num := range nums {
		// Create a fresh result list for the current iteration
		nextResult := [][]int{}
		
		// For each existing permutation
		for _, perm := range result {
			// Create a new permutation by inserting the current number at every possible position
			for i := 0; i <= len(perm); i++ {
				// Make a copy of the permutation
				newPerm := make([]int, len(perm)+1)
				
				// Copy elements before the insertion point
				copy(newPerm[:i], perm[:i])
				// Insert the new element
				newPerm[i] = num
				// Copy elements after the insertion point
				copy(newPerm[i+1:], perm[i:])
				
				// Add the new permutation to the result
				nextResult = append(nextResult, newPerm)
			}
		}
		
		// Update the result with the new permutations
		result = nextResult
	}
	
	return result
}

// PermuteWithDuplicates generates all unique permutations when the input may contain duplicates
// Time Complexity: O(N!), where N is the number of elements in nums
// Space Complexity: O(N!) for storing all permutations
func PermuteWithDuplicates(nums []int) [][]int {
	result := [][]int{}
	
	// Sort the input to handle duplicates (optional but makes duplicate detection easier)
	// sort.Ints(nums)
	
	// Track which elements we've used
	used := make([]bool, len(nums))
	currentPerm := []int{}
	
	generateUniquePermutations(nums, used, currentPerm, &result)
	
	return result
}

// generateUniquePermutations recursively generates all unique permutations
func generateUniquePermutations(nums []int, used []bool, currentPerm []int, result *[][]int) {
	// If current permutation is complete, add it to the result
	if len(currentPerm) == len(nums) {
		permutation := make([]int, len(currentPerm))
		copy(permutation, currentPerm)
		*result = append(*result, permutation)
		return
	}
	
	// Use a map to track which values we've used at the current position
	// This ensures we don't generate duplicate permutations
	usedValues := make(map[int]bool)
	
	for i := 0; i < len(nums); i++ {
		// Skip if this element is already used in the current permutation
		if used[i] {
			continue
		}
		
		// Skip if we've already used this value at the current position
		if usedValues[nums[i]] {
			continue
		}
		
		// Mark this value as used at the current position
		usedValues[nums[i]] = true
		
		// Add the current element to the permutation
		used[i] = true
		currentPerm = append(currentPerm, nums[i])
		
		// Recursively generate permutations with the current element included
		generateUniquePermutations(nums, used, currentPerm, result)
		
		// Backtrack
		currentPerm = currentPerm[:len(currentPerm)-1]
		used[i] = false
	}
}

// Example usage:
//
// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// Explanation: All possible permutations of [1,2,3]
//
// Input: nums = [0,1]
// Output: [[0,1],[1,0]]
// Explanation: All possible permutations of [0,1]
//
// Input: nums = [1,1,2]
// Output: [[1,1,2],[1,2,1],[2,1,1]]
// Explanation: All unique permutations of [1,1,2] 