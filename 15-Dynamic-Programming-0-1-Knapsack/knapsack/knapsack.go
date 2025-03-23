package knapsack

// KnapsackRecursive solves the 0/1 Knapsack problem using a recursive approach with memoization
// Time Complexity: O(N * C) where N is the number of items and C is the capacity
// Space Complexity: O(N * C) for the memoization table + O(N) for recursion stack
func KnapsackRecursive(profits []int, weights []int, capacity int) int {
	n := len(profits)
	if n == 0 || capacity <= 0 {
		return 0
	}

	// Initialize memoization table with -1
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, capacity+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return knapsackHelper(profits, weights, capacity, 0, memo)
}

// knapsackHelper is a recursive helper function with memoization
func knapsackHelper(profits []int, weights []int, capacity int, currentIndex int, memo [][]int) int {
	// Base cases
	if capacity <= 0 || currentIndex >= len(profits) {
		return 0
	}

	// Check if we have already solved this subproblem
	if memo[currentIndex][capacity] != -1 {
		return memo[currentIndex][capacity]
	}

	// Decision 1: Skip the current item
	profit1 := knapsackHelper(profits, weights, capacity, currentIndex+1, memo)

	// Decision 2: Include the current item if it fits
	profit2 := 0
	if weights[currentIndex] <= capacity {
		profit2 = profits[currentIndex] +
			knapsackHelper(profits, weights, capacity-weights[currentIndex], currentIndex+1, memo)
	}

	// Take the maximum profit
	memo[currentIndex][capacity] = max(profit1, profit2)
	return memo[currentIndex][capacity]
}

// KnapsackDP solves the 0/1 Knapsack problem using bottom-up dynamic programming (tabulation)
// Time Complexity: O(N * C) where N is the number of items and C is the capacity
// Space Complexity: O(N * C) for the DP table
func KnapsackDP(profits []int, weights []int, capacity int) int {
	n := len(profits)
	if n == 0 || capacity <= 0 {
		return 0
	}

	// Initialize DP table
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Fill the dp table
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			// Skip the item (inherit value from the previous item)
			dp[i][w] = dp[i-1][w]

			// Take the item if it fits
			if weights[i-1] <= w {
				valueWithCurrent := profits[i-1] + dp[i-1][w-weights[i-1]]
				dp[i][w] = max(dp[i][w], valueWithCurrent)
			}
		}
	}

	return dp[n][capacity]
}

// KnapsackDP1D solves the 0/1 Knapsack problem with an optimized 1D DP array
// Time Complexity: O(N * C) where N is the number of items and C is the capacity
// Space Complexity: O(C) for the DP array
func KnapsackDP1D(profits []int, weights []int, capacity int) int {
	n := len(profits)
	if n == 0 || capacity <= 0 {
		return 0
	}

	// We only need a 1D array of size capacity+1
	dp := make([]int, capacity+1)

	// Initialize for the first item
	for w := weights[0]; w <= capacity; w++ {
		dp[w] = profits[0]
	}

	// Process all items
	for i := 1; i < n; i++ {
		// Process the capacity in reverse to avoid using the updated value
		for w := capacity; w >= weights[i]; w-- {
			dp[w] = max(dp[w], profits[i]+dp[w-weights[i]])
		}
	}

	return dp[capacity]
}

// KnapsackWithItems returns both the maximum value and the selected items
// Time Complexity: O(N * C) where N is the number of items and C is the capacity
// Space Complexity: O(N * C) for the DP table
func KnapsackWithItems(profits []int, weights []int, capacity int) (int, []int) {
	n := len(profits)
	if n == 0 || capacity <= 0 {
		return 0, []int{}
	}

	// Initialize DP table
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Fill the dp table
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			// Skip the item
			dp[i][w] = dp[i-1][w]

			// Take the item if it fits
			if weights[i-1] <= w {
				valueWithCurrent := profits[i-1] + dp[i-1][w-weights[i-1]]
				dp[i][w] = max(dp[i][w], valueWithCurrent)
			}
		}
	}

	// Backtrack to find the selected items
	selectedItems := []int{}
	w := capacity
	for i := n; i > 0; i-- {
		// If the value came from including this item
		if dp[i][w] != dp[i-1][w] {
			selectedItems = append(selectedItems, i-1) // Add the item index
			w -= weights[i-1]                          // Reduce the capacity
		}
	}

	// Reverse the selected items to get them in the original order
	for i, j := 0, len(selectedItems)-1; i < j; i, j = i+1, j-1 {
		selectedItems[i], selectedItems[j] = selectedItems[j], selectedItems[i]
	}

	return dp[n][capacity], selectedItems
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example usage:
//
// Profits: [1, 6, 10, 16]
// Weights: [1, 2, 3, 5]
// Capacity: 7
// Maximum profit: 22 (by selecting items at indices 1, 2, and 3 with values 6, 10, and 16)
//
// Profits: [4, 5, 3, 7]
// Weights: [2, 3, 1, 4]
// Capacity: 5
// Maximum profit: 10 (by selecting items at indices 0 and 3 with values 4 and 7)
