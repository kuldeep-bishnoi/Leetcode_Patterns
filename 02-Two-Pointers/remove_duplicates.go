package twopointers

// RemoveDuplicates removes duplicates from a sorted array in-place
// Returns the new length of the array after removing duplicates
// Time Complexity: O(n) where n is the length of the array
// Space Complexity: O(1) as we modify the array in-place
func RemoveDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// 'nextNonDuplicate' is the index where the next non-duplicate element should go
	nextNonDuplicate := 1

	for i := 1; i < len(arr); i++ {
		if arr[nextNonDuplicate-1] != arr[i] {
			arr[nextNonDuplicate] = arr[i]
			nextNonDuplicate++
		}
	}

	return nextNonDuplicate // Length of the array after removing duplicates
}

// RemoveDuplicatesAllowingTwice removes duplicates from a sorted array, allowing at most two occurrences of each element
// Returns the new length of the array
// Time Complexity: O(n) where n is the length of the array
// Space Complexity: O(1)
func RemoveDuplicatesAllowingTwice(arr []int) int {
	if len(arr) <= 2 {
		return len(arr)
	}

	// 'nextPosition' is the index where the next element should go
	nextPosition := 2

	for i := 2; i < len(arr); i++ {
		// Only proceed if the current element is different from the element two positions back
		if arr[i] != arr[nextPosition-2] {
			arr[nextPosition] = arr[i]
			nextPosition++
		}
	}

	return nextPosition
}

// Example cases:
// Input: [2, 3, 3, 3, 6, 9, 9]
// Output: 4
// Explanation: The first four elements after removing duplicates will be [2, 3, 6, 9]

// Input: [2, 2, 2, 11]
// Output: 2
// Explanation: The first two elements after removing duplicates will be [2, 11]
