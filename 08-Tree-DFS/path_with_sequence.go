package treedfs

// FindPath determines if there exists a root-to-leaf path with the given sequence
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (call stack)
func FindPath(root *TreeNode, sequence []int) bool {
	if root == nil {
		return len(sequence) == 0
	}

	return findPathRecursive(root, sequence, 0)
}

// findPathRecursive is a helper function to find the path with the given sequence
func findPathRecursive(node *TreeNode, sequence []int, index int) bool {
	if node == nil {
		return false
	}

	// If the sequence is shorter than the current index, or the current node's value
	// doesn't match the sequence at the current index, the path doesn't match
	if index >= len(sequence) || node.Val != sequence[index] {
		return false
	}

	// If we've reached a leaf node and the end of the sequence, we've found a match
	if node.Left == nil && node.Right == nil && index == len(sequence)-1 {
		return true
	}

	// Recursively check the left and right subtrees
	return findPathRecursive(node.Left, sequence, index+1) ||
		findPathRecursive(node.Right, sequence, index+1)
}

// Example usage:
//
// Input: root = [1,7,9,2,9], sequence = [1,9,9]
//
//       1
//      / \
//     7   9
//    /     \
//   2       9
//
// Output: true
// Explanation: The path 1->9->9 matches the given sequence.
//
// Input: root = [1,0,1,1,6,5], sequence = [1,1,6]
//
//       1
//      / \
//     0   1
//    / \
//   1   6
//  /
// 5
//
// Output: true
// Explanation: The path 1->0->1->6 matches the given sequence.
//
// Input: root = [1,0,1,1,6,5], sequence = [1,0,7]
//
//       1
//      / \
//     0   1
//    / \
//   1   6
//  /
// 5
//
// Output: false
// Explanation: There is no path that matches the given sequence.
