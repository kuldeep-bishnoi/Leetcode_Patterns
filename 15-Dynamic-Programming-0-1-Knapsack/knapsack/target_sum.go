package knapsack

import "strconv"

// FindTargetSumWays counts the number of ways to assign + and - signs to
// elements in the array to achieve the target sum
// Time Complexity: O(N * S) where N is the number of elements and S is the sum of all elements
// Space Complexity: O(S) with the optimized approach
func FindTargetSumWays(nums []int, target int) int {
	// Calculate the total sum of the array
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If target is out of range or (totalSum + target) is odd, no valid assignment exists
	if target > totalSum || target < -totalSum || (totalSum+target)%2 != 0 {
		return 0
	}

	// This problem can be reduced to finding count of subsets with a specific sum
	// Let P be the sum of elements with + sign and N be the sum of elements with - sign
	// We have: P + N = totalSum and P - N = target
	// Solving these equations: P = (totalSum + target) / 2
	sum := (totalSum + target) / 2

	// Now we need to find the count of subsets that sum up to 'sum'
	return CountSubsetSum(nums, sum)
}

// FindTargetSumWaysDP uses a 2D dynamic programming approach
// Time Complexity: O(N * (2*S + 1)) where N is the number of elements and S is the sum of all elements
// Space Complexity: O(N * (2*S + 1)) for the DP table
func FindTargetSumWaysDP(nums []int, target int) int {
	n := len(nums)

	// Calculate the total sum to determine the range of possible sums
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If target is out of range, no valid assignment exists
	if target > totalSum || target < -totalSum {
		return 0
	}

	// The range of possible sums is [-totalSum, totalSum]
	// We'll shift this to [0, 2*totalSum] to use as indices in our DP table
	offset := totalSum

	// DP table: dp[i][j] represents count of ways to get sum j using first i elements
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2*totalSum+1)
	}

	// Base case: there is one way to get sum 0 using 0 elements (empty set)
	dp[0][offset] = 1

	// Fill the DP table
	for i := 1; i <= n; i++ {
		for j := 0; j <= 2*totalSum; j++ {
			// Skip this sum if it's not achievable
			if dp[i-1][j] == 0 {
				continue
			}

			// Add current element with positive sign
			if j+nums[i-1] <= 2*totalSum {
				dp[i][j+nums[i-1]] += dp[i-1][j]
			}

			// Add current element with negative sign
			if j-nums[i-1] >= 0 {
				dp[i][j-nums[i-1]] += dp[i-1][j]
			}
		}
	}

	// Return the count of ways to get the target sum
	return dp[n][target+offset]
}

// FindTargetSumWaysRecursive uses recursion with memoization
// Time Complexity: O(N * (2*S + 1)) where N is the number of elements and S is the sum of all elements
// Space Complexity: O(N * (2*S + 1)) for the memoization table
func FindTargetSumWaysRecursive(nums []int, target int) int {
	n := len(nums)

	// Calculate the total sum to determine the range of possible sums
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If target is out of range, no valid assignment exists
	if target > totalSum || target < -totalSum {
		return 0
	}

	// Initialize memoization table
	// The range of sums is [-totalSum, totalSum], shifted to [0, 2*totalSum]
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 2*totalSum+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// Use a helper function for the recursive calculation
	return findTargetSumHelper(nums, target, 0, 0, totalSum, memo)
}

// findTargetSumHelper is a recursive helper function with memoization
func findTargetSumHelper(nums []int, target, index, currentSum, totalSum int, memo [][]int) int {
	// Base case: we've processed all elements
	if index == len(nums) {
		if currentSum == target {
			return 1
		}
		return 0
	}

	// Shift currentSum to use as an index in the memo table
	shiftedSum := currentSum + totalSum

	// Check if we have already computed this state
	if memo[index][shiftedSum] != -1 {
		return memo[index][shiftedSum]
	}

	// Try adding the current element with a positive sign
	positive := findTargetSumHelper(nums, target, index+1, currentSum+nums[index], totalSum, memo)

	// Try adding the current element with a negative sign
	negative := findTargetSumHelper(nums, target, index+1, currentSum-nums[index], totalSum, memo)

	// Total ways is the sum of both options
	memo[index][shiftedSum] = positive + negative
	return memo[index][shiftedSum]
}

// FindTargetSumWaysDP1D uses a 1D dynamic programming approach for space optimization
// Time Complexity: O(N * S) where N is the number of elements and S is the sum of all elements
// Space Complexity: O(S) for the DP array
func FindTargetSumWaysDP1D(nums []int, target int) int {
	// Calculate the total sum of the array
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If target is out of range or (totalSum + target) is odd, no valid assignment exists
	if target > totalSum || target < -totalSum || (totalSum+target)%2 != 0 {
		return 0
	}

	// Calculate the required subset sum
	sum := (totalSum + target) / 2

	// DP array: dp[j] represents count of subsets that sum to j
	dp := make([]int, sum+1)
	dp[0] = 1 // Empty subset can sum to 0

	// Fill the DP array
	for i := 0; i < len(nums); i++ {
		// Skip 0 elements to avoid counting them multiple times
		if nums[i] == 0 {
			dp[0] *= 2
			continue
		}

		// Process in reverse to avoid counting subsets multiple times
		for j := sum; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[sum]
}

// FindAllTargetSumExpressions returns all possible expressions that evaluate to the target
// Time Complexity: O(N * 2^N) where N is the number of elements
// Space Complexity: O(N * 2^N) for storing all expressions
func FindAllTargetSumExpressions(nums []int, target int) []string {
	result := []string{}
	currentExpression := ""

	// Start the backtracking process
	findAllExpressionsHelper(nums, target, 0, 0, currentExpression, &result)

	return result
}

// findAllExpressionsHelper is a recursive backtracking helper function
func findAllExpressionsHelper(nums []int, target int, index int, currentSum int,
	currentExpression string, result *[]string) {

	// Base case: we've processed all elements
	if index == len(nums) {
		if currentSum == target {
			// Remove the leading operator if it exists
			if len(currentExpression) > 0 && currentExpression[0] == '+' {
				*result = append(*result, currentExpression[1:])
			} else {
				*result = append(*result, currentExpression)
			}
		}
		return
	}

	// Add current number with positive sign
	posExpr := currentExpression
	if index > 0 {
		posExpr += "+"
	}
	posExpr += strconv.Itoa(nums[index])
	findAllExpressionsHelper(nums, target, index+1, currentSum+nums[index], posExpr, result)

	// Add current number with negative sign
	negExpr := currentExpression + "-" + strconv.Itoa(nums[index])
	findAllExpressionsHelper(nums, target, index+1, currentSum-nums[index], negExpr, result)
}

// Example usage:
//
// Input: nums = [1, 1, 1, 1, 1], target = 3
// Output: 5 (There are 5 ways: +1+1+1+1-1, +1+1+1-1+1, +1+1-1+1+1, +1-1+1+1+1, -1+1+1+1+1)
//
// Input: nums = [1], target = 1
// Output: 1 (There is 1 way: +1)
//
// Input: nums = [1, 2, 3], target = 0
// Output: 2 (There are 2 ways: +1-2+3, -1+2-3)
