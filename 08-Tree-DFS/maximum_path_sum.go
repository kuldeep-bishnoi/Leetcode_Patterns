package treedfs

import "math"

// MaxPathSum finds the maximum path sum in a binary tree
// A path is defined as any sequence of nodes from some starting node to any node
// in the tree along the parent-child connections
// The path must contain at least one node and does not need to go through the root
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (call stack)
func MaxPathSum(root *TreeNode) int {
	// Initialize with the minimum possible value to handle negative values
	maxSum := math.MinInt32
	calculateMaxPathSum(root, &maxSum)
	return maxSum
}

// calculateMaxPathSum is a helper function that computes the maximum path sum
// through each node and updates the global maximum
func calculateMaxPathSum(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	// Recursively compute the maximum path sum for left and right subtrees
	// If the sum is negative, we don't include that path (take 0 instead)
	leftSum := max(calculateMaxPathSum(node.Left, maxSum), 0)
	rightSum := max(calculateMaxPathSum(node.Right, maxSum), 0)

	// Current path value is the sum of the node value and the maximum
	// path sums from its left and right subtrees
	currentPathSum := node.Val + leftSum + rightSum

	// Update the maximum path sum if the current path is greater
	*maxSum = max(*maxSum, currentPathSum)

	// Return the maximum sum of a path that starts from this node
	// and goes down to one of its children (can't go both ways when
	// returning to the parent call)
	return node.Val + max(leftSum, rightSum)
}

// Example usage:
//
// Input: root = [1,2,3]
//
//      1
//     / \
//    2   3
//
// Output: 6
// Explanation: The optimal path is 2 -> 1 -> 3 with a sum of 6.
//
// Input: root = [-10,9,20,null,null,15,7]
//
//      -10
//      / \
//     9  20
//       /  \
//      15   7
//
// Output: 42
// Explanation: The optimal path is 20 -> 15 -> 7 with a sum of 42.
//
// Input: root = [-3]
//
//      -3
//
// Output: -3
// Explanation: The tree contains only one node, so the maximum path sum is -3.
