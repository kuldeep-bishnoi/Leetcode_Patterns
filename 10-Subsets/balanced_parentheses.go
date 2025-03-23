package subsets

// GenerateParenthesis generates all combinations of well-formed parentheses
// Time Complexity: O(4^n / sqrt(n)) - Catalan number
// Space Complexity: O(4^n / sqrt(n)) for storing all valid combinations
func GenerateParenthesis(n int) []string {
	result := []string{}
	backtrack("", 0, 0, n, &result)
	return result
}

// backtrack recursively generates valid parentheses combinations
// open: count of open parentheses used so far
// close: count of close parentheses used so far
// max: maximum number of pairs to use
func backtrack(current string, open, close, max int, result *[]string) {
	// Base case: if the combination is complete
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	// We can add an open parenthesis if we haven't used all of them
	if open < max {
		backtrack(current+"(", open+1, close, max, result)
	}

	// We can add a close parenthesis if it doesn't create an imbalance
	if close < open {
		backtrack(current+")", open, close+1, max, result)
	}
}

// GenerateParenthesisBFS generates all combinations of well-formed parentheses using BFS
// Time Complexity: O(4^n / sqrt(n)) - Catalan number
// Space Complexity: O(4^n / sqrt(n)) for storing all valid combinations
func GenerateParenthesisBFS(n int) []string {
	if n == 0 {
		return []string{""}
	}

	// Queue to store partial results during BFS
	type ParenState struct {
		str   string
		open  int
		close int
	}

	queue := []ParenState{{"", 0, 0}}
	result := []string{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// If we've used all parentheses, add to result
		if len(curr.str) == 2*n {
			result = append(result, curr.str)
			continue
		}

		// Try adding an open parenthesis
		if curr.open < n {
			queue = append(queue, ParenState{
				str:   curr.str + "(",
				open:  curr.open + 1,
				close: curr.close,
			})
		}

		// Try adding a close parenthesis if valid
		if curr.close < curr.open {
			queue = append(queue, ParenState{
				str:   curr.str + ")",
				open:  curr.open,
				close: curr.close + 1,
			})
		}
	}

	return result
}

// Example usage:
//
// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]
// Explanation: All valid combinations of 3 pairs of parentheses.
//
// Input: n = 1
// Output: ["()"]
// Explanation: Only one valid combination of 1 pair of parentheses.
