package inplace

// ReverseKGroup reverses every k-element sub-list in a linked list
// Time Complexity: O(n) where n is the number of nodes in the linked list
// Space Complexity: O(1)
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	// Create a dummy node to handle edge cases
	dummy := &ListNode{Next: head}
	beforeGroup := dummy

	count := 0
	current := head

	// Count total nodes to handle the case where the last group has fewer than k elements
	for current != nil {
		count++
		current = current.Next
	}

	current = head

	// Process each group of k elements
	for count >= k {
		// The first node of the current group, will be the last after reversal
		groupStart := current

		// Reverse k nodes
		var prev *ListNode = nil
		for i := 0; i < k; i++ {
			next := current.Next
			current.Next = prev
			prev = current
			current = next
		}

		// Connect the reversed group with the rest of the list
		groupStart.Next = current // Connect the last node of the reversed group to the next group
		beforeGroup.Next = prev   // Connect the previous group to the first node of the reversed group

		// Update beforeGroup to be the last node of the current group
		beforeGroup = groupStart

		// Update count
		count -= k
	}

	return dummy.Next
}

// Example:
// Input: 1->2->3->4->5->6->7->8->NULL, k = 3
// Output: 3->2->1->6->5->4->8->7->NULL
//
// Visualization:
// Initial state:
// 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> nil
//
// After first group reversal:
//           |-----------|
// dummy -> 3 -> 2 -> 1 -> 4 -> 5 -> 6 -> 7 -> 8 -> nil
//   |      |         |    |
//   beforeGroup    groupStart current
//
// After second group reversal:
//           |-----------|    |-----------|
// dummy -> 3 -> 2 -> 1 -> 6 -> 5 -> 4 -> 7 -> 8 -> nil
//                  |    |         |    |
//               beforeGroup    groupStart current
//
// After third group reversal (only 2 elements, so no reversal needed):
//           |-----------|    |-----------|    |-----|
// dummy -> 3 -> 2 -> 1 -> 6 -> 5 -> 4 -> 8 -> 7 -> nil
//
// Result: 3 -> 2 -> 1 -> 6 -> 5 -> 4 -> 8 -> 7 -> nil
