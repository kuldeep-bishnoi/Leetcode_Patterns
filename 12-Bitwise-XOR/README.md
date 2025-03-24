# Bitwise XOR Pattern

## What is Bitwise XOR?

Imagine you have a special magic trick where you can combine two numbers in a special way! XOR (Exclusive OR) is like a game where you compare two numbers and get a new number based on whether they're the same or different. It's like playing "Spot the Difference" with numbers!

## Real-Life Examples

1. **Light Switch**: When you flip a light switch twice, it goes back to its original state.
   - ON → OFF → ON
   - OFF → ON → OFF

2. **Matching Socks**: When you have pairs of socks.
   - If you have two red socks, they match (0)
   - If you have one red and one blue sock, they don't match (1)

3. **Secret Messages**: When you want to hide and show a message.
   - Use a key to hide the message
   - Use the same key to show the message again

## When Do We Use Bitwise XOR?

Use this technique when:
- You need to find a single number in a list of pairs
- You want to swap numbers without using extra space
- You need to find missing numbers
- You want to check if numbers are different
- You need to solve problems with binary operations

## How Does It Work?

1. **Step 1**: Compare two numbers bit by bit
2. **Step 2**: If bits are different, result is 1
3. **Step 3**: If bits are same, result is 0

Example:
```
1 XOR 1 = 0
0 XOR 0 = 0
1 XOR 0 = 1
0 XOR 1 = 1
```

## Simple Code Example

```go
func findSingleNumber(nums []int) int {
    result := 0
    
    // XOR all numbers
    for _, num := range nums {
        result ^= num
    }
    
    return result
}
```

## Common Mistakes to Avoid

1. **Order Matters**: XOR is associative and commutative
2. **Zero Handling**: XOR with 0 gives the same number
3. **Same Number**: XOR with same number gives 0
4. **Bit Operations**: Remember how bits work

## Fun Practice Problems

1. **Sock Matcher**: Find the unpaired sock
2. **Light Flipper**: Track light state after flips
3. **Number Finder**: Find the missing number
4. **Secret Keeper**: Hide and show messages
5. **Bit Swapper**: Swap numbers without extra space

## LeetCode Problems Using Bitwise XOR

Here are some popular LeetCode problems that can be solved using Bitwise XOR:

### Easy Problems

1. **[#136 Single Number](https://leetcode.com/problems/single-number/)** - Find single number in array.
   - **Approach**: XOR all numbers.

2. **[#268 Missing Number](https://leetcode.com/problems/missing-number/)** - Find missing number.
   - **Approach**: XOR with range.

### Medium Problems

1. **[#260 Single Number III](https://leetcode.com/problems/single-number-iii/)** - Find two single numbers.
   - **Approach**: Use XOR with grouping.

2. **[#371 Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/)** - Add without + or -.
   - **Approach**: Use XOR for sum, AND for carry.

### Hard Problems

1. **[#421 Maximum XOR of Two Numbers](https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/)** - Find max XOR pair.
   - **Approach**: Use trie or bit manipulation.

2. **[#1707 Maximum XOR With an Element](https://leetcode.com/problems/maximum-xor-with-an-element-from-array/)** - Find max XOR with limit.
   - **Approach**: Use trie with query optimization.

### Tips for Solving LeetCode Bitwise XOR Problems

1. **Properties**: Remember XOR properties
   - a XOR a = 0
   - a XOR 0 = a
   - a XOR b = b XOR a
   - (a XOR b) XOR c = a XOR (b XOR c)

2. **Bit Operations**: Understand bit operations
   - AND (&)
   - OR (|)
   - XOR (^)
   - NOT (~)

3. **Pattern Recognition**: Look for patterns
   - Pairs of numbers
   - Missing numbers
   - Single numbers
   - Binary operations

4. **Optimization**: Use XOR for space efficiency
   - No extra variables
   - In-place operations
   - Constant space

## Why Learn This Pattern?

The Bitwise XOR pattern is super useful because:
1. It's very efficient (O(n) time, O(1) space)
2. It's used in many real-world applications
3. It's a favorite in coding interviews
4. It teaches important concepts about binary operations
5. It helps solve many array-related problems

Once you master this pattern, you'll be able to solve many bit manipulation problems efficiently and impress your friends with your coding skills! 