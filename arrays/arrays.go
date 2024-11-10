package arrays

import (
	"math"
	"strings"
)

func MaximumArraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	max := nums[0]
	for i := range nums {
		max = int(math.Max(float64(max+nums[i]), float64(nums[i])))
		res = int(math.Max(float64(max), float64(res)))
	}
	return res
}

/*
26. Remove Duplicates from Sorted Array
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place
such that each unique element appears only once. The relative order of the elements
should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted,
you need to do the following things:
1.Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
2.Return k.
*/

func RemoveDuplicates(nums []int) int {
	index := 0
	for i := range nums {
		if nums[index] != nums[i] {
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
}

/*
27. Remove Element
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed.
Then return the number of elements in nums which are not equal to val.
Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do
the following things:

1. Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
The remaining elements of nums are not important as well as the size of nums.
2. Return k.

Example:

Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
*/
func RemoveElement(nums []int, val int) int {
	index := 0
	for i := range nums {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

/*
28. Find the Index of the First Occurrence in a String
Given two strings needle and haystack, return the index of the first occurrence of needle in
haystack, or -1 if needle is not part of haystack.

Example:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
*/

// Solution 1
func StrStr1(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		// for i := 0; i < len(haystack)+1-len(needle); i++ {
		for j := 0; j < len(haystack); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}

// Solution 2
func StrStr2(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if strings.Contains(haystack, needle) {
		return strings.Index(haystack, needle)
	}
	return -1
}

/*
35. Search Insert Position
Given a sorted array of distinct integers and a target value, return the index if the target is
found. If not, return the index where it would be if it were inserted in order.
You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1
*/

// Solution 1 O(n) time complexity
func SearchInsert1(nums []int, target int) int {
	i := 0
	for i := range nums {
		if nums[i] == target || target < nums[i] {
			return i
		}
	}
	return i
}

// Solution 2 O(log n) time complexity
func SearchInsert2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// Container With Most Water
// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.
// Notice that you may not slant the container.

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
// In this case, the max area of water (blue section) the container can contain is 49.

func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	var area float64
	for left < right {
		if height[left] < height[right] {
			area = math.Max(area, float64(height[left])*float64(right-left))
			left++
		} else {
			area = math.Max(area, float64(height[right])*float64(right-left))
			right--
		}
	}
	res := int(area)
	return res
}

func ReverseArray(data []int) []int {
	start := 0
	end := len(data) - 1

	for start < end {
		data[start] = data[start] + data[end]
		data[end] = data[start] - data[end]
		data[start] = data[start] - data[end]
		start++
		end--
	}
	return data
}

// Given an integer array nums, rotate the array to the left by k steps,
//
//	where k is non-negative.
func RotateArrayLeft(nums []int, k int) []int {
	start := 0
	end := len(nums) - 1

	for start < end {
		nums[start] = nums[start] + nums[end]
		nums[end] = nums[start] - nums[end]
		nums[start] = nums[start] - nums[end]
		start++
		end--
	}
	return nums
}

// Given an integer array nums, rotate the array to the right by k steps,
//
//	where k is non-negative.
func RotateArrayRight(data []int, k int) []int {
	start := 0
	end := len(data) - 1

	for start < end {
		data[start] = data[start] + data[end]
		data[end] = data[start] - data[end]
		data[start] = data[start] - data[end]
		start++
		end--
	}
	return data
}

// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
/*
Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
*/

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		diff := target - nums[i]
		_, ok := m[diff]
		if ok {
			return []int{m[diff], i}
		}
		m[nums[i]] = i
	}
	return nil
}
