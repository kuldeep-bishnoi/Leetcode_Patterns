package treedfs

// CountPaths finds the number of paths in the tree that sum to the target sum
// A path can start and end at any node (not necessarily root-to-leaf)
// Time Complexity: O(n^2) in the worst case, where n is the number of nodes
// Space Complexity: O(h) where h is the height of the tree (call stack)
func CountPaths(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// Count paths that start from the root
	pathsFromRoot := countPathsFromNode(root, targetSum, 0)

	// Recursively count paths from left and right subtrees
	pathsFromLeft := CountPaths(root.Left, targetSum)
	pathsFromRight := CountPaths(root.Right, targetSum)

	// Return the total count
	return pathsFromRoot + pathsFromLeft + pathsFromRight
}

// countPathsFromNode counts the number of paths starting from the given node that sum to the target
func countPathsFromNode(node *TreeNode, targetSum int, currentSum int) int {
	if node == nil {
		return 0
	}

	// Add the current node's value to the running sum
	currentSum += node.Val

	// Initialize path count
	pathCount := 0

	// If we've found a path that sums to the target, increment the count
	if currentSum == targetSum {
		pathCount = 1
	}

	// Continue counting paths through the left and right children
	pathCount += countPathsFromNode(node.Left, targetSum, currentSum)
	pathCount += countPathsFromNode(node.Right, targetSum, currentSum)

	return pathCount
}

// CountPathsOptimized provides a more efficient solution using a prefix sum technique
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (call stack + hash map)
func CountPathsOptimized(root *TreeNode, targetSum int) int {
	// Map to store prefix sums and their frequencies
	prefixSumCount := make(map[int]int)
	// Initialize with 0 sum having 1 occurrence (empty path)
	prefixSumCount[0] = 1

	return countPathsWithPrefixSum(root, targetSum, 0, prefixSumCount)
}

// countPathsWithPrefixSum uses a map to track prefix sums for more efficient counting
func countPathsWithPrefixSum(node *TreeNode, targetSum int, currentSum int, prefixSumCount map[int]int) int {
	if node == nil {
		return 0
	}

	// Add the current node's value to the running sum
	currentSum += node.Val

	// Initialize path count
	// Check if there are any paths that can be completed by adding the current node
	// This is equivalent to checking if (currentSum - targetSum) exists in our prefix sums
	pathCount := prefixSumCount[currentSum-targetSum]

	// Update the prefix sum map with the current sum
	prefixSumCount[currentSum]++

	// Recursively count paths in the left and right subtrees
	pathCount += countPathsWithPrefixSum(node.Left, targetSum, currentSum, prefixSumCount)
	pathCount += countPathsWithPrefixSum(node.Right, targetSum, currentSum, prefixSumCount)

	// Backtrack: remove the current path's sum when going back up the recursion tree
	// to avoid counting paths from other branches
	prefixSumCount[currentSum]--

	return pathCount
}

// Example usage:
//
// Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
//
//        10
//       /  \
//      5   -3
//     / \    \
//    3   2   11
//   / \   \
//  3  -2   1
//
// Output: 3
// Explanation: There are three paths that sum to 8:
// 1. 5 -> 3
// 2. 5 -> 2 -> 1
// 3. -3 -> 11
//
// Input: root = [1,null,2,null,3,null,4,null,5], targetSum = 3
//
//    1
//     \
//      2
//       \
//        3
//         \
//          4
//           \
//            5
//
// Output: 2
// Explanation: There are two paths that sum to 3:
// 1. 1 -> 2
// 2. 3
