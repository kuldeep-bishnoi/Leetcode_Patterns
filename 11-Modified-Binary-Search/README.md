# Modified Binary Search Pattern

## What is Modified Binary Search?

Imagine you're playing a number guessing game where you need to find a special number. Regular binary search is like playing "higher or lower" with sorted numbers. Modified binary search is like playing the same game, but with some special rules or twists! It's like finding a book in a library where the books are arranged in a special way.

## Real-Life Examples

1. **Finding a Book**: When looking for a book in a library.
   - Books are arranged by size
   - Some books are missing
   - You need to find the first book of a certain size

2. **Temperature Reading**: When checking temperature records.
   - Temperatures are recorded every hour
   - Some readings are missing
   - You need to find the first temperature above 30°C

3. **Game Score**: When looking at game scores.
   - Scores are arranged by date
   - Some scores are missing
   - You need to find the first score above 100

## When Do We Use Modified Binary Search?

Use this technique when:
- You need to find a specific element in a sorted array
- The array might have duplicates
- You need to find the first or last occurrence
- The array might be rotated
- You need to find the smallest or largest element

## How Does It Work?

1. **Step 1**: Find the middle element
2. **Step 2**: Compare with target:
   - If equal, check if it's the first/last occurrence
   - If smaller, look in the right half
   - If larger, look in the left half
3. **Step 3**: Keep going until you find what you're looking for

## Simple Code Example

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == target {
            // Found target, check if it's the first occurrence
            if mid == 0 || nums[mid-1] != target {
                return mid
            }
            // Look for first occurrence in left half
            right = mid - 1
        } else if nums[mid] < target {
            // Look in right half
            left = mid + 1
        } else {
            // Look in left half
            right = mid - 1
        }
    }
    
    return -1 // Target not found
}
```

## Common Mistakes to Avoid

1. **Infinite Loops**: Make sure to update left/right correctly
2. **Overflow**: Use `left + (right-left)/2` instead of `(left+right)/2`
3. **Edge Cases**: Handle empty arrays and single elements
4. **Duplicate Handling**: Consider how to handle duplicates

## Fun Practice Problems

1. **Book Finder**: Find the first book of a certain size
2. **Temperature Tracker**: Find the first temperature above 30°C
3. **Score Seeker**: Find the first score above 100
4. **Number Navigator**: Find the first occurrence of a number
5. **Rotation Finder**: Find the smallest number in a rotated array

## LeetCode Problems Using Modified Binary Search

Here are some popular LeetCode problems that can be solved using Modified Binary Search:

### Easy Problems

1. **[#704 Binary Search](https://leetcode.com/problems/binary-search/)** - Find target in sorted array.
   - **Approach**: Use standard binary search.

2. **[#35 Search Insert Position](https://leetcode.com/problems/search-insert-position/)** - Find insertion position.
   - **Approach**: Use binary search to find position.

### Medium Problems

1. **[#34 Find First and Last Position](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)** - Find first and last occurrence.
   - **Approach**: Use binary search twice.

2. **[#33 Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)** - Search in rotated array.
   - **Approach**: Use binary search with rotation handling.

3. **[#153 Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)** - Find minimum in rotated array.
   - **Approach**: Use binary search to find pivot.

4. **[#162 Find Peak Element](https://leetcode.com/problems/find-peak-element/)** - Find peak in array.
   - **Approach**: Use binary search with peak detection.

### Hard Problems

1. **[#4 Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)** - Find median of two arrays.
   - **Approach**: Use binary search on smaller array.

2. **[#154 Find Minimum in Rotated Sorted Array II](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/)** - Find minimum with duplicates.
   - **Approach**: Use binary search with duplicate handling.

### Tips for Solving LeetCode Modified Binary Search Problems

1. **Array Properties**: Check if array is sorted/rotated
2. **Duplicate Handling**: Consider how to handle duplicates
3. **Edge Cases**: Handle empty arrays and single elements
4. **Target Position**: Determine if you need first/last occurrence
5. **Mid Calculation**: Use safe mid calculation to avoid overflow

## Why Learn This Pattern?

The Modified Binary Search pattern is super useful because:
1. It's very efficient (O(log n) time)
2. It's used in many real-world applications
3. It's a favorite in coding interviews
4. It teaches important concepts about searching
5. It helps solve many array-related problems

Once you master this pattern, you'll be able to solve many search problems efficiently and impress your friends with your coding skills! 