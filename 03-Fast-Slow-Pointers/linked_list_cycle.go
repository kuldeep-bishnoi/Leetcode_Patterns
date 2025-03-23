package fastslow

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle determines if a linked list has a cycle
// Time Complexity: O(n) where n is the number of nodes in the linked list
// Space Complexity: O(1)
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// Initialize slow and fast pointers
	slow, fast := head, head

	// Move slow pointer one step and fast pointer two steps at a time
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Move one step
		fast = fast.Next.Next // Move two steps

		// If there's a cycle, the fast pointer will eventually meet the slow pointer
		if slow == fast {
			return true
		}
	}

	// If fast pointer reaches the end, there's no cycle
	return false
}

// FindCycleLength determines the length of a cycle in a linked list
// Returns 0 if no cycle is found
// Time Complexity: O(n) where n is the number of nodes in the linked list
// Space Complexity: O(1)
func FindCycleLength(head *ListNode) int {
	if head == nil || head.Next == nil {
		return 0
	}

	// Initialize slow and fast pointers
	slow, fast := head, head

	// Detect cycle
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast { // Cycle detected
			// Calculate cycle length
			current := slow
			cycleLength := 0

			// Count the length of the cycle
			for {
				current = current.Next
				cycleLength++

				// If we've come back to the same point, we've completed the cycle
				if current == slow {
					return cycleLength
				}
			}
		}
	}

	return 0 // No cycle found
}

// Example cases:
// For a linked list: 1->2->3->4->5->2 (connecting back to node with value 2)
// HasCycle returns: true
// FindCycleLength returns: 4 (nodes 2->3->4->5 form a cycle of length 4)

// For a linked list: 1->2->3->4->5 (no cycle)
// HasCycle returns: false
// FindCycleLength returns: 0
