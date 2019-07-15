package cal

// Square struct
type Square struct {
	origin int64
	result int64
}

//Calculate func
func (s *Square) Calculate(value int64) {
	s.origin = value
	s.result = value * value
}

//GetResult func
func (s *Square) GetResult() int64 {
	return s.result
}
