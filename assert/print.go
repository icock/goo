package assert

import "testing"

func PrintFailure(t *testing.T, i int, actual interface{}, expected interface{}) {
	t.Errorf(
		"#%d FAIL:\n  Actual Value: %v\nExpected Value: %v",
		i, actual, expected)
}
