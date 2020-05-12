// leetcode_0005_longestPalindrome
package main

import (
	"fmt"
)

/*
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：
输入: "cbbd"
输出: "bb"

*/
func main() {
	str, expect := "cbbd", "bb"
	result := longestPalindrome(str)
	if expect != result {
		fmt.Println("str=", str, ",result=", result, " (expect=", expect, ")")
		panic(expect)
	}

	str, expect = "babad", "bab"
	result = longestPalindrome(str)
	if expect != result {
		fmt.Println("str=", str, ",result=", result, " (expect=", expect, ")")
		panic(expect)
	}

}

func longestPalindrome(s string) string {
	palindrome := ""
	index := make(map[byte][]int)

	length := len(s)
	var i int
	for i = 0; i < length; i++ {
		v := s[i]
		var is []int
		var ok bool
		if is, ok = index[v]; !ok {
			is = []int{}
			index[v] = is
		}
		is = append(is, i)
		index[v] = is
	}

	for i = 0; i < length-len(palindrome); i++ {
		indexs := index[s[i]]
		iLen := len(indexs)
		for j := iLen - 1; j >= 0; j-- {
			end := indexs[j] + 1
			if end-i > len(palindrome) && isPalindrome(s[i:end]) {
				palindrome = s[i:end]
			}
		}
	}

	return palindrome
}

func isPalindrome(s string) bool {
	length := len(s)
	midLen := length/2 - 1

	for i := 0; i <= midLen; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
