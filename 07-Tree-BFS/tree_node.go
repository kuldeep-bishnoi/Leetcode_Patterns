package treebfs

import (
	"fmt"
)

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode creates a new TreeNode with the given value
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

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
