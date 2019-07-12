package main

import "log"

func main() {
	t1 := &square{}
	t2 := &square{}
	t1.calculate(4)
	t2.calculate(8)
	log.Println(t1.getResult())
	log.Println(t2.getResult())
}

type square struct {
	origin int64
	result int64
}

func (s *square) calculate(value int64) {
	s.origin = value
	s.result = value * value
}

func (s *square) getResult() int64 {
	return s.result
}
