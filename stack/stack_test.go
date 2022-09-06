package stack

import (
	"testing"

	"github.com/mdwhatcott/go-collections/internal/should"
)

func TestStack(t *testing.T) {
	s := New[int](0)
	should.So(t, s.Len(), should.Equal, 0)
	should.So(t, s.Empty(), should.BeTrue)
	s.Push(42)
	s.Push(43)
	s.Push(44)
	should.So(t, s.Len(), should.Equal, 3)
	should.So(t, s.Empty(), should.BeFalse)
	should.So(t, s.Peek(), should.Equal, 44)
	should.So(t, s.Slice(), should.Equal, []int{42, 43, 44})
	p1 := s.Pop()
	p2 := s.Pop()
	p3 := s.Pop()
	should.So(t, s.Len(), should.Equal, 0)
	should.So(t, s.Empty(), should.BeTrue)
	should.So(t, p1, should.Equal, 44)
	should.So(t, p2, should.Equal, 43)
	should.So(t, p3, should.Equal, 42)

	defer func() { recover() }()
	s.Pop()
	t.Error("pop should panic and be recovered (above)")
}
