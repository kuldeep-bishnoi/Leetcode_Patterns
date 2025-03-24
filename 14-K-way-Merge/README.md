# K-way Merge Pattern

## What is K-way Merge?

Imagine you have K different piles of cards, and each pile is already sorted from smallest to biggest. The K-way Merge pattern is like having a special helper that can combine all these piles into one big sorted pile! It's like having a magic sorting machine that can merge many sorted lists into one perfect list.

## Real-Life Examples

1. **Merging Decks**: Combining many sorted card decks.
   - Each deck is sorted
   - Need to keep them sorted
   - Take smallest card each time

2. **Class Lists**: Combining many sorted class lists.
   - Each list is sorted by grade
   - Need one big sorted list
   - Take highest grade each time

3. **Book Shelves**: Combining many sorted book shelves.
   - Each shelf is sorted by title
   - Need one big sorted shelf
   - Take first alphabetically each time

## When Do We Use K-way Merge?

Use this technique when:
- You have K sorted lists to merge
- You need to find smallest/largest elements
- You want to combine sorted data
- You need to merge multiple streams
- You want to find common elements

## How Does It Work?

1. **Step 1**: Create a special container (heap)
2. **Step 2**: Add first element from each list
3. **Step 3**: Take smallest/largest element
4. **Step 4**: Add next element from that list
5. **Step 5**: Repeat until all lists are empty

Example:
```
Lists: [1,4,7], [2,5,8], [3,6,9]
1. Add 1,2,3 → Take 1
2. Add 4,2,3 → Take 2
3. Add 4,5,3 → Take 3
4. Add 4,5,6 → Take 4
And so on...
```

## Simple Code Example

```go
func mergeKLists(lists []*ListNode) *ListNode {
    // Create min heap
    h := &ListNodeHeap{}
    heap.Init(h)
    
    // Add first node from each list
    for _, list := range lists {
        if list != nil {
            heap.Push(h, list)
        }
    }
    
    // Create result list
    dummy := &ListNode{}
    curr := dummy
    
    // Keep merging until heap is empty
    for h.Len() > 0 {
        node := heap.Pop(h).(*ListNode)
        curr.Next = node
        curr = curr.Next
        
        // Add next node from same list
        if node.Next != nil {
            heap.Push(h, node.Next)
        }
    }
    
    return dummy.Next
}
```

## Common Mistakes to Avoid

1. **Heap Type**: Use min heap for ascending, max heap for descending
2. **Empty Lists**: Handle empty lists properly
3. **Next Node**: Always add next node from same list
4. **Edge Cases**: Handle nil lists and single list

## Fun Practice Problems

1. **Deck Merger**: Merge K sorted card decks
2. **List Combiner**: Merge K sorted class lists
3. **Shelf Merger**: Merge K sorted book shelves
4. **Number Merger**: Merge K sorted number lists
5. **Stream Combiner**: Merge K sorted data streams

## LeetCode Problems Using K-way Merge

Here are some popular LeetCode problems that can be solved using K-way Merge:

### Easy Problems

1. **[#21 Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)** - Merge two lists.
   - **Approach**: Use two pointers.

2. **[#88 Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/)** - Merge two arrays.
   - **Approach**: Merge from end.

### Medium Problems

1. **[#23 Merge K Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)** - Merge K lists.
   - **Approach**: Use heap.

2. **[#378 Kth Smallest Element](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/)** - Find Kth smallest.
   - **Approach**: Use heap with matrix.

### Hard Problems

1. **[#632 Smallest Range](https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/)** - Find smallest range.
   - **Approach**: Use heap with range tracking.

2. **[#786 K-th Smallest Prime Fraction](https://leetcode.com/problems/k-th-smallest-prime-fraction/)** - Find Kth fraction.
   - **Approach**: Use heap with fractions.

### Tips for Solving LeetCode K-way Merge Problems

1. **Heap Selection**: Choose right heap type
   - Min heap for ascending
   - Max heap for descending

2. **Node Management**: Handle nodes properly
   - Add first node from each list
   - Add next node after removing
   - Update pointers correctly

3. **Complexity**: Understand time complexity
   - Building heap: O(k)
   - Each operation: O(log k)
   - Overall: O(n log k)

4. **Optimization**: Use appropriate approach
   - Heap for dynamic lists
   - Two pointers for two lists
   - Binary search for matrix

## Why Learn This Pattern?

The K-way Merge pattern is super useful because:
1. It's very efficient (O(n log k) time)
2. It's used in many real-world applications
3. It's a favorite in coding interviews
4. It teaches important concepts about heaps
5. It helps solve many merging problems

Once you master this pattern, you'll be able to solve many merging problems efficiently and impress your friends with your coding skills! 