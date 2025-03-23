package fastslow

// FindCycleStart finds the node where the cycle begins in a linked list
// Returns nil if no cycle is found
// Time Complexity: O(n) where n is the number of nodes in the linked list
// Space Complexity: O(1)
func FindCycleStart(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 1. Find if there is a cycle and the meeting point
	slow, fast := head, head
	hasCycle := false

	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Move one step
		fast = fast.Next.Next // Move two steps

		if slow == fast { // Cycle detected
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		return nil // No cycle
	}

	// 2. Find the start of the cycle
	// When a cycle is present, if we reset one pointer to the head and move both pointers
	// one step at a time, they will meet at the start of the cycle
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow // This is the start of the cycle
}

// Mathematical Explanation:
// Let's say:
// - Distance from head to cycle start = x
// - Distance from cycle start to meeting point = y
// - Cycle length = c

// When slow and fast pointers meet:
// Distance traveled by slow pointer = x + y
// Distance traveled by fast pointer = x + y + n*c (where n is some integer)

// Since fast pointer moves twice as fast:
// 2*(x + y) = x + y + n*c
// x + y = n*c
// x = n*c - y

// This means, from the meeting point, if we move x more steps, we'll be back at the start of the cycle.
// Also, from the head, if we move x steps, we'll also reach the start of the cycle.
// That's why we reset one pointer to head and move both pointers one step at a time.

// Example cases:
// For a linked list: 1->2->3->4->5->2 (connecting back to node with value 2)
// FindCycleStart returns: node with value 2

// For a linked list: 1->2->3->4->5 (no cycle)
// FindCycleStart returns: nil
