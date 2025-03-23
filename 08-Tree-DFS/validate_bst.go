package treedfs

import "math"

// IsValidBST checks if a binary tree is a valid binary search tree (BST)
// A valid BST has the property that all values in the left subtree are less than the node's value,
// and all values in the right subtree are greater than the node's value
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (call stack)
func IsValidBST(root *TreeNode) bool {
	return validateBST(root, math.MinInt64, math.MaxInt64)
}

// validateBST checks if the tree is a valid BST with values in the given range
func validateBST(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}

	// Check if the current node's value is within the valid range
	if int64(node.Val) <= min || int64(node.Val) >= max {
		return false
	}

	// Recursively validate left and right subtrees
	// For the left subtree, the maximum value should be the current node's value
	// For the right subtree, the minimum value should be the current node's value
	return validateBST(node.Left, min, int64(node.Val)) && validateBST(node.Right, int64(node.Val), max)
}

// IsValidBSTInOrder validates a BST using inorder traversal
// A valid BST will have its values in ascending order when traversed inorder
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (call stack)
func IsValidBSTInOrder(root *TreeNode) bool {
	// Start with the minimum possible value
	var prev int64 = math.MinInt64
	return validateBSTInOrder(root, &prev)
}

// validateBSTInOrder validates a BST using inorder traversal
func validateBSTInOrder(node *TreeNode, prev *int64) bool {
	if node == nil {
		return true
	}

	// First, validate the left subtree
	if !validateBSTInOrder(node.Left, prev) {
		return false
	}

	// Check if the current node's value is greater than the previous value
	if int64(node.Val) <= *prev {
		return false
	}

	// Update the previous value to the current node's value
	*prev = int64(node.Val)

	// Finally, validate the right subtree
	return validateBSTInOrder(node.Right, prev)
}

// Example usage:
//
// Input: root = [2,1,3]
//
//      2
//     / \
//    1   3
//
// Output: true
// Explanation: All nodes satisfy the BST property.
//
// Input: root = [5,1,4,null,null,3,6]
//
//      5
//     / \
//    1   4
//       / \
//      3   6
//
// Output: false
// Explanation: The value 4 is less than the value 5, but it's in the right subtree.
//
// Input: root = [5,4,6,null,null,3,7]
//
//      5
//     / \
//    4   6
//       / \
//      3   7
//
// Output: false
// Explanation: The value 3 is less than the value 5, but it's in the right subtree.
