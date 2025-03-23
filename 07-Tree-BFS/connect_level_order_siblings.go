package treebfs

// Node represents a node in a binary tree with a next pointer
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// ConnectLevelOrderSiblings connects each node with its level order successor
// Each node's next pointer is set to the node immediately to its right on the same level,
// or nil if it's the last node on its level
func ConnectLevelOrderSiblings(root *Node) *Node {
	if root == nil {
		return nil
	}

	// Start with the root node
	queue := []*Node{root}

	// Continue traversal until all levels are processed
	for len(queue) > 0 {
		levelSize := len(queue)

		// Process all nodes at the current level
		for i := 0; i < levelSize; i++ {
			// Dequeue the next node
			current := queue[0]
			queue = queue[1:]

			// If not the last node in the current level, connect to the next node
			if i < levelSize-1 {
				current.Next = queue[0]
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

	return root
}

// NewNode creates a new Node with the given value
func NewNode(val int) *Node {
	return &Node{Val: val}
}

// Example usage:
//
// Input: root = [1,2,3,4,5,null,7]
//
//       1                1 -> nil
//      / \              / \
//     2   3     =>     2 -> 3 -> nil
//    / \   \          / \   \
//   4   5   7        4 -> 5 -> 7 -> nil
//
// Output: [1,#,2,3,#,4,5,7,#]
// Explanation: The tree is connected level by level, with # representing the end of each level
//
// Input: root = [1,2,3,4,5,6,7]
//
//       1                1 -> nil
//      / \              / \
//     2   3     =>     2 -> 3 -> nil
//    / \ / \          / \ / \
//   4  5 6  7        4->5->6->7 -> nil
//
// Output: [1,#,2,3,#,4,5,6,7,#]
//
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(n) where n is the number of nodes in the tree
