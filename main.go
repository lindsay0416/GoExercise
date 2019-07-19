package main

import (
	"GoExercise/cal"
	"fmt"
	"math/rand"
)

func main() {
	// t1 := &cal.Square{}
	// t2 := &cal.Square{}
	// t3 := &cal.Square{}
	// t1.Calculate(4)
	// t2.Calculate(8)
	// t3.Calculate(16)
	// log.Println(t1.GetResult())
	// log.Println(t2.GetResult())
	// squareSlice := []*cal.Square{t1, t2, t3}
	// squareMap := make(map[int64]int64)

	// for _, s := range squareSlice {
	// 	squareMap[s.GetOrigin()] = s.GetResult()
	// }

	// m := t2.ToMap()
	// fmt.Printf("square of 8 is %v", m[8])

	// t1 := &cal.SquareRoot{}
	// t2 := &cal.SquareRoot{}
	// t1.Calculate(4)
	// t2.Calculate(16)
	// log.Println(t1.GetResult())
	// log.Println(t2.GetResult())
	for i := 1; i < 10; i++ {
		n := rand.Int63n(100-1) + 1
		s := &cal.Square{}
		s.Calculate(int64(n))
		m := s.ToMap()
		fmt.Printf("the square of %v is %v\n", n, m[int64(n)])
	}
}
