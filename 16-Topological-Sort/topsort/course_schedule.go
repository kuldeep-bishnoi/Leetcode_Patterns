package topsort

// CanFinish determines if it's possible to finish all courses
// numCourses: total number of courses
// prerequisites: array of course prerequisite pairs [course, prerequisite]
// 
// Time Complexity: O(V + E) where V is the number of courses and E is the number of prerequisites
// Space Complexity: O(V + E) for the graph representation and visited arrays
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// Build the graph
	graph := NewGraph(numCourses)
	for _, prereq := range prerequisites {
		course, prerequisite := prereq[0], prereq[1]
		// Add edge from prerequisite to course
		graph.AddEdge(prerequisite, course)
	}
	
	// Check if the graph is a DAG using topological sort
	_, isDAG := graph.TopologicalSortBFS()
	return isDAG
}

// FindOrder returns the order in which courses should be taken to finish all courses
// numCourses: total number of courses
// prerequisites: array of course prerequisite pairs [course, prerequisite]
// 
// Time Complexity: O(V + E) where V is the number of courses and E is the number of prerequisites
// Space Complexity: O(V + E) for the graph representation and visited arrays
func FindOrder(numCourses int, prerequisites [][]int) ([]int, bool) {
	// Build the graph
	graph := NewGraph(numCourses)
	for _, prereq := range prerequisites {
		course, prerequisite := prereq[0], prereq[1]
		// Add edge from prerequisite to course
		graph.AddEdge(prerequisite, course)
	}
	
	// Get the topological order using DFS
	return graph.TopologicalSortDFS()
}

// FindOrderBFS returns the order in which courses should be taken to finish all courses using BFS
// numCourses: total number of courses
// prerequisites: array of course prerequisite pairs [course, prerequisite]
// 
// Time Complexity: O(V + E) where V is the number of courses and E is the number of prerequisites
// Space Complexity: O(V + E) for the graph representation and visited arrays
func FindOrderBFS(numCourses int, prerequisites [][]int) ([]int, bool) {
	// Build the graph
	graph := NewGraph(numCourses)
	for _, prereq := range prerequisites {
		course, prerequisite := prereq[0], prereq[1]
		// Add edge from prerequisite to course
		graph.AddEdge(prerequisite, course)
	}
	
	// Get the topological order using BFS (Kahn's algorithm)
	return graph.TopologicalSortBFS()
}

// ScheduleAllCourses returns a schedule that minimizes the total time to complete all courses
// numCourses: total number of courses
// prerequisites: array of course prerequisite pairs [course, prerequisite]
// duration: array of course durations where duration[i] is the time needed to complete course i
// 
// Time Complexity: O(V + E) where V is the number of courses and E is the number of prerequisites
// Space Complexity: O(V + E) for the graph representation, earliest completion times, and result
func ScheduleAllCourses(numCourses int, prerequisites [][]int, duration []int) ([]int, int) {
	// Build the graph
	graph := NewGraph(numCourses)
	for _, prereq := range prerequisites {
		course, prerequisite := prereq[0], prereq[1]
		// Add edge from prerequisite to course
		graph.AddEdge(prerequisite, course)
	}
	
	// Check if the graph is a DAG
	order, isDAG := graph.TopologicalSortBFS()
	if !isDAG {
		return nil, -1 // Cannot schedule if there's a cycle
	}
	
	// Calculate earliest completion time for each course
	completionTime := make([]int, numCourses)
	
	for _, course := range order {
		// Initialize with the current course duration
		completionTime[course] = duration[course]
		
		// Check dependencies (in the original graph)
		for _, nextCourse := range graph.AdjList[course] {
			// Update the earliest completion time of the next course
			completionTime[nextCourse] = max(
				completionTime[nextCourse],
				completionTime[course] + duration[nextCourse],
			)
		}
	}
	
	// Find the maximum completion time (total schedule time)
	totalTime := 0
	for _, time := range completionTime {
		totalTime = max(totalTime, time)
	}
	
	return order, totalTime
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example usage:
//
// numCourses := 4
// prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
//
// canFinish := CanFinish(numCourses, prerequisites)
// // canFinish = true
//
// order, possible := FindOrder(numCourses, prerequisites)
// // order = [0 1 2 3], possible = true
//
// duration := []int{1, 2, 3, 4}
// schedule, totalTime := ScheduleAllCourses(numCourses, prerequisites, duration)
// // schedule = [0 1 2 3], totalTime = 7 