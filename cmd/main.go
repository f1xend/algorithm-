package main

import (
	intersectarr "algorithm/internal/leetcode/intersect_arr"
	"fmt"
)

func main() {

	fmt.Println(intersectarr.Intersect([]int{2, 1}, []int{1, 1}))
}

// import (
// 	"algorithm/internal/codewars"
// 	yandexp "algorithm/internal/yandexP"
// 	"fmt"
// 	"time"
// )

// type MyStruct struct {
// 	A   int
// 	Log func(s string)
// }

// func main() {
// 	fmt.Println("OK, let's try")

// 	if err := yandexp.HenleDeliveryStatus(yandexp.DeliveryState("pending")); err != nil {
// 		panic(err)
// 	}

// 	buf := yandexp.NewCircularBuffer(4)
// 	yandexp.Handle(1.0, buf.AddValue)
// 	yandexp.Handle(2.0, buf.AddValue)
// 	yandexp.Handle(3.0, buf.AddValue)
// 	yandexp.Handle(4.0, buf.AddValue)
// 	fmt.Printf("[%d]: %v\n", buf.GetCurrentSize(), buf.GetValues())

// 	buf.ForceSetValueByIdx(0, -1.0)
// 	buf.ForceSetValueByIdx(1, -2.0)
// 	// fmt.Println(buf.Values)

// 	for i := 0; i < 6; i++ {
// 		if i > 0 {
// 			buf.AddValue(float64(i))
// 		}
// 		// fmt.Printf("[%d]: %v\n", buf.GetCurrentSize(), buf.GetValues())
// 	}

// 	// s := make(yandexp.IntSlice, 0)
// 	// s.Add(1)
// 	// s.Add(2)
// 	// fmt.Println(s)

// 	var s = MyStruct{
// 		A:   1,
// 		Log: func(s string) { fmt.Println(s) },
// 	}

// 	s.Log("some string")

// 	fmt.Println(codewars.Greet("Ruslan"))

// 	sw := yandexp.Stopwatch{}
// 	sw.Start()

// 	time.Sleep(1 * time.Second)
// 	sw.SaveSplit()

// 	time.Sleep(500 * time.Millisecond)
// 	sw.SaveSplit()

// 	time.Sleep(300 * time.Millisecond)
// 	sw.SaveSplit()

// 	fmt.Println(sw.GetResults())
// 	// fmt.Println(sw.Current, sw.DurrList)
// }
