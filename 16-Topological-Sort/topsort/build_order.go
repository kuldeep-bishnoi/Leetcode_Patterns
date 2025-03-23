package topsort

import (
	"errors"
)

// Project represents a project in the build system
type Project struct {
	Name         string
	Dependencies []*Project
	Children     []*Project // Projects that depend on this project
	InDegree     int        // Number of dependencies
	Visited      bool       // For DFS-based approach
	Temporary    bool       // For cycle detection in DFS
}

// BuildManager manages the build order of projects
type BuildManager struct {
	Projects map[string]*Project
}

// NewBuildManager creates a new build manager
func NewBuildManager() *BuildManager {
	return &BuildManager{
		Projects: make(map[string]*Project),
	}
}

// AddProject adds a project to the build manager
func (bm *BuildManager) AddProject(name string) {
	if _, exists := bm.Projects[name]; !exists {
		bm.Projects[name] = &Project{
			Name:         name,
			Dependencies: []*Project{},
			Children:     []*Project{},
		}
	}
}

// AddDependency adds a dependency between projects
func (bm *BuildManager) AddDependency(project, dependency string) error {
	// Ensure both projects exist
	if _, exists := bm.Projects[project]; !exists {
		bm.AddProject(project)
	}
	if _, exists := bm.Projects[dependency]; !exists {
		bm.AddProject(dependency)
	}

	proj := bm.Projects[project]
	dep := bm.Projects[dependency]

	// Add dependency relationship
	proj.Dependencies = append(proj.Dependencies, dep)
	dep.Children = append(dep.Children, proj)
	proj.InDegree++

	return nil
}

// FindBuildOrderDFS finds a valid build order using DFS
func (bm *BuildManager) FindBuildOrderDFS() ([]string, error) {
	// Reset all project states
	for _, p := range bm.Projects {
		p.Visited = false
		p.Temporary = false
	}

	buildOrder := []string{}

	// Perform DFS on each unvisited project
	for _, p := range bm.Projects {
		if !p.Visited {
			if err := bm.dfs(p, &buildOrder); err != nil {
				return nil, err
			}
		}
	}

	// Reverse the order (DFS gives a reverse topological sort)
	for i, j := 0, len(buildOrder)-1; i < j; i, j = i+1, j-1 {
		buildOrder[i], buildOrder[j] = buildOrder[j], buildOrder[i]
	}

	return buildOrder, nil
}

// dfs performs a depth-first search to find topological order
func (bm *BuildManager) dfs(project *Project, buildOrder *[]string) error {
	// Check for cycle
	if project.Temporary {
		return errors.New("cycle detected in project dependencies")
	}

	if project.Visited {
		return nil
	}

	// Mark as temporarily visited for cycle detection
	project.Temporary = true

	// Visit all dependencies first
	for _, dep := range project.Dependencies {
		if err := bm.dfs(dep, buildOrder); err != nil {
			return err
		}
	}

	// Mark as permanently visited
	project.Visited = true
	project.Temporary = false

	// Add to build order
	*buildOrder = append(*buildOrder, project.Name)

	return nil
}

// FindBuildOrderBFS finds a valid build order using BFS (Kahn's algorithm)
func (bm *BuildManager) FindBuildOrderBFS() ([]string, error) {
	// Count of projects
	n := len(bm.Projects)
	if n == 0 {
		return []string{}, nil
	}

	// Create a copy of in-degrees to avoid modifying the original
	inDegree := make(map[string]int)
	for name, proj := range bm.Projects {
		inDegree[name] = proj.InDegree
	}

	// Queue for BFS - start with all projects that have no dependencies
	queue := []string{}
	for name, count := range inDegree {
		if count == 0 {
			queue = append(queue, name)
		}
	}

	buildOrder := []string{}

	// Process the queue
	for len(queue) > 0 {
		// Dequeue a project
		current := queue[0]
		queue = queue[1:]

		// Add to build order
		buildOrder = append(buildOrder, current)

		// Update in-degrees of dependent projects
		for _, child := range bm.Projects[current].Children {
			inDegree[child.Name]--
			// If in-degree becomes 0, add to queue
			if inDegree[child.Name] == 0 {
				queue = append(queue, child.Name)
			}
		}
	}

	// Check if all projects were included
	if len(buildOrder) != n {
		return nil, errors.New("cycle detected in project dependencies")
	}

	return buildOrder, nil
}

// FindBuildOrder provides a valid order to build all projects
// projects: list of project names
// dependencies: list of [project, dependency] pairs where project depends on dependency
// Returns the build order as a list of project names
//
// Time Complexity: O(P + D) where P is the number of projects and D is the number of dependencies
// Space Complexity: O(P + D)
func FindBuildOrder(projects []string, dependencies [][]string) ([]string, error) {
	// Create build manager
	bm := NewBuildManager()

	// Add all projects
	for _, p := range projects {
		bm.AddProject(p)
	}

	// Add all dependencies
	for _, dep := range dependencies {
		if len(dep) != 2 {
			return nil, errors.New("invalid dependency format")
		}
		// Note: dep[0] depends on dep[1]
		if err := bm.AddDependency(dep[0], dep[1]); err != nil {
			return nil, err
		}
	}

	// Find build order using BFS (Kahn's algorithm)
	return bm.FindBuildOrderBFS()
}

// Example usage:
//
// projects := []string{"a", "b", "c", "d", "e", "f"}
// dependencies := [][]string{{"a", "d"}, {"f", "b"}, {"b", "d"}, {"f", "a"}, {"d", "c"}}
// order, err := FindBuildOrder(projects, dependencies)
// // order could be one of several valid orderings, e.g., ["e", "c", "d", "b", "a", "f"]
//
// bm := NewBuildManager()
// bm.AddProject("a")
// bm.AddProject("b")
// bm.AddProject("c")
// bm.AddDependency("a", "b") // a depends on b
// bm.AddDependency("b", "c") // b depends on c
// order, err := bm.FindBuildOrderDFS()
// // order = ["c", "b", "a"] - must build c first, then b, then a
