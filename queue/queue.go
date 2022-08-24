package queue

import "github.com/mdwhatcott/go-collections/list"

type Queue[T comparable] struct {
	list *list.List[T]
}

func New[T comparable](size int) *Queue[T] {
	return &Queue[T]{list: list.New[T](size)}
}
func (s *Queue[T]) Len() int    { return s.list.Len() }
func (s *Queue[T]) Empty() bool { return s.Len() == 0 }
func (s *Queue[T]) Enqueue(t T) { s.list.Append(t) }
func (s *Queue[T]) Dequeue() T  { return s.list.PopAt(0) }
func (s *Queue[T]) Peek() T     { return s.list.At(0) }
func (s *Queue[T]) Slice() []T  { return s.list.Slice() }
