package arrays

import "fmt"

func Input() {
	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// area := MaxArea(height)
	// fmt.Println(area)
	// reversed := ReverseArray(height)
	// fmt.Println(reversed)
	// nums := []int{3, 2, 4}
	// target := 6
	// fmt.Println(TwoSum(nums, target))

	// nums := []int{1, 2, 2, 3, 4}
	// k := RemoveDuplicates(nums)
	// fmt.Println(nums[:k])

	// nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// val := 2
	// k := RemoveElement(nums, val)
	// fmt.Println(nums[:k])
	// haystack := "sadbutsad"
	// needle := "sad"
	// k := StrStr1(haystack, needle)
	// fmt.Println(haystack[k:len(needle)])
	// k = StrStr2(haystack, needle)
	// fmt.Println(haystack[k:len(needle)])
	nums := []int{1, 3, 5, 6}
	target := 5
	k := SearchInsert1(nums, target)
	fmt.Println(k)
	target = 2
	k = SearchInsert1(nums, target)
	fmt.Println(k)
	k = SearchInsert2(nums, target)
	fmt.Println(k)

}
