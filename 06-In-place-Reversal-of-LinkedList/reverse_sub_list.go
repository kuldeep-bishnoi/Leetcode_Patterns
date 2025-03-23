package inplace

// ReverseBetween reverses a sub-list between positions p and q (1-indexed)
// Time Complexity: O(n) where n is the number of nodes in the linked list
// Space Complexity: O(1)
func ReverseBetween(head *ListNode, p, q int) *ListNode {
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

	// current points to the pth node (this will be the last node after reversal)
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
	beforeP.Next.Next = current // Link the pth node (now the last in the reversed portion) to the node after q
	beforeP.Next = prev         // Link the node before p to the qth node (now the first in the reversed portion)

	return dummy.Next
}

// Example:
// Input: 1->2->3->4->5->NULL, p = 2, q = 4
// Output: 1->4->3->2->5->NULL
//
// Visualization:
// Initial state:
// 1 -> 2 -> 3 -> 4 -> 5 -> nil
//      p         q
//
// After finding beforeP (node before p):
// beforeP = 1, current = 2
// 1 -> 2 -> 3 -> 4 -> 5 -> nil
// |    |
// beforeP current
//
// Reversing nodes between p and q:
// 1    2 <- 3 <- 4    5 -> nil
// |    |         |    |
// beforeP       prev current
//
// Connecting the reversed portion with the rest of the list:
// beforeP.Next.Next = current: 2->5
// beforeP.Next = prev: 1->4
// 1 -> 4 -> 3 -> 2 -> 5 -> nil
//
// Result: 1 -> 4 -> 3 -> 2 -> 5 -> nil
