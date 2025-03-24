# Tree Breadth-First Search (BFS) Pattern

## What is Tree BFS?

Imagine you're exploring a tree house with many levels. BFS is like exploring each level completely before moving to the next level. It's like checking all the rooms on the first floor before going up to the second floor, then checking all rooms on the second floor before going to the third floor, and so on!

## Real-Life Examples

1. **School Assembly**: When students line up by grade level.
   - First, all first graders line up
   - Then, all second graders line up
   - Then, all third graders line up
   - And so on...

2. **Family Tree**: When you want to find all your cousins.
   - First, look at your siblings
   - Then, look at your cousins (your parents' siblings' children)
   - Then, look at your second cousins
   - And so on...

3. **Building Floors**: When checking each floor of a building.
   - First, check all rooms on the first floor
   - Then, check all rooms on the second floor
   - Then, check all rooms on the third floor
   - And so on...

## When Do We Use Tree BFS?

Use this technique when:
- You need to process nodes level by level
- You need to find the shortest path in a tree
- You need to find the minimum depth of a tree
- You need to find the largest value in each level
- You need to connect nodes at the same level

## How Does It Work?

1. **Step 1**: Start with the root node
2. **Step 2**: Add the root to a queue
3. **Step 3**: While the queue is not empty:
   - Take out the first node from the queue
   - Process that node
   - Add its children to the queue
4. **Step 4**: Keep going until the queue is empty

## Simple Code Example

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    // Create a queue to store nodes
    queue := []*TreeNode{root}
    result := [][]int{}
    
    // Process each level
    for len(queue) > 0 {
        levelSize := len(queue)
        currentLevel := []int{}
        
        // Process all nodes at current level
        for i := 0; i < levelSize; i++ {
            // Get the first node from queue
            node := queue[0]
            queue = queue[1:]
            
            // Add node's value to current level
            currentLevel = append(currentLevel, node.Val)
            
            // Add children to queue
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        // Add current level to result
        result = append(result, currentLevel)
    }
    
    return result
}
```

## Common Mistakes to Avoid

1. **Forgetting Queue**: Always use a queue for BFS
2. **Level Tracking**: Make sure to track level size correctly
3. **Empty Tree**: Handle the case when the tree is empty
4. **Child Check**: Always check if children exist before adding them

## Fun Practice Problems

1. **Level Explorer**: Find the largest number in each level
2. **Cousin Finder**: Find all cousins of a given node
3. **Level Sum**: Calculate the sum of all numbers in each level
4. **Level Average**: Find the average of numbers in each level
5. **Level Connect**: Connect all nodes at the same level

## LeetCode Problems Using Tree BFS

Here are some popular LeetCode problems that can be solved using Tree BFS:

### Easy Problems

1. **[#102 Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)** - Return the level order traversal of a binary tree.
   - **Approach**: Use queue to process nodes level by level.

2. **[#637 Average of Levels in Binary Tree](https://leetcode.com/problems/average-of-levels-in-binary-tree/)** - Find the average value of nodes at each level.
   - **Approach**: Use BFS to calculate sum and count at each level.

### Medium Problems

1. **[#116 Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)** - Connect nodes at the same level.
   - **Approach**: Use BFS to connect nodes at each level.

2. **[#199 Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/)** - Return the rightmost node of each level.
   - **Approach**: Use BFS and take the last node of each level.

3. **[#515 Find Largest Value in Each Tree Row](https://leetcode.com/problems/find-largest-value-in-each-tree-row/)** - Find the largest value in each level.
   - **Approach**: Use BFS to track maximum at each level.

4. **[#429 N-ary Tree Level Order Traversal](https://leetcode.com/problems/n-ary-tree-level-order-traversal/)** - Level order traversal for N-ary tree.
   - **Approach**: Use BFS with multiple children.

### Hard Problems

1. **[#297 Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/)** - Convert tree to string and back.
   - **Approach**: Use BFS to serialize and deserialize level by level.

2. **[#987 Vertical Order Traversal of a Binary Tree](https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/)** - Traverse tree vertically.
   - **Approach**: Use BFS with column tracking.

### Tips for Solving LeetCode Tree BFS Problems

1. **Queue Usage**: Always use a queue for BFS
2. **Level Tracking**: Keep track of level size
3. **Node Processing**: Process all nodes at current level
4. **Child Handling**: Handle multiple children correctly
5. **Result Structure**: Organize results by level

## Why Learn This Pattern?

The Tree BFS pattern is super useful because:
1. It's perfect for level-by-level tree traversal
2. It helps find shortest paths in trees
3. It's a favorite in coding interviews
4. It teaches important concepts about tree traversal
5. It's used in many real-world applications
