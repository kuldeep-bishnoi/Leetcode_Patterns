package binarysearch

// NextLetter finds the smallest letter in the array that is lexicographically
// greater than the target. The letters wrap around, so if all letters are less
// than or equal to the target, return the first letter.
// Time Complexity: O(log n)
// Space Complexity: O(1)
func NextLetter(letters []byte, target byte) byte {
	n := len(letters)

	// Edge case: empty array
	if n == 0 {
		return 0 // Invalid input
	}

	// If target is greater than or equal to the last letter, wrap around
	if target >= letters[n-1] {
		return letters[0]
	}

	start, end := 0, n-1

	for start <= end {
		mid := start + (end-start)/2

		if letters[mid] <= target {
			start = mid + 1 // Look in the right half
		} else {
			end = mid - 1 // Look in the left half
		}
	}

	// At this point, 'start' is the index of the next letter
	return letters[start]
}

// NextLetterIgnoreCase finds the smallest letter in the array that is lexicographically
// greater than the target, ignoring case. The letters wrap around.
// Time Complexity: O(log n)
// Space Complexity: O(1)
func NextLetterIgnoreCase(letters []byte, target byte) byte {
	n := len(letters)

	// Edge case: empty array
	if n == 0 {
		return 0 // Invalid input
	}

	// Convert target to lowercase for case-insensitive comparison
	target = toLower(target)

	// If target is greater than or equal to the last letter, wrap around
	if target >= toLower(letters[n-1]) {
		return letters[0]
	}

	start, end := 0, n-1

	for start <= end {
		mid := start + (end-start)/2

		// Compare lowercase versions
		if toLower(letters[mid]) <= target {
			start = mid + 1 // Look in the right half
		} else {
			end = mid - 1 // Look in the left half
		}
	}

	// At this point, 'start' is the index of the next letter
	return letters[start]
}

// Helper function to convert a byte to lowercase
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

// NextGreaterElement finds the first element in the array that is greater than the target
// Time Complexity: O(log n)
// Space Complexity: O(1)
func NextGreaterElement(arr []int, target int) int {
	n := len(arr)

	if n == 0 {
		return -1 // Invalid input
	}

	// If target is greater than or equal to the last element, no next greater element
	if target >= arr[n-1] {
		return -1
	}

	start, end := 0, n-1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] <= target {
			start = mid + 1 // Look in the right half
		} else {
			end = mid - 1 // Look in the left half
		}
	}

	// At this point, 'start' is the index of the next greater element
	return arr[start]
}

// Example usage:
//
// Next Letter:
// Input: letters = ['a', 'c', 'f', 'h'], target = 'f'
// Output: 'h' (smallest letter greater than 'f')
//
// Input: letters = ['a', 'c', 'f', 'h'], target = 'h'
// Output: 'a' (wrap around to the first letter)
//
// Input: letters = ['a', 'c', 'f', 'h'], target = 'b'
// Output: 'c' (smallest letter greater than 'b')
//
// Next Letter Ignore Case:
// Input: letters = ['a', 'C', 'f', 'H'], target = 'F'
// Output: 'H' (smallest letter greater than 'F', ignoring case)
//
// Next Greater Element:
// Input: arr = [1, 3, 8, 10, 15], target = 8
// Output: 10 (first element greater than 8)
//
// Input: arr = [1, 3, 8, 10, 15], target = 12
// Output: 15 (first element greater than 12)
//
// Input: arr = [1, 3, 8, 10, 15], target = 15
// Output: -1 (no element greater than 15)
