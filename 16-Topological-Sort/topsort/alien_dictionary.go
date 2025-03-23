package topsort

// AlienOrder determines the order of characters in an alien language
// words: array of strings sorted lexicographically in the alien language
// Returns the order of characters as a string, or an empty string if no valid order exists
//
// Time Complexity: O(C) where C is the total length of all words
// Space Complexity: O(1) since the alphabet size is constant (26 characters)
func AlienOrder(words []string) string {
	// Build a set of all characters in the alien language
	charSet := make(map[byte]bool)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			charSet[word[i]] = true
		}
	}

	// Count the number of unique characters
	n := len(charSet)
	if n == 0 {
		return ""
	}

	// Map each character to a unique vertex ID
	charToID := make(map[byte]int)
	idToChar := make(map[int]byte)

	id := 0
	for char := range charSet {
		charToID[char] = id
		idToChar[id] = char
		id++
	}

	// Create the graph
	graph := NewGraph(n)

	// Add edges based on character order
	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]

		// Find the first differing character
		minLen := min(len(word1), len(word2))

		// Check if word2 is a prefix of word1, which is invalid in a sorted dictionary
		if len(word1) > len(word2) && word1[:len(word2)] == word2 {
			return ""
		}

		// Find the first differing character
		for j := 0; j < minLen; j++ {
			if word1[j] != word2[j] {
				// Add edge from word1[j] to word2[j]
				graph.AddEdge(charToID[word1[j]], charToID[word2[j]])
				break
			}
		}
	}

	// Perform topological sort
	order, isDAG := graph.TopologicalSortDFS()
	if !isDAG {
		return "" // No valid ordering exists
	}

	// Convert the order of IDs back to characters
	result := make([]byte, n)
	for i, id := range order {
		result[i] = idToChar[id]
	}

	return string(result)
}

// FindAlienCycles identifies conflicting character orderings in alien words
// words: array of strings claimed to be sorted lexicographically in the alien language
// Returns a list of character pairs that form cycles (conflicts)
//
// Time Complexity: O(C) where C is the total length of all words
// Space Complexity: O(1) since the alphabet size is constant (26 characters)
func FindAlienCycles(words []string) [][2]byte {
	// Build a set of all characters in the alien language
	charSet := make(map[byte]bool)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			charSet[word[i]] = true
		}
	}

	// Count the number of unique characters
	n := len(charSet)
	if n == 0 {
		return nil
	}

	// Map each character to a unique vertex ID
	charToID := make(map[byte]int)
	idToChar := make(map[int]byte)

	id := 0
	for char := range charSet {
		charToID[char] = id
		idToChar[id] = char
		id++
	}

	// Create the graph
	graph := NewGraph(n)

	// Add edges based on character order
	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]

		// Find the first differing character
		minLen := min(len(word1), len(word2))

		// Check if word2 is a prefix of word1, which is invalid in a sorted dictionary
		if len(word1) > len(word2) && word1[:len(word2)] == word2 {
			// This is a conflict but doesn't form a cycle in the graph
			continue
		}

		// Find the first differing character
		for j := 0; j < minLen; j++ {
			if word1[j] != word2[j] {
				// Add edge from word1[j] to word2[j]
				graph.AddEdge(charToID[word1[j]], charToID[word2[j]])
				break
			}
		}
	}

	// Find cycles in the graph
	cycles := findCycles(graph, idToChar)
	return cycles
}

// findCycles identifies cycles in the graph
func findCycles(g *Graph, idToChar map[int]byte) [][2]byte {
	visited := make([]bool, g.Vertices)
	recursionStack := make([]bool, g.Vertices)
	cycles := [][2]byte{}

	for i := 0; i < g.Vertices; i++ {
		if !visited[i] {
			findCyclesUtil(g, i, visited, recursionStack, &cycles, idToChar)
		}
	}

	return cycles
}

// findCyclesUtil is a recursive utility function to find cycles
func findCyclesUtil(g *Graph, v int, visited, recursionStack []bool,
	cycles *[][2]byte, idToChar map[int]byte) {

	visited[v] = true
	recursionStack[v] = true

	// Check all adjacent vertices
	for _, u := range g.AdjList[v] {
		if !visited[u] {
			findCyclesUtil(g, u, visited, recursionStack, cycles, idToChar)
		} else if recursionStack[u] {
			// Found a back edge (cycle)
			cycle := [2]byte{idToChar[v], idToChar[u]}
			*cycles = append(*cycles, cycle)
		}
	}

	recursionStack[v] = false
}

// AlienOrderBFS determines the order of characters in an alien language using BFS
// words: array of strings sorted lexicographically in the alien language
// Returns the order of characters as a string, or an empty string if no valid order exists
//
// Time Complexity: O(C) where C is the total length of all words
// Space Complexity: O(1) since the alphabet size is constant (26 characters)
func AlienOrderBFS(words []string) string {
	// Build a set of all characters in the alien language
	charSet := make(map[byte]bool)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			charSet[word[i]] = true
		}
	}

	// Count the number of unique characters
	n := len(charSet)
	if n == 0 {
		return ""
	}

	// Map each character to a unique vertex ID
	charToID := make(map[byte]int)
	idToChar := make(map[int]byte)

	id := 0
	for char := range charSet {
		charToID[char] = id
		idToChar[id] = char
		id++
	}

	// Create the graph
	graph := NewGraph(n)

	// Add edges based on character order
	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]

		// Find the first differing character
		minLen := min(len(word1), len(word2))

		// Check if word2 is a prefix of word1, which is invalid in a sorted dictionary
		if len(word1) > len(word2) && word1[:len(word2)] == word2 {
			return ""
		}

		// Find the first differing character
		for j := 0; j < minLen; j++ {
			if word1[j] != word2[j] {
				// Add edge from word1[j] to word2[j]
				graph.AddEdge(charToID[word1[j]], charToID[word2[j]])
				break
			}
		}
	}

	// Perform topological sort using BFS (Kahn's algorithm)
	order, isDAG := graph.TopologicalSortBFS()
	if !isDAG {
		return "" // No valid ordering exists
	}

	// Convert the order of IDs back to characters
	result := make([]byte, n)
	for i, id := range order {
		result[i] = idToChar[id]
	}

	return string(result)
}

// Example usage:
//
// words := []string{"wrt", "wrf", "er", "ett", "rftt"}
// order := AlienOrder(words)
// // order = "wertf"
//
// words := []string{"z", "x"}
// order := AlienOrder(words)
// // order = "zx"
//
// words := []string{"z", "x", "z"}
// order := AlienOrder(words)
// // order = "" (invalid dictionary)
//
// words := []string{"abc", "ab"}
// order := AlienOrder(words)
// // order = "" (invalid dictionary - "abc" should come after "ab")
