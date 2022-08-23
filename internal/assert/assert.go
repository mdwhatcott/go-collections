package assert

import (
	"reflect"
	"testing"
)

// That allows assertions as in: assert.That(t, actual).Equals(expected)

func That(t *testing.T, actual interface{}) *assertion {
	return &assertion{t: t, actual: actual}
}

type assertion struct {
	t      *testing.T
	actual interface{}
}

func (a *assertion) IsNil()   { a.t.Helper(); a.Equals(nil) }
func (a *assertion) IsTrue()  { a.t.Helper(); a.Equals(true) }
func (a *assertion) IsFalse() { a.t.Helper(); a.Equals(false) }
func (a *assertion) Equals(expected interface{}) {
	a.t.Helper()
	if !reflect.DeepEqual(a.actual, expected) {
		a.t.Errorf("\n"+
			"Expected: %#v\n"+
			"Actual:   %#v",
			expected,
			a.actual,
		)
	}
}
