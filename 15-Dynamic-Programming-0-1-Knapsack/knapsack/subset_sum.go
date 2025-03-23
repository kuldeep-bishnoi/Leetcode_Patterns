package knapsack

// SubsetSumRecursive determines if there is a subset of the given array
// that adds up to the target sum using recursion with memoization
// Time Complexity: O(N * S) where N is the number of elements and S is the target sum
// Space Complexity: O(N * S) for the memoization table + O(N) for recursion stack
func SubsetSumRecursive(nums []int, sum int) bool {
	n := len(nums)
	if n == 0 || sum < 0 {
		return false
	}

	// Initialize memoization table
	// -1: not calculated, 0: false, 1: true
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, sum+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return subsetSumHelper(nums, sum, 0, memo)
}

// subsetSumHelper is a recursive helper function with memoization
func subsetSumHelper(nums []int, sum int, currentIndex int, memo [][]int) bool {
	// Base cases
	if sum == 0 {
		return true
	}

	if currentIndex >= len(nums) || sum < 0 {
		return false
	}

	// Check if we have already computed this state
	if memo[currentIndex][sum] != -1 {
		return memo[currentIndex][sum] == 1
	}

	// Decision 1: Include the current element
	if nums[currentIndex] <= sum {
		if subsetSumHelper(nums, sum-nums[currentIndex], currentIndex+1, memo) {
			memo[currentIndex][sum] = 1
			return true
		}
	}

	// Decision 2: Skip the current element
	result := subsetSumHelper(nums, sum, currentIndex+1, memo)
	if result {
		memo[currentIndex][sum] = 1
	} else {
		memo[currentIndex][sum] = 0
	}

	return result
}

// SubsetSumDP determines if there is a subset of the given array
// that adds up to the target sum using dynamic programming
// Time Complexity: O(N * S) where N is the number of elements and S is the target sum
// Space Complexity: O(N * S) for the DP table
func SubsetSumDP(nums []int, sum int) bool {
	n := len(nums)
	if n == 0 || sum < 0 {
		return false
	}

	// DP table: dp[i][j] represents whether a subset of the first i elements can sum to j
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		// There is always a subset (empty set) that sums to 0
		dp[i][0] = true
	}

	// Fill the DP table
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			// Skip the current element (inherit from previous state)
			dp[i][j] = dp[i-1][j]

			// Include the current element if it's not bigger than the target sum
			if nums[i-1] <= j {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][sum]
}

// SubsetSumDP1D determines if there is a subset of the given array
// that adds up to the target sum using dynamic programming with a 1D array
// Time Complexity: O(N * S) where N is the number of elements and S is the target sum
// Space Complexity: O(S) for the DP array
func SubsetSumDP1D(nums []int, sum int) bool {
	n := len(nums)
	if n == 0 || sum < 0 {
		return false
	}

	// DP array: dp[j] represents whether a subset can sum to j
	dp := make([]bool, sum+1)
	dp[0] = true // Empty subset can sum to 0

	// Fill the DP array
	for i := 0; i < n; i++ {
		// Process in reverse to avoid using the updated value
		for j := sum; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[sum]
}

// FindSubsetSum not only determines if a subset exists but also returns the subset
// Time Complexity: O(N * S) where N is the number of elements and S is the target sum
// Space Complexity: O(N * S) for the DP table
func FindSubsetSum(nums []int, sum int) (bool, []int) {
	n := len(nums)
	if n == 0 || sum < 0 {
		return false, nil
	}

	// DP table
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true // Empty subset can sum to 0
	}

	// Fill the DP table
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			// Skip the current element
			dp[i][j] = dp[i-1][j]

			// Include the current element if it's not bigger than the target sum
			if nums[i-1] <= j {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	// Check if a subset exists
	if !dp[n][sum] {
		return false, nil
	}

	// Reconstruct the subset
	subset := []int{}
	currentSum := sum
	for i := n; i > 0; i-- {
		// If including this element gives us the current sum
		if !dp[i-1][currentSum] && currentSum >= nums[i-1] {
			subset = append(subset, nums[i-1])
			currentSum -= nums[i-1]
		}
	}

	// The subset will be in reverse order, so let's fix that
	for i, j := 0, len(subset)-1; i < j; i, j = i+1, j-1 {
		subset[i], subset[j] = subset[j], subset[i]
	}

	return true, subset
}

// FindAllSubsetSums finds all possible sums that can be formed by any subset of the array
// Time Complexity: O(N * sum(nums)) where N is the number of elements
// Space Complexity: O(sum(nums)) for the DP set
func FindAllSubsetSums(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{0} // Empty set sums to 0
	}

	// Calculate the total sum to determine the DP array size
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// DP set to track all possible sums
	// Using map as a set for simplicity
	possibleSums := make(map[int]bool)
	possibleSums[0] = true // Empty subset sums to 0

	// Process each number
	for _, num := range nums {
		// Create a copy of the current sums to avoid modifying while iterating
		newSums := make(map[int]bool)
		for sum := range possibleSums {
			newSums[sum] = true     // Keep the original sum
			newSums[sum+num] = true // Add the current number to the sum
		}
		// Update the possible sums
		possibleSums = newSums
	}

	// Convert the map keys to a slice
	result := make([]int, 0, len(possibleSums))
	for sum := range possibleSums {
		result = append(result, sum)
	}

	return result
}

// Example usage:
//
// Input: nums = [1, 2, 3, 7], sum = 6
// Output: true (subset [1, 2, 3] sums to 6)
//
// Input: nums = [1, 3, 4, 8], sum = 6
// Output: false (no subset sums to 6)
//
// Input: nums = [2, 3, 5], sum = 10
// Output: true (subset [2, 3, 5] sums to 10)
//
// Input: nums = [1, 3, 5, 7], getAllSums = true
// Output: [0, 1, 3, 4, 5, 6, 8, 9, 10, 11, 12, 16] (all possible subset sums)
