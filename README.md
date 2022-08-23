# github.com/mdwhatcott/go-collections



	package list // import "container/list"
	
	Package list implements a doubly linked list.
	
	To iterate over a list (where l is a *List):
	
	    for e := l.Front(); e != nil; e = e.Next() {
	    	// do something with e.Value
	    }
	
	TYPES
	
	type Element struct {
	
		// The value stored with this element.
		Value any
		// Has unexported fields.
	}
	    Element is an element of a linked list.
	
	func (e *Element) Next() *Element
	    Next returns the next list element or nil.
	
	func (e *Element) Prev() *Element
	    Prev returns the previous list element or nil.
	
	type List struct {
		// Has unexported fields.
	}
	    List represents a doubly linked list. The zero value for List is an empty
	    list ready to use.
	
	func New() *List
	    New returns an initialized list.
	
	func (l *List) Back() *Element
	    Back returns the last element of list l or nil if the list is empty.
	
	func (l *List) Front() *Element
	    Front returns the first element of list l or nil if the list is empty.
	
	func (l *List) Init() *List
	    Init initializes or clears list l.
	
	func (l *List) InsertAfter(v any, mark *Element) *Element
	    InsertAfter inserts a new element e with value v immediately after mark
	    and returns e. If mark is not an element of l, the list is not modified.
	    The mark must not be nil.
	
	func (l *List) InsertBefore(v any, mark *Element) *Element
	    InsertBefore inserts a new element e with value v immediately before mark
	    and returns e. If mark is not an element of l, the list is not modified.
	    The mark must not be nil.
	
	func (l *List) Len() int
	    Len returns the number of elements of list l. The complexity is O(1).
	
	func (l *List) MoveAfter(e, mark *Element)
	    MoveAfter moves element e to its new position after mark. If e or mark is
	    not an element of l, or e == mark, the list is not modified. The element and
	    mark must not be nil.
	
	func (l *List) MoveBefore(e, mark *Element)
	    MoveBefore moves element e to its new position before mark. If e or mark is
	    not an element of l, or e == mark, the list is not modified. The element and
	    mark must not be nil.
	
	func (l *List) MoveToBack(e *Element)
	    MoveToBack moves element e to the back of list l. If e is not an element of
	    l, the list is not modified. The element must not be nil.
	
	func (l *List) MoveToFront(e *Element)
	    MoveToFront moves element e to the front of list l. If e is not an element
	    of l, the list is not modified. The element must not be nil.
	
	func (l *List) PushBack(v any) *Element
	    PushBack inserts a new element e with value v at the back of list l and
	    returns e.
	
	func (l *List) PushBackList(other *List)
	    PushBackList inserts a copy of another list at the back of list l. The lists
	    l and other may be the same. They must not be nil.
	
	func (l *List) PushFront(v any) *Element
	    PushFront inserts a new element e with value v at the front of list l and
	    returns e.
	
	func (l *List) PushFrontList(other *List)
	    PushFrontList inserts a copy of another list at the front of list l.
	    The lists l and other may be the same. They must not be nil.
	
	func (l *List) Remove(e *Element) any
	    Remove removes e from l if e is an element of list l. It returns the element
	    value e.Value. The element must not be nil.
	
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

