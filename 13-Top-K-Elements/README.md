# Top K Elements Pattern

## Introduction

The Top K Elements pattern is used to solve problems that involve finding the top K largest or smallest elements in a given set. This pattern leverages data structures like heaps (priority queues) to efficiently track and retrieve these extreme elements without the need to fully sort the entire collection.

## How It Works

The core idea behind this pattern is to maintain a heap of K elements while processing the input data:

1. For finding the top K largest elements:
   - Create a min heap to track the K largest elements
   - Process each element in the collection
   - If the heap size is less than K, add the element to the heap
   - If the element is larger than the smallest element in the heap (the root), remove the root and add the new element
   - After processing all elements, the heap contains the K largest elements

2. For finding the top K smallest elements:
   - Use a max heap with similar logic, but reversed comparisons

## Time and Space Complexity

- **Time Complexity**: O(N log K) where N is the number of elements in the input collection and K is the number of top elements we need to find. This is because for each of the N elements, we perform at most one heap insertion/deletion operation which costs O(log K).
- **Space Complexity**: O(K) for storing the heap of K elements.

## When to Use Top K Elements Pattern

This pattern is useful when:

1. You need to find a small subset of extreme values (largest/smallest) in a large collection
2. You need to maintain a "running" set of top K elements in a stream
3. You need to find elements that are closest, most frequent, or most similar
4. The problem involves partial sorting rather than complete sorting

## Common Problem Patterns

1. **Top K Numbers**: Find the K largest/smallest elements in an array
2. **K Closest Points**: Find the K closest points to the origin (or any reference point)
3. **Top K Frequent Elements**: Find the K most frequent elements in a collection
4. **Sort Characters By Frequency**: Rearrange characters in a string by their frequency
5. **Kth Largest Element**: Find the Kth largest element in an unsorted array
6. **Connect Ropes**: Connect ropes with minimum cost
7. **K Closest Numbers**: Find K numbers that are closest to a given number
8. **Maximum Distinct Elements**: Maximize the count of distinct elements after removing K elements
9. **Sum of Elements**: Find the sum of elements between K1'th and K2'th smallest elements
10. **Rearrange String**: Rearrange a string so that no two adjacent characters are the same

## Implementation in Golang

Golang does not have a built-in heap implementation, but the standard library provides a `container/heap` package that allows us to implement a heap. Here's a basic template for using a min-heap to find the top K largest elements:

```go
import (
    "container/heap"
)

// MinHeap implementation
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Function to find top K largest elements
func findTopKLargest(nums []int, k int) []int {
    h := &MinHeap{}
    heap.Init(h)
    
    for _, num := range nums {
        if h.Len() < k {
            heap.Push(h, num)
        } else if num > (*h)[0] {
            heap.Pop(h)
            heap.Push(h, num)
        }
    }
    
    result := make([]int, h.Len())
    for i := h.Len() - 1; i >= 0; i-- {
        result[i] = heap.Pop(h).(int)
    }
    
    return result
}
```

## Example Problems

1. **Kth Largest Element in an Array**: Find the kth largest element in an unsorted array.
2. **Top K Frequent Elements**: Given an array, return the k most frequent elements.
3. **Sort Characters By Frequency**: Sort characters in a string by decreasing frequency.
4. **K Closest Points to Origin**: Find the k closest points to the origin in a 2D plane.
5. **Connect Ropes**: Connect n ropes with minimum cost.
6. **Top K Frequent Words**: Find the k most frequent words in a corpus.
7. **Kth Smallest Element in a Sorted Matrix**: Find the kth smallest element in a sorted matrix.
8. **Reorganize String**: Rearrange characters so no two adjacent are the same.
9. **K Closest Numbers**: Find k numbers that are closest to a given number.
10. **Maximum Frequency Stack**: Implement a stack supporting the push, pop, and popMax operations. 