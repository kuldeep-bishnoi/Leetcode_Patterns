# Topological Sort Pattern

## What is Topological Sort?

Imagine you're building a LEGO castle! You can't put the roof on before building the walls, and you can't build the walls before laying the foundation. Topological Sort is like creating a step-by-step plan to build your castle in the right order. It helps us arrange things in a sequence where each step only happens after all the steps it depends on are completed.

## Real-Life Examples

1. **Building a Castle**:
   - Foundation must be built first
   - Then walls
   - Then windows and doors
   - Finally, the roof
   - Each part depends on parts below it

2. **Getting Ready for School**:
   - Wake up first
   - Then brush teeth
   - Then get dressed
   - Then eat breakfast
   - Finally, pack bag and leave

3. **Making a Sandwich**:
   - Get bread first
   - Then add filling
   - Then add toppings
   - Finally, close sandwich
   - Each step needs previous steps done

## When Do We Use Topological Sort?

Use this technique when:
- You have tasks with dependencies
- You need to find a valid order
- You have a directed graph
- You want to detect cycles
- You need to schedule tasks

## How Does It Work?

1. **Step 1**: Find tasks with no dependencies
2. **Step 2**: Add these tasks to our order
3. **Step 3**: Remove these tasks and their connections
4. **Step 4**: Repeat until all tasks are done

Example:
```
Tasks: A, B, C, D
Dependencies:
A → B
A → C
B → D
C → D

Order: A → B → C → D
or
Order: A → C → B → D
```

## Simple Code Example

```go
func topologicalSort(numCourses int, prerequisites [][]int) []int {
    // Create graph and count dependencies
    graph := make(map[int][]int)
    inDegree := make([]int, numCourses)
    
    // Build graph
    for _, pre := range prerequisites {
        graph[pre[1]] = append(graph[pre[1]], pre[0])
        inDegree[pre[0]]++
    }
    
    // Find nodes with no dependencies
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }
    
    // Process nodes
    result := []int{}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        result = append(result, node)
        
        // Update dependencies
        for _, next := range graph[node] {
            inDegree[next]--
            if inDegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }
    
    return result
}
```

## Common Mistakes to Avoid

1. **Cycle Detection**: Check for cycles
2. **Dependency Count**: Update counts correctly
3. **Queue Management**: Handle queue properly
4. **Result Order**: Maintain correct order

## Fun Practice Problems

1. **Castle Builder**: Plan castle building steps
2. **School Schedule**: Plan morning routine
3. **Sandwich Maker**: Plan sandwich making
4. **Game Levels**: Plan game level order
5. **Story Writer**: Plan story chapter order

## LeetCode Problems Using Topological Sort

Here are some popular LeetCode problems that can be solved using Topological Sort:

### Easy Problems

1. **[#207 Course Schedule](https://leetcode.com/problems/course-schedule/)** - Check if courses can be completed.
   - **Approach**: Use topological sort to detect cycles.

2. **[#210 Course Schedule II](https://leetcode.com/problems/course-schedule-ii/)** - Find course order.
   - **Approach**: Use topological sort to get order.

### Medium Problems

1. **[#269 Alien Dictionary](https://leetcode.com/problems/alien-dictionary/)** - Find alien language order.
   - **Approach**: Use topological sort with graph building.

2. **[#310 Minimum Height Trees](https://leetcode.com/problems/minimum-height-trees/)** - Find center nodes.
   - **Approach**: Use topological sort with leaf removal.

### Hard Problems

1. **[#329 Longest Increasing Path](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/)** - Find longest path.
   - **Approach**: Use topological sort with path tracking.

2. **[#1203 Sort Items by Groups](https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/)** - Sort with groups.
   - **Approach**: Use topological sort with group handling.

### Tips for Solving LeetCode Topological Sort Problems

1. **Graph Building**: Build graph properly
   - Create adjacency list
   - Count dependencies
   - Handle edge cases

2. **Processing**: Process nodes correctly
   - Start with no dependencies
   - Update counts properly
   - Handle cycles

3. **Optimization**: Use efficient structures
   - Queue for processing
   - Map for graph
   - Array for counts

4. **Edge Cases**: Handle special cases
   - Empty graph
   - Single node
   - Multiple solutions

## Why Learn This Pattern?

The Topological Sort pattern is super useful because:
1. It helps solve scheduling problems
2. It's used in build systems
3. It helps detect cycles
4. It's important in project planning
5. It's a fundamental graph algorithm