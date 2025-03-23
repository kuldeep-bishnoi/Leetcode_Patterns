package treebfs

// LevelAverages calculates the average value of nodes at each level of a binary tree
// and returns a slice where each element is the average of a level
func LevelAverages(root *TreeNode) []float64 {
	result := []float64{}
	if root == nil {
		return result
	}

	// Use a queue to keep track of nodes at the current level
	queue := []*TreeNode{root}

	// Continue traversal until the queue is empty
	for len(queue) > 0 {
		// Get the size of the current level
		levelSize := len(queue)
		levelSum := 0

		// Process all nodes at the current level
		for i := 0; i < levelSize; i++ {
			// Dequeue the next node
			currentNode := queue[0]
			queue = queue[1:]

			// Add the node's value to the level sum
			levelSum += currentNode.Val

			// Enqueue the node's children for the next level
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}

		// Calculate the average and add it to the result
		levelAverage := float64(levelSum) / float64(levelSize)
		result = append(result, levelAverage)
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
// Output: [3.0, 14.5, 11.0]
// Explanation:
// Level 0 average: 3 / 1 = 3.0
// Level 1 average: (9 + 20) / 2 = 14.5
// Level 2 average: (15 + 7) / 2 = 11.0
//
// Input: root = [1]
//
//       1
//
// Output: [1.0]
//
// Input: root = []
//
// Output: []
//
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(n) where n is the number of nodes in the tree
