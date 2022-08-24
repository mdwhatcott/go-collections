package list

import (
	"testing"

	"github.com/mdwhatcott/go-collections/internal/assert"
)

func TestCreation(t *testing.T) {
	assert.That(t, New[int](0).Len()).Equals(0)
	assert.That(t, From[int]().Len()).Equals(0)
	assert.That(t, From[int]().Len()).Equals(0)
}
func TestContains(t *testing.T) {
	assert.That(t, From[int](1, 2, 3).Contains(1)).IsTrue()
	assert.That(t, From[int](2, 3, 4).Contains(1)).IsFalse()
}
func TestAppend(t *testing.T) {
	list := New[int](0)
	list.Append(1, 2, 3, 3)
	assert.That(t, list.Contains(1)).IsTrue()
	assert.That(t, list.Contains(2)).IsTrue()
	assert.That(t, list.Contains(3)).IsTrue()
	assert.That(t, list.Len()).Equals(4)
}
func TestRemove(t *testing.T) {
	list := From[int](1, 2, 3)
	list.Remove(2)
	assert.That(t, list.Contains(1)).IsTrue()
	assert.That(t, list.Contains(2)).IsFalse()
	assert.That(t, list.Contains(3)).IsTrue()
	assert.That(t, list.Len()).Equals(2)
}
func TestClear(t *testing.T) {
	list := From[int](1, 2, 3)
	list.Clear()
	assert.That(t, list.Contains(1)).IsFalse()
	assert.That(t, list.Contains(2)).IsFalse()
	assert.That(t, list.Contains(3)).IsFalse()
	assert.That(t, list.Len()).Equals(0)
}
func TestSlice(t *testing.T) {
	list := From[int](1, 2, 3, 4, 5)
	items := list.Slice()
	assert.That(t, list.Len()).Equals(5)
	assert.That(t, len(items)).Equals(5)
	assert.That(t, items).Equals([]int{1, 2, 3, 4, 5})
}
func TestEqual(t *testing.T) {
	assert.That(t, From[int](1, 2, 3).Equal(From[int](1, 2, 3))).IsTrue()
	assert.That(t, From[int](1, 2, 3).Equal(From[int](3, 2, 1))).IsFalse()
	assert.That(t, From[int](1, 2).Equal(From[int](3, 2, 1))).IsFalse()
	assert.That(t, From[int](1, 2, 2).Equal(From[int](1, 2, 3))).IsFalse()
	assert.That(t, From[int](1, 2, 3).Equal(From[int](1, 2, 4))).IsFalse()
}
func TestUpdate(t *testing.T) {
	list := From[int](1, 2, 3)
	list.Update(1, 42)
	assert.That(t, list.At(0)).Equals(1)
	assert.That(t, list.At(1)).Equals(42)
	assert.That(t, list.At(2)).Equals(3)
}
func TestPopAt(t *testing.T) {
	list := From[int](1, 2, 1)
	pop := list.PopAt(2)
	assert.That(t, pop).Equals(1)
	assert.That(t, list.Slice()).Equals([]int{1, 2})
}
func TestPop(t *testing.T) {
	list := From[int](1, 2, 3)
	pop := list.Pop()
	assert.That(t, pop).Equals(3)
	assert.That(t, list.Slice()).Equals([]int{1, 2})
}
func TestIndex(t *testing.T) {
	list := From[int](1, 2, 3)
	assert.That(t, list.Index(2)).Equals(1)
	assert.That(t, list.Index(4)).Equals(-1)
}
