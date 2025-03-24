# Merge Intervals Pattern

## What are Merge Intervals?

Imagine you have a bunch of time slots on your calendar, and some of them overlap. The Merge Intervals pattern helps you combine these overlapping time slots into simpler, non-overlapping ones. It's like cleaning up your messy schedule by combining overlapping appointments into single blocks of busy time!

## Real-Life Examples

1. **Calendar Appointments**: When you have multiple meetings that overlap, you combine them into one block of busy time.
   - 9:00-10:00 Meeting with Team
   - 9:30-11:00 Project Review
   - These merge into: 9:00-11:00 (one busy block)

2. **TV Show Schedule**: When TV shows overlap, you might want to know the total time you'll be watching TV.
   - Show 1: 8:00-9:00
   - Show 2: 8:30-9:30
   - These merge into: 8:00-9:30 (one continuous block)

3. **Class Schedule**: When classes have overlapping times, you need to know your total class time.
   - Math: 10:00-11:30
   - Science: 11:00-12:30
   - These merge into: 10:00-12:30 (one continuous block)

## When Do We Use Merge Intervals?

Use this technique when you need to:
- Combine overlapping time periods
- Find free time slots between busy periods
- Calculate total busy time
- Schedule meetings without conflicts
- Find the minimum number of rooms needed for meetings

## Types of Intervals

1. **Overlapping Intervals**: When one interval's end time is greater than or equal to another interval's start time
   ```
   Interval 1: [1, 5]
   Interval 2: [3, 7]
   Result: [1, 7]
   ```

2. **Non-overlapping Intervals**: When intervals don't overlap at all
   ```
   Interval 1: [1, 3]
   Interval 2: [4, 6]
   Result: [1, 3], [4, 6]
   ```

3. **Contained Intervals**: When one interval is completely inside another
   ```
   Interval 1: [1, 7]
   Interval 2: [3, 5]
   Result: [1, 7]
   ```

## How Does It Work?

1. **Step 1**: Sort all intervals by their start time
2. **Step 2**: Take the first interval as the current interval
3. **Step 3**: For each next interval:
   - If it overlaps with the current interval, merge them
   - If it doesn't overlap, start a new current interval
4. **Step 4**: Keep track of the merged intervals

## Simple Code Examples

### Basic Merge Intervals Example
Merging overlapping intervals:

```go
type Interval struct {
    Start int
    End   int
}

func mergeIntervals(intervals []Interval) []Interval {
    // If we have less than 2 intervals, no need to merge
    if len(intervals) <= 1 {
        return intervals
    }
    
    // Sort intervals by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    
    // Store merged intervals
    result := []Interval{intervals[0]}
    
    // Check each interval
    for i := 1; i < len(intervals); i++ {
        // Get the last merged interval
        last := &result[len(result)-1]
        
        // If current interval overlaps with last merged interval
        if intervals[i].Start <= last.End {
            // Update the end time if current interval ends later
            if intervals[i].End > last.End {
                last.End = intervals[i].End
            }
        } else {
            // No overlap, add as new interval
            result = append(result, intervals[i])
        }
    }
    
    return result
}
```

### Insert New Interval Example
Inserting a new interval into existing intervals:

```go
func insertInterval(intervals []Interval, newInterval Interval) []Interval {
    result := []Interval{}
    
    // Add all intervals that come before newInterval
    i := 0
    for i < len(intervals) && intervals[i].End < newInterval.Start {
        result = append(result, intervals[i])
        i++
    }
    
    // Merge overlapping intervals
    for i < len(intervals) && intervals[i].Start <= newInterval.End {
        newInterval.Start = min(newInterval.Start, intervals[i].Start)
        newInterval.End = max(newInterval.End, intervals[i].End)
        i++
    }
    
    // Add the merged interval
    result = append(result, newInterval)
    
    // Add remaining intervals
    for i < len(intervals) {
        result = append(result, intervals[i])
        i++
    }
    
    return result
}
```

## Common Mistakes to Avoid

