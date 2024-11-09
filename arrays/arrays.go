package arrays

import (
	"math"
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
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
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
