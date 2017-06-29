package slice

import (
	"testing"
	"notabug.org/icock/goo/assert"
)


var nullSlice []string = nil
var rejectsNullTests = []struct {
	this []string
	that []string
} {
	{nullSlice, nullSlice},
	{nullSlice, []string{}},
	{[]string{"s"}, nullSlice},
}


func TestOfStringIsNotNull(t *testing.T) {
	OfStringNotNull([]string{"non empty slice"})
	OfStringNotNull([]string{}) // empty slice
	assert.Panic(t, func() { OfStringNotNull(nullSlice) },
		"FAIL OfStringNotNull(nullSlice) should panic, but it did not. ")
}

func TestOfStringIsEmptyRejectsNull(t *testing.T) {
	assert.RejectNull(t, func() { OfStringIsEmpty(nullSlice) }, "OfStringIsEmpty")
}
var ofStringIsEmptyTests = []struct {
	in []string
	out bool
}{
	{[]string{}, true},
	{[]string{"false"}, false},
	{make([]string, 0), true},
	{make([]string, 0, 10), true},
	{make([]string, 10), false},
}
func TestOfStringIsEmpty(t *testing.T) {
	for i, tt := range ofStringIsEmptyTests {
		if out := OfStringIsEmpty(tt.in); out != tt.out {
			assert.PrintFailure(t, i, out, tt.out)
		}
	}
}

func TestOfStringEqualsRejectsNull(t *testing.T) {
	for _, tt := range rejectsNullTests {
		assert.RejectNull(
			t,
			func() { OfStringEquals(tt.this, tt.that) },
			"OfStringEquals")
	}
}
var ofStringEqualsTests = []struct {
	this []string
	that []string
	out bool
}{
	{[]string{"s"}, []string{"s"}, true},
	{[]string{"s"}, []string{"s", "t"}, false},
	{[]string{"s", "t"}, []string{"s", "t"}, true},
	{[]string{"t", "s"}, []string{"s", "t"}, false},
	{[]string{}, []string{}, true},
	{[]string{}, []string{"s"}, false},
	{[]string{"s"}, []string{}, false},
}
func TestOfStringEquals(t *testing.T) {
	for i, tt := range ofStringEqualsTests {
		if out := OfStringEquals(tt.this, tt.that); out != tt.out {
			assert.PrintFailure(t, i, out, tt.out)
		}
	}
}

func TestOfStringSizeRejectsNull(t *testing.T) {
	assert.RejectNull(t, func() { OfStringSize(nullSlice) }, "OfStringSize")
}
var ofStringSizeTests = []struct {
	in []string
	out int
}{
	{[]string{}, 0},
	{[]string{""}, 1},
	{[]string{"a", "b", "c"}, 3},
}
func TestOfStringSize(t *testing.T) {
	for i, tt := range ofStringSizeTests {
		if out := OfStringSize(tt.in); out != tt.out {
			assert.PrintFailure(t, i, out, tt.out)
		}
	}
}

func TestAddStringRejectsNull(t *testing.T) {
	assert.RejectNull(
		t,
		func() { AddString(nullSlice, "will panic") },
		"AddString")
}
func TestAddString(t *testing.T) {
	expected := []string{"s", "str"}
	if result := AddString(AddString([]string{}, "s"), "str"); !OfStringEquals(result, expected) {
		t.Errorf("AddString FAIL:\n  Actual Value: %v\nExpected Value: %v",
		result, expected)
	}
}


func TestAddAllStringsRejectsNull(t *testing.T) {
	for _, tt := range rejectsNullTests {
		assert.RejectNull(
			t,
			func() { AddAllStrings(tt.this, tt.that) },
			"AddAllStrings")
	}
}
var addAllStringsTests = []struct {
	this []string
	that []string
	out []string
}{
	{[]string{}, []string{}, []string{}},
	{[]string{}, []string{"s", "abc"}, []string{"s", "abc"}},
	{[]string{"s", "abc"}, []string{},[]string{"s", "abc"}},
	{[]string{"s", "abc"}, []string{"abc", "s"}, []string{"s", "abc", "abc", "s"}},
}
func TestAddAllStrings(t *testing.T) {
	for i, tt := range addAllStringsTests {
		if out := AddAllStrings(tt.this, tt.that); !OfStringEquals(out, tt.out) {
			assert.PrintFailure(t, i, out, tt.out)
		}
	}
}

func TestClearStringsRejectsNull(t *testing.T) {
	assert.RejectNull(t, func() { ClearStrings(nullSlice) }, "ClearStrings")
}
var clearStringsTests = []struct {
	in []string
	capacity int
}{
	{[]string{}, 0},
	{[]string{"s"}, 1},
	{[]string{"s", "str"}, 2},
	{make([]string, 10, 10), 10},
	{make([]string, 3, 10), 10},
}
func TestClearStrings(t *testing.T) {
	for i, tt := range clearStringsTests {
		if out := ClearStrings(tt.in); len(out) != 0 {
			t.Errorf("#d FAIL: length is not zero.", i)
		} else if cap(out) != tt.capacity {
			t.Errorf("#d FAIL:\n  Actual capacity: %d\nExpected capacity: %d",
				cap(out), tt.capacity)
		}
	}
}

