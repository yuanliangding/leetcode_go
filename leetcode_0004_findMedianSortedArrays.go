// leetcode_0005_findMedianSortedArrays
package main

import (
	"fmt"
)

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
		n2L := len(nums2)
		nums2 = append(nums2, nums2[n2L-1])
		i := n2L - 1
		for i >= 0 && nums2[i] >= nums1[0] {
			nums2[i+1] = nums2[i]
			i--
		}
		nums2[i+1] = nums1[0]

		nums1 = nums1[1:]
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
