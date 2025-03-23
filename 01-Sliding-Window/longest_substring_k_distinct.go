package slidingwindow

import "math"

// LongestSubstringKDistinct finds the length of the longest substring with at most 'k' distinct characters
// Time Complexity: O(n) where n is the length of the string
// Space Complexity: O(k) for storing at most k characters in the frequency map
func LongestSubstringKDistinct(str string, k int) int {
	if len(str) == 0 || k == 0 {
		return 0
	}

	windowStart := 0
	maxLength := 0
	charFrequency := make(map[byte]int)

	// Try to extend the range [windowStart, windowEnd]
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		// Add the current character to the frequency map
		charFrequency[rightChar]++

		// Shrink the sliding window, until we have at most 'k' distinct characters in the frequency map
		for len(charFrequency) > k {
			leftChar := str[windowStart]
			// Decrement the frequency of the character going out of the window
			charFrequency[leftChar]--
			if charFrequency[leftChar] == 0 {
				delete(charFrequency, leftChar) // Remove the character if its frequency becomes zero
			}
			windowStart++ // Shrink the window
		}

		// Remember the maximum length so far
		maxLength = int(math.Max(float64(maxLength), float64(windowEnd-windowStart+1)))
	}

	return maxLength
}

// Example cases:
// Input: "araaci", k=2
// Output: 4
// Explanation: The longest substring with at most 2 distinct characters is "araa"

// Input: "araaci", k=1
// Output: 2
// Explanation: The longest substring with at most 1 distinct character is "aa"

// Input: "cbbebi", k=3
// Output: 5
// Explanation: The longest substrings with at most 3 distinct characters are "cbbeb" and "bbebi"
