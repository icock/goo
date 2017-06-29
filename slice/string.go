package slice

import "notabug.org/icock/goo/exception"

// OfStringNotNull panics when slice is nil.
func OfStringNotNull(stringSlice []string) {
	if stringSlice == nil {
		panic(&exception.NullReference{})
	}
}
// Returns true if this slice contains no elements.
func OfStringIsEmpty(stringSlice []string) bool {
	OfStringNotNull(stringSlice)
	if len(stringSlice) == 0 {
		return true
	} else {
		return false
	}
}
// Equals returns true if both slice contains the same strings,
// capacity not counted.
func OfStringEquals(stringSlice []string, other []string) bool {
	OfStringNotNull(stringSlice)
	OfStringNotNull(other)
	if len(stringSlice) != len(other) {
		return false
	} else {
		for i, s := range stringSlice {
			if s != other[i] {
				return false
			}
		}
		return true
	}
}
// OfStringSize returns the number of strings in this slice.
func OfStringSize(stringSlice []string) int {
	OfStringNotNull(stringSlice)
	return len(stringSlice)
}

// AddString adds the specified string to this slice.
func AddString(stringSlice []string, str string) []string {
	OfStringNotNull(stringSlice)
	return append(stringSlice, str)
}
// AddAllStrings adds all of the strings in the specified slice to this slice.
func AddAllStrings(stringSlice []string, other []string) []string {
	OfStringNotNull(stringSlice)
	OfStringNotNull(other)
	return append(stringSlice, other...)
}

// Clear removes all of the strings from this slice,
// returning an empty slice with the same capacity.
func ClearStrings(stringSlice []string) []string {
	OfStringNotNull(stringSlice)
	capacity := cap(stringSlice)
	return make([]string, 0, capacity)
}

// Contains returns true if this slice contains the specified string.
func ContainsString(stringSlice []string, str string) bool {
	OfStringNotNull(stringSlice)
	for _, s := range stringSlice {
		if s == str {
			return true
		}
	}
	return false
}
// ContainsAll returns true if this slice contains all of the strings in the specified slice.
func ContainsAllStrings(stringSlice []string, other []string) bool {
	isStringSliceEmpty := OfStringIsEmpty(stringSlice)
	isOtherEmpty := OfStringIsEmpty(other)
	if isStringSliceEmpty && isOtherEmpty {
		return true
	} else if isStringSliceEmpty || isOtherEmpty {
		return false
	} else {
		for _, str := range other {
			if !ContainsString(stringSlice, str) {
				return false
			}
		}
		return true
	}
}

// RemoveString removes the specified string from this slice.
func RemoveString(stringSlice []string, str string) []string {
	OfStringNotNull(stringSlice)
	for i, s := range stringSlice {
		if s == str {
			stringSlice = append(stringSlice[:i], stringSlice[i+1:]...)
		}
	}
	return stringSlice
}
// RemoveAllStrings removes from this slice all of its strings that are contained in the specified slice.
func RemoveAllStrings(stringSlice []string, other []string) []string {
	isStringSliceEmpty := OfStringIsEmpty(stringSlice)
	isOtherEmpty := OfStringIsEmpty(other)
	if isStringSliceEmpty || isOtherEmpty {
		return stringSlice
	} else {
		for _, s := range other {
			stringSlice = RemoveString(stringSlice, s)
		}
		return stringSlice
	}
}
// RetainAllStrings retains only the strings in this slice that are contained in the specified slice.
func RetainAllStrings(stringSlice []string, other []string) []string {
	isStringSliceEmpty := OfStringIsEmpty(stringSlice)
	isOtherEmpty := OfStringIsEmpty(other)
	if isStringSliceEmpty || isOtherEmpty {
		return []string{}
	} else {
		result := make([]string, 0, cap(stringSlice))
		for _, s := range other {
			if ContainsString(stringSlice, s) {
				result = append(result, s)
			}
		}
		return result
	}
}

