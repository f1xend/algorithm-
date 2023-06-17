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
