package kwaymerge

import (
	"container/heap"
)

// Pair represents a pair of elements from two arrays
type Pair struct {
	First  int // Element from first array
	Second int // Element from second array
	Sum    int // Sum of the two elements
}

// PairMinHeap implementation
type PairMinHeap []Pair

func (h PairMinHeap) Len() int           { return len(h) }
func (h PairMinHeap) Less(i, j int) bool { return h[i].Sum < h[j].Sum } // Min heap based on sum
func (h PairMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairMinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// PairIndexMinHeap represents a min heap of pairs of indices
type PairIndexMinHeap []PairIndex

// PairIndex represents a pair of indices from two arrays
type PairIndex struct {
	I, J int // Indices in the first and second arrays
	Sum  int // Sum of the elements at these indices
}

func (h PairIndexMinHeap) Len() int           { return len(h) }
func (h PairIndexMinHeap) Less(i, j int) bool { return h[i].Sum < h[j].Sum }
func (h PairIndexMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairIndexMinHeap) Push(x interface{}) {
	*h = append(*h, x.(PairIndex))
}

func (h *PairIndexMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// KSmallestPairs finds k pairs with the smallest sums from two arrays
// Time Complexity: O(k log k) where k is the number of pairs to find
// Space Complexity: O(k) for storing the heap and the result
func KSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// Handle edge cases
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return [][]int{}
	}

	// Create a min heap
	h := &PairMinHeap{}
	heap.Init(h)

	// Add all possible pairs from nums1[0] to the heap
	for j := 0; j < len(nums2) && j < k; j++ {
		heap.Push(h, Pair{
			First:  nums1[0],
			Second: nums2[j],
			Sum:    nums1[0] + nums2[j],
		})
	}

	// Process the heap to find the k smallest pairs
	result := make([][]int, 0, k)
	for h.Len() > 0 && len(result) < k {
		// Get the pair with the smallest sum
		pair := heap.Pop(h).(Pair)
		result = append(result, []int{pair.First, pair.Second})

		// Find the index of the current element in nums1
		i := 0
		for i < len(nums1) && nums1[i] != pair.First {
			i++
		}

		// Find the index of the current element in nums2
		j := 0
		for j < len(nums2) && nums2[j] != pair.Second {
			j++
		}

		// If there's a next element in nums1, add a new pair
		if i+1 < len(nums1) {
			heap.Push(h, Pair{
				First:  nums1[i+1],
				Second: nums2[j],
				Sum:    nums1[i+1] + nums2[j],
			})
		}
	}

	return result
}

// KSmallestPairsOptimized finds k pairs with the smallest sums using a more efficient approach
// Time Complexity: O(k log k) where k is the number of pairs to find
// Space Complexity: O(k) for storing the heap and the result
func KSmallestPairsOptimized(nums1 []int, nums2 []int, k int) [][]int {
	// Handle edge cases
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return [][]int{}
	}

	// Limit k to the maximum possible number of pairs
	if k > len(nums1)*len(nums2) {
		k = len(nums1) * len(nums2)
	}

	// Create a min heap of pair indices
	h := &PairIndexMinHeap{}
	heap.Init(h)

	// Add the first pair (0,0) to the heap
	heap.Push(h, PairIndex{
		I:   0,
		J:   0,
		Sum: nums1[0] + nums2[0],
	})

	// Keep track of visited pairs to avoid duplicates
	visited := make(map[int]map[int]bool)
	visited[0] = make(map[int]bool)
	visited[0][0] = true

	// Process the heap to find the k smallest pairs
	result := make([][]int, 0, k)
	for h.Len() > 0 && len(result) < k {
		// Get the pair indices with the smallest sum
		pairIdx := heap.Pop(h).(PairIndex)
		i, j := pairIdx.I, pairIdx.J

		// Add the pair to the result
		result = append(result, []int{nums1[i], nums2[j]})

		// Try the next element in nums1
		if i+1 < len(nums1) {
			// Check if (i+1,j) has been visited
			if _, exists := visited[i+1]; !exists {
				visited[i+1] = make(map[int]bool)
			}
			if !visited[i+1][j] {
				heap.Push(h, PairIndex{
					I:   i + 1,
					J:   j,
					Sum: nums1[i+1] + nums2[j],
				})
				visited[i+1][j] = true
			}
		}

		// Try the next element in nums2
		if j+1 < len(nums2) {
			// Check if (i,j+1) has been visited
			if !visited[i][j+1] {
				heap.Push(h, PairIndex{
					I:   i,
					J:   j + 1,
					Sum: nums1[i] + nums2[j+1],
				})
				visited[i][j+1] = true
			}
		}
	}

	return result
}

// Example usage:
//
// Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// Output: [[1,2],[1,4],[1,6]]
// Explanation: The first 3 pairs are returned from the sequence:
//              [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//
// Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
// Output: [[1,1],[1,1]]
// Explanation: The first 2 pairs are returned from the sequence:
//              [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//
// Input: nums1 = [1,2], nums2 = [3], k = 3
// Output: [[1,3],[2,3]]
// Explanation: All possible pairs are returned from the sequence: [1,3],[2,3]
