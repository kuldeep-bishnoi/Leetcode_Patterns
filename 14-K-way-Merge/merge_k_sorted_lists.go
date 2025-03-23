package kwaymerge

import (
	"container/heap"
)

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// ListNodeElement represents an element in the heap for merging K sorted lists
type ListNodeElement struct {
	Node      *ListNode // Pointer to the current node
	ListIndex int       // Index of the list this node belongs to
}

// ListNodeMinHeap implementation
type ListNodeMinHeap []ListNodeElement

func (h ListNodeMinHeap) Len() int           { return len(h) }
func (h ListNodeMinHeap) Less(i, j int) bool { return h[i].Node.Val < h[j].Node.Val }
func (h ListNodeMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListNodeMinHeap) Push(x interface{}) {
	*h = append(*h, x.(ListNodeElement))
}

func (h *ListNodeMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MergeKLists merges k sorted linked lists into one sorted linked list
// Time Complexity: O(N log K) where N is the total number of nodes across all lists
// and K is the number of linked lists
// Space Complexity: O(K) for storing the heap with one node from each of the K lists
func MergeKLists(lists []*ListNode) *ListNode {
	// Handle edge cases
	if len(lists) == 0 {
		return nil
	}

	// Create a dummy head for the result list
	dummy := &ListNode{}
	current := dummy

	// Create a min heap
	h := &ListNodeMinHeap{}
	heap.Init(h)

	// Insert the first node from each list into the heap
	for i, list := range lists {
		if list != nil {
			heap.Push(h, ListNodeElement{
				Node:      list,
				ListIndex: i,
			})
		}
	}

	// Extract the minimum node and add the next node from the same list
	for h.Len() > 0 {
		// Get the minimum node
		minElement := heap.Pop(h).(ListNodeElement)

		// Add the node to the result list
		current.Next = minElement.Node
		current = current.Next

		// If there are more nodes in the same list, add the next one
		if minElement.Node.Next != nil {
			heap.Push(h, ListNodeElement{
				Node:      minElement.Node.Next,
				ListIndex: minElement.ListIndex,
			})
		}
	}

	return dummy.Next
}

// MergeKListsDivideConquer merges k sorted linked lists using divide and conquer
// Time Complexity: O(N log K) where N is the total number of nodes across all lists
// and K is the number of linked lists
// Space Complexity: O(log K) for recursion stack
func MergeKListsDivideConquer(lists []*ListNode) *ListNode {
	// Handle edge cases
	if len(lists) == 0 {
		return nil
	}

	return divideAndConquer(lists, 0, len(lists)-1)
}

// divideAndConquer is a helper function that recursively merges lists
func divideAndConquer(lists []*ListNode, start, end int) *ListNode {
	// Base case: single list
	if start == end {
		return lists[start]
	}

	// Base case: two lists
	if start+1 == end {
		return mergeTwoLists(lists[start], lists[end])
	}

	// Recursive case: divide the problem
	mid := start + (end-start)/2
	left := divideAndConquer(lists, start, mid)
	right := divideAndConquer(lists, mid+1, end)

	// Merge the results
	return mergeTwoLists(left, right)
}

// mergeTwoLists merges two sorted linked lists
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// Create a dummy head
	dummy := &ListNode{}
	current := dummy

	// Merge the two lists
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// Attach the remaining list
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}

// Example usage:
//
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked lists are:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6
//
// Input: lists = []
// Output: nil
//
// Input: lists = [[]]
// Output: nil
