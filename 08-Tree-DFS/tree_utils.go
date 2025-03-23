package treedfs

import (
	"fmt"
)

// CreateTreeFromArray creates a binary tree from a level-order array representation
// nil values in the array represent null nodes
func CreateTreeFromArray(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := NewTreeNode(arr[0].(int))
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(arr) && arr[i] != nil {
			node.Left = NewTreeNode(arr[i].(int))
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(arr) && arr[i] != nil {
			node.Right = NewTreeNode(arr[i].(int))
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// TreeToArray converts a binary tree to its array representation
func TreeToArray(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	result := []interface{}{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == nil {
				result = append(result, nil)
			} else {
				result = append(result, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
	}

	// Trim trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

// PrintTree prints the tree in a readable format
func PrintTree(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	arr := TreeToArray(root)
	result := "["
	for i, val := range arr {
		if val == nil {
			result += "null"
		} else {
			result += fmt.Sprintf("%d", val)
		}

		if i < len(arr)-1 {
			result += ", "
		}
	}
	result += "]"

	return result
}

// GetTreeHeight returns the height of the binary tree
// The height is the number of edges on the longest path from the root to a leaf
func GetTreeHeight(root *TreeNode) int {
	if root == nil {
		return -1
	}

	leftHeight := GetTreeHeight(root.Left)
	rightHeight := GetTreeHeight(root.Right)

	// Return the maximum height plus 1 for the current level
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// IsBalanced checks if the binary tree is height-balanced
// A height-balanced tree is a tree where the difference between the heights
// of the left and right subtree of every node is not more than 1
func IsBalanced(root *TreeNode) bool {
	return checkBalanced(root) != -1
}

// checkBalanced returns the height of the tree if it's balanced, or -1 if it's not
func checkBalanced(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := checkBalanced(node.Left)
	if left == -1 {
		return -1
	}

	right := checkBalanced(node.Right)
	if right == -1 {
		return -1
	}

	// If the difference in heights is more than 1, the tree is not balanced
	if abs(left-right) > 1 {
		return -1
	}

	// Return the height of the current subtree
	return max(left, right) + 1
}

// abs returns the absolute value of an integer
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// max returns the larger of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
