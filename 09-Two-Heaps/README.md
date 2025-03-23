# Two Heaps Pattern

## Introduction

The Two Heaps pattern is a technique that uses two heaps—a min heap and a max heap—to efficiently solve problems that require tracking medians or finding the middle elements of a dataset. This pattern is particularly useful when you need to find the smallest element from one part of the data and the largest element from the other part.

In this pattern, we divide elements into two parts and use:
- A max heap to store the smaller half of the elements
- A min heap to store the larger half of the elements

This arrangement allows us to efficiently find elements around the middle of a dataset, especially when data is continuously changing.

## How It Works

1. **Partitioning the data**: Elements are divided into two halves, with smaller elements in the max heap and larger elements in the min heap.
2. **Balancing**: We maintain a balance such that the size difference between the two heaps is not more than 1.
3. **Accessing medians/middle elements**: 
   - If both heaps have equal size, the median is the average of the top elements from both heaps.
   - If they have unequal size, the median is the top element of the heap with more elements.

## Time and Space Complexity

For most operations in the two heaps pattern:
- **Time Complexity**: O(log n) for insert operations and O(1) for peek operations
- **Space Complexity**: O(n) for storing all elements across both heaps

## When to Use Two Heaps

Use the Two Heaps pattern when:
- You need to track the median of a dynamic data stream
- You need to find the kth largest or smallest element
- You need to efficiently find median ranges or central tendencies
- You need to balance partitioning of elements in real-time

## Common Problem Patterns

1. **Find the Median of a Number Stream**
   - Maintain two heaps to keep track of the median as new numbers are added

2. **Sliding Window Median**
   - Find the median for each window of size k in a given array

3. **IPO (Initial Public Offering)**
   - Maximize capital by selecting projects with maximum profit while considering available capital

4. **Next Interval**
   - Find the next interval for each interval in an array

## Implementation in Golang

```go
// Basic implementation of the Two Heaps pattern for finding the median of a stream
type MedianFinder struct {
    maxHeap *MaxHeap // Stores the smaller half of numbers
    minHeap *MinHeap // Stores the larger half of numbers
}

func Constructor() MedianFinder {
    return MedianFinder{
        maxHeap: &MaxHeap{},
        minHeap: &MinHeap{},
    }
}

func (mf *MedianFinder) AddNum(num int) {
    // Step 1: Add the number to the appropriate heap
    if mf.maxHeap.Len() == 0 || num <= mf.maxHeap.Peek() {
        heap.Push(mf.maxHeap, num)
    } else {
        heap.Push(mf.minHeap, num)
    }
    
    // Step 2: Balance the heaps
    // If maxHeap has more than one extra element than minHeap
    if mf.maxHeap.Len() > mf.minHeap.Len()+1 {
        heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
    }
    // If minHeap has more elements than maxHeap
    if mf.minHeap.Len() > mf.maxHeap.Len() {
        heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
    }
}

func (mf *MedianFinder) FindMedian() float64 {
    // If both heaps have the same size, return the average of their tops
    if mf.maxHeap.Len() == mf.minHeap.Len() {
        return float64(mf.maxHeap.Peek()+mf.minHeap.Peek()) / 2.0
    }
    // Otherwise, the median is the top of maxHeap (which has one more element)
    return float64(mf.maxHeap.Peek())
}
```

## Example Problems

1. **Find the Median of a Number Stream**
2. **Sliding Window Median**
3. **Maximize Capital (IPO)**
4. **Find Right Interval**

Each of these problems has a dedicated solution file in this directory. 