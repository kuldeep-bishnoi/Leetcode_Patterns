# Cyclic Sort Pattern

## What is Cyclic Sort?

Imagine you have a bunch of numbered cards (1, 2, 3, etc.) that are all mixed up. Cyclic Sort is like playing a matching game where you try to put each card in its correct position. It's called "cyclic" because you keep going in circles until everything is in the right place!

## Real-Life Examples

1. **Library Books**: When books are out of order on the shelf, you put each book in its correct numbered spot.
   - Book 1 goes in spot 1
   - Book 2 goes in spot 2
   - And so on...

2. **Classroom Seats**: When students need to sit in numbered seats, they find their number and sit there.
   - Student 1 sits in seat 1
   - Student 2 sits in seat 2
   - And so on...

3. **Toy Box Organization**: When organizing toys by number, you put each toy in its numbered box.
   - Toy 1 goes in box 1
   - Toy 2 goes in box 2
   - And so on...

## When Do We Use Cyclic Sort?

Use this technique when:
- You have numbers from 1 to N (where N is the length of the array)
- You need to find missing or duplicate numbers
- You need to sort numbers in a specific range
- You want to solve problems with numbers in a fixed range

## How Does It Work?

1. **Step 1**: Start with the first number
2. **Step 2**: If the number is not in its correct position (number ≠ position + 1):
   - Swap it with the number that should be in its position
   - Keep swapping until the current number is in the right place
3. **Step 3**: Move to the next position
4. **Step 4**: Repeat until all numbers are in their correct positions

## Simple Code Example

```go
func cyclicSort(nums []int) {
    i := 0
    for i < len(nums) {
        // If number is not in its correct position
        if nums[i] != i+1 {
            // Swap with the number that should be in this position
            correct := nums[i] - 1
            nums[i], nums[correct] = nums[correct], nums[i]
        } else {
            // Move to next position
            i++
        }
    }
}
```

## Common Mistakes to Avoid

1. **Off-by-one errors**: Remember that position 0 should have number 1
2. **Infinite loops**: Make sure you're moving to the next position when needed
3. **Array bounds**: Check that you're not accessing positions beyond the array

## Fun Practice Problems

1. **Missing Number Game**: Find which number is missing from 1 to N
2. **Duplicate Finder**: Find which number appears twice in 1 to N
3. **First Missing Positive**: Find the first positive number that's missing
4. **Find All Duplicates**: Find all numbers that appear twice
5. **Find All Missing**: Find all numbers that are missing from 1 to N

## LeetCode Problems Using Cyclic Sort

Here are some popular LeetCode problems that can be solved using Cyclic Sort:

### Easy Problems

1. **[#268 Missing Number](https://leetcode.com/problems/missing-number/)** - Find the missing number in an array of numbers from 0 to n.
   - **Approach**: Use cyclic sort and find the first number not in its correct position.

2. **[#448 Find All Numbers Disappeared in an Array](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/)** - Find all numbers missing from 1 to n.
   - **Approach**: Use cyclic sort and check which positions don't have correct numbers.

3. **[#645 Set Mismatch](https://leetcode.com/problems/set-mismatch/)** - Find the duplicate and missing number.
   - **Approach**: Use cyclic sort to find both the duplicate and missing numbers.

### Medium Problems

1. **[#41 First Missing Positive](https://leetcode.com/problems/first-missing-positive/)** - Find the first missing positive integer.
   - **Approach**: Use cyclic sort ignoring negative numbers and numbers larger than array length.

2. **[#442 Find All Duplicates in an Array](https://leetcode.com/problems/find-all-duplicates-in-an-array/)** - Find all numbers that appear twice.
   - **Approach**: Use cyclic sort and find numbers that don't match their positions.

3. **[#287 Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)** - Find the duplicate number in an array of n+1 integers.
   - **Approach**: Use cyclic sort or Floyd's cycle detection algorithm.

4. **[#645 Set Mismatch](https://leetcode.com/problems/set-mismatch/)** - Find the duplicate and missing number.
   - **Approach**: Use cyclic sort to find both numbers.

### Hard Problems

1. **[#41 First Missing Positive](https://leetcode.com/problems/first-missing-positive/)** - Find the first missing positive integer.
   - **Approach**: Use cyclic sort with special handling for out-of-range numbers.

2. **[#765 Couples Holding Hands](https://leetcode.com/problems/couples-holding-hands/)** - Arrange couples to sit together.
   - **Approach**: Use cyclic sort to arrange couples in correct positions.

3. **[#41 First Missing Positive](https://leetcode.com/problems/first-missing-positive/)** - Find the first missing positive integer.
   - **Approach**: Use cyclic sort with special handling for out-of-range numbers.

### Tips for Solving LeetCode Cyclic Sort Problems

1. **Range Check**: Always verify if numbers are within valid range
2. **Position Check**: Make sure to check if number matches its position
3. **Swap Logic**: Be careful with swap operations to avoid infinite loops
4. **Edge Cases**: Handle special cases like negative numbers or numbers larger than array length
5. **Optimization**: Look for ways to reduce number of swaps

## Why Learn This Pattern?

The Cyclic Sort pattern is super useful because:
1. It's perfect for sorting numbers in a specific range (1 to N)
2. It helps find missing or duplicate numbers efficiently
3. It's a favorite in coding interviews
4. It teaches important concepts about array manipulation and sorting
5. It's space-efficient (uses O(1) extra space)