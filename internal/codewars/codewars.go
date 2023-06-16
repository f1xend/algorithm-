package codewars

import (
	"strconv"
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

func Greet() string {
	return "hello world!"
}

func StringToNumber(str string) int {
	res, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}
	return res
}
