package list

import (
	"testing"

	"github.com/mdwhatcott/go-collections/internal/should"
)

func TestCreation(t *testing.T) {
	should.So(t, New[int](0).Len(), should.Equal, 0)
	should.So(t, From[int]().Len(), should.Equal, 0)
	should.So(t, From[int]().Empty(), should.BeTrue)
	should.So(t, From[int](1).Empty(), should.BeFalse)
}
func TestContains(t *testing.T) {
	should.So(t, From[int](1, 2, 3).Contains(1), should.BeTrue)
	should.So(t, From[int](2, 3, 4).Contains(1), should.BeFalse)
}
func TestAppend(t *testing.T) {
	list := New[int](0)
	list.Append(1, 2, 3, 3)
	should.So(t, list.Contains(1), should.BeTrue)
	should.So(t, list.Contains(2), should.BeTrue)
	should.So(t, list.Contains(3), should.BeTrue)
	should.So(t, list.Len(), should.Equal, 4)
}
func TestRemove(t *testing.T) {
	list := From[int](1, 2, 3)
	list.Remove(2)
	should.So(t, list.Contains(1), should.BeTrue)
	should.So(t, list.Contains(2), should.BeFalse)
	should.So(t, list.Contains(3), should.BeTrue)
	should.So(t, list.Len(), should.Equal, 2)
}
func TestClear(t *testing.T) {
	list := From[int](1, 2, 3)
	list.Clear()
	should.So(t, list.Contains(1), should.BeFalse)
	should.So(t, list.Contains(2), should.BeFalse)
	should.So(t, list.Contains(3), should.BeFalse)
	should.So(t, list.Len(), should.Equal, 0)
}
func TestSlice(t *testing.T) {
	list := From[int](1, 2, 3, 4, 5)
	items := list.Slice()
	should.So(t, list.Len(), should.Equal, 5)
	should.So(t, len(items), should.Equal, 5)
	should.So(t, items, should.Equal, []int{1, 2, 3, 4, 5})
}
func TestEqual(t *testing.T) {
	should.So(t, From[int](1, 2, 3).Equal(From[int](1, 2, 3)), should.BeTrue)
	should.So(t, From[int](1, 2, 3).Equal(From[int](3, 2, 1)), should.BeFalse)
	should.So(t, From[int](1, 2).Equal(From[int](3, 2, 1)), should.BeFalse)
	should.So(t, From[int](1, 2, 2).Equal(From[int](1, 2, 3)), should.BeFalse)
	should.So(t, From[int](1, 2, 3).Equal(From[int](1, 2, 4)), should.BeFalse)
}
func TestUpdate(t *testing.T) {
	list := From[int](1, 2, 3)
	list.Update(1, 42)
	should.So(t, list.At(0), should.Equal, 1)
	should.So(t, list.At(1), should.Equal, 42)
	should.So(t, list.At(2), should.Equal, 3)
}
func TestPopAt(t *testing.T) {
	list := From[int](1, 2, 1)
	pop := list.PopAt(2)
	should.So(t, pop, should.Equal, 1)
	should.So(t, list.Slice(), should.Equal, []int{1, 2})
}
func TestPop(t *testing.T) {
	list := From[int](1, 2, 3)
	pop := list.Pop()
	should.So(t, pop, should.Equal, 3)
	should.So(t, list.Slice(), should.Equal, []int{1, 2})
}
func TestIndex(t *testing.T) {
	list := From[int](1, 2, 3)
	should.So(t, list.Index(2), should.Equal, 1)
	should.So(t, list.Index(4), should.Equal, -1)
}
