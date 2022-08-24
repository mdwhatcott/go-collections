package stack

import "github.com/mdwhatcott/go-collections/list"

type Stack[T comparable] struct {
	list *list.List[T]
}

func New[T comparable](size int) *Stack[T] {
	return &Stack[T]{list: list.New[T](size)}
}
func (s *Stack[T]) Len() int    { return s.list.Len() }
func (s *Stack[T]) Empty() bool { return s.Len() == 0 }
func (s *Stack[T]) Push(t T)    { s.list.Append(t) }
func (s *Stack[T]) Pop() T      { return s.list.Pop() }
func (s *Stack[T]) Peek() T     { return s.list.At(s.Len() - 1) }
func (s *Stack[T]) Slice() []T  { return s.list.Slice() }
