package subsets

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// GenerateTrees generates all structurally unique BSTs that store values 1...n
// Time Complexity: O(4^n / sqrt(n)) - Catalan number
// Space Complexity: O(4^n / sqrt(n)) for storing all unique trees
func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTreesHelper(1, n)
}

// generateTreesHelper generates all unique BSTs with values from start to end
func generateTreesHelper(start, end int) []*TreeNode {
	result := []*TreeNode{}

	// Base case
	if start > end {
		// Return null node for empty subtree
		result = append(result, nil)
		return result
	}

	// Try each value as the root
	for i := start; i <= end; i++ {
		// Generate all possible left subtrees
		leftSubtrees := generateTreesHelper(start, i-1)

		// Generate all possible right subtrees
		rightSubtrees := generateTreesHelper(i+1, end)

		// Connect left and right subtrees to the root i
		for _, left := range leftSubtrees {
			for _, right := range rightSubtrees {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				result = append(result, root)
			}
		}
	}

	return result
}

// NumTrees calculates the number of structurally unique BSTs that store values 1...n
// using dynamic programming
// Time Complexity: O(n^2)
// Space Complexity: O(n)
func NumTrees(n int) int {
	// dp[i] = number of unique BST with i nodes
	dp := make([]int, n+1)
	dp[0] = 1 // Empty tree is one valid BST

	for i := 1; i <= n; i++ {
		// Calculate dp[i] using the Catalan number recurrence relation
		for j := 1; j <= i; j++ {
			// j as root: dp[j-1] * dp[i-j]
			// where dp[j-1] = number of unique left subtrees
			// and dp[i-j] = number of unique right subtrees
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}

// NumTreesMath calculates the number of structurally unique BSTs using the Catalan number formula
// Time Complexity: O(n)
// Space Complexity: O(1)
func NumTreesMath(n int) int {
	// Catalan number formula: C(n) = (2n)! / ((n+1)! * n!)
	// Computed iteratively to avoid overflow

	// Start with C(0) = 1
	catalan := 1

	for i := 0; i < n; i++ {
		// C(n+1) = C(n) * 2*(2n+1)/(n+2)
		catalan = catalan * 2 * (2*i + 1) / (i + 2)
	}

	return catalan
}

// Example usage:
//
// Input: n = 3
// Output for NumTrees: 5
// Explanation: There are 5 structurally unique BST's that store values 1, 2, 3.
//
// For GenerateTrees(3), it would return all 5 unique tree structures:
// 1. [1,null,2,null,3]
// 2. [1,null,3,2]
// 3. [2,1,3]
// 4. [3,1,null,null,2]
// 5. [3,2,null,1]
