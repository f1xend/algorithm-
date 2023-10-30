package codewars

import (
	"fmt"
	"math"
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

/*
Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.
The output should be two capital letters with a dot separating them.
It should look like this:
Sam Harris => S.H
patrick feeney => P.F
*/

/*
BEST PRACTIES:

func AbbrevName(name string) string {
    firstLetterToUpper := func(str string) string {
        return strings.ToUpper(string([]rune(str)[0]))
    }

    var names = strings.Fields(name)
    return firstLetterToUpper(names[0]) + "." + firstLetterToUpper(names[1])
}
*/

// my solution This solution fails on Unicode strings, for example "Иван Петров" should yeild "И.П", but yeilds "Ð.Ð" instead.
// To work correctly, this function needs to use runes (see my comment on the top solution for the code)
func AbbrevName(name string) string {

	builder := strings.Builder{}
	n := strings.Index(name, " ") + 1

	builder.WriteString(strings.ToUpper(name[0:1]) + "." + strings.ToUpper(name[n:n+1]))

	return builder.String()
}

/*
Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.
*/

func SumMix(arr []any) int {
	res := 0

	for _, value := range arr {
		switch v := value.(type) {
		case int:
			res += v
		case string:
			number, _ := strconv.Atoi(v)
			res += number
		}
	}
	return res
}

/*
There is a bus moving in the city which takes and drops some people at each bus stop.
You are provided with a list (or array) of integer pairs. Elements of each pair represent the number of people that get on the bus (the first item) and the number of people that get off the bus (the second item) at a bus stop.
Your task is to return the number of people who are still on the bus after the last bus stop (after the last array). Even though it is the last bus stop, the bus might not be empty and some people might still be inside the bus, they are probably sleeping there :D
Take a look on the test cases.
Please keep in mind that the test cases ensure that the number of people in the bus is always >= 0. So the returned integer can't be negative.
The second value in the first pair in the array is 0, since the bus is empty in the first bus stop.
*/

func Number(busStops [][2]int) int {
	res := 0
	for _, v := range busStops {
		res = res + v[0] - v[1]
	}
	return res
}


/*
Implement a function that accepts 3 integer values a, b, c. The function should return true if a triangle can be built with the sides of given length and false in any other case.

(In this case, all triangles must have surface greater than 0 to be accepted).
*/

func IsTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

/*
Count the number of divisors of a positive integer n.

Random tests go up to n = 500000.

Examples (input --> output)
4 --> 3 // we have 3 divisors - 1, 2 and 4
5 --> 2 // we have 2 divisors - 1 and 5
12 --> 6 // we have 6 divisors - 1, 2, 3, 4, 6 and 12
30 --> 8 // we have 8 divisors - 1, 2, 3, 5, 6, 10, 15 and 30
Note you should only return a number, the count of divisors. The numbers between parentheses are shown only for you to see which numbers are counted in each case.
*/

func Divisors(n int)int{
	count := 0
	sqrt := int(math.Sqrt(float64(n)))

	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			count += 2
		}
	}

	if sqrt*sqrt == n {
		count--
	}

	return count
  }