package main

import (
	"fmt"
	"sort"
	"strings"

	"knapsack-demo/knapsack"
)

func main() {
	fmt.Println("========== 0/1 Knapsack Problem Implementations ==========")

	// Basic 0/1 Knapsack Problem
	fmt.Println("\n1. Basic 0/1 Knapsack Problem:")
	profits := []int{1, 6, 10, 16}
	weights := []int{1, 2, 3, 5}
	capacity := 7

	fmt.Printf("Profits: %v\n", profits)
	fmt.Printf("Weights: %v\n", weights)
	fmt.Printf("Capacity: %d\n", capacity)

	maxProfitRecursive := knapsack.KnapsackRecursive(profits, weights, capacity)
	fmt.Printf("Maximum profit (recursive): %d\n", maxProfitRecursive)

	maxProfitDP := knapsack.KnapsackDP(profits, weights, capacity)
	fmt.Printf("Maximum profit (DP): %d\n", maxProfitDP)

	maxProfitDP1D := knapsack.KnapsackDP1D(profits, weights, capacity)
	fmt.Printf("Maximum profit (DP 1D): %d\n", maxProfitDP1D)

	maxProfit, selectedItems := knapsack.KnapsackWithItems(profits, weights, capacity)
	fmt.Printf("Maximum profit with items: %d\n", maxProfit)
	fmt.Printf("Selected item indices: %v\n", selectedItems)

	// Subset Sum Problem
	fmt.Println("\n2. Subset Sum Problem:")
	nums := []int{1, 2, 3, 7}
	targetSum := 6

	fmt.Printf("Array: %v\n", nums)
	fmt.Printf("Target sum: %d\n", targetSum)

	canFormSubsetRecursive := knapsack.SubsetSumRecursive(nums, targetSum)
	fmt.Printf("Can form subset with sum %d (recursive): %t\n", targetSum, canFormSubsetRecursive)

	canFormSubsetDP := knapsack.SubsetSumDP(nums, targetSum)
	fmt.Printf("Can form subset with sum %d (DP): %t\n", targetSum, canFormSubsetDP)

	canFormSubsetDP1D := knapsack.SubsetSumDP1D(nums, targetSum)
	fmt.Printf("Can form subset with sum %d (DP 1D): %t\n", targetSum, canFormSubsetDP1D)

	exists, subset := knapsack.FindSubsetSum(nums, targetSum)
	if exists {
		fmt.Printf("Subset with sum %d: %v\n", targetSum, subset)
	} else {
		fmt.Printf("No subset exists with sum %d\n", targetSum)
	}

	// Equal Subset Sum Partition
	fmt.Println("\n3. Equal Subset Sum Partition Problem:")
	nums = []int{1, 5, 11, 5}

	fmt.Printf("Array: %v\n", nums)

	canPartition := knapsack.CanPartition(nums)
	fmt.Printf("Can be partitioned into equal sum subsets (DP): %t\n", canPartition)

	canPartitionRecursive := knapsack.CanPartitionRecursive(nums)
	fmt.Printf("Can be partitioned into equal sum subsets (recursive): %t\n", canPartitionRecursive)

	canPartition, subset1, subset2 := knapsack.FindEqualSubsetPartition(nums)
	if canPartition {
		fmt.Printf("Subset 1: %v (sum: %d)\n", subset1, sumArray(subset1))
		fmt.Printf("Subset 2: %v (sum: %d)\n", subset2, sumArray(subset2))
	} else {
		fmt.Println("Cannot be partitioned into equal sum subsets")
	}

	// Partition to K Equal Sum Subsets
	fmt.Println("\n4. Partition to K Equal Sum Subsets:")
	nums = []int{4, 3, 2, 3, 5, 2, 1}
	k := 4

	fmt.Printf("Array: %v\n", nums)
	fmt.Printf("k: %d\n", k)

	canPartitionKSubsets := knapsack.CanPartitionKSubsets(nums, k)
	fmt.Printf("Can be partitioned into %d equal sum subsets: %t\n", k, canPartitionKSubsets)

	// Minimum Subset Sum Difference
	fmt.Println("\n5. Minimum Subset Sum Difference Problem:")
	nums = []int{1, 2, 3, 9}

	fmt.Printf("Array: %v\n", nums)

	minDiff := knapsack.MinimumSubsetSumDifference(nums)
	fmt.Printf("Minimum subset sum difference (DP): %d\n", minDiff)

	minDiffRecursive := knapsack.MinimumSubsetSumDifferenceRecursive(nums)
	fmt.Printf("Minimum subset sum difference (recursive): %d\n", minDiffRecursive)

	subset1, subset2, minDifference := knapsack.FindMinimumSubsetSumPartition(nums)
	fmt.Printf("Subset 1: %v (sum: %d)\n", subset1, sumArray(subset1))
	fmt.Printf("Subset 2: %v (sum: %d)\n", subset2, sumArray(subset2))
	fmt.Printf("Minimum difference: %d\n", minDifference)

	// Count of Subset Sum
	fmt.Println("\n6. Count of Subset Sum Problem:")
	nums = []int{1, 1, 2, 3}
	targetSum = 4

	fmt.Printf("Array: %v\n", nums)
	fmt.Printf("Target sum: %d\n", targetSum)

	countDP := knapsack.CountSubsetSum(nums, targetSum)
	fmt.Printf("Count of subsets with sum %d (DP): %d\n", targetSum, countDP)

	countDP2D := knapsack.CountSubsetSumDP(nums, targetSum)
	fmt.Printf("Count of subsets with sum %d (DP 2D): %d\n", targetSum, countDP2D)

	countRecursive := knapsack.CountSubsetSumRecursive(nums, targetSum)
	fmt.Printf("Count of subsets with sum %d (recursive): %d\n", targetSum, countRecursive)

	subsets := knapsack.FindSubsetsWithSum(nums, targetSum)
	fmt.Printf("Subsets with sum %d: %v\n", targetSum, subsets)

	// Count of Subsets with Target Difference
	fmt.Println("\n7. Count of Subsets with Target Difference:")
	nums = []int{1, 1, 2, 3}
	diff := 1

	fmt.Printf("Array: %v\n", nums)
	fmt.Printf("Target difference: %d\n", diff)

	countWithDiff := knapsack.CountSubsetsWithTargetDifference(nums, diff)
	fmt.Printf("Count of subsets with difference %d: %d\n", diff, countWithDiff)

	// Target Sum Problem
	fmt.Println("\n8. Target Sum Problem:")
	nums = []int{1, 1, 1, 1, 1}
	target := 3

	fmt.Printf("Array: %v\n", nums)
	fmt.Printf("Target: %d\n", target)

	ways := knapsack.FindTargetSumWays(nums, target)
	fmt.Printf("Number of ways to achieve target %d (reduced to subset sum): %d\n", target, ways)

	waysDP := knapsack.FindTargetSumWaysDP(nums, target)
	fmt.Printf("Number of ways to achieve target %d (DP): %d\n", target, waysDP)

	waysRecursive := knapsack.FindTargetSumWaysRecursive(nums, target)
	fmt.Printf("Number of ways to achieve target %d (recursive): %d\n", target, waysRecursive)

	waysDP1D := knapsack.FindTargetSumWaysDP1D(nums, target)
	fmt.Printf("Number of ways to achieve target %d (DP 1D): %d\n", target, waysDP1D)

	expressions := knapsack.FindAllTargetSumExpressions(nums, target)
	sort.Strings(expressions)
	fmt.Printf("Expressions to achieve target %d: %s\n", target, strings.Join(expressions, ", "))
}

// Helper function to calculate the sum of an array
func sumArray(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}
	return total
}
