# K-way Merge Pattern

## Introduction

The K-way Merge pattern is used to efficiently merge multiple sorted arrays, lists or portions of data. It's an extension of the classic merge operation in merge sort, but instead of merging just two sorted arrays, it handles merging K sorted arrays simultaneously. This pattern is particularly useful when dealing with large datasets that can't fit entirely in memory and are split across multiple sorted chunks or when merging data from multiple sources.

## How It Works

The K-way Merge pattern works as follows:

1. Take K sorted arrays/lists as input
2. Create a min-heap (priority queue) to track the smallest elements
3. Initially, insert the first element from each array into the heap
4. Repeatedly:
   - Extract the minimum element from the heap
   - Add it to the result array
   - Insert the next element from the same array where the min element came from
   - Continue until all elements are processed

## Time and Space Complexity

- **Time Complexity**: O(N log K) where N is the total number of elements across all arrays and K is the number of arrays. We perform N heap operations, each taking O(log K) time.
- **Space Complexity**: O(K) for storing the heap with one element from each of the K arrays.

## When to Use K-way Merge Pattern

This pattern is useful when:

1. You need to merge multiple sorted arrays or lists
2. Working with external sorting where data doesn't fit in memory
3. Processing streams of sorted data from multiple sources
4. Dealing with distributed merge operations
5. Implementing certain algorithms like merging K sorted linked lists

## Common Problem Patterns

1. **Merge K Sorted Arrays**: Merge K sorted arrays into a single sorted array
2. **Merge K Sorted Lists**: Merge K sorted linked lists into a single sorted list
3. **K Pairs with Smallest Sums**: Find K pairs with the smallest sums from two arrays
4. **Kth Smallest Element in M Sorted Arrays**: Find the Kth smallest element across multiple sorted arrays
5. **Median of a Stream**: Find the median of a continuous stream of numbers
6. **Find K Closest Elements**: Find the K closest elements to a given value from a sorted array
7. **Sort a Nearly Sorted Array**: Sort an array where each element is at most K positions away from its target position

## Implementation in Golang

Here's how the basic K-way merge operation can be implemented in Golang using a heap (priority queue):

```go
import (
    "container/heap"
)

// Element represents an element in one of the arrays
type Element struct {
    Value     int // The actual value
    ArrayIndex int // Index of the array this element belongs to
    ElementIndex int // Index of this element in its array
}

// MinHeap implementation
type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// MergeKSortedArrays merges K sorted arrays into one sorted array
func MergeKSortedArrays(arrays [][]int) []int {
    result := []int{}
    
    // Create a min heap
    h := &MinHeap{}
    heap.Init(h)
    
    // Insert the first element from each array into the heap
    for i, array := range arrays {
        if len(array) > 0 {
            heap.Push(h, Element{
                Value:        array[0],
                ArrayIndex:   i,
                ElementIndex: 0,
            })
        }
    }
    
    // Extract the minimum element and add the next element from the same array
    for h.Len() > 0 {
        // Get the minimum element
        minElement := heap.Pop(h).(Element)
        result = append(result, minElement.Value)
        
        // If there are more elements in the same array, add the next one
        if minElement.ElementIndex+1 < len(arrays[minElement.ArrayIndex]) {
            heap.Push(h, Element{
                Value:        arrays[minElement.ArrayIndex][minElement.ElementIndex+1],
                ArrayIndex:   minElement.ArrayIndex,
                ElementIndex: minElement.ElementIndex+1,
            })
        }
    }
    
    return result
}
```

## Example Problems

1. **Merge K Sorted Arrays**: Merge K sorted arrays into a single sorted array.
2. **Merge K Sorted Lists**: Merge K sorted linked lists into a single sorted list.
3. **Find K Pairs with Smallest Sums**: Given two sorted arrays, find the pairs with the smallest sums.
4. **Kth Smallest Element in Sorted Matrix**: Find the kth smallest element in a sorted matrix.
5. **Smallest Range Covering Elements from K Lists**: Find the smallest range that includes at least one number from each of the K lists.
6. **Find K Closest Numbers**: Find the K closest integers to a given number in a sorted array.
7. **Stream of Medians**: Find the median of a stream of numbers. 