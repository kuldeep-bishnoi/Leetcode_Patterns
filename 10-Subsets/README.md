# Subsets Pattern

## What are Subsets?

Imagine you have a box of toys, and you want to find all possible ways to play with them. The Subsets pattern helps you find all possible combinations of your toys. It's like making different teams with your friends - you can have teams of 1, 2, 3, or more friends!

## Real-Life Examples

1. **Toy Combinations**: When you want to play with different toys.
   - You can play with just one toy
   - You can play with two toys together
   - You can play with three toys together
   - And so on...

2. **Ice Cream Toppings**: When choosing toppings for your ice cream.
   - You can have no toppings
   - You can have one topping
   - You can have two toppings
   - And so on...

3. **School Teams**: When forming teams for a game.
   - You can have teams of one student
   - You can have teams of two students
   - You can have teams of three students
   - And so on...

## When Do We Use Subsets?

Use this technique when:
- You need to find all possible combinations
- You need to generate all possible subsets
- You need to find all possible permutations
- You need to solve problems with multiple choices
- You need to find all possible ways to select items

## How Does It Work?

1. **Step 1**: Start with an empty subset
2. **Step 2**: For each number/item:
   - Include it in the current subset
   - Recursively process remaining items
   - Remove it from the current subset
   - Recursively process remaining items
3. **Step 3**: Keep track of all subsets

## Simple Code Example

```go
func subsets(nums []int) [][]int {
    result := [][]int{}
    
    // Helper function to generate subsets
    var generate func([]int, int)
    generate = func(current []int, index int) {
        // Add current subset to result
        subset := make([]int, len(current))
        copy(subset, current)
        result = append(result, subset)
        
        // Try adding each remaining number
        for i := index; i < len(nums); i++ {
            // Include current number
            current = append(current, nums[i])
            
            // Generate subsets with current number
            generate(current, i+1)
            
            // Remove current number (backtrack)
            current = current[:len(current)-1]
        }
    }
    
    // Start with empty subset
    generate([]int{}, 0)
    return result
}
```

## Common Mistakes to Avoid

1. **Duplicate Subsets**: Make sure to avoid duplicate combinations
2. **Backtracking**: Remember to remove items after processing
3. **Order**: Consider if order matters in your subsets
4. **Empty Set**: Don't forget to include the empty subset

## Fun Practice Problems

1. **Toy Combinations**: Find all ways to play with different toys
2. **Ice Cream Toppings**: Find all possible topping combinations
3. **Team Formations**: Find all possible team combinations
4. **Letter Combinations**: Find all possible letter combinations
5. **Number Combinations**: Find all possible number combinations

## LeetCode Problems Using Subsets

Here are some popular LeetCode problems that can be solved using Subsets:

### Easy Problems

1. **[#78 Subsets](https://leetcode.com/problems/subsets/)** - Generate all possible subsets.
   - **Approach**: Use backtracking to generate all combinations.

2. **[#784 Letter Case Permutation](https://leetcode.com/problems/letter-case-permutation/)** - Generate all letter case permutations.
   - **Approach**: Use backtracking to try different cases.

### Medium Problems

1. **[#90 Subsets II](https://leetcode.com/problems/subsets-ii/)** - Generate subsets with duplicates.
   - **Approach**: Use backtracking with duplicate handling.

2. **[#46 Permutations](https://leetcode.com/problems/permutations/)** - Generate all permutations.
   - **Approach**: Use backtracking to try all positions.

3. **[#47 Permutations II](https://leetcode.com/problems/permutations-ii/)** - Generate permutations with duplicates.
   - **Approach**: Use backtracking with duplicate handling.

4. **[#39 Combination Sum](https://leetcode.com/problems/combination-sum/)** - Find combinations that sum to target.
   - **Approach**: Use backtracking with sum tracking.

### Hard Problems

1. **[#51 N-Queens](https://leetcode.com/problems/n-queens/)** - Place n queens on n×n board.
   - **Approach**: Use backtracking to try all positions.

2. **[#37 Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)** - Solve a Sudoku puzzle.
   - **Approach**: Use backtracking to try all numbers.

### Tips for Solving LeetCode Subsets Problems

1. **Backtracking**: Use backtracking to try all possibilities
2. **State Tracking**: Keep track of current subset
3. **Duplicate Handling**: Handle duplicates if present
4. **Order Matters**: Consider if order is important
5. **Constraints**: Check for any size or sum constraints

## Why Learn This Pattern?

The Subsets pattern is super useful because:
1. It helps find all possible combinations
2. It's used in many real-world problems
3. It's a favorite in coding interviews
4. It teaches important concepts about recursion and backtracking
5. It helps solve complex combinatorial problems

Once you master this pattern, you'll be able to solve many combination and permutation problems efficiently and impress your friends with your coding skills! 