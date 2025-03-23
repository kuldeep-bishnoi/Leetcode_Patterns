package topsort

// Graph represents a directed graph using adjacency list
type Graph struct {
	Vertices int
	AdjList  [][]int
}

// NewGraph creates a new graph with the given number of vertices
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

// AddEdge adds a directed edge from u to v
func (g *Graph) AddEdge(u, v int) {
	g.AdjList[u] = append(g.AdjList[u], v)
}

// TopologicalSortBFS performs topological sorting using Kahn's algorithm (BFS-based)
// Returns the topological order and a boolean indicating if a valid ordering exists
// Time Complexity: O(V + E) where V is the number of vertices and E is the number of edges
// Space Complexity: O(V) for storing the in-degree array, queue, and result list
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
	// then there is a cycle in the graph (not a DAG)
	if visitedCount != g.Vertices {
		return nil, false
	}

	return topOrder, true
}

// TopologicalSortDFS performs topological sorting using Depth-First Search
// Returns the topological order and a boolean indicating if a valid ordering exists
// Time Complexity: O(V + E) where V is the number of vertices and E is the number of edges
// Space Complexity: O(V) for the visited array, recursion stack, and result stack
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

// topologicalSortUtil is a recursive utility function for DFS-based topological sort
// Returns true if a cycle is detected, false otherwise
func (g *Graph) topologicalSortUtil(v int, visited, recursionStack []bool, stack *[]int) bool {
	// Mark the current node as visited and part of recursion stack
	visited[v] = true
	recursionStack[v] = true

	// Recur for all adjacent vertices
	for _, u := range g.AdjList[v] {
		// If the vertex is not visited, then call the function recursively
		if !visited[u] {
			if g.topologicalSortUtil(u, visited, recursionStack, stack) {
				return true // Cycle detected
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

	return false // No cycle detected
}

// AllAncestors returns all ancestors for each node in the graph
// An ancestor of a node is any node that can reach that node through a directed path
// Time Complexity: O(V * (V + E)) as we run BFS/DFS for each vertex
// Space Complexity: O(V^2) in worst case for storing all ancestors
func (g *Graph) AllAncestors() [][]int {
	// Create transpose of the graph (reverse all edges)
	transpose := NewGraph(g.Vertices)
	for u := 0; u < g.Vertices; u++ {
		for _, v := range g.AdjList[u] {
			transpose.AddEdge(v, u)
		}
	}

	// For each vertex, find all vertices reachable from it in the transpose graph
	// These are the ancestors in the original graph
	ancestors := make([][]int, g.Vertices)

	for v := 0; v < g.Vertices; v++ {
		// Run DFS from vertex v in the transpose graph
		visited := make([]bool, g.Vertices)
		transpose.dfsForAncestors(v, visited)

		// Add all visited vertices (except v itself) as ancestors of v
		for u := 0; u < g.Vertices; u++ {
			if visited[u] && u != v {
				ancestors[v] = append(ancestors[v], u)
			}
		}
	}

	return ancestors
}

// dfsForAncestors is a utility function for the AllAncestors method
func (g *Graph) dfsForAncestors(v int, visited []bool) {
	// Mark the current node as visited
	visited[v] = true

	// Recur for all adjacent vertices
	for _, u := range g.AdjList[v] {
		if !visited[u] {
			g.dfsForAncestors(u, visited)
		}
	}
}

// Example usage:
//
// g := NewGraph(6)
// g.AddEdge(5, 2)
// g.AddEdge(5, 0)
// g.AddEdge(4, 0)
// g.AddEdge(4, 1)
// g.AddEdge(2, 3)
// g.AddEdge(3, 1)
//
// order, isDAG := g.TopologicalSortBFS()
// // order = [4 5 0 2 3 1], isDAG = true
//
// // Graph with cycle
// g := NewGraph(3)
// g.AddEdge(0, 1)
// g.AddEdge(1, 2)
// g.AddEdge(2, 0)
//
// order, isDAG := g.TopologicalSortDFS()
// // order = nil, isDAG = false
