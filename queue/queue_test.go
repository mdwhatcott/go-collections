package queue

import (
	"testing"

	"github.com/mdwhatcott/go-collections/internal/should"
)

func TestQueue(t *testing.T) {
	s := New[int](0)
	should.So(t, s.Len(), should.Equal, 0)
	should.So(t, s.Empty(), should.BeTrue)
	s.Enqueue(42)
	s.Enqueue(43)
	s.Enqueue(44)
	should.So(t, s.Len(), should.Equal, 3)
	should.So(t, s.Empty(), should.BeFalse)
	should.So(t, s.Peek(), should.Equal, 42)
	should.So(t, s.Slice(), should.Equal, []int{42, 43, 44})
	p1 := s.Dequeue()
	p2 := s.Dequeue()
	p3 := s.Dequeue()
	should.So(t, s.Len(), should.Equal, 0)
	should.So(t, s.Empty(), should.BeTrue)
	should.So(t, p1, should.Equal, 42)
	should.So(t, p2, should.Equal, 43)
	should.So(t, p3, should.Equal, 44)

	defer func() { recover() }()
	s.Dequeue()
	t.Error("pop should panic and be recovered (above)")
}
