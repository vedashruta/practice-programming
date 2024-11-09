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

	nums := []int{1, 2, 2, 3, 4}
	k := RemoveDuplicates(nums)
	fmt.Println(nums[:k])

}
