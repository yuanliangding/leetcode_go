// t
package main

// import (
// 	"fmt"
// )

/**
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: "pwwkew"
输出: 3

解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
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
