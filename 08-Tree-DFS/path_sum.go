package treedfs

// HasPathSum determines if the tree has a root-to-leaf path that sums to the target sum
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (call stack)
func HasPathSum(root *TreeNode, targetSum int) bool {
	// Base case: empty tree
	if root == nil {
		return false
	}

	// If we've reached a leaf node, check if the target sum is equal to the current node's value
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	// Recursively check left and right subtrees with the remaining sum
	remainingSum := targetSum - root.Val
	return HasPathSum(root.Left, remainingSum) || HasPathSum(root.Right, remainingSum)
}

// AllPathSum finds all root-to-leaf paths where the sum of node values equals the target sum
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (call stack + path storage)
func AllPathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	findPaths(root, targetSum, []int{}, &result)
	return result
}

// findPaths is a helper function to find all paths that sum to the target
func findPaths(node *TreeNode, targetSum int, currentPath []int, result *[][]int) {
	if node == nil {
		return
	}

	// Add the current node to the path
	currentPath = append(currentPath, node.Val)

	// If we've reached a leaf node and the sum equals the target, add the path to the result
	if node.Left == nil && node.Right == nil && targetSum == node.Val {
		// Create a copy of the current path to avoid modifying it later
		pathCopy := make([]int, len(currentPath))
		copy(pathCopy, currentPath)
		*result = append(*result, pathCopy)
		return
	}

	// Recursively traverse the left and right subtrees with the remaining sum
	remainingSum := targetSum - node.Val
	findPaths(node.Left, remainingSum, currentPath, result)
	findPaths(node.Right, remainingSum, currentPath, result)
}

// Example usage:
//
// Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
//
//       5
//      / \
//     4   8
//    /   / \
//   11  13  4
//  / \      \
// 7   2      1
//
// HasPathSum output: true
// AllPathSum output: [[5,4,11,2], [5,8,4,5]]
// Explanation: There are two paths whose sum of node values is 22:
// 1. 5 -> 4 -> 11 -> 2
// 2. 5 -> 8 -> 4 -> 5
//
// Input: root = [1,2,3], targetSum = 5
//
//     1
//    / \
//   2   3
//
// HasPathSum output: false
// AllPathSum output: []
// Explanation: There is no path whose sum of node values is 5.
//
// Input: root = [1,2], targetSum = 1
//
//     1
//    /
//   2
//
// HasPathSum output: false
// AllPathSum output: []
// Explanation: The only path is 1 -> 2 with sum 3, not 1.
