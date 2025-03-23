package treedfs

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

// PreorderTraversal performs a root-left-right traversal of the binary tree
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (call stack)
func PreorderTraversal(root *TreeNode) []int {
	result := []int{}
	preorderDFS(root, &result)
	return result
}

func preorderDFS(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	// Visit the current node (Root)
	*result = append(*result, node.Val)

	// Traverse left subtree
	preorderDFS(node.Left, result)

	// Traverse right subtree
	preorderDFS(node.Right, result)
}

// InorderTraversal performs a left-root-right traversal of the binary tree
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (call stack)
func InorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorderDFS(root, &result)
	return result
}

func inorderDFS(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	// Traverse left subtree
	inorderDFS(node.Left, result)

	// Visit the current node (Root)
	*result = append(*result, node.Val)

	// Traverse right subtree
	inorderDFS(node.Right, result)
}

// PostorderTraversal performs a left-right-root traversal of the binary tree
// Time Complexity: O(n) where n is the number of nodes in the tree
// Space Complexity: O(h) where h is the height of the tree (call stack)
func PostorderTraversal(root *TreeNode) []int {
	result := []int{}
	postorderDFS(root, &result)
	return result
}

func postorderDFS(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	// Traverse left subtree
	postorderDFS(node.Left, result)

	// Traverse right subtree
	postorderDFS(node.Right, result)

	// Visit the current node (Root)
	*result = append(*result, node.Val)
}

// Example usage:
//
// Input: root = [1,null,2,3]
//
//      1
//       \
//        2
//       /
//      3
//
// Preorder output: [1,2,3]
// Inorder output: [1,3,2]
// Postorder output: [3,2,1]
//
// Input: root = [1]
//
//      1
//
// Preorder output: [1]
// Inorder output: [1]
// Postorder output: [1]
//
// Input: root = []
//
// Preorder output: []
// Inorder output: []
// Postorder output: []
