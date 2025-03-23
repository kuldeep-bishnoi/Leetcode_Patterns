package treebfs

// LevelOrderSuccessor finds the level order successor of a given key in a binary tree
// The level order successor is the node that appears immediately after the given key in a level order traversal
// If the key is not found or has no successor, nil is returned
func LevelOrderSuccessor(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// Use a queue to perform level order traversal
	queue := []*TreeNode{root}

	// Continue traversal until the queue is empty
	for len(queue) > 0 {
		// Dequeue the next node
		current := queue[0]
		queue = queue[1:]

		// Enqueue the node's children
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}

		// If the current node is the key, return the next node in the queue (if any)
		if current.Val == key && len(queue) > 0 {
			return queue[0]
		}
	}

	// Key not found or has no successor
	return nil
}

// Example usage:
//
// Input: root = [1,2,3,4,5,null,6], key = 3
//
//       1
//      / \
//     2   3
//    / \   \
//   4   5   6
//
// Output: 4
// Explanation: The level order traversal is [1,2,3,4,5,6]. The level order successor of 3 is 4.
//
// Input: root = [12,7,1,9,10,5], key = 9
//
//       12
//      /  \
//     7    1
//    / \  /
//   9  10 5
//
// Output: 10
// Explanation: The level order traversal is [12,7,1,9,10,5]. The level order successor of 9 is 10.
//
// Input: root = [12,7,1,9,null,10,5], key = 5
//
//       12
//      /  \
//     7    1
//    /    / \
//   9    10  5
//
// Output: null
// Explanation: The key 5 has no successor in the level order traversal.
//
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(n) where n is the number of nodes in the tree
