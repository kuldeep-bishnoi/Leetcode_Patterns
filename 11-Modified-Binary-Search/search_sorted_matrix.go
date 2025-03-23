package binarysearch

// SearchMatrix searches for a target in a matrix where:
// - Each row is sorted in ascending order from left to right
// - Each column is sorted in ascending order from top to bottom
// Returns true if the target is found, false otherwise
// Time Complexity: O(m + n) where m is the number of rows and n is the number of columns
// Space Complexity: O(1)
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// Start from the top-right corner
	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true // Target found
		} else if matrix[row][col] > target {
			// Current value is greater than target, move left
			col--
		} else {
			// Current value is less than target, move down
			row++
		}
	}

	return false // Target not found
}

// SearchMatrixPosition searches for a target in a sorted matrix and returns its position
// Returns [row, col] if found, [-1, -1] if not found
// Time Complexity: O(m + n)
// Space Complexity: O(1)
func SearchMatrixPosition(matrix [][]int, target int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{-1, -1}
	}

	// Start from the top-right corner
	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return []int{row, col} // Target found
		} else if matrix[row][col] > target {
			// Current value is greater than target, move left
			col--
		} else {
			// Current value is less than target, move down
			row++
		}
	}

	return []int{-1, -1} // Target not found
}

// SearchStrictlySortedMatrix searches for a target in a matrix where:
// - Each row is sorted in ascending order
// - First integer of each row is greater than the last integer of the previous row
// Returns true if the target is found, false otherwise
// Time Complexity: O(log(m*n))
// Space Complexity: O(1)
func SearchStrictlySortedMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	start, end := 0, rows*cols-1

	for start <= end {
		mid := start + (end-start)/2

		// Convert the 1D index to 2D coordinates
		midRow, midCol := mid/cols, mid%cols

		if matrix[midRow][midCol] == target {
			return true // Target found
		} else if matrix[midRow][midCol] < target {
			start = mid + 1 // Look in the right half
		} else {
			end = mid - 1 // Look in the left half
		}
	}

	return false // Target not found
}

// SearchStrictlySortedMatrixPosition searches for a target in a strictly sorted matrix
// Returns [row, col] if found, [-1, -1] if not found
// Time Complexity: O(log(m*n))
// Space Complexity: O(1)
func SearchStrictlySortedMatrixPosition(matrix [][]int, target int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{-1, -1}
	}

	rows, cols := len(matrix), len(matrix[0])
	start, end := 0, rows*cols-1

	for start <= end {
		mid := start + (end-start)/2

		// Convert the 1D index to 2D coordinates
		midRow, midCol := mid/cols, mid%cols

		if matrix[midRow][midCol] == target {
			return []int{midRow, midCol} // Target found
		} else if matrix[midRow][midCol] < target {
			start = mid + 1 // Look in the right half
		} else {
			end = mid - 1 // Look in the left half
		}
	}

	return []int{-1, -1} // Target not found
}

// Example usage:
//
// Search in Row-wise and Column-wise Sorted Matrix:
// Input:
// matrix = [
//   [1,  4,  7, 11, 15],
//   [2,  5,  8, 12, 19],
//   [3,  6,  9, 16, 22],
//   [10, 13, 14, 17, 24]
// ], target = 5
// Output: true (5 is found in the matrix)
//
// Search Position in Matrix:
// Input: same matrix as above, target = 5
// Output: [1, 1] (5 is found at row 1, column 1)
//
// Search in Strictly Sorted Matrix:
// Input:
// matrix = [
//   [1,  3,  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ], target = 3
// Output: true (3 is found in the matrix)
//
// Search Position in Strictly Sorted Matrix:
// Input: same strictly sorted matrix as above, target = 3
// Output: [0, 1] (3 is found at row 0, column 1)
