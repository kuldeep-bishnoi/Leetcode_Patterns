# Tree Depth First Search (DFS) Pattern

## Introduction

The Tree Depth First Search (DFS) pattern is a technique used to traverse or search tree data structures by exploring as far as possible along each branch before backtracking. Unlike Breadth First Search (BFS) which explores all nodes at the current level before moving to the next level, DFS explores a branch to its leaf before moving to the next branch.

This pattern is particularly useful for problems where you need to explore paths, find specific nodes, or check properties that involve traversing from root to leaf.

## How It Works

The Tree DFS pattern can be implemented using three main traversal methods:

1. **Preorder Traversal (Root → Left → Right)**
   - Visit the current node
   - Recursively traverse the left subtree
   - Recursively traverse the right subtree

2. **Inorder Traversal (Left → Root → Right)**
   - Recursively traverse the left subtree
   - Visit the current node
   - Recursively traverse the right subtree

3. **Postorder Traversal (Left → Right → Root)**
   - Recursively traverse the left subtree
   - Recursively traverse the right subtree
   - Visit the current node

All of these traversals can be implemented either recursively or iteratively (using a stack).

## Time and Space Complexity

- **Time Complexity**: O(n) - where n is the number of nodes in the tree
- **Space Complexity**: 
  - Best case: O(h) where h is the height of the tree (call stack space)
  - Worst case: O(n) for a skewed tree where h = n

## When to Use Tree DFS

Use the Tree DFS pattern when:
- You need to explore all paths from root to leaf nodes
- You need to search for a specific node or path in the tree
- You need to check or collect information about all the tree paths
- You need to determine properties related to tree depth or height
- You need to validate tree structures or properties

## Common Problem Patterns

1. **Path Sum Problems**
   - Find if there exists a root-to-leaf path that sums to a target
   - Find all root-to-leaf paths that sum to a target
   - Find paths (not necessarily root-to-leaf) that sum to a target
   
2. **Tree Properties Validation**
   - Check if a binary tree is balanced
   - Validate if a tree is a valid Binary Search Tree (BST)
   - Find the diameter of a binary tree
   
3. **Tree Transformation Problems**
   - Flatten a binary tree to a linked list
   - Convert a binary tree to its mirror image
   - Serialize and deserialize a binary tree
   
4. **Lowest Common Ancestor (LCA) Problems**
   - Find the LCA of two nodes in a binary tree
   - Find the LCA in a binary search tree

## Implementation in Golang

```go
// Binary Tree DFS traversals

// Preorder Traversal (Root → Left → Right)
func preorderTraversal(root *TreeNode) []int {
    result := []int{}
    preorderDFS(root, &result)
    return result
}

func preorderDFS(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    
    // Visit the current node
    *result = append(*result, node.Val)
    
    // Traverse left subtree
    preorderDFS(node.Left, result)
    
    // Traverse right subtree
    preorderDFS(node.Right, result)
}

// Inorder Traversal (Left → Root → Right)
func inorderTraversal(root *TreeNode) []int {
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
    
    // Visit the current node
    *result = append(*result, node.Val)
    
    // Traverse right subtree
    inorderDFS(node.Right, result)
}

// Postorder Traversal (Left → Right → Root)
func postorderTraversal(root *TreeNode) []int {
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
    
    // Visit the current node
    *result = append(*result, node.Val)
}
```

## Example Problems

1. **Binary Tree Path Sum**
2. **All Paths for a Sum**
3. **Sum of Path Numbers**
4. **Path With Given Sequence**
5. **Count Paths for a Sum**
6. **Tree Diameter**
7. **Binary Tree Maximum Path Sum**
8. **Validate Binary Search Tree (BST)**

Each of these problems has a dedicated solution file in this directory. 