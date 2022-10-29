// Package set implements a generic set type. Finally!
// https://en.wikipedia.org/wiki/Set_(mathematics)
package set

type Set[T comparable] map[T]nothing

func New[T comparable](size int) Set[T] {
	return make(Set[T], size)
}
func From[T comparable](items ...T) (result Set[T]) {
	result = New[T](len(items))
	result.Add(items...)
	return result
}
func FromMapKeys[K comparable, V any](m map[K]V) (result Set[K]) {
	result = New[K](len(m))
	for key := range m {
		result.Add(key)
	}
	return result
}
func (s Set[T]) Empty() bool {
	return s.Len() == 0
}
func (s Set[T]) Len() int {
	return len(s)
}
func (s Set[T]) Slice() (result []T) {
	result = make([]T, 0, len(s))
	for item := range s {
		result = append(result, item)
	}
	return result
}
func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = nothing{}
	}
}
func (s Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}
func (s Set[T]) Clear() {
	for item := range s {
		delete(s, item)
	}
}
func (s Set[T]) Contains(item T) bool {
	_, found := s[item]
	return found
}
func (s Set[T]) Equal(that Set[T]) bool {
	if len(s) != len(that) {
		return false
	}
	for item := range s {
		if !that.Contains(item) {
			return false
		}
	}
	return true
}
func (s Set[T]) IsSubset(that Set[T]) bool {
	for item := range s {
		if !that.Contains(item) {
			return false
		}
	}
	return true
}
func (s Set[T]) IsSuperset(that Set[T]) bool {
	for item := range that {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}
func (s Set[T]) Union(that Set[T]) (result Set[T]) {
	result = make(Set[T])
	for item := range s {
		result.Add(item)
	}
	for item := range that {
		result.Add(item)
	}
	return result
}
func (s Set[T]) Intersection(that Set[T]) (result Set[T]) {
	result = make(Set[T])
	for item := range s {
		if that.Contains(item) {
			result.Add(item)
		}
	}
	for item := range that {
		if s.Contains(item) {
			result.Add(item)
		}
	}
	return result
}
func (s Set[T]) Difference(that Set[T]) (result Set[T]) {
	result = make(Set[T])
	for item := range s {
		if !that.Contains(item) {
			result.Add(item)
		}
	}
	return result
}
func (s Set[T]) SymmetricDifference(that Set[T]) (result Set[T]) {
	result = make(Set[T])
	for item := range s {
		if !that.Contains(item) {
			result.Add(item)
		}
	}
	for item := range that {
		if !s.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

type nothing struct{}
