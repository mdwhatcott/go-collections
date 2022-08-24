// Package list implements a generic array list type.
// The API was inspired by Python's built-in list.
// https://en.wikipedia.org/wiki/Dynamic_array
// https://docs.python.org/3/tutorial/introduction.html#lists
// https://docs.python.org/3/tutorial/datastructures.html#more-on-lists
package list

type List[T comparable] struct{ items []T }

func New[T comparable](size int) *List[T] {
	return &List[T]{items: make([]T, 0, size)}
}
func From[T comparable](items ...T) *List[T] {
	return &List[T]{items: items}
}
func (l *List[T]) Len() int {
	return len(l.items)
}
func (l *List[T]) Contains(needle T) bool {
	return l.Index(needle) >= 0
}
func (l *List[T]) Append(items ...T) {
	l.items = append(l.items, items...)
}
func (l *List[T]) Remove(needle T) {
	l.remove(l.Index(needle))
}
func (l *List[T]) Clear() {
	var zero T
	for i := range l.items {
		l.items[i] = zero
	}
	l.items = l.items[0:0]
}
func (l *List[T]) Slice() (results []T) {
	results = make([]T, 0, len(l.items))
	for _, item := range l.items {
		results = append(results, item)
	}
	return results
}
func (l *List[T]) Equal(o *List[T]) bool {
	if l.Len() != o.Len() {
		return false
	}
	for x := 0; x < l.Len(); x++ {
		if l.At(x) != o.At(x) {
			return false
		}
	}
	return true
}
func (l *List[T]) At(i int) T {
	return l.items[i]
}
func (l *List[T]) Update(i int, item T) {
	l.items[i] = item
}
func (l *List[T]) PopAt(i int) T {
	pop := l.items[i]
	defer l.remove(i)
	return pop
}
func (l *List[T]) Pop() T {
	return l.PopAt(l.Len() - 1)
}
func (l *List[T]) Index(needle T) int {
	for i, item := range l.items {
		if item == needle {
			return i
		}
	}
	return -1
}
func (l *List[T]) remove(i int) {
	var zero T
	l.items[i] = zero
	l.items = append(l.items[:i], l.items[i+1:]...)
}
