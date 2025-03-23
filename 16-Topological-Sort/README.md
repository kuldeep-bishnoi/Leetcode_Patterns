# Topological Sort Pattern

## Introduction

Topological Sort is a linear ordering of vertices in a directed acyclic graph (DAG) such that for every directed edge (u, v), vertex u comes before vertex v in the ordering. In simpler terms, if there is a path from vertex u to vertex v, then vertex u comes before vertex v in the ordering.

This pattern is particularly useful for solving problems related to:
- Scheduling tasks with dependencies
- Course prerequisites
- Build systems
- Package dependencies
- Task sequencing
- Any problem where certain things must happen before others

A key insight: Topological sort is only possible if the graph has no directed cycles (i.e., it is a DAG). If a cycle exists, no valid sequencing is possible.

## How It Works

There are two main algorithms for topological sorting:

### 1. Kahn's Algorithm (BFS-based)

1. Calculate the in-degree (number of incoming edges) for each vertex
2. Identify vertices with an in-degree of 0 (no dependencies) and add them to a queue
3. While the queue is not empty:
   - Remove a vertex from the queue
   - Add it to the result list
   - Reduce the in-degree of all adjacent vertices by 1
   - If any adjacent vertex's in-degree becomes 0, add it to the queue
4. If all vertices are in the result list, return the list; otherwise, the graph has a cycle

### 2. DFS-based Algorithm

1. Initialize a boolean array to track visited vertices
2. Initialize a stack (or list) to store the result
3. For each unvisited vertex:
   - Perform a DFS traversal
   - After exploring all neighbors of a vertex, push it onto the stack
4. The topological sort is obtained by popping the stack

## Time and Space Complexity

For both algorithms:
- **Time Complexity**: O(V + E) where V is the number of vertices and E is the number of edges
- **Space Complexity**: O(V) for storing the in-degree array, queue, and result list

## When to Use Topological Sort

Consider using topological sort when:
1. You have a directed graph with dependencies
2. You need to find a valid ordering that respects those dependencies
3. You're dealing with sequential tasks where some must finish before others can start
4. You need to detect if a valid ordering exists (cycle detection)

## Common Problem Patterns

The Topological Sort pattern is commonly used to solve the following types of problems:

1. **Task Scheduling**: Given a set of tasks and dependencies, find a valid execution order
2. **Course Scheduling**: Determine if a student can finish all courses given prerequisites
3. **Build Order**: Find the build order of projects with dependencies
4. **Package Dependencies**: Resolve package installation order based on dependencies
5. **Cycle Detection**: Determine if a valid ordering exists or if there's a circular dependency
6. **Alien Dictionary**: Given a sorted dictionary of an alien language, find the order of characters

## Implementation in Golang

Here are the implementations of both approaches in Go:

```go
// Graph representation using adjacency list
type Graph struct {
    Vertices int
    AdjList  [][]int
}

// Create a new graph
func NewGraph(vertices int) *Graph {
    adjList := make([][]int, vertices)
    for i := range adjList {
        adjList[i] = []int{}
    }
    return &Graph{
        Vertices: vertices,
        AdjList:  adjList,
    }
}

// Add directed edge from u to v
func (g *Graph) AddEdge(u, v int) {
    g.AdjList[u] = append(g.AdjList[u], v)
}

// Kahn's Algorithm (BFS-based)
func (g *Graph) TopologicalSortBFS() ([]int, bool) {
    // Calculate in-degree of all vertices
    inDegree := make([]int, g.Vertices)
    for u := 0; u < g.Vertices; u++ {
        for _, v := range g.AdjList[u] {
            inDegree[v]++
        }
    }
    
    // Create a queue and enqueue all vertices with in-degree 0
    queue := []int{}
    for i := 0; i < g.Vertices; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }
    
    // Initialize count of visited vertices
    visitedCount := 0
    
    // Store topological order
    topOrder := []int{}
    
    // Process until queue is empty
    for len(queue) > 0 {
        // Dequeue a vertex
        u := queue[0]
        queue = queue[1:]
        
        // Add it to topological order
        topOrder = append(topOrder, u)
        
        // Increment count of visited vertices
        visitedCount++
        
        // Reduce in-degree of adjacent vertices
        for _, v := range g.AdjList[u] {
            inDegree[v]--
            // If in-degree becomes 0, add it to queue
            if inDegree[v] == 0 {
                queue = append(queue, v)
            }
        }
    }
    
    // If count of visited vertices is not equal to the number of vertices,
    // then there is a cycle in the graph
    if visitedCount != g.Vertices {
        return nil, false
    }
    
    return topOrder, true
}

// DFS-based topological sort
func (g *Graph) TopologicalSortDFS() ([]int, bool) {
    // Create a stack to store result
    stack := []int{}
    
    // Mark all vertices as not visited
    visited := make([]bool, g.Vertices)
    
    // To detect cycle
    recursionStack := make([]bool, g.Vertices)
    
    // Call the recursive helper function for all unvisited vertices
    for i := 0; i < g.Vertices; i++ {
        if !visited[i] {
            if g.topologicalSortUtil(i, visited, recursionStack, &stack) {
                // If topologicalSortUtil returns true, there is a cycle
                return nil, false
            }
        }
    }
    
    // Reverse the stack to get the correct order
    for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
        stack[i], stack[j] = stack[j], stack[i]
    }
    
    return stack, true
}

// Recursive utility function for DFS-based topological sort
func (g *Graph) topologicalSortUtil(v int, visited, recursionStack []bool, stack *[]int) bool {
    // Mark the current node as visited and part of recursion stack
    visited[v] = true
    recursionStack[v] = true
    
    // Recur for all adjacent vertices
    for _, u := range g.AdjList[v] {
        // If the vertex is not visited, then call the function recursively
        if !visited[u] {
            if g.topologicalSortUtil(u, visited, recursionStack, stack) {
                return true
            }
        } else if recursionStack[u] {
            // If the vertex is visited and is also part of the recursion stack,
            // then there is a cycle
            return true
        }
    }
    
    // Remove the vertex from recursion stack
    recursionStack[v] = false
    
    // Push current vertex to stack which stores the result
    *stack = append(*stack, v)
    
    return false
}
```

## Example Problems

1. **Course Schedule**: Given the total number of courses and a list of prerequisite pairs, determine if it's possible to finish all courses.
2. **Alien Dictionary**: Given a sorted dictionary of an alien language, find the order of characters.
3. **Minimum Height Trees**: Find the minimum height trees in a graph.
4. **Task Scheduler**: Schedule tasks with dependencies to minimize the total time.
5. **Build Order**: Find a valid build order for projects with dependencies.
6. **All Ancestors of a Node in DAG**: Find all ancestors of each node in a DAG.

Each of these problems leverages the topological sort pattern to solve complex dependency-based scenarios effectively. 