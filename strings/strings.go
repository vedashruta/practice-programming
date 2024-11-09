package strings

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// The following operations can be performed on a string in go
// 1. Contains
// 2. Count
// 3. HasPrefix
// 4. HasSuffix
// 5. Index
// 6. Join
// 7. Repeat
// 8. Replace
// 9. ToLower
// 10. ToUpper
func BasicOperations() {
	p := fmt.Println

	p(strings.Contains("test", "t"))
	p(strings.Count("test", "t"))
	p(strings.HasPrefix("test", "s"))
	p(strings.HasSuffix("test", "e"))
	p(strings.Index("test", "e"))
	p(strings.Join([]string{"t", "e", "s", "t"}, ""))
	p(strings.Repeat("t", 5))
	p(strings.Replace("test", "t", "v", -1))
	p(strings.Replace("test", "t", "v", 0))
	p(strings.Replace("test", "t", "v", 1))
	p(strings.Split("test", ""))

	p(strings.ToUpper("teSt"))
	p(strings.ToLower("TEst"))
	p(strings.Trim("  test ", " "))
	p(strings.TrimLeft("  test ", " "))
	p(strings.TrimRight("  test ", "  "))
}

// Logic 1 :
// func Reverse(str string) (res string) {
// 	length := len(str) - 1
// 	for i := range str {
// 		res += fmt.Sprintf("%[1]s", string(str[length-i]))
// 	}
// 	return
// }

// Logic 2: Using rune
// func Reverse(str string) (res string) {
// 	runes := []rune(str)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	res = string(runes)
// 	return
// }

// Logic 3: Using stringBuilder
// [Optimized]
func Reverse(str string) (res string) {
	var sb strings.Builder
	sb.Grow(len(str))
	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteByte(str[i])
	}
	res = sb.String()
	return
}

// Returns the rune in the specified index of a string,
// only if index>=len(str)-1
func GetChar(str string, index int) (char rune) {
	if index <= len(str)-1 {
		char = rune(str[index])
	}
	return
}

// Returns true if str is a palindrome
// Logic 1 : Implement a reverse function
// func IsPalindrome(str string) bool {
// 	var sb strings.Builder
// 	sb.Grow(len(str))
// 	for i := len(str) - 1; i >= 0; i-- {
// 		sb.WriteByte(str[i])
// 	}
// 	res := strings.Compare(str, sb.String())
// 	return res == 0
// }

// Logic 2:
func IsPalindrome(str string) bool {
	var sb strings.Builder
	length := len(str)
	sb.Grow(length)
	for i := length - 1; i >= length/2; i-- {
		sb.WriteByte(str[i])
	}
	bytes := []byte(str)
	index := 0
	switch length % 2 {
	case 0:
		index = (length / 2)
	default:
		index = (length / 2) + 1

	}
	res := strings.Compare(string(bytes[:index]), sb.String())
	return res == 0
}

// Seven different symbols represent Roman numerals with the following values:
// I	1
// V	5
// X	10
// L	50
// C	100
// D	500
// M	1000
/*
Constraints:
1 <= num <= 3999
*/
func IntToRoman(number int) string {
	if number < 0 || number > 3999 {
		return ""
	}
	sb := &strings.Builder{}
	for number > 0 {
		remainder := number % 10
		if (remainder > 0 && remainder < 4) || (remainder > 5 && remainder < 9) {
			sb.WriteString(getRoman(remainder))
		}
		if remainder == 4 {
			sb.WriteString("I" + getRoman(remainder))
		}
		if remainder == 9 {
			sb.WriteString("I" + getRoman(remainder))
		}
		n := remainder
		for n > 0 {
			n = n - remainder
			if (n > 0 && n < 4) || (n > 5 && n < 9) {
				sb.WriteString(getRoman(n))
			}
			if n == 4 {
				sb.WriteString("I" + getRoman(n))
			}
			if n == 9 {
				sb.WriteString("I" + getRoman(n))
			}
		}
		number /= 10
	}
	return sb.String()
}

func getRoman(number int) string {
	switch number {
	case 1:
		return "I"
	case 5:
		return "V"
	case 10:
		return "X"
	case 50:
		return "L"
	case 100:
		return "C"
	case 500:
		return "D"
	case 1000:
		return "M"
	default:
		return ""
	}
}

/*
	Regular Expressions
*/

func RegexOperations() {
	ok, err := regexp.MatchString("([a-zA-Z]+)ch", "peAch")
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		fmt.Println("String matched")
	}
	regex, err := regexp.Compile(`([a-zA-Z]+.pdf)"`)
	if err != nil {
		log.Fatal(err)
	}
	ok = regex.MatchString("abc.pdf")
	if ok {
		fmt.Println("String matched")
	}
}

/*
	Regular Expressions
	1. If the string contains "key" extract its value until u reach ","
	2. Extract urls from a ssample jumbled string
*/
/*
	13. Roman to Integer
		Symbol       Value
		I             1
		V             5
		X             10
		L             50
		C             100
		D             500
		M             1000
	Roman numerals are usually written largest to smallest from left to right.
	However, the numeral for four is not IIII. Instead, the number four is written as IV.
	Because the one is before the five we subtract it making four.
	The same principle applies to the number nine, which is written as IX.
*/
func RomanToInteger(roman string) int {
	m := make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000
	return 0
}

/*
	14. Longest Common Prefix
	Write a function to find the longest common prefix string amongst an array of strings.
	If there is no common prefix, return an empty string "".
	Example 1:

	Input: strs = ["flower","flow","flight"]
	Output: "fl"
	Example 2:

	Input: strs = ["dog","racecar","car"]
	Output: ""
	Explanation: There is no common prefix among the input strings.
*/

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	first := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], first) {
			first = first[0 : len(first)-1]
		}
	}
	return first
}
