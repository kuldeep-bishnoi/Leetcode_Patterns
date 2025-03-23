package treedfs

// SumPathNumbers finds the sum of all numbers formed by root-to-leaf paths in a binary tree
// Each path represents a number where each node is a digit
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (call stack)
func SumPathNumbers(root *TreeNode) int {
	return calculatePathSum(root, 0)
}

// calculatePathSum is a helper function that calculates the path sum recursively
func calculatePathSum(node *TreeNode, pathSum int) int {
	if node == nil {
		return 0
	}

	// Calculate the current path sum by multiplying the previous sum by 10 and adding the current value
	// This effectively shifts the digits left and adds the new digit
	currentSum := pathSum*10 + node.Val

	// If we've reached a leaf node, return the current path sum
	if node.Left == nil && node.Right == nil {
		return currentSum
	}

	// Recursively calculate the sum for left and right subtrees
	leftSum := calculatePathSum(node.Left, currentSum)
	rightSum := calculatePathSum(node.Right, currentSum)

	// Return the total sum from both subtrees
	return leftSum + rightSum
}

// Example usage:
//
// Input: root = [1,2,3]
//
//     1
//    / \
//   2   3
//
// Output: 25
// Explanation:
// Path 1->2 forms the number 12
// Path 1->3 forms the number 13
// The sum of all root-to-leaf path numbers is 12 + 13 = 25
//
// Input: root = [4,9,0,5,1]
//
//       4
//      / \
//     9   0
//    / \
//   5   1
//
// Output: 1026
// Explanation:
// Path 4->9->5 forms the number 495
// Path 4->9->1 forms the number 491
// Path 4->0 forms the number 40
// The sum of all root-to-leaf path numbers is 495 + 491 + 40 = 1026
