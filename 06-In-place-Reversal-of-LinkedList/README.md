# In-place Reversal of Linked List Pattern

## What is In-place Reversal?

Imagine you have a chain of paper clips connected together. In-place reversal is like turning this chain around so that the last paper clip becomes the first one, and the first becomes the last. We do this without breaking the chain or using any extra paper clips!

## Real-Life Examples

1. **Train Cars**: When a train needs to go in the opposite direction, the engine moves to the other end.
   - Last car becomes first
   - First car becomes last
   - All cars stay connected

2. **Bead Necklace**: When you want to wear a necklace backwards.
   - Last bead becomes first
   - First bead becomes last
   - All beads stay connected

3. **Line of Students**: When students need to face the opposite direction.
   - Last student becomes first
   - First student becomes last
   - Everyone stays in line

## When Do We Use In-place Reversal?

Use this technique when:
- You need to reverse a linked list
- You need to reverse parts of a linked list
- You need to reverse nodes in groups
- You need to reverse nodes between certain positions

## How Does It Work?

1. **Step 1**: Keep track of three pointers:
   - `prev`: points to the previous node
   - `curr`: points to the current node
   - `next`: points to the next node

2. **Step 2**: For each node:
   - Save the next node
   - Make current node point to previous node
   - Move previous to current
   - Move current to next

3. **Step 3**: Keep going until we reach the end

## Simple Code Example

```go
type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    // Initialize pointers
    var prev *ListNode = nil
    curr := head
    
    // Keep reversing until we reach the end
    for curr != nil {
        // Save the next node
        next := curr.Next
        
        // Reverse the link
        curr.Next = prev
        
        // Move pointers forward
        prev = curr
        curr = next
    }
    
    // Return the new head (which was the last node)
    return prev
}
```

## Common Mistakes to Avoid

1. **Losing References**: Always save the next node before changing links
2. **Wrong Order**: Make sure to update pointers in the correct order
3. **Edge Cases**: Handle empty lists and single-node lists carefully

## Fun Practice Problems

1. **Reverse Train**: Reverse the order of train cars
2. **Reverse Beads**: Reverse the order of beads in a necklace
3. **Reverse Students**: Reverse the order of students in a line
4. **Reverse Groups**: Reverse students in groups of 3
5. **Reverse Between**: Reverse students between positions 2 and 5

## LeetCode Problems Using In-place Reversal

Here are some popular LeetCode problems that can be solved using In-place Reversal:

### Easy Problems

1. **[#206 Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)** - Reverse a singly linked list.
   - **Approach**: Use three pointers to reverse links.

2. **[#234 Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)** - Check if a linked list is a palindrome.
   - **Approach**: Reverse second half and compare with first half.

### Medium Problems

1. **[#92 Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/)** - Reverse nodes between positions m and n.
   - **Approach**: Find start and end positions, then reverse that portion.

2. **[#25 Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/)** - Reverse nodes in groups of k.
   - **Approach**: Reverse each group of k nodes.

3. **[#143 Reorder List](https://leetcode.com/problems/reorder-list/)** - Reorder list in L0→Ln→L1→Ln-1→L2→Ln-2→... order.
   - **Approach**: Find middle, reverse second half, and merge.

4. **[#24 Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/)** - Swap every two adjacent nodes.
   - **Approach**: Reverse nodes in pairs.

### Hard Problems

1. **[#25 Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/)** - Reverse nodes in groups of k.
   - **Approach**: Reverse each group of k nodes, handle remaining nodes.

2. **[#23 Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)** - Merge k sorted linked lists.
   - **Approach**: Use priority queue or divide and conquer.

### Tips for Solving LeetCode In-place Reversal Problems

1. **Pointer Management**: Keep track of all necessary pointers
2. **Link Updates**: Update links in the correct order
3. **Edge Cases**: Handle empty lists and single nodes
4. **Group Reversal**: For k-group problems, handle remaining nodes
5. **Visualization**: Draw the list before and after each step

## Why Learn This Pattern?

The In-place Reversal pattern is super useful because:
1. It's a fundamental operation on linked lists
2. It's space-efficient (uses O(1) extra space)
3. It's a favorite in coding interviews
4. It teaches important concepts about pointer manipulation
5. It helps solve many linked list problems efficiently