package topk

import (
	"container/heap"
	"strings"
)

// CharFrequency represents a character with its frequency
type CharFrequency struct {
	Char      rune
	Frequency int
}

// MaxCharFrequencyHeap is a max heap based on character frequency
type MaxCharFrequencyHeap []CharFrequency

func (h MaxCharFrequencyHeap) Len() int           { return len(h) }
func (h MaxCharFrequencyHeap) Less(i, j int) bool { return h[i].Frequency > h[j].Frequency } // Max heap
func (h MaxCharFrequencyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxCharFrequencyHeap) Push(x interface{}) {
	*h = append(*h, x.(CharFrequency))
}

func (h *MaxCharFrequencyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// FrequencySort sorts characters in a string by decreasing frequency
// Time Complexity: O(N log K) where N is string length and K is number of unique characters
// Space Complexity: O(K) for the frequency map and heap
func FrequencySort(s string) string {
	// Count frequency of each character
	frequencyMap := make(map[rune]int)
	for _, char := range s {
		frequencyMap[char]++
	}

	// Create a max heap based on frequency
	h := &MaxCharFrequencyHeap{}
	heap.Init(h)

	// Add all characters with their frequencies to the heap
	for char, frequency := range frequencyMap {
		heap.Push(h, CharFrequency{Char: char, Frequency: frequency})
	}

	// Build the result string by extracting from the heap
	var result strings.Builder
	for h.Len() > 0 {
		item := heap.Pop(h).(CharFrequency)
		result.WriteString(strings.Repeat(string(item.Char), item.Frequency))
	}

	return result.String()
}

// FrequencySortBucket uses bucket sort to sort characters by frequency
// Time Complexity: O(N) where N is the string length
// Space Complexity: O(N) for the frequency map and buckets
func FrequencySortBucket(s string) string {
	// Count frequency of each character
	frequencyMap := make(map[rune]int)
	for _, char := range s {
		frequencyMap[char]++
	}

	// Find the maximum frequency
	maxFrequency := 0
	for _, frequency := range frequencyMap {
		if frequency > maxFrequency {
			maxFrequency = frequency
		}
	}

	// Create buckets where index represents frequency
	buckets := make([][]rune, maxFrequency+1)
	for char, frequency := range frequencyMap {
		if buckets[frequency] == nil {
			buckets[frequency] = []rune{}
		}
		buckets[frequency] = append(buckets[frequency], char)
	}

	// Build the result string by traversing buckets from highest frequency
	var result strings.Builder
	for i := maxFrequency; i > 0; i-- {
		for _, char := range buckets[i] {
			result.WriteString(strings.Repeat(string(char), i))
		}
	}

	return result.String()
}

// Example usage:
//
// Input: s = "tree"
// Output: "eert" or "eetr"
// Explanation: 'e' appears twice while 'r' and 't' both appear once.
// So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
//
// Input: s = "cccaaa"
// Output: "cccaaa" or "aaaccc"
// Explanation: Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
// Note that "cacaca" is incorrect, as the same characters must be together.
//
// Input: s = "Aabb"
// Output: "bbAa" or "bbaA"
// Explanation: "bbaA" is also a valid answer, but "Aabb" is incorrect.
// Note that 'A' and 'a' are treated as two different characters.
