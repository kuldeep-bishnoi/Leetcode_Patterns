# Fast & Slow Pointers Pattern

## What are Fast & Slow Pointers?

Imagine you're playing a game of "Catch Me If You Can" with a friend. You run at normal speed while your friend runs twice as fast. If you're running in a circle, your friend will eventually catch up to you! This is exactly how the Fast & Slow Pointers technique works in programming - one pointer moves faster than the other to help us solve special problems.

## Real-Life Examples

1. **Tortoise and Hare Race**: Like the famous story, when a tortoise and hare race on a circular track, the faster hare will eventually catch up to the slower tortoise.

2. **School Bus and Sports Car**: If a school bus and a sports car drive on a circular road, the sports car will eventually catch up to the bus because it's moving faster.

3. **Walking and Running**: When walking and running on a circular track, the runner will eventually catch up to the walker.

## When Do We Use Fast & Slow Pointers?

Use this technique when you need to:
- Find if there's a cycle in a linked list
- Find the middle of a linked list
- Find if a number is happy
- Find the start of a cycle
- Find if a linked list is a palindrome

## Types of Fast & Slow Pointers

1. **Cycle Detection**: Fast pointer moves twice as fast as slow pointer
   - Example: "Check if a linked list has a cycle"
   
   ![Cycle Detection Visualization]
   ```
   List: 1 → 2 → 3 → 4 → 5 → 3 (cycle back to 3)
   
   Start: 1 → 2 → 3 → 4 → 5 → 3
          ↑
          slow
          ↑
          fast
   
   Step 1: 1 → 2 → 3 → 4 → 5 → 3
             ↑
             slow
                ↑
             fast
   
   Step 2: 1 → 2 → 3 → 4 → 5 → 3
                ↑
                slow
                   ↑
                fast
   
   Step 3: 1 → 2 → 3 → 4 → 5 → 3
                   ↑
                   slow
                      ↑
                   fast
   
   Step 4: 1 → 2 → 3 → 4 → 5 → 3
                      ↑
                      slow
                      ↑
                      fast
   
   They meet at 3, so there's a cycle!
   ```

2. **Finding Middle**: Fast pointer moves twice as fast to find the middle
   - Example: "Find the middle of a linked list"
   
   ![Middle Finding Visualization]
   ```
   List: 1 → 2 → 3 → 4 → 5
   
   Start: 1 → 2 → 3 → 4 → 5
          ↑
          slow
          ↑
          fast
   
   Step 1: 1 → 2 → 3 → 4 → 5
             ↑
             slow
                ↑
             fast
   
   Step 2: 1 → 2 → 3 → 4 → 5
                ↑
                slow
                   ↑
                fast
   
   Step 3: 1 → 2 → 3 → 4 → 5
                   ↑
                   slow
                      ↑
                   fast
   
   When fast reaches the end, slow is at the middle (3)!
   ```

## How Does It Work?

1. **For Cycle Detection**:
   - Start both pointers at the beginning
   - Move slow pointer one step at a time
   - Move fast pointer two steps at a time
   - If they meet, there's a cycle!
   - If fast reaches the end, there's no cycle

2. **For Finding Middle**:
   - Start both pointers at the beginning
   - Move slow pointer one step at a time
   - Move fast pointer two steps at a time
   - When fast reaches the end, slow is at the middle

## Simple Code Examples

### Cycle Detection Example
Checking if a linked list has a cycle:

```go
type ListNode struct {
    Val  int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    // If list is empty or has only one node, no cycle
    if head == nil || head.Next == nil {
        return false
    }
    
    // Start both pointers at the beginning
    slow := head
    fast := head
    
    // Keep moving until fast reaches the end or meets slow
    for fast != nil && fast.Next != nil {
        // Move slow one step
        slow = slow.Next
        // Move fast two steps
        fast = fast.Next.Next
        
        // If they meet, there's a cycle
        if slow == fast {
            return true
        }
    }
    
    // If fast reached the end, no cycle
    return false
}
```

### Finding Middle Example
Finding the middle of a linked list:

