package topsort

// FindMinHeightTrees finds the nodes that, if used as roots, would result in trees with minimum height
// n: number of nodes (labeled from 0 to n-1)
// edges: list of undirected edges between nodes
// Returns a list of node IDs that can be roots of minimum height trees
//
// Time Complexity: O(n) where n is the number of nodes
// Space Complexity: O(n)
func FindMinHeightTrees(n int, edges [][]int) []int {
	if n == 0 {
		return []int{}
	}

	if n == 1 {
		return []int{0} // Only one node, so it's the only possible root
	}

	// Build the undirected graph using adjacency lists
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}

	// Track the degree of each node (number of connections)
	degrees := make([]int, n)

	// Add undirected edges
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
		degrees[u]++
		degrees[v]++
	}

	// Add all leaf nodes (nodes with only one connection) to the queue
	queue := []int{}
	for i := 0; i < n; i++ {
		if degrees[i] == 1 {
			queue = append(queue, i)
		}
	}

	// Keep trimming leaf nodes until we have 1 or 2 nodes left
	// These will be the centroids of the graph and thus the roots of MHTs
	remainingNodes := n
	for remainingNodes > 2 {
		size := len(queue)
		remainingNodes -= size

		// Process all current leaf nodes
		for i := 0; i < size; i++ {
			leaf := queue[0]
			queue = queue[1:] // Dequeue

			// Remove the leaf from its only neighbor
			for _, neighbor := range graph[leaf] {
				degrees[neighbor]--
				// If the neighbor becomes a leaf, add it to the queue
				if degrees[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	// The remaining nodes (1 or 2) are the MHT roots
	return queue
}

// FindHeights calculates the height of a tree rooted at each node
// n: number of nodes (labeled from 0 to n-1)
// edges: list of undirected edges between nodes
// Returns a slice where index i contains the height of the tree rooted at node i
//
// Time Complexity: O(n²) where n is the number of nodes
// Space Complexity: O(n)
func FindHeights(n int, edges [][]int) []int {
	// Build the undirected graph
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	heights := make([]int, n)

	// For each possible root
	for i := 0; i < n; i++ {
		// BFS to find the farthest node
		visited := make([]bool, n)
		queue := []int{i}
		depth := -1
		visited[i] = true

		for len(queue) > 0 {
			size := len(queue)
			depth++

			for j := 0; j < size; j++ {
				node := queue[0]
				queue = queue[1:] // Dequeue

				for _, neighbor := range graph[node] {
					if !visited[neighbor] {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}
		}

		heights[i] = depth
	}

	return heights
}

// FindFarthestNode finds the farthest node from the start node in an undirected graph
// graph: adjacency list representation of the graph
// start: starting node
// Returns the farthest node and its distance from the start node
//
// Time Complexity: O(n) where n is the number of nodes
// Space Complexity: O(n)
func FindFarthestNode(graph [][]int, start int) (int, int) {
	n := len(graph)
	visited := make([]bool, n)
	queue := []int{start}
	visited[start] = true

	distance := -1
	var lastNode int

	for len(queue) > 0 {
		size := len(queue)
		distance++

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:] // Dequeue
			lastNode = node

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return lastNode, distance
}

// FindTreeDiameter finds the diameter of a tree (longest path between any two nodes)
// n: number of nodes (labeled from 0 to n-1)
// edges: list of undirected edges between nodes
// Returns the diameter of the tree
//
// Time Complexity: O(n) where n is the number of nodes
// Space Complexity: O(n)
func FindTreeDiameter(n int, edges [][]int) int {
	if n <= 1 {
		return 0
	}

	// Build the undirected graph
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Step 1: Start from any node (e.g., node 0) and find the farthest node x
	farthestNode, _ := FindFarthestNode(graph, 0)

	// Step 2: Start from node x and find the farthest node y
	// The distance between x and y is the diameter
	_, diameter := FindFarthestNode(graph, farthestNode)

	return diameter
}

// Example usage:
//
// n := 4
// edges := [][]int{{1, 0}, {1, 2}, {1, 3}}
// roots := FindMinHeightTrees(n, edges)
// // roots = [1] (node 1 is the only MHT root)
//
// n := 6
// edges := [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}}
// roots := FindMinHeightTrees(n, edges)
// // roots = [3, 4] (nodes 3 and 4 can be MHT roots)
//
// heights := FindHeights(6, [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}})
// // heights = [2, 2, 2, 1, 2, 3] (trees rooted at node 3 have minimum height)
//
// diameter := FindTreeDiameter(6, [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}})
// // diameter = 3 (longest path is from node 5 to node 0/1/2)
