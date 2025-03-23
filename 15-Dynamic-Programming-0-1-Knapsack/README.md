# Dynamic Programming - 0/1 Knapsack Pattern

## Introduction

The 0/1 Knapsack pattern is a classic problem in dynamic programming. It addresses situations where we need to make optimal selections from a set of items, each with its own value and weight, under a constraint (such as maximum weight capacity). The "0/1" refers to the binary choice for each item - either take it (1) or leave it (0); we cannot take a fraction of an item.

This pattern serves as the foundation for solving many optimization problems that involve making selections under constraints.

## How It Works

The 0/1 Knapsack approach typically follows these steps:

1. Define a state that represents a subproblem (often using indices for items and remaining capacity)
2. Create a recurrence relation that connects subproblems
3. Use either top-down (memoization) or bottom-up (tabulation) approach to solve the problem
4. Track decisions to reconstruct the solution (if needed)

### State Definition

In the classic 0/1 Knapsack problem, the state is defined by:
- `dp[i][j]` = maximum value that can be obtained from the first `i` items with a capacity of `j`

### Recurrence Relation

For each item, we have two choices:
1. Skip the item: `dp[i][j] = dp[i-1][j]`
2. Include the item (if it fits): `dp[i][j] = value[i-1] + dp[i-1][j-weight[i-1]]`

The optimal value is the maximum of these two options.

## Time and Space Complexity

- **Time Complexity**: O(N × W), where N is the number of items and W is the weight capacity
- **Space Complexity**: O(N × W) for standard implementation, though it can be optimized to O(W) with a 1D array

## When to Use 0/1 Knapsack Pattern

This pattern is useful when:

1. You need to make a binary decision (take/leave) for each item
2. You have a constraint (like weight capacity)
3. You want to maximize or minimize an objective (like total value)
4. The items cannot be split or taken more than once

## Common Problem Patterns

The 0/1 Knapsack pattern can be adapted to solve many problems, including:

1. **Classic 0/1 Knapsack**: Maximize value under a weight constraint
2. **Subset Sum**: Determine if a subset of elements can sum to a target
3. **Equal Subset Sum Partition**: Split array into two equal sum subsets
4. **Minimum Subset Sum Difference**: Partition array to minimize difference between subset sums
5. **Count of Subset Sum**: Count the number of subsets with a given sum
6. **Target Sum**: Assign +/- signs to numbers to achieve a target sum
7. **Coin Change**: Count the number of ways to make change for a target amount

## Implementation in Golang

Here's how the basic 0/1 Knapsack problem can be implemented in Go:

```go
// Recursive approach with memoization
func knapsackRecursive(profits []int, weights []int, capacity int) int {
    n := len(profits)
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

// Bottom-up Dynamic Programming (Tabulation)
func knapsackDP(profits []int, weights []int, capacity int) int {
    n := len(profits)
    if capacity <= 0 || n == 0 {
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
            // Skip the item
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

// Space-optimized version using 1D array
func knapsackDP1D(profits []int, weights []int, capacity int) int {
    n := len(profits)
    if capacity <= 0 || n == 0 {
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

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

## Example Problems

1. **0/1 Knapsack**: Given weights and values of N items, put these items in a knapsack of capacity W to get the maximum value.
2. **Subset Sum**: Given a set of non-negative integers, determine if there is a subset of the given set with sum equal to a given sum.
3. **Equal Subset Sum Partition**: Partition the array into two subsets such that the sum of elements in both subsets is equal.
4. **Minimum Subset Sum Difference**: Partition the array into two subsets such that the difference between the subset sums is minimized.
5. **Count of Subset Sum**: Count the number of subsets with a given sum.
6. **Target Sum**: Assign + or - signs to each element to make the sum equal to a target value.
7. **Minimum Coin Change**: Find the minimum number of coins needed to make a given amount. 