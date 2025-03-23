package binarysearch

// SearchInRotatedArray finds the index of a target value in a rotated sorted array
// Time Complexity: O(log n)
// Space Complexity: O(1)
func SearchInRotatedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)/2

		// If we found the target at mid, return mid
		if nums[mid] == target {
			return mid
		}

		// Check if left half is sorted
		if nums[start] <= nums[mid] {
			// Target is in the left sorted half
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				// Target is in the right half
				start = mid + 1
			}
		} else { // Right half is sorted
			// Target is in the right sorted half
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				// Target is in the left half
				end = mid - 1
			}
		}
	}

	return -1 // Target not found
}

// SearchInRotatedArrayWithDuplicates finds the index of a target in a rotated sorted array with duplicates
// Time Complexity: O(log n) in average case, O(n) in worst case with many duplicates
// Space Complexity: O(1)
func SearchInRotatedArrayWithDuplicates(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)/2

		// If we found the target at mid, return true
		if nums[mid] == target {
			return true
		}

		// Handle duplicates - skip duplicates from both ends
		if nums[start] == nums[mid] && nums[mid] == nums[end] {
			start++
			end--
			continue
		}

		// Check if left half is sorted
		if nums[start] <= nums[mid] {
			// Target is in the left sorted half
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				// Target is in the right half
				start = mid + 1
			}
		} else { // Right half is sorted
			// Target is in the right sorted half
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				// Target is in the left half
				end = mid - 1
			}
		}
	}

	return false // Target not found
}

// FindMinInRotatedArray finds the minimum element in a rotated sorted array
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindMinInRotatedArray(nums []int) int {
	if len(nums) == 0 {
		return -1 // Invalid input
	}

	start, end := 0, len(nums)-1

	// Not rotated or only one element
	if nums[start] < nums[end] || len(nums) == 1 {
		return nums[start]
	}

	for start < end {
		mid := start + (end-start)/2

		// If mid is greater than end, the minimum is in the right half
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			// Minimum is in the left half or at mid
			end = mid
		}
	}

	// At this point, start == end, pointing to the minimum element
	return nums[start]
}

// FindMinInRotatedArrayWithDuplicates finds the minimum element in a rotated sorted array with duplicates
// Time Complexity: O(log n) in average case, O(n) in worst case with many duplicates
// Space Complexity: O(1)
func FindMinInRotatedArrayWithDuplicates(nums []int) int {
	if len(nums) == 0 {
		return -1 // Invalid input
	}

	start, end := 0, len(nums)-1

	// Not rotated or only one element
	if nums[start] < nums[end] || len(nums) == 1 {
		return nums[start]
	}

	for start < end {
		mid := start + (end-start)/2

		// Handle duplicates - if mid equals end, reduce the search space by 1
		if nums[mid] == nums[end] {
			end--
			continue
		}

		// If mid is greater than end, the minimum is in the right half
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			// Minimum is in the left half or at mid
			end = mid
		}
	}

	// At this point, start == end, pointing to the minimum element
	return nums[start]
}

// Example usage:
//
// Search in rotated array:
// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4 (index of 0 in the array)
//
// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1 (target not found)
//
// Search with duplicates:
// Input: nums = [2,5,6,0,0,1,2], target = 0
// Output: true (0 exists in the array)
//
// Find minimum:
// Input: nums = [3,4,5,1,2]
// Output: 1 (minimum element in the array)
//
// Find minimum with duplicates:
// Input: nums = [2,2,2,0,1]
// Output: 0 (minimum element in the array)
