package assert

import (
	"testing"
)


// Panic helps to test panics.
//
// Usage:
//
//     func TestVoidFunction(t *testing.T) {
//         assert.Panic(t, VoidFunction,
// 				"VoidFunction should panic, but it did not.")
//     }
//
//     func TestNonVoid(t *testing.T) {
//         assert.Panic(t, func() { NonVoid(foo) }(),
// 				"NonVoid(foo) should panic, but it did not.")
func Panic(t *testing.T, f func(), message string) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(message)
		}
	}()
	f()
}

// RejectNull helps to test function not accepting nil as parameter.
// See also Panic.
func RejectNull(t *testing.T, f func(), functionName string) {
	Panic(t, f, functionName + "should have panic when taking nil, but it did not.")
}