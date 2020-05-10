// t
package main

// import (
// 	"fmt"
// )

// func main() {
// 	var c1 int = lengthOfLongestSubstring("abc")
// 	fmt.Println(c1)
// }

func lengthOfLongestSubstring(s string) int {
	maxLen, begin, end, sLen := 0, 0, 0, len(s)
	index := make(map[byte]int)
	for ; end < sLen; end++ {
		c := s[end]
		if i, ok := index[c]; ok {
			for j := begin; j < i; j++ {
				delete(index, s[j])
			}
			begin = i + 1
		}
		index[c] = end
		if maxLen < len(index) {
			maxLen = len(index)
		}
	}

	return maxLen
}

// func lengthOfLongestSubstring(s string) int {
// 	var maxLen int = 0
// 	// var begin int = 0
// 	var end int = 0

// 	index := make(map[byte]int)
// 	indexPtr := &index

// 	for ; end < len(s); end++ {
// 		var c byte = s[end]
// 		if i, ok := (*indexPtr)[c]; ok {
// 			indexPtr, _, end = tryMove(&s, i+1, maxLen+1)
// 			// begin = moveResult.start
// 			c = s[end]
// 		}
// 		(*indexPtr)[c] = end

// 		if maxLen < len(*indexPtr) {
// 			maxLen = len(*indexPtr)
// 		}
// 	}

// 	return maxLen
// }

func tryMove(s *string, start int, length int) (*map[byte]int, int, int) {
	index := make(map[byte]int)

	var last int = start + length - 1
	var l int = len(*s)
	if last >= l {
		return &index, l - 1, l - 1
	}
	for i := last; i >= start; i-- {
		var c byte = (*s)[i]
		if _, ok := index[c]; ok {
			return tryMove(s, i+1, length)
		} else {
			index[c] = i
		}
	}

	return &index, start, last
}
