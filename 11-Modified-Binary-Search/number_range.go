package binarysearch

// SearchRange finds the starting and ending position of a given target value
// Returns [-1, -1] if the target is not found in the array
// Time Complexity: O(log n)
// Space Complexity: O(1)
func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}

	if len(nums) == 0 {
		return result
	}

	// Find the first occurrence of the target
	result[0] = findBound(nums, target, true)

	// If the target doesn't exist in the array
	if result[0] == -1 {
		return result
	}

	// Find the last occurrence of the target
	result[1] = findBound(nums, target, false)

	return result
}

// findBound is a helper function to find the first or last occurrence of the target
// isFirst: if true, find the first occurrence; otherwise, find the last occurrence
func findBound(nums []int, target int, isFirst bool) int {
	start, end := 0, len(nums)-1
	index := -1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			// Found the target
			index = mid

			if isFirst {
				// Continue searching on the left to find the first occurrence
				end = mid - 1
			} else {
				// Continue searching on the right to find the last occurrence
				start = mid + 1
			}
		}
	}

	return index
}

// CountOccurrences counts the occurrences of a target number in a sorted array
// Time Complexity: O(log n)
// Space Complexity: O(1)
func CountOccurrences(nums []int, target int) int {
	range_ := SearchRange(nums, target)

	// If target is not found
	if range_[0] == -1 {
		return 0
	}

	// Count = last index - first index + 1
	return range_[1] - range_[0] + 1
}

// FindFirstEqualsK finds the first occurrence of an element equal to k in the array
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindFirstEqualsK(nums []int, k int) int {
	return findBound(nums, k, true)
}

// FindLastEqualsK finds the last occurrence of an element equal to k in the array
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindLastEqualsK(nums []int, k int) int {
	return findBound(nums, k, false)
}

// FindFirstGreaterThanK finds the first element greater than k in the array
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindFirstGreaterThanK(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	index := -1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] > k {
			index = mid
			end = mid - 1 // Continue searching on the left
		} else {
			start = mid + 1 // Look in the right half
		}
	}

	return index
}

// FindLastLessThanK finds the last element less than k in the array
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindLastLessThanK(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	index := -1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] < k {
			index = mid
			start = mid + 1 // Continue searching on the right
		} else {
			end = mid - 1 // Look in the left half
		}
	}

	return index
}

// Example usage:
//
// Search Range:
// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4] (first and last position of 8)
//
// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1] (target not found)
//
// Count Occurrences:
// Input: nums = [5,7,7,8,8,10], target = 8
// Output: 2 (8 appears twice)
//
// First Equal to K:
// Input: nums = [5,7,7,8,8,10], k = 7
// Output: 1 (first position of 7)
//
// Last Equal to K:
// Input: nums = [5,7,7,8,8,10], k = 7
// Output: 2 (last position of 7)
//
// First Greater Than K:
// Input: nums = [5,7,7,8,8,10], k = 7
// Output: 3 (first element greater than 7 is at index 3)
//
// Last Less Than K:
// Input: nums = [5,7,7,8,8,10], k = 8
// Output: 2 (last element less than 8 is at index 2)
