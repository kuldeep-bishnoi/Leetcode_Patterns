# Sliding Window Pattern

## What is a Sliding Window?

Imagine you're looking through a small window frame at a long line of toys. You can only see a few toys at a time through this window. Now, you start moving the window slowly from left to right. As you move it, new toys come into view from the right side, while toys on the left side disappear from view. This is exactly how the sliding window technique works in programming!

## Real-Life Examples

1. **School Bus Window**: When you sit by a school bus window, you can only see a portion of the scenery. As the bus moves, new scenes come into view while old ones disappear.

2. **Reading with a Ruler**: When you place a ruler under a line of text to help you read, you focus on just a few words at a time, then move the ruler to read the next few words.

3. **Playing Card Hand**: In a card game, you might discard one card and pick up a new one, maintaining a hand of the same size but with different cards.

## When Do We Use Sliding Window?

Use this technique when you need to:
- Look at a continuous chunk of items (like numbers in a list or characters in a word)
- Find the best chunk that follows certain rules
- Process items one after another in order

## Types of Sliding Windows

1. **Fixed Size Window**: The window always shows the same number of items
   - Example: "Find the highest sum of any 3 consecutive numbers in a list"
   
   ![Fixed Window Visualization]
   ```
   Array: [1, 3, 2, 6, 8, 1, 4, 2]
   
   Window 1: [1, 3, 2] = Sum: 6
   Window 2:    [3, 2, 6] = Sum: 11
   Window 3:       [2, 6, 8] = Sum: 16 (highest!)
   Window 4:          [6, 8, 1] = Sum: 15
   Window 5:             [8, 1, 4] = Sum: 13
   Window 6:                [1, 4, 2] = Sum: 7
   ```

2. **Variable Size Window**: The window can grow or shrink as needed
   - Example: "Find the smallest group of consecutive numbers that add up to at least 10"
   
   ![Variable Window Visualization]
   ```
   Array: [1, 3, 2, 6, 8, 1, 4, 2]
   
   Window starts with [1]
   Grow to [1, 3] = Sum: 4 (not enough)
   Grow to [1, 3, 2] = Sum: 6 (not enough)
   Grow to [1, 3, 2, 6] = Sum: 12 (enough!)
   Try to shrink: [3, 2, 6] = Sum: 11 (still enough)
   Try to shrink: [2, 6] = Sum: 8 (not enough)
   So smallest window is [3, 2, 6] with length 3
   ```

## How Does It Work?

1. **Step 1**: Place your window at the beginning of your list
2. **Step 2**: Look at the items in your window and do whatever calculation you need
3. **Step 3**: Slide the window one step to the right:
   - Add the new item that enters from the right
   - Remove the item that leaves from the left
4. **Step 4**: Repeat steps 2-3 until you reach the end of the list

## Simple Code Examples

### Fixed Size Window Example
Finding the largest sum of 3 consecutive numbers:

```go
func largestSumOfThreeNumbers(numbers []int) int {
    // If we don't have at least 3 numbers, we can't make a window of size 3
    if len(numbers) < 3 {
        return 0
    }
    
    // First, calculate the sum of the first 3 numbers
    currentSum := numbers[0] + numbers[1] + numbers[2]
    
    // This is our largest sum so far
    largestSum := currentSum
    
    // Now slide the window from position 1 to the end
    for i := 3; i < len(numbers); i++ {
        // Add the new number entering the window
        // Remove the number leaving the window
        currentSum = currentSum + numbers[i] - numbers[i-3]
        
        // Update the largest sum if we found a better one
        if currentSum > largestSum {
            largestSum = currentSum
        }
    }
    
    return largestSum
}
```

### Variable Size Window Example
Finding the smallest subarray with a sum of at least 8:

```go
func smallestSubarrayWithSumAtLeast(numbers []int, targetSum int) int {
    currentSum := 0
    windowStart := 0
    smallestLength := len(numbers) + 1 // Start with an impossible length
    
    for windowEnd := 0; windowEnd < len(numbers); windowEnd++ {
        // Add the new number to our sum
        currentSum += numbers[windowEnd]
        
        // Shrink the window as small as possible while keeping sum >= targetSum
        for currentSum >= targetSum {
            // Current window size
            windowLength := windowEnd - windowStart + 1
            
            // Update the smallest length if we found a smaller valid window
            if windowLength < smallestLength {
                smallestLength = windowLength
            }
            
            // Remove the leftmost number and shrink the window
            currentSum -= numbers[windowStart]
            windowStart++
        }
    }
    
    // If we couldn't find any subarray that sums to at least targetSum
    if smallestLength > len(numbers) {
        return 0
    }
    
    return smallestLength
}
```

## Common Mistakes to Avoid

