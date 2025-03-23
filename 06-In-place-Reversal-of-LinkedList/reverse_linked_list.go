package inplace

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseLinkedList reverses a linked list in-place and returns the new head
// Time Complexity: O(n) where n is the number of nodes in the linked list
// Space Complexity: O(1)
func ReverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head

	for current != nil {
		next := current.Next // Store next node
		current.Next = prev  // Reverse the link
		prev = current       // Move prev to current
		current = next       // Move current to next
	}

	// prev is the new head of the reversed list
	return prev
}

// Example:
// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
//
// Visualization:
// Initial state:
// prev: nil, current: 1, next: nil
// 1 -> 2 -> 3 -> 4 -> 5 -> nil
//
// First iteration:
// next = 2, current.Next = nil, prev = 1, current = 2
// nil <- 1    2 -> 3 -> 4 -> 5 -> nil
//      prev  current
//
// Second iteration:
// next = 3, current.Next = 1, prev = 2, current = 3
// nil <- 1 <- 2    3 -> 4 -> 5 -> nil
//           prev  current
//
// Third iteration:
// next = 4, current.Next = 2, prev = 3, current = 4
// nil <- 1 <- 2 <- 3    4 -> 5 -> nil
//                prev  current
//
// Fourth iteration:
// next = 5, current.Next = 3, prev = 4, current = 5
// nil <- 1 <- 2 <- 3 <- 4    5 -> nil
//                     prev  current
//
// Fifth iteration:
// next = nil, current.Next = 4, prev = 5, current = nil
// nil <- 1 <- 2 <- 3 <- 4 <- 5    nil
//                          prev  current
//
// Result: return prev = 5 (new head)
// 5 -> 4 -> 3 -> 2 -> 1 -> nil
