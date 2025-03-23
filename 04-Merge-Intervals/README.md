# Merge Intervals Pattern

## Introduction

The Merge Intervals pattern is a technique to deal with overlapping intervals. This pattern is useful for problems that involve merging intervals, finding intersections, or handling any other operation that requires understanding how intervals relate to each other.

An interval is typically represented as a pair of numbers (start, end) and can be used to represent various concepts such as time ranges, numerical ranges, or periods.

## How It Works

1. Sort the intervals based on the start time
2. Iterate through the sorted intervals and compare each interval with the next
3. Merge overlapping intervals or perform other required operations based on their relationships

## Time and Space Complexity

- **Time Complexity**: O(n log n) - where n is the number of intervals (mainly due to sorting)
- **Space Complexity**: O(n) - for storing the result

## When to Use Merge Intervals

Use the Merge Intervals pattern when:
- You're dealing with overlapping intervals
- You need to find free time slots in a schedule
- You need to find the intersection between intervals
- You need to determine if intervals can be merged

## Common Problem Patterns

1. **Merge Overlapping Intervals**
   - Given a collection of intervals, merge all overlapping intervals
   - Sort by start time and then merge intervals if they overlap

2. **Insert Interval**
   - Given a set of non-overlapping intervals and a new interval, insert the new interval at the correct position and merge if necessary
   - Identify intervals that come before, overlap with, and come after the new interval

3. **Interval Intersection**
   - Given two lists of intervals, find the intersection of these intervals
   - For each pair of intervals from the two lists, find their intersection if it exists

4. **Conflicting Appointments**
   - Given a set of intervals representing appointments, find if a person can attend all of them
   - Check if any two intervals overlap

## Implementation in Golang

```go
// Interval structure
type Interval struct {
    Start int
    End   int
}

// Merge overlapping intervals
func mergeIntervals(intervals []Interval) []Interval {
    if len(intervals) <= 1 {
        return intervals
    }
    
    // Sort intervals based on start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    
    merged := []Interval{intervals[0]}
    
    for i := 1; i < len(intervals); i++ {
        current := intervals[i]
        lastMerged := &merged[len(merged)-1]
        
        // If current interval overlaps with the last merged interval, merge them
        if current.Start <= lastMerged.End {
            // Update end of the last merged interval if necessary
            if current.End > lastMerged.End {
                lastMerged.End = current.End
            }
        } else {
            // No overlap, add current interval to result
            merged = append(merged, current)
        }
    }
    
    return merged
}

// Insert an interval into a list of non-overlapping intervals
func insertInterval(intervals []Interval, newInterval Interval) []Interval {
    result := []Interval{}
    i := 0
    n := len(intervals)
    
    // Add all intervals that come before newInterval
    for i < n && intervals[i].End < newInterval.Start {
        result = append(result, intervals[i])
        i++
    }
    
    // Merge all overlapping intervals
    for i < n && intervals[i].Start <= newInterval.End {
        newInterval.Start = min(newInterval.Start, intervals[i].Start)
        newInterval.End = max(newInterval.End, intervals[i].End)
        i++
    }
    
    // Add the merged interval
    result = append(result, newInterval)
    
    // Add all intervals that come after newInterval
    for i < n {
        result = append(result, intervals[i])
        i++
    }
    
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

## Example Problems

1. **Merge Intervals**
2. **Insert Interval**
3. **Intervals Intersection**
4. **Conflicting Appointments**
5. **Minimum Meeting Rooms**
6. **Maximum CPU Load**
7. **Employee Free Time**

Each of these problems has a dedicated solution file in this directory. 