func TestContainsStringRejectsNull(t *testing.T) {
	assert.RejectNull(
		t,
		func() { ContainsString(nullSlice, "") },
		"ContainsString")
}
var containsStringTests = []struct {
	s []string
	str string
	out bool
}{
	{[]string{}, "s", false},
	{[]string{"s", "str"}, "s", true},
	{[]string{"s", "str"}, "str", true},
	{[]string{"a", "b", "c"}, "s", false},
}
func TestContainsString(t *testing.T) {
	for i, tt := range containsStringTests {
		if out := ContainsString(tt.s, tt.str); out != tt.out {
			assert.PrintFailure(t, i, out, tt.out)
		}
	}
}

func TestContainsAllStringsRejectsNull(t *testing.T) {
	for _, tt := range rejectsNullTests {
		assert.RejectNull(
			t,
			func() { ContainsAllStrings(tt.this, tt.that) },
			"ContainsAllStrings")
	}
}
var containsAllStringsTests = []struct {
	this []string
	that []string
	out bool
}{
	{[]string{}, []string{}, true},
	{[]string{}, []string{""}, false},
	{[]string{""}, []string{}, false},
	{[]string{"s", "str"}, []string{"s"}, true},
	{[]string{"s", "str"}, []string{"s", "str"}, true},
	{[]string{"s", "str"}, []string{"str"}, true},
	{[]string{"s"}, []string{"str"}, false},
}
func TestContainsAllStrings(t *testing.T) {
	for i, tt := range containsAllStringsTests {
		if out := ContainsAllStrings(tt.this, tt.that); out != tt.out {
			assert.PrintFailure(t, i, out, tt.out)
		}
	}
}

func TestRemoveStringRejectsNull(t *testing.T) {
	assert.RejectNull(
		t,
		func () { RemoveString(nullSlice, "") }, "RemoveString")
}
var removeStringTests = []struct {
	s []string
	str string
	out []string
}{
	{[]string{}, "s", []string{}},
	{[]string{"a", "b", "abc"}, "a", []string{"b", "abc"}},
	{[]string{"a", "b", "abc"}, "abc", []string{"a", "b"}},
	{[]string{"a", "b", "abc"}, "d", []string{"a", "b", "abc"}},
}
func TestRemoveString(t *testing.T) {
	for i, tt := range removeStringTests {
		if out := RemoveString(tt.s, tt.str); !OfStringEquals(out, tt.out) {
			assert.PrintFailure(t, i, out, tt.out)
		}
	}
}

func TestRemoveAllStringsRejectsNull(t *testing.T) {
	for _, tt := range rejectsNullTests {
		assert.RejectNull(
			t,
			func() { RemoveAllStrings(tt.this, tt.that) }, "RemoveAllStrings")
	}
}
var removeAllStringsTests = []struct {
	this []string
	that []string
	out []string
}{
	{[]string{}, []string{}, []string{}},
	{[]string{}, []string{"s"}, []string{}},
	{[]string{"a", "b", "c"}, []string{}, []string{"a", "b", "c"}},
	{[]string{"a", "b", "c"}, []string{"s"}, []string{"a", "b", "c"}},
	{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, []string{}},
	{[]string{"a", "b", "c"}, []string{"a"}, []string{"b", "c"}},
	{[]string{"a", "b", "c"}, []string{"a", "c"}, []string{"b"}},
}
func TestRemoveAllStrings(t *testing.T) {
	for i, tt := range removeAllStringsTests {
		if out := RemoveAllStrings(tt.this, tt.that); !OfStringEquals(out, tt.out) {
			assert.PrintFailure(t, i, out, tt.out)
		}
	}
}

func TestRetainAllStringsRejectsNull(t *testing.T) {
	for _, tt := range rejectsNullTests {
		assert.RejectNull(
			t,
			func() { RetainAllStrings(tt.this, tt.that) }, "RetainAllStrings")
	}
}
var retainAllStringsTests = []struct {
	this []string
	that []string
	out []string
}{
	{[]string{}, []string{}, []string{}},
	{[]string{}, []string{"s"}, []string{}},
	{[]string{"a", "b", "c"}, []string{}, []string{}},
	{[]string{"a", "b", "c"}, []string{"s"}, []string{}},
	{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, []string{"a", "b", "c"}},
	{[]string{"a", "b", "c"}, []string{"a"}, []string{"a"}},
	{[]string{"a", "b", "c"}, []string{"a", "c"}, []string{"a", "c"}},
}
func TestRetainAllStrings(t *testing.T) {
	for i, tt := range retainAllStringsTests {
		if out := RetainAllStrings(tt.this, tt.that); !OfStringEquals(out, tt.out) {
			assert.PrintFailure(t, i, out, tt.out)
		}
	}
}