1. **Forgetting to update the window**: Always make sure you add the new element and remove the old one when sliding.
2. **Off-by-one errors**: Be careful with your start and end positions.
3. **Not handling edge cases**: What if the list is empty? What if it's too small for your window?

## Fun Practice Problems

1. **Lemonade Stand**: You have sales data for a lemonade stand for 30 days. Find the 7 consecutive days with the highest total sales.

2. **Word Finder**: In the word "MISSISSIPPI", find the longest substring without repeating letters.

3. **Temperature Watch**: Given a list of daily temperatures, find the shortest sequence of days where the average temperature is above 75°F.

4. **Candy Collection**: You walk past houses and collect different types of candy. At any time, you can only carry 2 types of candy. What's the maximum number of candies you can collect?

5. **Book Reading**: You're reading a book and can only focus on 5 pages at a time. If each page has a "fun rating," which 5 consecutive pages are the most fun to read?

## LeetCode Problems Using Sliding Window

Here are some popular LeetCode problems that can be solved using the sliding window technique:

### Easy Problems

1. **[#53 Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)** - Find the subarray with the largest sum.
   - **Approach**: Use a sliding window that expands when adding elements gives a positive sum and resets when the sum becomes negative.

2. **[#121 Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)** - Find the maximum profit from buying and selling a stock once.
   - **Approach**: Slide through the prices, keeping track of the minimum price seen so far and maximum profit possible.

3. **[#643 Maximum Average Subarray I](https://leetcode.com/problems/maximum-average-subarray-i/)** - Find the subarray of length k with the highest average.
   - **Approach**: Use a fixed-size window of length k to find the maximum sum, then divide by k.

4. **[#1984 Minimum Difference Between Highest and Lowest of K Scores](https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/)** - Find the minimum difference between highest and lowest scores in any k-sized window.
   - **Approach**: Sort the array, then use a fixed-size window to find the minimum difference.

### Medium Problems

1. **[#3 Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)** - Find the longest substring without repeating characters.
   - **Approach**: Use a variable-size window and a set/map to track characters in the current window.

2. **[#209 Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/)** - Find the smallest subarray with a sum at least equal to the target.
   - **Approach**: Use a variable-size window that expands to add elements and contracts when the sum exceeds the target.

3. **[#438 Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/)** - Find all anagrams of a pattern in a string.
   - **Approach**: Use a fixed-size window equal to the pattern length, and a frequency counter to check for anagrams.

4. **[#567 Permutation in String](https://leetcode.com/problems/permutation-in-string/)** - Check if a string contains a permutation of another string.
   - **Approach**: Similar to finding anagrams, use a fixed-size window and frequency counter.

5. **[#1004 Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/)** - Find the longest subarray of 1s after flipping at most k 0s.
   - **Approach**: Use a variable-size window that can include at most k zeros.

6. **[#424 Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/)** - Find the longest substring where you can replace at most k characters to make all characters the same.
   - **Approach**: Use a variable-size window, tracking the frequency of each character to determine if the window is valid.

7. **[#904 Fruit Into Baskets](https://leetcode.com/problems/fruit-into-baskets/)** - Find the longest subarray containing at most 2 distinct types of fruits (numbers).
   - **Approach**: Use a variable-size window and a map to track at most 2 distinct types.

### Hard Problems

1. **[#76 Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/)** - Find the smallest substring that contains all characters from the target string.
   - **Approach**: Use a variable-size window and a frequency counter to track when all required characters are included.

2. **[#239 Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)** - Find the maximum element in each sliding window of size k.
   - **Approach**: Use a deque (double-ended queue) to efficiently track the maximum element in each window.

3. **[#992 Subarrays with K Different Integers](https://leetcode.com/problems/subarrays-with-k-different-integers/)** - Count the number of subarrays with exactly k different integers.
   - **Approach**: Use the difference between the count of subarrays with at most k different integers and the count of subarrays with at most (k-1) different integers.

4. **[#1074 Number of Submatrices That Sum to Target](https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/)** - Count submatrices that sum to the target value.
   - **Approach**: Combine sliding window with prefix sum technique, reducing a 2D problem to multiple 1D problems.

### Tips for Solving LeetCode Sliding Window Problems

1. **Identify the Window**: Determine if you need a fixed or variable-size window
2. **Track State**: Decide what information you need to track within the window (sums, counts, etc.)
3. **Window Movement**: Implement the logic for expanding and shrinking the window
4. **Optimization**: Pay attention to time complexity, especially for large inputs
5. **Edge Cases**: Always consider empty arrays, single elements, or when k exceeds the array length

## Why Learn This Pattern?

The sliding window pattern is super useful because:
1. It's much faster than checking every possible subarray or substring
2. It's used in many real-world applications like data analysis, text processing, and network monitoring
3. It's a favorite in coding interviews and competitions