```go
func findMiddle(head *ListNode) *ListNode {
    // If list is empty, return nil
    if head == nil {
        return nil
    }
    
    // Start both pointers at the beginning
    slow := head
    fast := head
    
    // Keep moving until fast reaches the end
    for fast != nil && fast.Next != nil {
        // Move slow one step
        slow = slow.Next
        // Move fast two steps
        fast = fast.Next.Next
    }
    
    // When fast reaches the end, slow is at the middle
    return slow
}
```

## Common Mistakes to Avoid

1. **Forgetting to check for nil**: Always make sure to check if the list is empty or has only one node
2. **Incorrect pointer movement**: Be careful about moving fast pointer two steps
3. **Not handling edge cases**: What if the list is empty? What if it has only one node?

## Fun Practice Problems

1. **Race Track**: You have a circular race track with checkpoints. Two runners start at the same point, one running twice as fast as the other. Will they ever meet again?

2. **Happy Number Game**: A number is happy if you can make it 1 by repeatedly squaring its digits. For example, 19 is happy because:
   - 1² + 9² = 82
   - 8² + 2² = 68
   - 6² + 8² = 100
   - 1² + 0² + 0² = 1

3. **Palindrome Checker**: Check if a linked list reads the same forwards and backwards (like "racecar").

4. **Cycle Finder**: Given a linked list, find if it has a cycle and where the cycle starts.

5. **Middle Finder**: Find the middle of a linked list in one pass.

## LeetCode Problems Using Fast & Slow Pointers

Here are some popular LeetCode problems that can be solved using the Fast & Slow Pointers technique:

### Easy Problems

1. **[#141 Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)** - Check if a linked list has a cycle.
   - **Approach**: Use fast and slow pointers to detect cycle.

2. **[#876 Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/)** - Find the middle node of a linked list.
   - **Approach**: Use fast and slow pointers to find middle.

3. **[#234 Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)** - Check if a linked list is a palindrome.
   - **Approach**: Use fast and slow pointers to find middle, then reverse second half.

4. **[#202 Happy Number](https://leetcode.com/problems/happy-number/)** - Check if a number is happy.
   - **Approach**: Use fast and slow pointers to detect cycle in digit square sum sequence.

### Medium Problems

1. **[#142 Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)** - Find the start of a cycle in a linked list.
   - **Approach**: Use fast and slow pointers to detect cycle, then find cycle start.

2. **[#287 Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)** - Find the duplicate number in an array.
   - **Approach**: Treat array as linked list and use fast and slow pointers.

3. **[#143 Reorder List](https://leetcode.com/problems/reorder-list/)** - Reorder linked list in L0→Ln→L1→Ln-1→L2→Ln-2→... order.
   - **Approach**: Use fast and slow pointers to find middle, then reverse second half.

4. **[#328 Odd Even Linked List](https://leetcode.com/problems/odd-even-linked-list/)** - Group odd and even nodes together.
   - **Approach**: Use multiple pointers to separate odd and even nodes.

5. **[#61 Rotate List](https://leetcode.com/problems/rotate-list/)** - Rotate linked list to the right by k places.
   - **Approach**: Use fast and slow pointers to find rotation point.

### Hard Problems

1. **[#25 Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/)** - Reverse nodes in groups of k.
   - **Approach**: Use multiple pointers to handle group reversal.

2. **[#23 Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)** - Merge k sorted linked lists.
   - **Approach**: Use fast and slow pointers for efficient merging.

3. **[#148 Sort List](https://leetcode.com/problems/sort-list/)** - Sort a linked list in O(n log n) time.
   - **Approach**: Use fast and slow pointers for merge sort implementation.

4. **[#138 Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/)** - Deep copy a linked list with random pointers.
   - **Approach**: Use fast and slow pointers for efficient copying.

### Tips for Solving LeetCode Fast & Slow Pointer Problems

1. **Identify the Need**: Determine if you need cycle detection or middle finding
2. **Pointer Movement**: Implement correct movement speeds for fast and slow pointers
3. **Cycle Detection**: Remember to check for nil and handle edge cases
4. **Optimization**: Look for ways to solve in one pass
5. **Edge Cases**: Consider empty lists, single nodes, and no cycles

## Why Learn This Pattern?

The Fast & Slow Pointers pattern is super useful because:
1. It helps solve complex problems in a simple and efficient way
2. It's used in many real-world applications like network cycle detection
3. It's a favorite in coding interviews and competitions
4. It can solve problems that seem impossible with a single pointer