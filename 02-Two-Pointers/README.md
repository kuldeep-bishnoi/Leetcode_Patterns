# Two Pointers Pattern

## What are Two Pointers?

Imagine you're playing a game where you have two fingers, and you need to find something in a line of items. You can place one finger at the start and one at the end, then move them towards each other or in the same direction to find what you're looking for. This is exactly how the two pointers technique works in programming!

## Real-Life Examples

1. **Finding a Book**: When looking for a book in a library shelf, you might start with one finger at the beginning and one at the end, then move them towards each other until you find the right book.

2. **Measuring Height**: When measuring someone's height, you might use two fingers to mark the top and bottom points, then bring them together to find the exact height.

3. **Playing Catch**: When playing catch, you need to track both the thrower and the catcher's positions to know where the ball should go.

## When Do We Use Two Pointers?

Use this technique when you need to:
- Find pairs of items in a sorted list
- Compare items from different positions
- Move through a list in two different directions
- Find the middle of something
- Check if something is a palindrome

## Types of Two Pointers

1. **Opposite Direction Pointers**: Pointers start at opposite ends and move towards each other
   - Example: "Find two numbers in a sorted list that add up to 10"
   
   ![Opposite Direction Visualization]
   ```
   Array: [1, 3, 5, 7, 9, 11]
   
   Start: [1, 3, 5, 7, 9, 11]
           ↑           ↑
           left       right
   
   Step 1: 1 + 11 = 12 (too big, move right pointer left)
   [1, 3, 5, 7, 9, 11]
    ↑        ↑
   
   Step 2: 1 + 9 = 10 (found it!)
   ```

2. **Same Direction Pointers**: Pointers move in the same direction at different speeds
   - Example: "Remove duplicates from a sorted list"
   
   ![Same Direction Visualization]
   ```
   Array: [1, 1, 2, 2, 3, 4, 4]
   
   Start: [1, 1, 2, 2, 3, 4, 4]
           ↑
           slow
           ↑
           fast
   
   Step 1: [1, 1, 2, 2, 3, 4, 4]
            ↑
            slow
             ↑
            fast
   
   Step 2: [1, 2, 2, 2, 3, 4, 4]
             ↑
             slow
              ↑
             fast
   ```

## How Does It Work?

1. **Opposite Direction**:
   - Place one pointer at the start and one at the end
   - Move pointers towards each other based on the comparison
   - Stop when pointers meet or find what you're looking for

2. **Same Direction**:
   - Place both pointers at the start
   - Move the fast pointer ahead to find new elements
   - Use the slow pointer to build the result

## Simple Code Examples

### Opposite Direction Example
Finding two numbers that add up to a target:

```go
func findPairWithSum(numbers []int, target int) []int {
    left := 0
    right := len(numbers) - 1
    
    // Keep checking until pointers meet
    for left < right {
        currentSum := numbers[left] + numbers[right]
        
        if currentSum == target {
            // Found a pair!
            return []int{numbers[left], numbers[right]}
        } else if currentSum < target {
            // Sum is too small, move left pointer right
            left++
        } else {
            // Sum is too big, move right pointer left
            right--
        }
    }
    
    // No pair found
    return nil
}
```

### Same Direction Example
Removing duplicates from a sorted array:

```go
func removeDuplicates(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }
    
    // Slow pointer keeps track of where to place next unique number
    slow := 1
    
    // Fast pointer finds new unique numbers
    for fast := 1; fast < len(numbers); fast++ {
        // If we find a new unique number
        if numbers[fast] != numbers[fast-1] {
            // Place it at the slow pointer position
            numbers[slow] = numbers[fast]
            // Move slow pointer forward
            slow++
        }
    }
    
    return slow
}
```

## Common Mistakes to Avoid

1. **Forgetting to check pointer bounds**: Always make sure your pointers don't go out of the array
2. **Incorrect pointer movement**: Be careful about which pointer to move and when
3. **Not handling empty or single-element cases**: What if the array is empty or has only one element?

