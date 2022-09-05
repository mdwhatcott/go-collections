# github.com/mdwhatcott/go-collections



	package list // import "github.com/mdwhatcott/go-collections/list"
	
	Package list implements a generic array list type. The API was inspired
	by Python's built-in list. https://en.wikipedia.org/wiki/Dynamic_array
	https://docs.python.org/3/tutorial/introduction.html#lists
	https://docs.python.org/3/tutorial/datastructures.html#more-on-lists
	
	TYPES
	
	type List[T comparable] struct {
		// Has unexported fields.
	}
	
	func From[T comparable](items ...T) *List[T]
	
	func New[T comparable](size int) *List[T]
	
	func (l *List[T]) Append(items ...T)
	
	func (l *List[T]) At(i int) T
	
	func (l *List[T]) Clear()
	
	func (l *List[T]) Contains(needle T) bool
	
	func (l *List[T]) Empty() bool
	
	func (l *List[T]) Equal(o *List[T]) bool
	
	func (l *List[T]) Index(needle T) int
	
	func (l *List[T]) Len() int
	
	func (l *List[T]) Pop() T
	
	func (l *List[T]) PopAt(i int) T
	
	func (l *List[T]) Remove(needle T)
	
	func (l *List[T]) Slice() (results []T)
	
	func (l *List[T]) Update(i int, item T)
	
---

	package queue // import "github.com/mdwhatcott/go-collections/queue"
	
	
	TYPES
	
	type Queue[T comparable] struct {
		// Has unexported fields.
	}
	
	func New[T comparable](size int) *Queue[T]
	
	func (s *Queue[T]) Dequeue() T
	
	func (s *Queue[T]) Empty() bool
	
	func (s *Queue[T]) Enqueue(t T)
	
	func (s *Queue[T]) Len() int
	
	func (s *Queue[T]) Peek() T
	
	func (s *Queue[T]) Slice() []T
	
---

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
	
	func (s Set[T]) Empty() bool
	
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

	package stack // import "github.com/mdwhatcott/go-collections/stack"
	
	
	TYPES
	
	type Stack[T comparable] struct {
		// Has unexported fields.
	}
	
	func New[T comparable](size int) *Stack[T]
	
	func (s *Stack[T]) Empty() bool
	
	func (s *Stack[T]) Len() int
	
	func (s *Stack[T]) Peek() T
	
	func (s *Stack[T]) Pop() T
	
	func (s *Stack[T]) Push(t T)
	
	func (s *Stack[T]) Slice() []T
	
---

