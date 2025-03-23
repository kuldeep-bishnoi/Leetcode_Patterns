# In-place Reversal of a LinkedList Pattern

## Introduction

The In-place Reversal of a LinkedList pattern is a technique used to reverse the links between nodes in a linked list without using additional data structures. This approach is space-efficient as it modifies the existing structure to reverse the list.

This pattern is particularly useful for problems that require reversing a linked list, either completely or in parts, such as reversing a sublist between specific nodes or reversing alternate k-element groups.

## How It Works

1. Initialize three pointers: previous (initially null), current (head of the list), and next
2. Iterate through the list
3. For each node, save the next node, then reverse the current node's link to point to the previous node
4. Update pointers for the next iteration

## Time and Space Complexity

- **Time Complexity**: O(n) - where n is the number of nodes in the linked list
- **Space Complexity**: O(1) - since we're only using a constant amount of extra space

## When to Use In-place Reversal of a LinkedList

Use the In-place Reversal of a LinkedList pattern when:
- You need to reverse a linked list or a portion of it
- You're asked to reorder elements in a linked list
- You need to modify the structure of a linked list without using extra space
- You're dealing with problems like reversing every k elements, alternating reversals, etc.

## Common Problem Patterns

1. **Reverse a LinkedList**
   - Reverse the entire linked list
   - Update head to point to the new first node (previously the last node)

2. **Reverse a Sub-list**
   - Reverse a portion of the linked list between positions p and q
   - Connect the nodes before and after the reversed portion correctly

3. **Reverse Every K-element Sub-list**
   - Divide the linked list into groups of k nodes
   - Reverse each group
   - Connect the reversed groups in the correct order

4. **Alternating Reverse**
   - Reverse alternate groups of k elements
   - Leave other groups as they are

## Implementation in Golang

```go
// ListNode represents a node in a linked list
type ListNode struct {
    Val  int
    Next *ListNode
}

// Reverse an entire linked list
func reverseLinkedList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    current := head
    
    for current != nil {
        next := current.Next       // Store next node
        current.Next = prev        // Reverse the link
        prev = current             // Move prev to current
        current = next             // Move current to next
    }
    
    // prev is the new head of the reversed list
    return prev
}

// Reverse a sub-list between positions p and q
func reverseBetween(head *ListNode, p, q int) *ListNode {
    if head == nil || p == q {
        return head
    }
    
    // Create a dummy node to handle edge case when p = 1
    dummy := &ListNode{Next: head}
    beforeP := dummy
    
    // Move beforeP to the node just before the pth node
    for i := 1; i < p; i++ {
        beforeP = beforeP.Next
    }
    
    // current points to the pth node
    current := beforeP.Next
    
    // Reverse nodes between p and q
    var prev *ListNode = nil
    for i := 0; i <= q-p; i++ {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }
    
    // Connect the reversed portion with the rest of the list
    beforeP.Next.Next = current  // Link the pth node (now the last in the reversed portion) to the node after q
    beforeP.Next = prev         // Link the node before p to the qth node (now the first in the reversed portion)
    
    return dummy.Next
}
```

## Example Problems

1. **Reverse a LinkedList**
2. **Reverse a Sub-list**
3. **Reverse Every K-element Sub-list**
4. **Reverse Alternating K-element Sub-list**
5. **Rotate a LinkedList**
6. **Palindrome LinkedList**

Each of these problems has a dedicated solution file in this directory. 