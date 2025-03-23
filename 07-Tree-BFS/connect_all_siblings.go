package treebfs

// ConnectAllSiblings connects each node with its level order successor across all levels
// Each node's next pointer points to the node that comes next in the level order traversal,
// regardless of whether it's on the same level or not
func ConnectAllSiblings(root *Node) *Node {
	if root == nil {
		return nil
	}

	// Start with the root node
	queue := []*Node{root}
	var prev *Node = nil

	// Continue traversal until all nodes are processed
	for len(queue) > 0 {
		// Dequeue the next node
		current := queue[0]
		queue = queue[1:]

		// Connect the previous node to the current node
		if prev != nil {
			prev.Next = current
		}
		prev = current

		// Enqueue the node's children for processing
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	// The last node's next pointer will be nil by default
	return root
}

// Example usage:
//
// Input: root = [1,2,3,4,5,null,7]
//
//       1                1 -> 2 -> 3 -> 4 -> 5 -> 7 -> nil
//      / \
//     2   3     =>
//    / \   \
//   4   5   7
//
// Output: [1,2,3,4,5,7,#]
// Explanation: All nodes are connected in level order
//
// Input: root = [1,2,3,4,5,6,7]
//
//       1                1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> nil
//      / \
//     2   3     =>
//    / \ / \
//   4  5 6  7
//
// Output: [1,2,3,4,5,6,7,#]
//
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(n) where n is the number of nodes in the tree
