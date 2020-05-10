// leetcode_0005_findMedianSortedArrays
package main

import (
	"fmt"
)

/**
4. 寻找两个正序数组的中位数
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。

示例 1:
nums1 = [1, 3]
nums2 = [2]
则中位数是 2.0

示例 2:
nums1 = [1, 2]
nums2 = [3, 4]
则中位数是 (2 + 3)/2 = 2.5

*/
func main() {

	nums1, nums2, expect := []int{1, 3}, []int{2}, 2.0
	result := findMedianSortedArrays(nums1, nums2)
	if expect != result {
		fmt.Println("nums1=", nums1, ",nums2=", nums2, ",result=", result, " (expect=", expect, ")")
		panic(expect)
	}

	nums1, nums2, expect = []int{1, 2}, []int{3, 4}, 2.5
	result = findMedianSortedArrays(nums1, nums2)
	if expect != result {
		fmt.Println("nums1=", nums1, ",nums2=", nums2, ",result=", result, " (expect=", expect, ")")
		panic(expect)
	}

	nums1, nums2, expect = []int{}, []int{1}, 1
	result = findMedianSortedArrays(nums1, nums2)
	if expect != result {
		fmt.Println("nums1=", nums1, ",nums2=", nums2, ",result=", result, " (expect=", expect, ")")
		panic(expect)
	}

	nums1, nums2, expect = []int{}, []int{2, 3}, 2.5
	result = findMedianSortedArrays(nums1, nums2)
	if expect != result {
		fmt.Println("nums1=", nums1, ",nums2=", nums2, ",result=", result, " (expect=", expect, ")")
		panic(expect)
	}

	nums1, nums2, expect = []int{2, 3}, []int{1}, 2.0
	result = findMedianSortedArrays(nums1, nums2)
	if expect != result {
		fmt.Println("nums1=", nums1, ",nums2=", nums2, ",result=", result, " (expect=", expect, ")")
		panic(expect)
	}

	nums1, nums2, expect = []int{2, 3, 4}, []int{1}, 2.5
	result = findMedianSortedArrays(nums1, nums2)
	if expect != result {
		fmt.Println("nums1=", nums1, ",nums2=", nums2, ",result=", result, " (expect=", expect, ")")
		panic(expect)
	}

	nums1, nums2, expect = []int{1, 2, 4, 5}, []int{3}, 3.0
	result = findMedianSortedArrays(nums1, nums2)
	if expect != result {
		fmt.Println("nums1=", nums1, ",nums2=", nums2, ",result=", result, " (expect=", expect, ")")
		panic(expect)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		nums1, nums2 = nums2, nums1
	}
	if len(nums2) == 0 {
		i := len(nums1) / 2
		a := float64(nums1[i])
		if len(nums1)%2 == 0 {
			a = (a + float64(nums1[i-1])) / 2
		}
		return a
	}
	if len(nums1)%2 == 0 {
		n1L, n2L := len(nums1), len(nums2)
		n1Last := nums1[n1L-1]
		nums2 = append(nums2, n1Last)
		i := n2L - 1
		for i >= 0 && nums2[i] >= n1Last {
			nums2[i+1] = nums2[i]
			i--
		}
		nums2[i+1] = n1Last

		nums1 = nums1[:n1L-1]
	}

	ai, bi := len(nums1)/2, len(nums2)/2
	mid := true
	if len(nums2)%2 == 0 {
		bi -= 1
		mid = false
	}

	a, b := nums1[ai], nums2[bi]
	if a > b {
		a, b = b, a
		ai, bi = bi, ai
		nums1, nums2 = nums2, nums1
	}
	for ai+1 < len(nums1) && bi >= 1 && nums1[ai+1] <= nums2[bi-1] {
		a, b = nums1[ai+1], nums2[bi-1]
		ai++
		bi--
	}

	if bi >= 1 && nums2[bi-1] > a {
		a = nums2[bi-1]
	}
	if ai+1 < len(nums1) && nums1[ai+1] < b {
		b = nums1[ai+1]
	}

	var result float64 = 0.0
	if mid {
		result = float64(a+b) / 2
	} else {
		result = float64(b)
	}

	return result
}
