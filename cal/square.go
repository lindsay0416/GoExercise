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

//GetOrigin func
func (s *Square) GetOrigin() int64 {
	return s.origin
}

//ToMap func
func (s *Square) ToMap() map[int64]int64 {
	newMap := make(map[int64]int64)
	newMap[s.origin] = s.result
	return newMap
}