## Fun Practice Problems

1. **Rainbow Sorter**: You have a line of colored balls (red, white, blue). Sort them so all red balls are on the left, white in the middle, and blue on the right.

2. **Container Water**: You have a line of walls of different heights. Find which two walls, when used as container sides, can hold the most water.

3. **Palindrome Checker**: Check if a word reads the same forwards and backwards (like "racecar").

4. **Height Matcher**: Given a list of student heights, find two students whose heights add up to exactly 6 feet.

5. **Color Sorter**: Sort a line of red and blue marbles so all red marbles are on the left and blue on the right.

## LeetCode Problems Using Two Pointers

Here are some popular LeetCode problems that can be solved using the two pointers technique:

### Easy Problems

1. **[#125 Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)** - Check if a string is a palindrome.
   - **Approach**: Use opposite direction pointers to compare characters from both ends.

2. **[#26 Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)** - Remove duplicates in-place from a sorted array.
   - **Approach**: Use same direction pointers to track unique elements.

3. **[#27 Remove Element](https://leetcode.com/problems/remove-element/)** - Remove all instances of a value in-place.
   - **Approach**: Use same direction pointers to track non-target elements.

4. **[#283 Move Zeroes](https://leetcode.com/problems/move-zeroes/)** - Move all zeros to the end while maintaining relative order.
   - **Approach**: Use same direction pointers to track non-zero elements.

### Medium Problems

1. **[#15 3Sum](https://leetcode.com/problems/3sum/)** - Find all unique triplets that sum to zero.
   - **Approach**: Fix one number and use two pointers to find pairs that sum to its negative.

2. **[#11 Container With Most Water](https://leetcode.com/problems/container-with-most-water/)** - Find two lines that together with x-axis form a container that holds the most water.
   - **Approach**: Use opposite direction pointers, moving the shorter line inward.

3. **[#42 Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)** - Calculate how much water can be trapped between bars.
   - **Approach**: Use opposite direction pointers with left and right max heights.

4. **[#75 Sort Colors](https://leetcode.com/problems/sort-colors/)** - Sort an array of 0s, 1s, and 2s in-place.
   - **Approach**: Use three pointers to partition the array into three sections.

5. **[#167 Two Sum II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)** - Find two numbers in a sorted array that sum to target.
   - **Approach**: Use opposite direction pointers to find the pair.

6. **[#344 Reverse String](https://leetcode.com/problems/reverse-string/)** - Reverse a string in-place.
   - **Approach**: Use opposite direction pointers to swap characters.

7. **[#977 Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/)** - Return array of squares in sorted order.
   - **Approach**: Use opposite direction pointers to handle negative numbers.

### Hard Problems

1. **[#42 Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)** - Calculate trapped water between bars.
   - **Approach**: Use opposite direction pointers with dynamic programming.

2. **[#76 Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/)** - Find smallest substring containing all characters from target string.
   - **Approach**: Use same direction pointers with a sliding window.

3. **[#30 Substring with Concatenation of All Words](https://leetcode.com/problems/substring-with-concatenation-of-all-words/)** - Find all starting indices of concatenated words.
   - **Approach**: Use sliding window with two pointers for word matching.

4. **[#76 Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/)** - Find smallest substring containing all characters.
   - **Approach**: Use same direction pointers with character frequency tracking.

### Tips for Solving LeetCode Two Pointer Problems

1. **Identify Pointer Type**: Determine if you need opposite or same direction pointers
2. **Track State**: Decide what information you need to track with each pointer
3. **Pointer Movement**: Implement the logic for when and how to move each pointer
4. **Optimization**: Look for ways to reduce unnecessary comparisons
5. **Edge Cases**: Consider empty arrays, single elements, or when pointers meet

## Why Learn This Pattern?

The two pointers pattern is super useful because:
1. It's often faster than using a single pointer that has to scan the entire array multiple times
2. It's used in many real-world applications like text processing and data analysis
3. It's a favorite in coding interviews and competitions
4. It helps solve problems that seem complex in a simple and efficient way