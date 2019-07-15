package main

import (
	"GoExercise/cal"
	"log"
)

func main() {
	// t1 := &cal.Square{}
	// t2 := &cal.Square{}
	// t1.Calculate(4)
	// t2.Calculate(8)
	// log.Println(t1.GetResult())
	// log.Println(t2.GetResult())

	t1 := &cal.SquareRoot{}
	t2 := &cal.SquareRoot{}
	t1.Calculate(4)
	t2.Calculate(16)
	log.Println(t1.GetResult())
	log.Println(t2.GetResult())
}
