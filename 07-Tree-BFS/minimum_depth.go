package treebfs

// MinimumDepth finds the minimum depth of a binary tree
// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node
func MinimumDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Use a queue to keep track of nodes and their depths
	type NodeWithDepth struct {
		node  *TreeNode
		depth int
	}

	queue := []NodeWithDepth{{node: root, depth: 1}}

	// Continue traversal until the queue is empty
	for len(queue) > 0 {
		// Dequeue the next node and its depth
		current := queue[0]
		queue = queue[1:]

		// If we find a leaf node, return its depth
		if current.node.Left == nil && current.node.Right == nil {
			return current.depth
		}

		// Enqueue the node's children with their depths
		if current.node.Left != nil {
			queue = append(queue, NodeWithDepth{node: current.node.Left, depth: current.depth + 1})
		}
		if current.node.Right != nil {
			queue = append(queue, NodeWithDepth{node: current.node.Right, depth: current.depth + 1})
		}
	}

	// This line should not be reached for a valid binary tree with at least one node
	return 0
}

// MinimumDepthAlternative is another way to find the minimum depth of a binary tree
// It processes nodes level by level without using a separate struct
func MinimumDepthAlternative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	// Continue traversal until the queue is empty
	for len(queue) > 0 {
		depth++
		levelSize := len(queue)

		// Process all nodes at the current level
		for i := 0; i < levelSize; i++ {
			// Dequeue the next node
			current := queue[0]
			queue = queue[1:]

			// If we find a leaf node, return the current depth
			if current.Left == nil && current.Right == nil {
				return depth
			}

			// Enqueue the node's children for the next level
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}

	// This line should not be reached for a valid binary tree with at least one node
	return 0
}

// Example usage:
//
// Input: root = [3,9,20,null,null,15,7]
//
//       3
//      / \
//     9  20
//       /  \
//      15   7
//
// Output: 2
// Explanation: The shortest path from the root node to a leaf node is: 3 -> 9, which has a depth of 2.
//
// Input: root = [2,null,3,null,4,null,5,null,6]
//
//       2
//        \
//         3
//          \
//           4
//            \
//             5
//              \
//               6
//
// Output: 5
// Explanation: The shortest path from the root node to a leaf node is: 2 -> 3 -> 4 -> 5 -> 6, which has a depth of 5.
//
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(n) where n is the number of nodes in the tree
