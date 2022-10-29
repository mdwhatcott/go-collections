package set

import (
	"sort"
	"testing"

	"github.com/mdwhatcott/go-collections/internal/should"
)

func TestCreation(t *testing.T) {
	should.So(t, len(New[int](0)), should.Equal, 0)
	should.So(t, len(From[int]()), should.Equal, 0)
	should.So(t, From[int]().Len(), should.Equal, 0)
	should.So(t, From[int]().Empty(), should.BeTrue)
	should.So(t, From[int](1).Empty(), should.BeFalse)
	should.So(t, FromMapKeys[int](map[int]int{}).Len(), should.Equal, 0)
	should.So(t, FromMapKeys[int](map[int]int{1: 4, 2: 5, 3: 6}), should.Equal, From[int](1, 2, 3))
}
func TestContains(t *testing.T) {
	should.So(t, From[int](1).Contains(1), should.BeTrue)
	should.So(t, From[int]().Contains(1), should.BeFalse)
}
func TestAdd(t *testing.T) {
	set := New[int](0)
	set.Add(1, 2, 3)
	should.So(t, set.Contains(1), should.BeTrue)
	should.So(t, set.Contains(2), should.BeTrue)
	should.So(t, set.Contains(3), should.BeTrue)
	should.So(t, set.Len(), should.Equal, 3)
}
func TestRemove(t *testing.T) {
	set := From[int](1, 2, 3)
	set.Remove(2)
	should.So(t, set.Contains(1), should.BeTrue)
	should.So(t, set.Contains(2), should.BeFalse)
	should.So(t, set.Contains(3), should.BeTrue)
	should.So(t, set.Len(), should.Equal, 2)
}
func TestClear(t *testing.T) {
	set := From[int](1, 2, 3)
	set.Clear()
	should.So(t, set.Contains(1), should.BeFalse)
	should.So(t, set.Contains(2), should.BeFalse)
	should.So(t, set.Contains(3), should.BeFalse)
	should.So(t, set.Len(), should.Equal, 0)
}
func TestSlice(t *testing.T) {
	set := From[int](1, 2, 3, 4, 5)
	items := set.Slice()
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	should.So(t, set.Len(), should.Equal, 5)
	should.So(t, len(items), should.Equal, 5)
	should.So(t, items, should.Equal, []int{1, 2, 3, 4, 5})
}
func TestEqual(t *testing.T) {
	should.So(t, From[int](1, 2, 3).Equal(From[int](3, 2, 1)), should.BeTrue)
	should.So(t, From[int](1, 2).Equal(From[int](3, 2, 1)), should.BeFalse)
	should.So(t, From[int](1, 2, 2).Equal(From[int](1, 2, 3)), should.BeFalse)
	should.So(t, From[int](1, 2, 3).Equal(From[int](1, 2, 4)), should.BeFalse)
}
func TestIsSubset(t *testing.T) {
	should.So(t, From[int](1, 2, 3).IsSubset(From[int](1, 2, 3, 4, 5)), should.BeTrue)
	should.So(t, From[int](4, 5, 6).IsSubset(From[int](1, 2, 3, 4, 5)), should.BeFalse)
}
func TestIsSuperset(t *testing.T) {
	should.So(t, From[int](1, 2, 3, 4, 5).IsSuperset(From[int](1, 2, 3)), should.BeTrue)
	should.So(t, From[int](1, 2, 3, 4, 5).IsSuperset(From[int](4, 5, 6)), should.BeFalse)
}
func TestUnion(t *testing.T) {
	should.So(t, From[int](1, 2, 3).Union(From[int](1, 2, 3)), should.Equal, From[int](1, 2, 3))
	should.So(t, From[int](1, 2, 3).Union(From[int](2, 3, 4)), should.Equal, From[int](1, 2, 3, 4))
	should.So(t, From[int](1, 2, 3).Union(From[int](4, 5, 6)), should.Equal, From[int](1, 2, 3, 4, 5, 6))
}
func TestIntersection(t *testing.T) {
	should.So(t, From[int](1, 2, 3).Intersection(From[int](4, 5, 6)), should.Equal, From[int]())
	should.So(t, From[int](1, 2, 3).Intersection(From[int](2, 3, 4)), should.Equal, From[int](2, 3))
}
func TestDifference(t *testing.T) {
	should.So(t, From[int](1, 2, 3).Difference(From[int](4, 5, 6)), should.Equal, From[int](1, 2, 3))
	should.So(t, From[int](1, 2, 3).Difference(From[int](2, 3)), should.Equal, From[int](1))
}
func TestSymmetricDifference(t *testing.T) {
	should.So(t, From[int](1, 2, 3).SymmetricDifference(From[int](4, 5, 6)), should.Equal, From[int](1, 2, 3, 4, 5, 6))
	should.So(t, From[int](1, 2, 3).SymmetricDifference(From[int](2, 3, 4)), should.Equal, From[int](1, 4))
}
