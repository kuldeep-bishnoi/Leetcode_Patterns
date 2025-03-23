package knapsack

// CanPartition determines if the given array can be partitioned into two subsets
// with equal sum using dynamic programming
// Time Complexity: O(N * S) where N is the number of elements and S is half of the total sum
// Space Complexity: O(S) with the optimized 1D DP approach
func CanPartition(nums []int) bool {
	n := len(nums)

	// Edge case: If the array has less than 2 elements, it can't be partitioned
	if n < 2 {
		return false
	}

	// Calculate the total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If the total sum is odd, we can't partition into two equal subsets
	if totalSum%2 != 0 {
		return false
	}

	// Target sum is half of the total sum
	targetSum := totalSum / 2

	// Use the 1D DP approach for finding a subset with sum equal to targetSum
	return SubsetSumDP1D(nums, targetSum)
}

// CanPartitionRecursive determines if the given array can be partitioned into two
// subsets with equal sum using recursion with memoization
// Time Complexity: O(N * S) where N is the number of elements and S is half of the total sum
// Space Complexity: O(N * S) for the memoization table + O(N) for recursion stack
func CanPartitionRecursive(nums []int) bool {
	n := len(nums)

	// Edge case: If the array has less than 2 elements, it can't be partitioned
	if n < 2 {
		return false
	}

	// Calculate the total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If the total sum is odd, we can't partition into two equal subsets
	if totalSum%2 != 0 {
		return false
	}

	// Target sum is half of the total sum
	targetSum := totalSum / 2

	// Use the recursive approach for finding a subset with sum equal to targetSum
	return SubsetSumRecursive(nums, targetSum)
}

// FindEqualSubsetPartition not only determines if a partition is possible but also
// returns the two subsets with equal sum if a partition exists
// Time Complexity: O(N * S) where N is the number of elements and S is half of the total sum
// Space Complexity: O(N * S) for the DP table
func FindEqualSubsetPartition(nums []int) (bool, []int, []int) {
	n := len(nums)

	// Edge case: If the array has less than 2 elements, it can't be partitioned
	if n < 2 {
		return false, nil, nil
	}

	// Calculate the total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If the total sum is odd, we can't partition into two equal subsets
	if totalSum%2 != 0 {
		return false, nil, nil
	}

	// Target sum is half of the total sum
	targetSum := totalSum / 2

	// Find a subset that sums to targetSum
	canPartition, subset1 := FindSubsetSum(nums, targetSum)

	if !canPartition {
		return false, nil, nil
	}

	// Build the second subset from the remaining elements
	subset2 := []int{}
	subset1Map := make(map[int]int) // Maps value to count of occurrences

	// Count occurrences of elements in subset1
	for _, num := range subset1 {
		subset1Map[num]++
	}

	// Add elements not in subset1 (or with fewer occurrences) to subset2
	for _, num := range nums {
		if count, exists := subset1Map[num]; !exists || count == 0 {
			subset2 = append(subset2, num)
		} else {
			subset1Map[num]--
		}
	}

	return true, subset1, subset2
}

// CanPartitionKSubsets determines if the given array can be partitioned into k
// subsets with equal sum
// Time Complexity: O(k * 2^n) where n is the number of elements
// Space Complexity: O(n) for recursion stack
func CanPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)

	// Edge cases
	if k <= 0 || n < k {
		return false
	}

	// Calculate the total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If the total sum is not divisible by k, we cannot partition
	if totalSum%k != 0 {
		return false
	}

	// Each subset should have this sum
	targetSum := totalSum / k

	// Sort in descending order to optimize backtracking
	// This helps to fail earlier for invalid partitions
	quickSortDesc(nums, 0, n-1)

	// If the largest element is greater than the target sum, we cannot partition
	if nums[0] > targetSum {
		return false
	}

	// Used to track which elements have been used in subsets
	used := make([]bool, n)

	// Try to find k subsets each with sum equal to targetSum
	return canPartitionKSubsetsHelper(nums, used, k, targetSum, 0, 0)
}

// canPartitionKSubsetsHelper is a recursive helper for the k-partition problem
func canPartitionKSubsetsHelper(nums []int, used []bool, k int, targetSum int, currentSum int, startIndex int) bool {
	// If we've found k-1 valid subsets, the last one must be valid too
	if k == 1 {
		return true
	}

	// If current subset sum equals target, find the next subset
	if currentSum == targetSum {
		return canPartitionKSubsetsHelper(nums, used, k-1, targetSum, 0, 0)
	}

	// Try adding each unused element to the current subset
	for i := startIndex; i < len(nums); i++ {
		// Skip if already used or would exceed target sum
		if used[i] || currentSum+nums[i] > targetSum {
			continue
		}

		// Skip duplicates to avoid redundant exploration
		if i > startIndex && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		// Include current element
		used[i] = true

		// Recursively try to build valid subset
		if canPartitionKSubsetsHelper(nums, used, k, targetSum, currentSum+nums[i], i+1) {
			return true
		}

		// Backtrack
		used[i] = false

		// If adding the first element of a subset fails, the whole approach fails
		if currentSum == 0 {
			break
		}
	}

	return false
}

// quickSortDesc sorts the array in descending order
func quickSortDesc(nums []int, low, high int) {
	if low < high {
		pivotIndex := partitionDesc(nums, low, high)
		quickSortDesc(nums, low, pivotIndex-1)
		quickSortDesc(nums, pivotIndex+1, high)
	}
}

// partitionDesc is the partition function for quicksort in descending order
func partitionDesc(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1

	for j := low; j < high; j++ {
		if nums[j] >= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

// Example usage:
//
// Input: nums = [1, 5, 11, 5]
// Output: true (can be partitioned into [1, 5, 5] and [11])
//
// Input: nums = [1, 2, 3, 5]
// Output: false (cannot be partitioned into equal sum subsets)
//
// Input: nums = [2, 2, 2, 2, 3, 4, 5]
// k = 4
// Output: true (can be partitioned into [5], [4, 2], [3, 3], [2, 2, 2])
