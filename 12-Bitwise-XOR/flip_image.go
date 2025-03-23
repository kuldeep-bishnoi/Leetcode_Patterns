package bitwisexor

// FlipAndInvertImage flips an image horizontally, then inverts it
// Time Complexity: O(n*m) where n is the number of rows and m is the number of columns
// Space Complexity: O(1) if we don't count the output as extra space
func FlipAndInvertImage(image [][]int) [][]int {
	n := len(image)
	
	for i := 0; i < n; i++ {
		// Two-pointer approach to flip the row horizontally
		left, right := 0, n-1
		for left <= right {
			// XOR with 1 to invert (0->1, 1->0) and swap
			image[i][left], image[i][right] = image[i][right]^1, image[i][left]^1
			left++
			right--
		}
	}
	
	return image
}

// FlipAndInvertImageAlternative does the same operation but with a different approach
// Time Complexity: O(n*m)
// Space Complexity: O(1) if we don't count the output as extra space
func FlipAndInvertImageAlternative(image [][]int) [][]int {
	n := len(image)
	
	for i := 0; i < n; i++ {
		for j := 0; j < (n+1)/2; j++ {
			// Calculate the opposite column
			k := n - 1 - j
			// Swap and invert in one step
			image[i][j], image[i][k] = image[i][k]^1, image[i][j]^1
		}
	}
	
	return image
}

// FlipImage flips an image horizontally without inverting
// Time Complexity: O(n*m)
// Space Complexity: O(1) if we don't count the output as extra space
func FlipImage(image [][]int) [][]int {
	n := len(image)
	
	for i := 0; i < n; i++ {
		// Two-pointer approach to flip the row horizontally
		left, right := 0, n-1
		for left < right {
			image[i][left], image[i][right] = image[i][right], image[i][left]
			left++
			right--
		}
	}
	
	return image
}

// InvertImage inverts an image without flipping
// Time Complexity: O(n*m)
// Space Complexity: O(1) if we don't count the output as extra space
func InvertImage(image [][]int) [][]int {
	n := len(image)
	m := len(image[0])
	
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// XOR with 1 to invert (0->1, 1->0)
			image[i][j] ^= 1
		}
	}
	
	return image
}

// Example usage:
//
// Flip and Invert Image:
// Input: image = [[1,1,0],[1,0,1],[0,0,0]]
// Output: [[1,0,0],[0,1,0],[1,1,1]]
// Explanation:
// Step 1: Flip horizontally: [[0,1,1],[1,0,1],[0,0,0]]
// Step 2: Invert: [[1,0,0],[0,1,0],[1,1,1]]
//
// Input: image = [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
// Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]] 