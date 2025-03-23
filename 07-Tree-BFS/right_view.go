package treebfs

// RightView returns the right side view of a binary tree
// The right side view of a binary tree is the set of nodes visible when looking at the tree from the right side
// It includes the rightmost node at each level of the tree
func RightView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	// Use a queue to perform level order traversal
	queue := []*TreeNode{root}

	// Continue traversal until the queue is empty
	for len(queue) > 0 {
		levelSize := len(queue)

		// Process all nodes at the current level
		for i := 0; i < levelSize; i++ {
			// Dequeue the next node
			current := queue[0]
			queue = queue[1:]

			// If this is the last node in the current level, add it to the result
			if i == levelSize-1 {
				result = append(result, current.Val)
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

	return result
}

// Example usage:
//
// Input: root = [1,2,3,null,5,null,4]
//
//       1
//      / \
//     2   3
//      \   \
//       5   4
//
// Output: [1,3,4]
// Explanation: The right side view includes the rightmost node at each level: 1, 3, 4
//
// Input: root = [1,null,3]
//
//       1
//        \
//         3
//
// Output: [1,3]
//
// Input: root = []
//
// Output: []
//
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(n) where n is the number of nodes in the tree
