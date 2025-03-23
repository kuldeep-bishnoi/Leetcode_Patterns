package treedfs

// DiameterOfBinaryTree calculates the diameter of a binary tree
// The diameter is the length of the longest path between any two nodes in the tree
// This path may or may not pass through the root
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (call stack)
func DiameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	calculateHeight(root, &diameter)
	return diameter
}

// calculateHeight is a helper function that computes the height of each node
// and updates the diameter if a longer path is found
func calculateHeight(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}

	// Recursively compute the height of left and right subtrees
	leftHeight := calculateHeight(node.Left, diameter)
	rightHeight := calculateHeight(node.Right, diameter)

	// Update the diameter if the path through the current node is longer
	// The path length is the sum of the heights of left and right subtrees
	*diameter = max(*diameter, leftHeight+rightHeight)

	// Return the height of the current node
	return 1 + max(leftHeight, rightHeight)
}

// Example usage:
//
// Input: root = [1,2,3,4,5]
//
//      1
//     / \
//    2   3
//   / \
//  4   5
//
// Output: 3
// Explanation: The longest path is [4,2,1,3] or [5,2,1,3], with a length of 3.
//
// Input: root = [1,2]
//
//    1
//   /
//  2
//
// Output: 1
// Explanation: The longest path is [1,2], with a length of 1.
