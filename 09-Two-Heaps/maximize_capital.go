package twoheaps

import (
	"container/heap"
)

// Project represents a project with capital required and profit gained
type Project struct {
	Capital int
	Profit  int
}

// MinCapitalHeap implements a min heap for projects, sorted by capital
type MinCapitalHeap []Project

func (h MinCapitalHeap) Len() int           { return len(h) }
func (h MinCapitalHeap) Less(i, j int) bool { return h[i].Capital < h[j].Capital }
func (h MinCapitalHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinCapitalHeap) Push(x interface{}) {
	*h = append(*h, x.(Project))
}

func (h *MinCapitalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxProfitHeap implements a max heap for profits
type MaxProfitHeap []int

func (h MaxProfitHeap) Len() int           { return len(h) }
func (h MaxProfitHeap) Less(i, j int) bool { return h[i] > h[j] } // We want max heap
func (h MaxProfitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxProfitHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxProfitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// FindMaximizedCapital finds the maximum capital after selecting k projects
// Time Complexity: O(N log N + K log N) where N is the number of projects and K is the number of projects we can select
// Space Complexity: O(N) for storing all projects in heaps
func FindMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)

	// Initialize heaps
	minCapitalHeap := &MinCapitalHeap{}
	maxProfitHeap := &MaxProfitHeap{}
	heap.Init(minCapitalHeap)
	heap.Init(maxProfitHeap)

	// Create projects array
	for i := 0; i < n; i++ {
		heap.Push(minCapitalHeap, Project{Capital: capital[i], Profit: profits[i]})
	}

	// Find k projects to maximize capital
	availableCapital := w
	for i := 0; i < k; i++ {
		// Add all projects that can be financed with current capital to the maxHeap
		for minCapitalHeap.Len() > 0 && (*minCapitalHeap)[0].Capital <= availableCapital {
			project := heap.Pop(minCapitalHeap).(Project)
			heap.Push(maxProfitHeap, project.Profit)
		}

		// If we can't afford any project, break
		if maxProfitHeap.Len() == 0 {
			break
		}

		// Select the project with the maximum profit
		availableCapital += heap.Pop(maxProfitHeap).(int)
	}

	return availableCapital
}

// Example usage:
//
// Input: k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
// Output: 4
// Explanation:
// Since your initial capital is 0, you can only start the project with capital requirement 0.
// After finishing it, your capital becomes 1. With capital 1, you can either start the project
// with capital requirement 0 or 1.
// Since you can choose at most 2 projects, you need to finish the project with capital requirement 0
// first and then the one with capital requirement 1.
// Therefore, the final capital is 0 + 1 + 3 = 4.
//
// Input: k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
// Output: 6
// Explanation:
// Since your initial capital is 0, you can only start the project with capital requirement 0.
// After finishing it, your capital becomes 1. With capital 1, you can start the project with
// capital requirement 1.
// After finishing it, your capital becomes 3. With capital 3, you can start the project with
// capital requirement 2.
// After finishing it, your capital becomes 6.
