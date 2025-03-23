package binarysearch

// FindCeiling finds the smallest element in the array that is greater than or equal to the target
// Returns the element if found, or -1 if no ceiling exists
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindCeiling(arr []int, target int) int {
	n := len(arr)

	// If the target is greater than the largest element, no ceiling exists
	if n == 0 || target > arr[n-1] {
		return -1
	}

	start, end := 0, n-1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return arr[mid] // Found exact match
		}

		if arr[mid] < target {
			start = mid + 1 // Look in the right half
		} else {
			end = mid - 1 // Look in the left half
		}
	}

	// At this point, 'start' is the index of the ceiling
	return arr[start]
}

// FindCeilingIndex finds the index of the smallest element in the array
// that is greater than or equal to the target
// Returns the index if found, or -1 if no ceiling exists
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindCeilingIndex(arr []int, target int) int {
	n := len(arr)

	// If the target is greater than the largest element, no ceiling exists
	if n == 0 || target > arr[n-1] {
		return -1
	}

	start, end := 0, n-1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return mid // Found exact match
		}

		if arr[mid] < target {
			start = mid + 1 // Look in the right half
		} else {
			end = mid - 1 // Look in the left half
		}
	}

	// At this point, 'start' is the index of the ceiling
	return start
}

// FindFloor finds the largest element in the array that is less than or equal to the target
// Returns the element if found, or -1 if no floor exists
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindFloor(arr []int, target int) int {
	n := len(arr)

	// If the target is less than the smallest element, no floor exists
	if n == 0 || target < arr[0] {
		return -1
	}

	start, end := 0, n-1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return arr[mid] // Found exact match
		}

		if arr[mid] < target {
			start = mid + 1 // Look in the right half
		} else {
			end = mid - 1 // Look in the left half
		}
	}

	// At this point, 'end' is the index of the floor
	return arr[end]
}

// FindFloorIndex finds the index of the largest element in the array
// that is less than or equal to the target
// Returns the index if found, or -1 if no floor exists
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindFloorIndex(arr []int, target int) int {
	n := len(arr)

	// If the target is less than the smallest element, no floor exists
	if n == 0 || target < arr[0] {
		return -1
	}

	start, end := 0, n-1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return mid // Found exact match
		}

		if arr[mid] < target {
			start = mid + 1 // Look in the right half
		} else {
			end = mid - 1 // Look in the left half
		}
	}

	// At this point, 'end' is the index of the floor
	return end
}

// Example usage:
//
// Ceiling:
// Input: arr = [1, 3, 8, 10, 15], target = 12
// Output: 15 (smallest element greater than or equal to 12)
//
// Ceiling Index:
// Input: arr = [1, 3, 8, 10, 15], target = 12
// Output: 4 (index of 15, which is the ceiling of 12)
//
// Floor:
// Input: arr = [1, 3, 8, 10, 15], target = 12
// Output: 10 (largest element less than or equal to 12)
//
// Floor Index:
// Input: arr = [1, 3, 8, 10, 15], target = 12
// Output: 3 (index of 10, which is the floor of 12)
//
// No Ceiling:
// Input: arr = [1, 3, 8, 10, 15], target = 20
// Output: -1 (no element greater than or equal to 20)
//
// No Floor:
// Input: arr = [1, 3, 8, 10, 15], target = 0
// Output: -1 (no element less than or equal to 0)
