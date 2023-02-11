package stack

import "fmt"

var ErrorNoSuchElement = fmt.Errorf("No such element")

func New() *Stack {
	return &Stack{
		arr: make([]int, 0),
	}
}

type Stack struct {
	arr []int
}

func (s Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Push(val int) error {
	s.arr = append(s.arr, val)
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrorNoSuchElement
	}
	l := s.Size() - 1
	val := s.arr[l]
	s.arr = s.arr[:l]
	return val, nil
}

func (s Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrorNoSuchElement
	}
	return s.getLastItem(), nil
}

func (s Stack) getLastItem() int {
	return s.arr[s.Size()-1]
}
