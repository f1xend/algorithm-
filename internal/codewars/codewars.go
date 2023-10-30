package codewars

import (
	"fmt"
	"strconv"
	"strings"
)

func CheckForFacror(base, factor int) bool {
	// if base%factor == 0 {
	// 	return true
	// }
	return base%factor == 0
}

func ReverseString(word string) string {
	// var result = ""
	// for _, c := range word {
	// 	result = string(c) + resultS
	// 	println(result)
	// }
	// return result
	strArr := []rune(word)

	for i, j := 0, len(strArr)-1; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	return string(strArr)
}

func Grow(arr []int) int {
	summ := 1
	for _, v := range arr {
		summ *= v
	}
	return summ
}

func Greeting() string {
	return "hello world!"
}

func StringToNumber(str string) int {
	res, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}
	return res
}

func Greet(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}

func SetAlarm(employed, vacation bool) bool {
	return !vacation && employed
}

func RepeatStr(repetitions int, value string) string {
	resStr := ""
	for repetitions > 0 {
		resStr += value
		repetitions--
	}
	return resStr
}

func StringToArray(str string) []string {
	return strings.Split(str, " ")
}

func DoubleSliceInt(x []int) []int {
	res := make([]int, (len(x)))
	if len(x) != 0 {
		for i, v := range x {
			res[i] = v * 2
		}
	}
	return res
}

func Hero(bullets, dragons int) bool {
	return bullets-dragons*2 >= 0
}

//String ends with
//Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

func StrEndsWith(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

// TODO: Sum of odd numbers
/*
Given the triangle of consecutive odd numbers:

             1
          3     5
       7     9    11
   13    15    17    19
21    23    25    27    29
...
Calculate the sum of the numbers in the nth row of this triangle (starting at index 1) e.g.: (Input --> Output)

1 -->  1
2 --> 3 + 5 = 8
*/

//Easy to solve this task = n^3 = haha...
func RowSumOddNumbers(n int) int {

	if n == 1 {
		return 1
	}

	start, finish := FindIndexes(n)
	sum := 0
	for i := start; i <= finish; i++ {
		sum = sum + (i * 2) + 1
	}
	return sum
}

func FindIndexes(n int) (start, finish int) {
	if n == 1 {
		return 0, 1
	}
	for i := n - 1; i > 0; i-- {
		start = start + i
	}

	return start, start + n - 1
}
