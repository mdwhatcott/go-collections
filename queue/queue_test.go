package queue

import (
	"testing"

	"github.com/mdwhatcott/go-collections/internal/assert"
)

func TestQueue(t *testing.T) {
	s := New[int](0)
	assert.That(t, s.Len()).Equals(0)
	assert.That(t, s.Empty()).IsTrue()
	s.Enqueue(42)
	s.Enqueue(43)
	s.Enqueue(44)
	assert.That(t, s.Len()).Equals(3)
	assert.That(t, s.Empty()).IsFalse()
	assert.That(t, s.Peek()).Equals(42)
	assert.That(t, s.Slice()).Equals([]int{42, 43, 44})
	p1 := s.Dequeue()
	p2 := s.Dequeue()
	p3 := s.Dequeue()
	assert.That(t, s.Len()).Equals(0)
	assert.That(t, s.Empty()).IsTrue()
	assert.That(t, p1).Equals(42)
	assert.That(t, p2).Equals(43)
	assert.That(t, p3).Equals(44)

	defer func() { recover() }()
	s.Dequeue()
	t.Error("pop should panic and be recovered (above)")
}
