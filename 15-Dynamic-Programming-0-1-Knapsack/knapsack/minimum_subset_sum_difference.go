package knapsack

import "math"

// MinimumSubsetSumDifference finds the minimum absolute difference between
// the sums of two subsets that partition the given array
// Time Complexity: O(N * S) where N is the number of elements and S is the total sum
// Space Complexity: O(S) with the optimized DP approach
func MinimumSubsetSumDifference(nums []int) int {
	n := len(nums)

	// Edge case: If array has only one element, return that element
	if n == 1 {
		return nums[0]
	}

	// Calculate the total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// We need to find a subset with sum closest to totalSum/2
	// Build a DP array to find all possible subset sums
	dp := make([]bool, totalSum+1)
	dp[0] = true // Empty set has sum 0

	// Calculate all possible subset sums
	currentSum := 0
	for _, num := range nums {
		currentSum += num
		// Process in reverse to avoid using updated values
		for j := min(totalSum, currentSum); j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	// Find the subset sum closest to half of the total sum
	minDiff := totalSum
	for sum := 0; sum <= totalSum/2; sum++ {
		if dp[sum] {
			// If sum is possible, calculate the difference
			diff := totalSum - 2*sum
			minDiff = min(minDiff, diff)
		}
	}

	return minDiff
}

// MinimumSubsetSumDifferenceRecursive finds the minimum difference between
// the sums of two subsets using recursion with memoization
// Time Complexity: O(N * S) where N is the number of elements and S is the total sum
// Space Complexity: O(N * S) for the memoization table
func MinimumSubsetSumDifferenceRecursive(nums []int) int {
	n := len(nums)

	// Edge case: If array has only one element, return that element
	if n == 1 {
		return nums[0]
	}

	// Calculate the total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// Initialize memoization table
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, totalSum+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// Find the closest subset sum to totalSum/2
	return totalSum - 2*findClosestSubsetSum(nums, totalSum, 0, 0, memo)
}

// findClosestSubsetSum finds the subset sum closest to totalSum/2
func findClosestSubsetSum(nums []int, totalSum int, currentIndex int, currentSum int, memo [][]int) int {
	// Base case: We've processed all elements
	if currentIndex == len(nums) {
		// Return the sum that minimizes |totalSum - 2*sum|
		if currentSum <= totalSum/2 {
			return currentSum
		}
		return 0
	}

	// Check if we have already computed this state
	if memo[currentIndex][currentSum] != -1 {
		return memo[currentIndex][currentSum]
	}

	// Decision 1: Include the current element
	sum1 := findClosestSubsetSum(nums, totalSum, currentIndex+1, currentSum+nums[currentIndex], memo)

	// Decision 2: Skip the current element
	sum2 := findClosestSubsetSum(nums, totalSum, currentIndex+1, currentSum, memo)

	// Choose the sum that is closer to totalSum/2
	if math.Abs(float64(totalSum-2*sum1)) < math.Abs(float64(totalSum-2*sum2)) {
		memo[currentIndex][currentSum] = sum1
	} else {
		memo[currentIndex][currentSum] = sum2
	}

	return memo[currentIndex][currentSum]
}

// FindMinimumSubsetSumPartition returns the two subsets with minimum difference
// Time Complexity: O(N * S) where N is the number of elements and S is the total sum
// Space Complexity: O(N * S) for the DP table
func FindMinimumSubsetSumPartition(nums []int) ([]int, []int, int) {
	n := len(nums)

	// Edge case: If array has only one element, return it as subset1
	if n == 1 {
		return nums, []int{}, nums[0]
	}

	// Calculate the total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// DP table: dp[i][j] represents whether a subset of the first i elements can sum to j
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, totalSum+1)
		dp[i][0] = true // Empty set has sum 0
	}

	// Fill the DP table
	for i := 1; i <= n; i++ {
		for j := 1; j <= totalSum; j++ {
			// Skip the current element
			dp[i][j] = dp[i-1][j]

			// Include the current element if possible
			if nums[i-1] <= j {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	// Find the closest sum to totalSum/2
	closestSum := 0
	for j := totalSum / 2; j >= 0; j-- {
		if dp[n][j] {
			closestSum = j
			break
		}
	}

	// Calculate the minimum difference
	minDiff := totalSum - 2*closestSum

	// Reconstruct the first subset
	subset1 := []int{}
	i, j := n, closestSum
	for i > 0 && j > 0 {
		// If the current element was not included
		if dp[i-1][j] {
			i--
		} else {
			// Current element was included
			subset1 = append(subset1, nums[i-1])
			j -= nums[i-1]
			i--
		}
	}

	// Build the second subset
	subset2 := []int{}
	subset1Map := make(map[int]int) // Value to count mapping
	for _, num := range subset1 {
		subset1Map[num]++
	}

	for _, num := range nums {
		if count, exists := subset1Map[num]; !exists || count == 0 {
			subset2 = append(subset2, num)
		} else {
			subset1Map[num]--
		}
	}

	return subset1, subset2, minDiff
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Example usage:
//
// Input: nums = [1, 2, 3, 9]
// Output: 3 (partitioned into [1, 2, 3] with sum 6 and [9] with sum 9)
//
// Input: nums = [1, 2, 7, 1, 5]
// Output: 0 (partitioned into [1, 2, 5] and [7, 1], both with sum 8)
//
// Input: nums = [10, 20, 15, 5, 25]
// Output: 5 (partitioned into [10, 20, 5] with sum 35 and [15, 25] with sum 40)
