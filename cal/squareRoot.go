package cal

import "math"

type SquareRoot struct {
	origin float64
	result float64
}

func (s *SquareRoot) Calculate(value float64) {
	s.origin = value
	s.result = math.Sqrt(value)

}

func (s *SquareRoot) GetResult() float64 {
	return s.result
}
