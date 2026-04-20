package internal

import "fmt"

type Stack[T any] interface {
	Size() int
	Pop() (T, bool)
	Push(value T) 
}

type internalStack[T any] struct {
	items []T
}

func NewStack[T any]() Stack[T] {
	return &internalStack[T]{}
}

func (s *internalStack[T]) String() string {
	return fmt.Sprintf("Stack: %v", s.items)
}

func (s *internalStack[T]) Size() int {
	return len(s.items)
}

func (s *internalStack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *internalStack[T]) Pop() (T, bool) {
	if s.Size() == 0 {
		var zero T
		return zero, false
	}

	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value, true
}