package stack

// Stack simplest implementation
type Stack struct {
	data []interface{}
	Size int
}

// Init initializes the stack
func (s *Stack) Init() {
	s.Size = 0
}

// Push pushed element to the stack
func (s *Stack) Push(value interface{}) {
	s.data = append(s.data, value)
	s.Size++
}

// Pop pops an element from the stack
func (s *Stack) Pop() {
	s.data = s.data[:len(s.data)-1]
	s.Size--
}

// Peek returns top element of the stack
func (s *Stack) Peek() (out interface{}) {
	out = s.data[s.Size-1]
	return
}
