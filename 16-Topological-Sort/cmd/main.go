package main

import (
	"fmt"
	"strings"

	"Topological-Sort/topsort"
)

func main() {
	fmt.Println("Topological Sort Pattern Examples")
	fmt.Println("=================================")

	// 1. Basic Topological Sort
	fmt.Println("\n1. Basic Topological Sort Example:")
	graph := topsort.NewGraph(6)
	graph.AddEdge(5, 2)
	graph.AddEdge(5, 0)
	graph.AddEdge(4, 0)
	graph.AddEdge(4, 1)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 1)

	fmt.Println("Graph with edges: 5→2, 5→0, 4→0, 4→1, 2→3, 3→1")

	// BFS-based topological sort (Kahn's algorithm)
	order, isDAG := graph.TopologicalSortBFS()
	if isDAG {
		fmt.Printf("BFS Topological Sort: %v\n", order)
	} else {
		fmt.Println("Graph contains a cycle, no valid topological ordering exists.")
	}

	// DFS-based topological sort
	order, isDAG = graph.TopologicalSortDFS()
	if isDAG {
		fmt.Printf("DFS Topological Sort: %v\n", order)
	} else {
		fmt.Println("Graph contains a cycle, no valid topological ordering exists.")
	}

	// 2. Graph with Cycle
	fmt.Println("\n2. Graph with Cycle Example:")
	cycleGraph := topsort.NewGraph(3)
	cycleGraph.AddEdge(0, 1)
	cycleGraph.AddEdge(1, 2)
	cycleGraph.AddEdge(2, 0)

	fmt.Println("Graph with edges: 0→1, 1→2, 2→0")

	// BFS-based topological sort
	order, isDAG = cycleGraph.TopologicalSortBFS()
	if isDAG {
		fmt.Printf("BFS Topological Sort: %v\n", order)
	} else {
		fmt.Println("Graph contains a cycle, no valid topological ordering exists.")
	}

	// 3. Course Schedule Problem
	fmt.Println("\n3. Course Schedule Problem:")
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}

	fmt.Printf("Number of courses: %d\n", numCourses)
	fmt.Printf("Prerequisites: %v\n", prerequisites)

	canFinish := topsort.CanFinish(numCourses, prerequisites)
	if canFinish {
		fmt.Println("Can finish all courses: Yes")

		// Get the course order
		courseOrder, possible := topsort.FindOrder(numCourses, prerequisites)
		if !possible {
			fmt.Println("No valid course order exists.")
		} else {
			fmt.Printf("Course Order: %v\n", courseOrder)
		}
	} else {
		fmt.Println("Can finish all courses: No")
	}

	// 4. Alien Dictionary Problem
	fmt.Println("\n4. Alien Dictionary Problem:")
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}

	fmt.Printf("Dictionary words: %s\n", strings.Join(words, ", "))

	alienOrder := topsort.AlienOrder(words)
	if alienOrder == "" {
		fmt.Println("No valid character ordering exists.")
	} else {
		fmt.Printf("Character Order: %s\n", alienOrder)
	}

	// Try with an invalid dictionary
	invalidWords := []string{"z", "x", "z"}
	fmt.Printf("\nInvalid dictionary words: %s\n", strings.Join(invalidWords, ", "))

	alienOrder = topsort.AlienOrder(invalidWords)
	if alienOrder == "" {
		fmt.Println("No valid character ordering exists.")
	} else {
		fmt.Printf("Character Order: %s\n", alienOrder)
	}

	// 5. Minimum Height Trees Problem
	fmt.Println("\n5. Minimum Height Trees Problem:")
	n := 6
	edges := [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}}

	fmt.Printf("Number of nodes: %d\n", n)
	fmt.Printf("Edges: %v\n", edges)

	roots := topsort.FindMinHeightTrees(n, edges)
	fmt.Printf("MHT Roots: %v\n", roots)

	heights := topsort.FindHeights(n, edges)
	fmt.Printf("Heights of trees rooted at each node: %v\n", heights)

	diameter := topsort.FindTreeDiameter(n, edges)
	fmt.Printf("Tree Diameter: %d\n", diameter)

	// 6. Build Order Problem
	fmt.Println("\n6. Build Order Problem:")
	projects := []string{"a", "b", "c", "d", "e", "f"}
	dependencies := [][]string{{"a", "d"}, {"f", "b"}, {"b", "d"}, {"f", "a"}, {"d", "c"}}

	fmt.Printf("Projects: %s\n", strings.Join(projects, ", "))
	fmt.Printf("Dependencies: ")
	for i, dep := range dependencies {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%s→%s", dep[0], dep[1])
	}
	fmt.Println()

	buildOrder, err := topsort.FindBuildOrder(projects, dependencies)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Build Order: %v\n", buildOrder)
	}

	// 7. Task Scheduler Problem
	fmt.Println("\n7. Task Scheduler Problem:")
	numTasks := 4
	durations := []int{2, 3, 1, 4}
	taskDeps := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}

	fmt.Printf("Number of tasks: %d\n", numTasks)
	fmt.Printf("Task durations: %v\n", durations)
	fmt.Printf("Task dependencies: %v\n", taskDeps)

	completionTime, _ := topsort.MinCompletionTime(numTasks, durations, taskDeps)
	fmt.Printf("Minimum completion time: %d\n", completionTime)

	startTimes, completionTime, _ := topsort.ScheduleTasks(numTasks, durations, taskDeps)
	fmt.Printf("Start times of each task: %v\n", startTimes)
	fmt.Printf("Total completion time: %d\n", completionTime)

	criticalPath, pathLength := topsort.CalculateCriticalPath(numTasks, durations, taskDeps)
	fmt.Printf("Critical path tasks: %v\n", criticalPath)
	fmt.Printf("Critical path length: %d\n", pathLength)
}