1. **Forgetting to sort**: Always sort intervals by start time first
2. **Incorrect overlap check**: Make sure to check both start and end times
3. **Not handling edge cases**: What if intervals are empty? What if they're all the same?

## Fun Practice Problems

1. **Movie Marathon**: You have a list of movie times. Find the total time you'll spend watching movies if you watch all of them.

2. **Playground Schedule**: Given different playtime slots for kids, find the total time the playground is busy.

3. **Library Hours**: A library has different opening hours on different days. Find the total hours it's open in a week.

4. **Bus Schedule**: Given arrival and departure times of buses at a station, find the maximum number of buses present at any time.

5. **Class Schedule**: Given class timings, find the minimum number of classrooms needed to schedule all classes.

## LeetCode Problems Using Merge Intervals

Here are some popular LeetCode problems that can be solved using the Merge Intervals technique:

### Easy Problems

1. **[#252 Meeting Rooms](https://leetcode.com/problems/meeting-rooms/)** - Check if a person can attend all meetings.
   - **Approach**: Sort intervals and check for overlaps.

2. **[#252 Meeting Rooms II](https://leetcode.com/problems/meeting-rooms-ii/)** - Find minimum number of meeting rooms needed.
   - **Approach**: Use priority queue to track room usage.

3. **[#56 Merge Intervals](https://leetcode.com/problems/merge-intervals/)** - Merge all overlapping intervals.
   - **Approach**: Sort and merge overlapping intervals.

4. **[#57 Insert Interval](https://leetcode.com/problems/insert-interval/)** - Insert a new interval into existing intervals.
   - **Approach**: Find insertion point and merge overlapping intervals.

### Medium Problems

1. **[#253 Meeting Rooms II](https://leetcode.com/problems/meeting-rooms-ii/)** - Find minimum number of meeting rooms needed.
   - **Approach**: Use priority queue to track room usage.

2. **[#435 Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)** - Find minimum number of intervals to remove to make all intervals non-overlapping.
   - **Approach**: Sort and count overlapping intervals.

3. **[#452 Minimum Number of Arrows to Burst Balloons](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)** - Find minimum arrows needed to burst all balloons.
   - **Approach**: Sort by end time and count non-overlapping intervals.

4. **[#763 Partition Labels](https://leetcode.com/problems/partition-labels/)** - Partition string into as many parts as possible.
   - **Approach**: Create intervals for each character and merge overlapping ones.

5. **[#986 Interval List Intersections](https://leetcode.com/problems/interval-list-intersections/)** - Find intersection of two lists of intervals.
   - **Approach**: Use two pointers to find overlapping intervals.

### Hard Problems

1. **[#759 Employee Free Time](https://leetcode.com/problems/employee-free-time/)** - Find common free time for all employees.
   - **Approach**: Merge all intervals and find gaps.

2. **[#715 Range Module](https://leetcode.com/problems/range-module/)** - Design a data structure to track ranges.
   - **Approach**: Use balanced tree to maintain non-overlapping intervals.

3. **[#352 Data Stream as Disjoint Intervals](https://leetcode.com/problems/data-stream-as-disjoint-intervals/)** - Design a data structure for streaming numbers.
   - **Approach**: Maintain sorted list of non-overlapping intervals.

4. **[#699 Falling Squares](https://leetcode.com/problems/falling-squares/)** - Track heights of falling squares.
   - **Approach**: Use interval tree to track overlapping squares.

### Tips for Solving LeetCode Merge Interval Problems

1. **Sort First**: Always sort intervals by start time
2. **Track State**: Keep track of current merged interval
3. **Overlap Check**: Implement correct overlap checking logic
4. **Edge Cases**: Handle empty lists and single intervals
5. **Optimization**: Look for ways to reduce unnecessary comparisons

## Why Learn This Pattern?

The Merge Intervals pattern is super useful because:
1. It helps solve real-world scheduling problems efficiently
2. It's used in many applications like calendar apps and meeting schedulers
3. It's a favorite in coding interviews and competitions
4. It teaches important concepts about sorting and merging