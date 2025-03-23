package treebfs

// LevelOrderTraversal performs a level order traversal of a binary tree
// and returns a slice of slices where each inner slice contains the values of nodes at the same level
func LevelOrderTraversal(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	// Use a queue to keep track of nodes at the current level
	queue := []*TreeNode{root}

	// Continue traversal until the queue is empty
	for len(queue) > 0 {
		// Get the size of the current level
		levelSize := len(queue)
		currentLevel := []int{}

		// Process all nodes at the current level
		for i := 0; i < levelSize; i++ {
			// Dequeue the next node
			currentNode := queue[0]
			queue = queue[1:]

			// Add the node's value to the current level
			currentLevel = append(currentLevel, currentNode.Val)

			// Enqueue the node's children for the next level
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}

		// Add the current level to the result
		result = append(result, currentLevel)
	}

	return result
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
// Output: [[3],[9,20],[15,7]]
//
// Input: root = [1]
//
//       1
//
// Output: [[1]]
//
// Input: root = []
//
// Output: []
//
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(n) where n is the number of nodes in the tree
