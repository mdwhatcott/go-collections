# github.com/mdwhatcott/go-collections



	package set // import "github.com/mdwhatcott/go-collections/set"
	
	Package set implements a generic set type. Finally!
	https://en.wikipedia.org/wiki/Set_(mathematics)
	
	TYPES
	
	type Set[T comparable] map[T]nothing
	
	func From[T comparable](items ...T) (result Set[T])
	
	func New[T comparable](size int) Set[T]
	
	func (s Set[T]) Add(items ...T)
	
	func (s Set[T]) Clear()
	
	func (s Set[T]) Contains(item T) bool
	
	func (s Set[T]) Difference(that Set[T]) (result Set[T])
	
	func (s Set[T]) Equal(that Set[T]) bool
	
	func (s Set[T]) Intersection(that Set[T]) (result Set[T])
	
	func (s Set[T]) IsSubset(that Set[T]) bool
	
	func (s Set[T]) IsSuperset(that Set[T]) bool
	
	func (s Set[T]) Len() int
	
	func (s Set[T]) Remove(items ...T)
	
	func (s Set[T]) Slice() (result []T)
	
	func (s Set[T]) SymmetricDifference(that Set[T]) (result Set[T])
	
	func (s Set[T]) Union(that Set[T]) (result Set[T])
	
---

