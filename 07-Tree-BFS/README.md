# Tree Breadth First Search (BFS) Pattern

## Introduction

The Tree Breadth First Search (BFS) pattern is used to traverse a tree in a level-by-level manner. This pattern explores all nodes at the current level before moving to the next level. Tree BFS uses a queue data structure to keep track of the nodes that need to be processed.

This pattern is particularly useful for problems where you need to find the shortest path, explore all neighbors, or process nodes in a level-by-level fashion.

## How It Works

1. Start with the root node and push it into a queue
2. While the queue is not empty:
   - Pop a node from the queue
   - Process the node
   - Push its children (left and right for a binary tree) into the queue
3. Continue until the queue is empty

## Time and Space Complexity

- **Time Complexity**: O(n) - where n is the number of nodes in the tree
- **Space Complexity**: O(w) - where w is the maximum width of the tree (the maximum number of nodes at any level)

## When to Use Tree BFS

Use the Tree BFS pattern when:
- You need to traverse a tree level by level
- You need to find the minimum depth of a tree
- You're looking for nodes closest to the root
- You need to connect nodes at the same level
- You're dealing with problems involving the zigzag traversal of trees

## Common Problem Patterns

1. **Level Order Traversal**
   - Traverse the tree level by level
   - Return nodes at each level as separate lists

2. **Zigzag Level Order Traversal**
   - Traverse the tree level by level
   - Alternate between traversing left to right and right to left

3. **Level Averages**
   - Find the average of values at each level of the tree

4. **Minimum Depth of a Binary Tree**
   - Find the minimum depth (shortest path from root to a leaf node)

5. **Connect Level Order Siblings**
   - Connect each node with its level order successor

## Implementation in Golang

```go
// TreeNode represents a node in a binary tree
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Level Order Traversal
func levelOrder(root *TreeNode) [][]int {
    result := [][]int{}
    if root == nil {
        return result
    }
    
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        currentLevel := []int{}
        
        for i := 0; i < levelSize; i++ {
            currentNode := queue[0]
            queue = queue[1:]
            
            // Add the node to the current level
            currentLevel = append(currentLevel, currentNode.Val)
            
            // Add the node's children to the queue
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

// Find Minimum Depth
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    queue := []*TreeNode{root}
    depth := 0
    
    for len(queue) > 0 {
        depth++
        levelSize := len(queue)
        
        for i := 0; i < levelSize; i++ {
            currentNode := queue[0]
            queue = queue[1:]
            
            // If we found a leaf node, return the current depth
            if currentNode.Left == nil && currentNode.Right == nil {
                return depth
            }
            
            // Add the node's children to the queue
            if currentNode.Left != nil {
                queue = append(queue, currentNode.Left)
            }
            if currentNode.Right != nil {
                queue = append(queue, currentNode.Right)
            }
        }
    }
    
    return depth
}
```

## Example Problems

1. **Binary Tree Level Order Traversal**
2. **Zigzag Traversal**
3. **Level Averages in a Binary Tree**
4. **Minimum Depth of a Binary Tree**
5. **Level Order Successor**
6. **Connect Level Order Siblings**
7. **Connect All Level Order Siblings**
8. **Right View of a Binary Tree**

Each of these problems has a dedicated solution file in this directory. 