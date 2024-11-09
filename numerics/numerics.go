package numerics

// Palindrome Number
// Given an integer x, return true if x is a
// palindrome and false otherwise.
/*
Input: x = 121
Output: true
*/

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	temp := x
	rev := 0
	for temp > 0 {
		remainder := temp % 10
		rev = rev*10 + remainder
		temp = temp / 10
	}
	return x == rev
}
