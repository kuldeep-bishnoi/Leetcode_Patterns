package topk

import (
	"container/heap"
)

// CharFrequencyPair represents a character with its frequency
// Re-used from frequency_sort.go but included here for completeness
type CharFrequencyPair struct {
	Char      byte
	Frequency int
}

// MaxFrequencyHeap is a max heap based on character frequency
type MaxFrequencyHeap []CharFrequencyPair

func (h MaxFrequencyHeap) Len() int           { return len(h) }
func (h MaxFrequencyHeap) Less(i, j int) bool { return h[i].Frequency > h[j].Frequency } // Max heap
func (h MaxFrequencyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxFrequencyHeap) Push(x interface{}) {
	*h = append(*h, x.(CharFrequencyPair))
}

func (h *MaxFrequencyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// RearrangeString rearranges a string so that no same characters are adjacent
// Time Complexity: O(N log K) where N is string length and K is unique characters
// Space Complexity: O(K) for the frequency map and heap
func RearrangeString(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Count frequency of each character
	frequencyMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		frequencyMap[s[i]]++
	}

	// Create a max heap based on frequency
	h := &MaxFrequencyHeap{}
	heap.Init(h)

	// Add all characters with their frequencies to the heap
	for char, frequency := range frequencyMap {
		heap.Push(h, CharFrequencyPair{Char: char, Frequency: frequency})
	}

	// Build the result by alternating characters
	result := make([]byte, 0, len(s))
	var prev CharFrequencyPair

	for h.Len() > 0 {
		// Get the most frequent character
		current := heap.Pop(h).(CharFrequencyPair)

		// Add the current character to result
		result = append(result, current.Char)

		// Decrease the frequency of the current character
		current.Frequency--

		// If the previous character still has frequency left, add it back to the heap
		if prev.Frequency > 0 {
			heap.Push(h, prev)
		}

		// Update the previous character
		prev = current
	}

	// If we couldn't use all characters, it's impossible to rearrange
	if len(result) != len(s) {
		return ""
	}

	return string(result)
}

// RearrangeStringKDistance rearranges a string so that same characters are at least k distance apart
// Time Complexity: O(N log K) where N is string length and K is unique characters
// Space Complexity: O(K) for the frequency map and heap
func RearrangeStringKDistance(s string, k int) string {
	if k <= 1 {
		return s // No rearrangement needed
	}

	// Count frequency of each character
	frequencyMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		frequencyMap[s[i]]++
	}

	// Create a max heap based on frequency
	h := &MaxFrequencyHeap{}
	heap.Init(h)

	// Add all characters with their frequencies to the heap
	for char, frequency := range frequencyMap {
		heap.Push(h, CharFrequencyPair{Char: char, Frequency: frequency})
	}

	// Build the result by alternating characters with k distance
	result := make([]byte, 0, len(s))

	// Queue to track characters that need to wait before being reused
	waitQueue := make([]CharFrequencyPair, 0)

	for h.Len() > 0 {
		// Get the most frequent character
		current := heap.Pop(h).(CharFrequencyPair)

		// Add the current character to result
		result = append(result, current.Char)

		// Decrease the frequency of the current character
		current.Frequency--

		// Add current character to wait queue if it still has frequency
		if current.Frequency > 0 {
			waitQueue = append(waitQueue, current)
		}

		// If wait queue has reached size k, add the front character back to heap
		if len(waitQueue) == k {
			heap.Push(h, waitQueue[0])
			waitQueue = waitQueue[1:] // Dequeue the front
		}
	}

	// If we couldn't use all characters, it's impossible to rearrange
	if len(result) != len(s) {
		return ""
	}

	return string(result)
}

// Example usage:
//
// Input: s = "aab"
// Output: "aba"
// Explanation: We can reorder s to "aba" such that no same characters are adjacent.
//
// Input: s = "aaab"
// Output: ""
// Explanation: We cannot rearrange s to avoid having two 'a's adjacent to each other.
//
// For K-distance rearrangement:
// Input: s = "mmpp", k = 2
// Output: "mpmp"
// Explanation: We can rearrange s to have each 'm' and 'p' at distance 2.
