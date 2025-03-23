package topsort

// Task represents a task in the scheduling system
type Task struct {
	ID           int
	Duration     int
	Dependencies []int
	EarliestTime int // Earliest start time
	LatestTime   int // Latest start time without delaying completion
}

// MinCompletionTime calculates the minimum time needed to complete all tasks
// n: number of tasks (labeled from 0 to n-1)
// durations: duration of each task
// dependencies: pairs of [task, dependency] where task depends on dependency
// Returns the minimum completion time for all tasks
//
// Time Complexity: O(V + E) where V is the number of tasks and E is the number of dependencies
// Space Complexity: O(V + E)
func MinCompletionTime(n int, durations []int, dependencies [][]int) (int, error) {
	if n == 0 {
		return 0, nil
	}

	// Create the graph
	graph := NewGraph(n)

	// Add all dependencies as edges
	for _, dep := range dependencies {
		// dep[0] depends on dep[1]
		graph.AddEdge(dep[1], dep[0]) // Add edge from dependency to dependent task
	}

	// Perform topological sort
	order, isDAG := graph.TopologicalSortBFS()
	if !isDAG {
		return -1, nil // Cycle detected, impossible to complete
	}

	// Calculate earliest completion time for each task
	earliestTime := make([]int, n)
	maxCompletionTime := 0

	for _, taskID := range order {
		for _, nextTask := range graph.AdjList[taskID] {
			// Update earliest start time for the dependent task
			earliestTime[nextTask] = max(earliestTime[nextTask],
				earliestTime[taskID]+durations[taskID])
		}

		// Update the overall completion time
		maxCompletionTime = max(maxCompletionTime,
			earliestTime[taskID]+durations[taskID])
	}

	return maxCompletionTime, nil
}

// ScheduleTasks creates a complete schedule for all tasks
// n: number of tasks (labeled from 0 to n-1)
// durations: duration of each task
// dependencies: pairs of [task, dependency] where task depends on dependency
// Returns the start time for each task and the overall completion time
//
// Time Complexity: O(V + E) where V is the number of tasks and E is the number of dependencies
// Space Complexity: O(V + E)
func ScheduleTasks(n int, durations []int, dependencies [][]int) ([]int, int, error) {
	if n == 0 {
		return []int{}, 0, nil
	}

	// Build the graph - taskID -> dependent tasks
	graph := make([][]int, n)
	// Build the reverse graph - taskID -> prerequisite tasks
	revGraph := make([][]int, n)
	// Initialize the graphs
	for i := 0; i < n; i++ {
		graph[i] = []int{}
		revGraph[i] = []int{}
	}

	// Track in-degree (number of dependencies) for each task
	inDegree := make([]int, n)

	// Add dependencies
	for _, dep := range dependencies {
		if len(dep) != 2 {
			return nil, 0, nil
		}
		task, prereq := dep[0], dep[1] // task depends on prereq
		if task < 0 || task >= n || prereq < 0 || prereq >= n {
			return nil, 0, nil
		}

		graph[prereq] = append(graph[prereq], task)     // prereq -> task
		revGraph[task] = append(revGraph[task], prereq) // task -> prereq
		inDegree[task]++
	}

	// Calculate earliest start times using Kahn's algorithm
	startTimes := make([]int, n)
	queue := []int{}

	// Add all tasks with no dependencies to the queue
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
			startTimes[i] = 0 // Can start immediately
		}
	}

	// Process tasks in topological order
	for len(queue) > 0 {
		// Dequeue a task
		current := queue[0]
		queue = queue[1:]

		// Update earliest start time for dependent tasks
		for _, next := range graph[current] {
			// Next task can't start until current task is complete
			startTimes[next] = max(startTimes[next],
				startTimes[current]+durations[current])

			// Decrease in-degree
			inDegree[next]--

			// If all dependencies satisfied, add to queue
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// Check if all tasks were scheduled
	for i := 0; i < n; i++ {
		if inDegree[i] > 0 {
			// Cycle detected
			return nil, -1, nil
		}
	}

	// Calculate completion time
	completionTime := 0
	for i := 0; i < n; i++ {
		completionTime = max(completionTime, startTimes[i]+durations[i])
	}

	return startTimes, completionTime, nil
}

// CalculateCriticalPath finds the critical path in the task dependency graph
// n: number of tasks (labeled from 0 to n-1)
// durations: duration of each task
// dependencies: pairs of [task, dependency] where task depends on dependency
// Returns the sequence of tasks in the critical path and its length
//
// Time Complexity: O(V + E) where V is the number of tasks and E is the number of dependencies
// Space Complexity: O(V + E)
func CalculateCriticalPath(n int, durations []int, dependencies [][]int) ([]int, int) {
	if n == 0 {
		return []int{}, 0
	}

	// Build the graph
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}

	// Track in-degree for each task
	inDegree := make([]int, n)

	// Add all dependencies
	for _, dep := range dependencies {
		task, prereq := dep[0], dep[1] // task depends on prereq
		graph[prereq] = append(graph[prereq], task)
		inDegree[task]++
	}

	// Calculate earliest start times
	earliestStart := make([]int, n)

	// Queue for topological sort
	queue := []int{}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Process tasks in topological order to find earliest start times
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, next := range graph[current] {
			earliestStart[next] = max(earliestStart[next],
				earliestStart[current]+durations[current])

			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// Find the completion time
	completionTime := 0
	for i := 0; i < n; i++ {
		finishTime := earliestStart[i] + durations[i]
		if finishTime > completionTime {
			completionTime = finishTime
		}
	}

	// Calculate latest start times (backward pass)
	latestStart := make([]int, n)
	for i := 0; i < n; i++ {
		latestStart[i] = completionTime
	}

	// Reset in-degree for backward pass
	outDegree := make([]int, n)
	for i := 0; i < n; i++ {
		outDegree[i] = len(graph[i])
	}

	// Queue for backward pass
	queue = []int{}
	for i := 0; i < n; i++ {
		if outDegree[i] == 0 {
			queue = append(queue, i)
			latestStart[i] = completionTime - durations[i]
		}
	}

	// Process tasks in reverse topological order
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for i := 0; i < n; i++ {
			for _, next := range graph[i] {
				if next == current {
					latestStart[i] = min(latestStart[i],
						latestStart[current]-durations[i])

					outDegree[i]--
					if outDegree[i] == 0 {
						queue = append(queue, i)
					}
					break
				}
			}
		}
	}

	// Find tasks in the critical path (where earliest = latest start time)
	criticalPath := []int{}
	for i := 0; i < n; i++ {
		if earliestStart[i] == latestStart[i] {
			criticalPath = append(criticalPath, i)
		}
	}

	return criticalPath, completionTime
}

// Example usage:
//
// n := 4
// durations := []int{2, 3, 1, 4}
// dependencies := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
// completionTime, _ := MinCompletionTime(n, durations, dependencies)
// // completionTime = 9
//
// startTimes, completionTime, _ := ScheduleTasks(n, durations, dependencies)
// // startTimes = [0, 2, 2, 5], completionTime = 9
//
// criticalPath, pathLength := CalculateCriticalPath(n, durations, dependencies)
// // criticalPath = [0, 1, 3], pathLength = 9
