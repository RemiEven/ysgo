package testutils

import (
	"strings"

	"github.com/go-test/deep"
)

// ErrorsEqual checks whether two error slices contains errors with the same messages and in the same order
func ErrorsEqual(actual, expected []error) bool {
	if len(actual) != len(expected) {
		return false
	}
	for i := range expected {
		if !ErrorEqual(actual[i], expected[i]) {
			return false
		}
	}
	return true
}

// ErrorEqual checks whether two errors have the same message (or are both nil)
func ErrorEqual(actual, expected error) bool {
	if actual == nil && expected == nil {
		return true
	}
	if (actual != nil && expected == nil) || (actual == nil && expected != nil) {
		return false
	}
	if actual.Error() != expected.Error() {
		return false
	}
	return true
}

// DeepEqual compares a and b and returns a formatted string explaining the differences if there are any
func DeepEqual(actual, expected interface{}) string {
	diff := deep.Equal(actual, expected)
	diffHeader := "difference(s) found between actual and expected:"
	switch len(diff) {
	case 0:
		return ""
	case 1:
		return diffHeader + " " + diff[0]
	default:
		return diffHeader + "\n- " + strings.Join(diff, "\n- ")
	}
}
