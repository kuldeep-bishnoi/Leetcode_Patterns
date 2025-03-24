# Tree Depth-First Search (DFS) Pattern

## What is Tree DFS?

Imagine you're exploring a maze. DFS is like following one path as far as you can go until you hit a dead end, then going back and trying another path. It's like exploring a tree house by going all the way up one branch until you can't go further, then coming back down and trying another branch!

## Real-Life Examples

1. **Maze Solving**: When you're trying to find your way through a maze.
   - Go straight until you hit a wall
   - Turn back and try another path
   - Keep going until you find the exit

2. **Family Tree**: When you want to find all your ancestors.
   - Go up one branch of your family tree
   - When you reach the top, go back down
   - Try another branch
   - Keep going until you've explored all branches

3. **Book Reading**: When you're reading a book with many chapters.
   - Read one chapter completely
   - If there's a reference to another chapter, read that chapter
   - Come back to the original chapter
   - Keep going until you've read everything

## When Do We Use Tree DFS?

Use this technique when:
- You need to explore all paths in a tree
- You need to find the maximum depth of a tree
- You need to check if a path exists
- You need to find all paths from root to leaves
- You need to validate a binary search tree

## How Does It Work?

1. **Step 1**: Start with the root node
2. **Step 2**: Choose a direction (left or right)
3. **Step 3**: Go as far as you can in that direction
4. **Step 4**: When you can't go further, go back
5. **Step 5**: Try another direction
6. **Step 6**: Keep going until you've explored everything

## Simple Code Example

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    
    // Helper function to do the traversal
    var traverse func(*TreeNode)
    traverse = func(node *TreeNode) {
        if node == nil {
            return
        }
        
        // First, go left
        traverse(node.Left)
        
        // Then, process current node
        result = append(result, node.Val)
        
        // Finally, go right
        traverse(node.Right)
    }
    
    // Start traversal from root
    traverse(root)
    return result
}
```

## Common Mistakes to Avoid

1. **Stack Overflow**: Don't forget base case (nil check)
2. **Wrong Order**: Remember the order of traversal (pre, in, post)
3. **Missing Path**: Make sure to explore all paths
4. **State Management**: Keep track of visited nodes if needed

## Fun Practice Problems

1. **Path Finder**: Find all paths from root to leaves
2. **Depth Explorer**: Find the maximum depth of the tree
3. **Sum Path**: Find if any path sums to a target value
4. **Mirror Check**: Check if a tree is symmetric
5. **BST Validator**: Check if a tree is a valid binary search tree

## LeetCode Problems Using Tree DFS

Here are some popular LeetCode problems that can be solved using Tree DFS:

### Easy Problems

1. **[#94 Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)** - Traverse tree in inorder order.
   - **Approach**: Use recursive DFS with left-root-right order.

2. **[#100 Same Tree](https://leetcode.com/problems/same-tree/)** - Check if two trees are identical.
   - **Approach**: Use DFS to compare nodes.

### Medium Problems

1. **[#98 Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)** - Check if a tree is a valid BST.
   - **Approach**: Use DFS with min/max bounds.

2. **[#236 Lowest Common Ancestor](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)** - Find the lowest common ancestor of two nodes.
   - **Approach**: Use DFS to find paths to both nodes.

3. **[#113 Path Sum II](https://leetcode.com/problems/path-sum-ii/)** - Find all paths that sum to target.
   - **Approach**: Use DFS to track current path and sum.

4. **[#129 Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers/)** - Sum all root-to-leaf numbers.
   - **Approach**: Use DFS to build numbers and sum them.

### Hard Problems

1. **[#124 Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/)** - Find maximum path sum.
   - **Approach**: Use DFS to track local and global maximum.

2. **[#297 Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)** - Convert tree to string and back.
   - **Approach**: Use DFS to serialize and deserialize.

### Tips for Solving LeetCode Tree DFS Problems

1. **Traversal Order**: Choose the right order (pre, in, post)
2. **State Tracking**: Keep track of necessary state (sum, path, etc.)
3. **Base Cases**: Handle nil nodes and leaf nodes
4. **Recursion**: Use helper functions for cleaner code
5. **Global State**: Use global variables when needed

## Why Learn This Pattern?

The Tree DFS pattern is super useful because:
1. It's perfect for exploring all paths in a tree
2. It helps solve many tree-related problems
3. It's a favorite in coding interviews
4. It teaches important concepts about recursion
5. It's used in many real-world applications