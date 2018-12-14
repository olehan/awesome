package stack

type (
	Stack interface {
		Push(v interface{}) Stack
		Pop() interface{}
		Peek() interface{}
		Count() int
	}

	stack []interface{}
)

func NewStack() Stack {
	return &stack{}
}

func (s *stack) Push(v interface{}) Stack {
	*s = append([]interface{}{v}, *s...)
	return s
}

func (s *stack) Pop() interface{} {
	last := (*s)[0]
	*s = (*s)[1:]
	return last
}

func (s *stack) Peek() interface{} {
	return (*s)[0]
}

func (s *stack) Count() int {
	return len(*s)
}
