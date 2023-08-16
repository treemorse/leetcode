package golang

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.queue = append(s.queue, x)
}

func (s *MyStack) Pop() int {
	if s.Empty() {
		return 0
	}

	elem := s.queue[len(s.queue)-1]

	s.queue = s.queue[:len(s.queue)-1]

	return elem
}

func (s *MyStack) Top() int {
	if s.Empty() {
		return 0
	}

	return s.queue[len(s.queue)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}
