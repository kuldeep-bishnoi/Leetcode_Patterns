package topk

import (
	"container/heap"
)

// TaskFrequency represents a task with its frequency
type TaskFrequency struct {
	Task      byte
	Frequency int
}

// MaxTaskHeap is a max heap based on task frequency
type MaxTaskHeap []TaskFrequency

func (h MaxTaskHeap) Len() int           { return len(h) }
func (h MaxTaskHeap) Less(i, j int) bool { return h[i].Frequency > h[j].Frequency } // Max heap
func (h MaxTaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxTaskHeap) Push(x interface{}) {
	*h = append(*h, x.(TaskFrequency))
}

func (h *MaxTaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// LeastInterval calculates the least number of intervals needed to execute all tasks
// with a cooling period of n between the same tasks
// Time Complexity: O(N) where N is the number of tasks
// Space Complexity: O(1) since there are at most 26 different tasks (A-Z)
func LeastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks) // No cooling period needed
	}

	// Count frequency of each task
	frequencyMap := make(map[byte]int)
	for _, task := range tasks {
		frequencyMap[task]++
	}

	// Find the maximum frequency
	maxFrequency := 0
	for _, frequency := range frequencyMap {
		if frequency > maxFrequency {
			maxFrequency = frequency
		}
	}

	// Find the number of tasks with maximum frequency
	maxCount := 0
	for _, frequency := range frequencyMap {
		if frequency == maxFrequency {
			maxCount++
		}
	}

	// Calculate the total intervals needed
	// Formula: (maxFrequency - 1) * (n + 1) + maxCount
	// Explanation:
	// - We need (maxFrequency - 1) chunks, each of size (n + 1)
	// - Plus maxCount for the last chunk
	intervals := (maxFrequency-1)*(n+1) + maxCount

	// Return the maximum of intervals and the number of tasks
	if intervals < len(tasks) {
		return len(tasks)
	}
	return intervals
}

// LeastIntervalHeap uses a max heap to calculate the least number of intervals
// This approach is more intuitive and can be extended for more complex scheduling
// Time Complexity: O(N log 26) which is effectively O(N)
// Space Complexity: O(1) since there are at most 26 different tasks
func LeastIntervalHeap(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks) // No cooling period needed
	}

	// Count frequency of each task
	frequencyMap := make(map[byte]int)
	for _, task := range tasks {
		frequencyMap[task]++
	}

	// Create a max heap based on frequency
	h := &MaxTaskHeap{}
	heap.Init(h)

	// Add all tasks with their frequencies to the heap
	for task, frequency := range frequencyMap {
		heap.Push(h, TaskFrequency{Task: task, Frequency: frequency})
	}

	totalTime := 0

	// Process until all tasks are scheduled
	for h.Len() > 0 {
		// Store tasks that need to wait for cooling period
		waitList := make([]TaskFrequency, 0)

		// Try to execute n+1 different tasks
		count := n + 1

		// Process up to n+1 tasks or until heap is empty
		for count > 0 && h.Len() > 0 {
			// Get the most frequent task
			current := heap.Pop(h).(TaskFrequency)

			// Decrease its frequency
			current.Frequency--

			// If it still has remaining executions, add to wait list
			if current.Frequency > 0 {
				waitList = append(waitList, current)
			}

			totalTime++
			count--
		}

		// Add tasks from wait list back to heap
		for _, task := range waitList {
			heap.Push(h, task)
		}

		// If heap still has tasks, we need to wait for cooling period
		if h.Len() > 0 {
			totalTime += count // Add idle time
		}
	}

	return totalTime
}

// Example usage:
//
// Input: tasks = ["A","A","A","B","B","B"], n = 2
// Output: 8
// Explanation: A -> B -> idle -> A -> B -> idle -> A -> B
// There is at least 2 units of time between any two same tasks.
//
// Input: tasks = ["A","A","A","B","B","B"], n = 0
// Output: 6
// Explanation: No cooling time required, so we can execute them in order: A,A,A,B,B,B
//
// Input: tasks = ["A","A","A","A","A","A","B","C","D","E","F","G"], n = 2
// Output: 16
// Explanation: One possible solution is
// A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> idle -> idle -> A -> idle -> idle -> A
