# Top 'K' Elements Pattern

## What are Top 'K' Elements?

Imagine you have a big box of toys, and you want to find your favorite ones! The Top 'K' Elements pattern is like having a special helper that can quickly find the best K things from a big list. It's like having a magic sorting hat that can pick out the top K students for a special class!

## Real-Life Examples

1. **Favorite Toys**: Finding your top 5 favorite toys from a big collection.
   - Look at all toys
   - Pick the 5 you like most
   - Keep them separate from others

2. **Class Grades**: Finding the top 10 students in a class.
   - Look at all grades
   - Pick the 10 highest scores
   - Make a special list

3. **Movie Ratings**: Finding the top 3 movies from a list.
   - Look at all ratings
   - Pick the 3 highest rated
   - Show them first

## When Do We Use Top 'K' Elements?

Use this technique when:
- You need to find K largest or smallest elements
- You want to find most frequent elements
- You need to sort partially
- You want to find closest elements
- You need to find top K pairs

## How Does It Work?

1. **Step 1**: Create a special container (heap)
2. **Step 2**: Add elements one by one
3. **Step 3**: Keep only K elements
4. **Step 4**: Remove smallest/largest as needed

Example:
```
Find top 3 numbers: [5, 2, 8, 1, 9, 3]
1. Add 5 → [5]
2. Add 2 → [2, 5]
3. Add 8 → [2, 5, 8]
4. Add 1 → [2, 5, 8] (1 is too small)
5. Add 9 → [5, 8, 9] (2 is removed)
6. Add 3 → [5, 8, 9] (3 is too small)
```

## Simple Code Example

```go
func findTopK(nums []int, k int) []int {
    // Create min heap
    h := &IntHeap{}
    heap.Init(h)
    
    // Add numbers to heap
    for _, num := range nums {
        heap.Push(h, num)
        // Keep only k elements
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    
    // Convert heap to slice
    result := make([]int, k)
    for i := k - 1; i >= 0; i-- {
        result[i] = heap.Pop(h).(int)
    }
    
    return result
}
```

## Common Mistakes to Avoid

1. **Heap Type**: Use min heap for largest K, max heap for smallest K
2. **Size Control**: Keep heap size at K
3. **Order**: Remember heap order
4. **Edge Cases**: Handle empty input and K > input size

## Fun Practice Problems

1. **Toy Picker**: Find top K favorite toys
2. **Grade Sorter**: Find top K students
3. **Movie Ranker**: Find top K movies
4. **Number Finder**: Find K largest numbers
5. **Word Counter**: Find K most frequent words

## LeetCode Problems Using Top 'K' Elements

Here are some popular LeetCode problems that can be solved using Top 'K' Elements:

### Easy Problems

1. **[#703 Kth Largest Element](https://leetcode.com/problems/kth-largest-element-in-a-stream/)** - Find Kth largest.
   - **Approach**: Use min heap.

2. **[#1046 Last Stone Weight](https://leetcode.com/problems/last-stone-weight/)** - Find remaining stone.
   - **Approach**: Use max heap.

### Medium Problems

1. **[#215 Kth Largest Element](https://leetcode.com/problems/kth-largest-element-in-an-array/)** - Find Kth largest.
   - **Approach**: Use heap or quickselect.

2. **[#347 Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/)** - Find K most frequent.
   - **Approach**: Use heap with frequency map.

### Hard Problems

1. **[#23 Merge K Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)** - Merge K lists.
   - **Approach**: Use heap with list heads.

2. **[#295 Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/)** - Find running median.
   - **Approach**: Use two heaps.

### Tips for Solving LeetCode Top 'K' Elements Problems

1. **Heap Selection**: Choose right heap type
   - Min heap for largest K
   - Max heap for smallest K

2. **Size Management**: Keep heap size at K
   - Remove smallest/largest
   - Add new elements
   - Maintain order

3. **Complexity**: Understand time complexity
   - Building heap: O(n)
   - Each operation: O(log k)
   - Overall: O(n log k)

4. **Optimization**: Use appropriate data structure
   - Heap for dynamic data
   - Quickselect for static data
   - Bucket sort for frequencies

## Why Learn This Pattern?

The Top 'K' Elements pattern is super useful because:
1. It's very efficient (O(n log k) time)
2. It's used in many real-world applications
3. It's a favorite in coding interviews
4. It teaches important concepts about heaps
5. It helps solve many sorting-related problems

Once you master this pattern, you'll be able to solve many selection problems efficiently and impress your friends with your coding skills! 