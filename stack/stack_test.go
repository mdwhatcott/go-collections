package stack

import (
	"testing"

	"github.com/mdwhatcott/go-collections/internal/assert"
)

func TestStack(t *testing.T) {
	s := New[int](0)
	assert.That(t, s.Len()).Equals(0)
	assert.That(t, s.Empty()).IsTrue()
	s.Push(42)
	s.Push(43)
	s.Push(44)
	assert.That(t, s.Len()).Equals(3)
	assert.That(t, s.Empty()).IsFalse()
	assert.That(t, s.Peek()).Equals(44)
	assert.That(t, s.Slice()).Equals([]int{42, 43, 44})
	p1 := s.Pop()
	p2 := s.Pop()
	p3 := s.Pop()
	assert.That(t, s.Len()).Equals(0)
	assert.That(t, s.Empty()).IsTrue()
	assert.That(t, p1).Equals(44)
	assert.That(t, p2).Equals(43)
	assert.That(t, p3).Equals(42)

	defer func() { recover() }()
	s.Pop()
	t.Error("pop should panic and be recovered (above)")
}
