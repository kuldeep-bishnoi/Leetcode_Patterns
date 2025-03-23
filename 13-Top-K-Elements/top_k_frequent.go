package topk

import (
	"container/heap"
)

// FrequencyItem represents an element with its frequency
type FrequencyItem struct {
	Value     int
	Frequency int
}

// MinFrequencyHeap is a min heap based on frequency
type MinFrequencyHeap []FrequencyItem

func (h MinFrequencyHeap) Len() int           { return len(h) }
func (h MinFrequencyHeap) Less(i, j int) bool { return h[i].Frequency < h[j].Frequency } // Min heap
func (h MinFrequencyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinFrequencyHeap) Push(x interface{}) {
	*h = append(*h, x.(FrequencyItem))
}

func (h *MinFrequencyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// TopKFrequent finds the k most frequent elements in the array
// Time Complexity: O(N log K) where N is the array size and K is the parameter
// Space Complexity: O(N) for the frequency map and O(K) for the heap
func TopKFrequent(nums []int, k int) []int {
	// Count frequency of each number
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	// Create a min heap based on frequency
	h := &MinFrequencyHeap{}
	heap.Init(h)

	// Process each unique number
	for num, frequency := range frequencyMap {
		// If heap size is less than k, add the element
		if h.Len() < k {
			heap.Push(h, FrequencyItem{Value: num, Frequency: frequency})
		} else if frequency > (*h)[0].Frequency {
			// If current element is more frequent than the least frequent in heap,
			// remove the least frequent and add the current element
			heap.Pop(h)
			heap.Push(h, FrequencyItem{Value: num, Frequency: frequency})
		}
	}

	// Extract the k most frequent elements from the heap
	result := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		item := heap.Pop(h).(FrequencyItem)
		result[i] = item.Value
	}

	return result
}

// TopKFrequentBucketSort uses bucket sort to find the k most frequent elements
// Time Complexity: O(N) where N is the array size
// Space Complexity: O(N) for the frequency map and buckets
func TopKFrequentBucketSort(nums []int, k int) []int {
	// Count frequency of each number
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	// Create buckets where index represents frequency
	// At most, an element can appear n times (all elements are the same)
	buckets := make([][]int, len(nums)+1)
	for num, frequency := range frequencyMap {
		if buckets[frequency] == nil {
			buckets[frequency] = []int{}
		}
		buckets[frequency] = append(buckets[frequency], num)
	}

	// Collect the top k frequent elements
	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		if buckets[i] != nil {
			// Add all elements with the current frequency
			for _, num := range buckets[i] {
				result = append(result, num)
				if len(result) == k {
					break
				}
			}
		}
	}

	return result
}

// Example usage:
//
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Explanation: The elements 1 and 2 appear the most. 1 appears 3 times and 2 appears 2 times.
//
// Input: nums = [1], k = 1
// Output: [1]
//
// Input: nums = [3,0,1,0], k = 1
// Output: [0]
// Explanation: 0 appears twice while 3 and 1 appear only once.
