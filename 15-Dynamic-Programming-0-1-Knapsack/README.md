# Dynamic Programming - 0/1 Knapsack Pattern

## What is 0/1 Knapsack?

Imagine you have a magical backpack (knapsack) that can only hold a certain weight, and you have many toys with different weights and values. The 0/1 Knapsack pattern is like playing a game where you need to choose which toys to put in your backpack to get the most fun (value) without making it too heavy! You can either take a toy (1) or leave it (0), but you can't split toys in half.

## Real-Life Examples

1. **Toy Selection**: When packing toys for a trip.
   - Each toy has a weight and fun value
   - Backpack can only hold 10 pounds
   - Need to choose toys that give most fun

2. **Snack Packing**: When packing snacks for school.
   - Each snack has calories and taste value
   - Lunch box can only hold 500 calories
   - Need to choose snacks that taste best

3. **Game Collection**: When choosing games to bring.
   - Each game has size and fun rating
   - Game bag can only hold 5 games
   - Need to choose games with highest ratings

## When Do We Use 0/1 Knapsack?

Use this technique when:
- You need to make yes/no choices
- You have a weight or size limit
- You want to maximize value
- You can't split items
- You need to solve optimization problems

## How Does It Work?

1. **Step 1**: Create a table for possibilities
2. **Step 2**: For each item:
   - Try taking it (if it fits)
   - Try leaving it
   - Choose the better option
3. **Step 3**: Keep track of best choices

Example:
```
Toys: [(2kg, $5), (3kg, $8), (4kg, $10)]
Backpack: 5kg

Table:
Weight: 0 1 2 3 4 5
Toy 1:  0 0 5 5 5 5
Toy 2:  0 0 5 8 8 13
Toy 3:  0 0 5 8 8 13

Best choice: Take Toy 2 and Toy 1 (3kg + 2kg = 5kg, $8 + $5 = $13)
```

## Simple Code Example

```go
func knapsack(weights []int, values []int, capacity int) int {
    n := len(weights)
    // Create table for possibilities
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, capacity+1)
    }
    
    // Fill table
    for i := 1; i <= n; i++ {
        for w := 1; w <= capacity; w++ {
            // Can't take this item
            dp[i][w] = dp[i-1][w]
            
            // Can take this item
            if weights[i-1] <= w {
                // Take maximum of taking or leaving
                dp[i][w] = max(dp[i][w], 
                             values[i-1] + dp[i-1][w-weights[i-1]])
            }
        }
    }
    
    return dp[n][capacity]
}
```

## Common Mistakes to Avoid

1. **Table Size**: Make table one size bigger
2. **Index Offsets**: Remember array indices start at 0
3. **Weight Check**: Always check if item fits
4. **Value Update**: Take maximum of choices

## Fun Practice Problems

1. **Toy Packer**: Pack toys in backpack
2. **Snack Selector**: Choose best snacks
3. **Game Collector**: Select best games
4. **Treasure Hunter**: Choose valuable items
5. **Lunch Planner**: Pack best lunch

## LeetCode Problems Using 0/1 Knapsack

Here are some popular LeetCode problems that can be solved using 0/1 Knapsack:

### Easy Problems

1. **[#416 Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/)** - Split array into equal parts.
   - **Approach**: Use knapsack with half sum.

2. **[#1049 Last Stone Weight II](https://leetcode.com/problems/last-stone-weight-ii/)** - Find minimum difference.
   - **Approach**: Use knapsack with sum/2.

### Medium Problems

1. **[#474 Ones and Zeroes](https://leetcode.com/problems/ones-and-zeroes/)** - Find maximum strings.
   - **Approach**: Use 2D knapsack.

2. **[#494 Target Sum](https://leetcode.com/problems/target-sum/)** - Find ways to reach target.
   - **Approach**: Use knapsack with sum difference.

### Hard Problems

1. **[#879 Profitable Schemes](https://leetcode.com/problems/profitable-schemes/)** - Find profitable combinations.
   - **Approach**: Use 2D knapsack with profit.

2. **[#956 Tallest Billboard](https://leetcode.com/problems/tallest-billboard/)** - Find tallest possible billboard.
   - **Approach**: Use knapsack with height difference.

### Tips for Solving LeetCode 0/1 Knapsack Problems

1. **State Definition**: Define states clearly
   - Current item
   - Remaining capacity
   - Current value

2. **Transitions**: Handle transitions properly
   - Take item
   - Leave item
   - Edge cases

3. **Optimization**: Use space optimization
   - 1D array for basic cases
   - Rolling array for memory
   - State compression

4. **Initialization**: Initialize properly
   - First row to 0
   - First column to 0
   - Rest to appropriate values

## Why Learn This Pattern?

The 0/1 Knapsack pattern is super useful because:
1. It's used in many optimization problems
2. It's a favorite in coding interviews
3. It teaches important concepts about dynamic programming
4. It helps solve many real-world problems
5. It's a fundamental algorithm pattern