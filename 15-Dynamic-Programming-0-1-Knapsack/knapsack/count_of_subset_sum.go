package knapsack

// CountSubsetSum counts the number of subsets of the given array
// that sum up to the target sum
// Time Complexity: O(N * S) where N is the number of elements and S is the target sum
// Space Complexity: O(S) with the optimized 1D DP approach
func CountSubsetSum(nums []int, sum int) int {
	n := len(nums)

	// Edge cases
	if sum < 0 {
		return 0
	}

	if n == 0 {
		if sum == 0 {
			return 1 // Empty set sums to 0
		}
		return 0
	}

	// DP array: dp[j] represents count of subsets that sum to j
	dp := make([]int, sum+1)
	dp[0] = 1 // Empty subset can sum to 0

	// Fill the DP array
	for i := 0; i < n; i++ {
		// Process in reverse to avoid counting subsets multiple times
		for j := sum; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[sum]
}

// CountSubsetSumDP uses a 2D dynamic programming approach
// Time Complexity: O(N * S) where N is the number of elements and S is the target sum
// Space Complexity: O(N * S) for the 2D DP table
func CountSubsetSumDP(nums []int, sum int) int {
	n := len(nums)

	// Edge cases
	if sum < 0 {
		return 0
	}

	if n == 0 {
		if sum == 0 {
			return 1 // Empty set sums to 0
		}
		return 0
	}

	// DP table: dp[i][j] represents count of subsets of first i elements that sum to j
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, sum+1)
		dp[i][0] = 1 // Empty set sums to 0
	}

	// Fill the DP table
	for i := 1; i <= n; i++ {
		for j := 0; j <= sum; j++ {
			// Skip the current element (inherit from previous state)
			dp[i][j] = dp[i-1][j]

			// Include the current element if it's not bigger than the target sum
			if nums[i-1] <= j {
				dp[i][j] += dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][sum]
}

// CountSubsetSumRecursive counts subsets with target sum using recursion with memoization
// Time Complexity: O(N * S) where N is the number of elements and S is the target sum
// Space Complexity: O(N * S) for the memoization table + O(N) for recursion stack
func CountSubsetSumRecursive(nums []int, sum int) int {
	n := len(nums)

	// Edge cases
	if sum < 0 {
		return 0
	}

	if n == 0 {
		if sum == 0 {
			return 1 // Empty set sums to 0
		}
		return 0
	}

	// Initialize memoization table with -1 (not calculated)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, sum+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return countSubsetSumHelper(nums, sum, 0, memo)
}

// countSubsetSumHelper is a recursive helper function with memoization
func countSubsetSumHelper(nums []int, sum int, currentIndex int, memo [][]int) int {
	// Base cases
	if sum == 0 {
		return 1
	}

	if currentIndex >= len(nums) || sum < 0 {
		return 0
	}

	// Check if we have already computed this state
	if memo[currentIndex][sum] != -1 {
		return memo[currentIndex][sum]
	}

	// Decision 1: Include the current element
	count1 := 0
	if nums[currentIndex] <= sum {
		count1 = countSubsetSumHelper(nums, sum-nums[currentIndex], currentIndex+1, memo)
	}

	// Decision 2: Skip the current element
	count2 := countSubsetSumHelper(nums, sum, currentIndex+1, memo)

	// Total count is the sum of both decisions
	memo[currentIndex][sum] = count1 + count2
	return memo[currentIndex][sum]
}

// FindSubsetsWithSum finds all the subsets that sum to the target
// Time Complexity: O(N * S * 2^N) in worst case where N is the number of elements and S is the target sum
// Space Complexity: O(2^N) for storing all potential subsets
func FindSubsetsWithSum(nums []int, sum int) [][]int {
	result := [][]int{}
	currentSubset := []int{}

	// Start the backtracking process
	findSubsetsWithSumHelper(nums, sum, 0, currentSubset, &result)

	return result
}

// findSubsetsWithSumHelper is a recursive backtracking helper function
func findSubsetsWithSumHelper(nums []int, remaining int, currentIndex int, currentSubset []int, result *[][]int) {
	// Base case: we found a valid subset
	if remaining == 0 {
		// Make a copy of the current subset and add it to the result
		subset := make([]int, len(currentSubset))
		copy(subset, currentSubset)
		*result = append(*result, subset)
		return
	}

	// Base case: we've exhausted the array or remaining sum is negative
	if currentIndex >= len(nums) || remaining < 0 {
		return
	}

	// Decision 1: Include the current element
	currentSubset = append(currentSubset, nums[currentIndex])
	findSubsetsWithSumHelper(nums, remaining-nums[currentIndex], currentIndex+1, currentSubset, result)

	// Backtrack: Remove the last element to try other combinations
	currentSubset = currentSubset[:len(currentSubset)-1]

	// Decision 2: Skip the current element
	findSubsetsWithSumHelper(nums, remaining, currentIndex+1, currentSubset, result)
}

// CountSubsetsWithTargetDifference counts subsets with a given difference
// This problem can be reduced to counting subsets with a specific sum
// Time Complexity: O(N * S) where N is the number of elements and S is sum of all elements
// Space Complexity: O(S) with the optimized approach
func CountSubsetsWithTargetDifference(nums []int, diff int) int {
	// Calculate the total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If totalSum + diff is odd or diff > totalSum, no valid partition exists
	if (totalSum+diff)%2 != 0 || diff > totalSum {
		return 0
	}

	// We need to find count of subsets with sum = (totalSum + diff) / 2
	sum := (totalSum + diff) / 2

	return CountSubsetSum(nums, sum)
}

// Example usage:
//
// Input: nums = [1, 1, 2, 3], sum = 4
// Output: 3 (The three subsets are: [1, 3], [1, 3], [1, 1, 2])
//
// Input: nums = [1, 2, 7, 1, 5], sum = 9
// Output: 3 (The three subsets are: [2, 7], [1, 1, 7], [1, 2, 1, 5])
//
// Input: nums = [1, 1, 1, 1], diff = 0
// Output: 6 (The subsets with difference 0 are: {}, {1,1,1,1}, {1}, {1,1,1}, {1,1}, {1,1})
