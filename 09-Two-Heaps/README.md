# Two Heaps Pattern

## What are Two Heaps?

Imagine you have two piles of cards - one for the biggest numbers and one for the smallest numbers. The Two Heaps pattern is like having two special containers where you can quickly find the biggest and smallest numbers. It's like having two baskets: one for your biggest toys and one for your smallest toys!

## Real-Life Examples

1. **Class Grades**: When you want to track the highest and lowest scores.
   - One pile for the highest scores
   - One pile for the lowest scores
   - Always know the top and bottom scores quickly

2. **Temperature Records**: When tracking daily temperatures.
   - One pile for the hottest days
   - One pile for the coldest days
   - Always know the extreme temperatures

3. **Shopping Prices**: When comparing product prices.
   - One pile for the most expensive items
   - One pile for the least expensive items
   - Always know the price range quickly

## When Do We Use Two Heaps?

Use this technique when:
- You need to find the median of a stream of numbers
- You need to track the highest and lowest values
- You need to find the middle value in a stream
- You need to balance two groups of numbers
- You need to find sliding window medians

## How Does It Work?

1. **Step 1**: Create two heaps:
   - Max Heap: stores smaller half of numbers
   - Min Heap: stores larger half of numbers

2. **Step 2**: When adding a number:
   - If it's smaller than max heap's top, add to max heap
   - If it's larger than min heap's top, add to min heap
   - Keep heaps balanced (size difference ≤ 1)

3. **Step 3**: To find median:
   - If heaps are equal size: average of both tops
   - If one heap is larger: top of larger heap

## Simple Code Example

```go
type MedianFinder struct {
    maxHeap *IntHeap // stores smaller half
    minHeap *IntHeap // stores larger half
}

func Constructor() MedianFinder {
    return MedianFinder{
        maxHeap: &IntHeap{},
        minHeap: &IntHeap{},
    }
}

func (this *MedianFinder) AddNum(num int) {
    // Add to max heap if empty or number is smaller
    if this.maxHeap.Len() == 0 || num <= (*this.maxHeap)[0] {
        heap.Push(this.maxHeap, num)
    } else {
        heap.Push(this.minHeap, num)
    }
    
    // Balance heaps
    if this.maxHeap.Len() > this.minHeap.Len()+1 {
        heap.Push(this.minHeap, heap.Pop(this.maxHeap))
    } else if this.minHeap.Len() > this.maxHeap.Len()+1 {
        heap.Push(this.maxHeap, heap.Pop(this.minHeap))
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.maxHeap.Len() == this.minHeap.Len() {
        return float64((*this.maxHeap)[0]+(*this.minHeap)[0]) / 2
    }
    if this.maxHeap.Len() > this.minHeap.Len() {
        return float64((*this.maxHeap)[0])
    }
    return float64((*this.minHeap)[0])
}
```

## Common Mistakes to Avoid

1. **Heap Balance**: Keep heaps balanced (size difference ≤ 1)
2. **Wrong Heap Type**: Use max heap for smaller numbers, min heap for larger
3. **Empty Heaps**: Handle cases when heaps are empty
4. **Heap Operations**: Use proper heap operations (Push, Pop)

## Fun Practice Problems

1. **Score Tracker**: Keep track of highest and lowest scores
2. **Temperature Monitor**: Track hottest and coldest temperatures
3. **Price Watcher**: Monitor highest and lowest prices
4. **Number Stream**: Find median of a stream of numbers
5. **Sliding Window**: Find median in a sliding window

## LeetCode Problems Using Two Heaps

Here are some popular LeetCode problems that can be solved using Two Heaps:

### Easy Problems

1. **[#295 Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/)** - Find median of stream of numbers.
   - **Approach**: Use two heaps to maintain smaller and larger halves.

2. **[#703 Kth Largest Element in a Stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/)** - Find kth largest element.
   - **Approach**: Use min heap to maintain k largest elements.

### Medium Problems

1. **[#480 Sliding Window Median](https://leetcode.com/problems/sliding-window-median/)** - Find median in sliding window.
   - **Approach**: Use two heaps and maintain window elements.

2. **[#502 IPO](https://leetcode.com/problems/ipo/)** - Maximize capital with k projects.
   - **Approach**: Use two heaps for available and unavailable projects.

3. **[#1046 Last Stone Weight](https://leetcode.com/problems/last-stone-weight/)** - Find last stone weight.
   - **Approach**: Use max heap to always get heaviest stones.

4. **[#973 K Closest Points to Origin](https://leetcode.com/problems/k-closest-points-to-origin/)** - Find k closest points.
   - **Approach**: Use max heap to maintain k closest points.

### Hard Problems

1. **[#295 Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/)** - Find median of stream.
   - **Approach**: Use two heaps with lazy removal.

2. **[#480 Sliding Window Median](https://leetcode.com/problems/sliding-window-median/)** - Find median in sliding window.
   - **Approach**: Use two heaps with delayed removal.

### Tips for Solving LeetCode Two Heaps Problems

1. **Heap Selection**: Choose right heap type for each half
2. **Balance**: Keep heaps balanced
3. **Operations**: Use proper heap operations
4. **Edge Cases**: Handle empty heaps and single element
5. **Cleanup**: Remove old elements when needed

## Why Learn This Pattern?

The Two Heaps pattern is super useful because:
1. It's perfect for finding medians in streams
2. It helps track highest and lowest values efficiently
3. It's a favorite in coding interviews
4. It teaches important concepts about heaps
5. It's used in many real-world applications