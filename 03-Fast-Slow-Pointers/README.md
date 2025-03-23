# Fast & Slow Pointers Pattern

## Introduction

The Fast & Slow Pointers pattern, also known as the "Tortoise and Hare" algorithm, is a pointer technique that uses two pointers moving at different speeds. This approach is particularly useful for problems involving cyclic linked lists or arrays.

The main idea is to have two pointers traversing through a data structure:
- The "slow" pointer moves one step at a time
- The "fast" pointer moves two steps at a time

By moving at different speeds, these pointers will eventually meet if there's a cycle in the data structure. This pattern is also useful for finding the middle element, detecting cycles, and identifying the starting point of cycles.

## How It Works

1. Initialize two pointers (fast and slow) to the start of the data structure
2. Move the slow pointer one step at a time
3. Move the fast pointer two steps at a time
4. If there's a cycle, the fast pointer will eventually catch up to the slow pointer
5. If there's no cycle, the fast pointer will reach the end of the data structure

## Time and Space Complexity

- **Time Complexity**: O(n) - where n is the number of elements in the data structure
- **Space Complexity**: O(1) - since we only use two pointers regardless of the input size

## When to Use Fast & Slow Pointers

Use the Fast & Slow Pointers pattern when:
- You need to detect cycles in a linked list or array
- You need to find the middle of a linked list in one pass
- You need to determine the length of a cycle
- You're solving problems like finding a happy number or determining if a linked list is a palindrome

## Common Problem Patterns

1. **Cycle Detection**
   - Determine if a linked list or sequence has a cycle
   - If the pointers meet, there's a cycle; otherwise, the fast pointer reaches the end

2. **Cycle Length**
   - Find the length of a cycle
   - Once a cycle is detected, count the steps needed for a pointer to come back to the same position

3. **Cycle Start**
   - Find the start node of a cycle
   - After cycle detection, reset one pointer to the beginning and move both one step at a time

4. **Middle of Linked List**
   - Find the middle of a linked list
   - When the fast pointer reaches the end, the slow pointer is at the middle

5. **Palindrome Linked List**
   - Check if a linked list is a palindrome
   - Find the middle, reverse the second half, and compare with the first half

## Implementation in Golang

```go
// LinkedList node structure
type ListNode struct {
    Val  int
    Next *ListNode
}

// Detect Cycle in a LinkedList
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    
    slow, fast := head, head
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next           // Move one step
        fast = fast.Next.Next      // Move two steps
        
        if slow == fast {          // Found the cycle
            return true
        }
    }
    
    return false                   // No cycle
}

// Find the start of the cycle in a LinkedList
func detectCycleStart(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    
    // Find if there's a cycle
    slow, fast := head, head
    hasCycle := false
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        
        if slow == fast {
            hasCycle = true
            break
        }
    }
    
    if !hasCycle {
        return nil
    }
    
    // Find the start of the cycle
    slow = head
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }
    
    return slow
}

// Find the middle of a LinkedList
func findMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    slow, fast := head, head
    
    // When fast reaches the end, slow will be at the middle
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    return slow
}
```

## Example Problems

1. **LinkedList Cycle**
2. **Start of LinkedList Cycle**
3. **Happy Number**
4. **Middle of the LinkedList**
5. **Palindrome LinkedList**
6. **Circular Array Loop**
7. **Find the Duplicate Number**
8. **Reorder List**

Each of these problems has a dedicated solution file in this directory. 