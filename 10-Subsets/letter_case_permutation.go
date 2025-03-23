package subsets

import (
	"unicode"
)

// LetterCasePermutation generates all permutations of a string by changing case of letters
// Time Complexity: O(2^L) where L is the number of letters in the string
// Space Complexity: O(2^L) for storing all permutations
func LetterCasePermutation(s string) []string {
	result := []string{""}

	// Process each character in the string
	for _, char := range s {
		n := len(result)

		// If the character is a digit, append it to all existing permutations
		if unicode.IsDigit(char) {
			for i := 0; i < n; i++ {
				result[i] += string(char)
			}
		} else {
			// For letters, create two permutations - one with lowercase and one with uppercase
			for i := 0; i < n; i++ {
				// Add lowercase version
				result = append(result, result[i]+string(unicode.ToLower(char)))
				// Add uppercase version
				result[i] += string(unicode.ToUpper(char))
			}
		}
	}

	return result
}

// LetterCasePermutationRecursive generates all permutations of a string by changing case of letters
// using a recursive approach
// Time Complexity: O(2^L) where L is the number of letters in the string
// Space Complexity: O(2^L) for storing all permutations
func LetterCasePermutationRecursive(s string) []string {
	result := []string{}
	generateLetterCasePermutations([]rune(s), 0, "", &result)
	return result
}

// generateLetterCasePermutations recursively generates letter case permutations
func generateLetterCasePermutations(chars []rune, index int, current string, result *[]string) {
	// If we've processed all characters, add the current permutation to the result
	if index == len(chars) {
		*result = append(*result, current)
		return
	}

	// Get the current character
	char := chars[index]

	// If the character is a digit, we only have one option
	if unicode.IsDigit(char) {
		generateLetterCasePermutations(chars, index+1, current+string(char), result)
	} else {
		// For letters, explore both lowercase and uppercase
		generateLetterCasePermutations(chars, index+1, current+string(unicode.ToLower(char)), result)
		generateLetterCasePermutations(chars, index+1, current+string(unicode.ToUpper(char)), result)
	}
}

// LetterCasePermutationBFS generates all permutations of a string by changing case of letters
// using a BFS approach with queues
// Time Complexity: O(2^L) where L is the number of letters in the string
// Space Complexity: O(2^L) for storing all permutations
func LetterCasePermutationBFS(s string) []string {
	result := []string{""}

	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		size := len(result)

		for j := 0; j < size; j++ {
			current := result[j]

			// If the character is a digit, we have only one choice
			if unicode.IsDigit(char) {
				result[j] = current + string(char)
			} else {
				// Replace the current entry with lowercase version
				result[j] = current + string(unicode.ToLower(char))

				// Add a new entry with uppercase version
				result = append(result, current+string(unicode.ToUpper(char)))
			}
		}
	}

	return result
}

// Example usage:
//
// Input: s = "a1b2"
// Output: ["a1b2", "a1B2", "A1b2", "A1B2"]
// Explanation: We can permute the case of 'a' and 'b' to get 4 different permutations.
//
// Input: s = "3z4"
// Output: ["3z4", "3Z4"]
// Explanation: We can only permute the case of 'z' since '3' and '4' are digits.
