package testutils

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestErrorEqual(t *testing.T) {
	testCases := map[string]struct {
		actualErr    error
		expectedErr  error
		shouldReturn bool
	}{
		"Both nil": {
			actualErr:    nil,
			expectedErr:  nil,
			shouldReturn: true,
		},
		"Both same messages": {
			actualErr:    errors.New("message"),
			expectedErr:  errors.New("message"),
			shouldReturn: true,
		},
		"Actual is nil": {
			actualErr:    nil,
			expectedErr:  errors.New("message"),
			shouldReturn: false,
		},
		"Expected is nil": {
			actualErr:    errors.New("message"),
			expectedErr:  nil,
			shouldReturn: false,
		},
		"Messages differ": {
			actualErr:    errors.New("a message"),
			expectedErr:  errors.New("another message"),
			shouldReturn: false,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			actualReturn := ErrorEqual(test.actualErr, test.expectedErr)
			if actualReturn != test.shouldReturn {
				t.Errorf("Got result [%v], expected [%v]", actualReturn, test.shouldReturn)
			}
		})
	}
}

func TestErrorsEqual(t *testing.T) {
	testCases := map[string]struct {
		actualErrs   []error
		expectedErrs []error
		shouldReturn bool
	}{
		"Both empty": {
			actualErrs:   []error{},
			expectedErrs: []error{},
			shouldReturn: true,
		},
		"more errors than expected": {
			actualErrs:   []error{errors.New("some error")},
			expectedErrs: []error{},
			shouldReturn: false,
		},
		"less errors than expected": {
			actualErrs:   []error{},
			expectedErrs: []error{errors.New("some error")},
			shouldReturn: false,
		},
		"Some errors have different messages": {
			actualErrs:   []error{errors.New("some error"), errors.New("some other error")},
			expectedErrs: []error{errors.New("some error"), errors.New("yet another error")},
			shouldReturn: false,
		},
		"Ordering is not the same": {
			actualErrs:   []error{errors.New("some error"), errors.New("some other error")},
			expectedErrs: []error{errors.New("some other error"), errors.New("some error")},
			shouldReturn: false,
		},
		"Many errors with same messages and ordering": {
			actualErrs:   []error{errors.New("some error"), errors.New("some other error")},
			expectedErrs: []error{errors.New("some error"), errors.New("some other error")},
			shouldReturn: true,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			actualReturn := ErrorsEqual(test.actualErrs, test.expectedErrs)
			if actualReturn != test.shouldReturn {
				t.Errorf("Got result [%v], expected [%v]", actualReturn, test.shouldReturn)
			}
		})
	}
}

type exampleStruct struct {
	FirstName, LastName string
	unexportedName      string
}

func TestDeepEqual(t *testing.T) {
	tests := map[string]struct {
		a, b           interface{}
		expectedResult string
	}{
		"Equal objects": {
			a: exampleStruct{
				FirstName:      "FirstName",
				LastName:       "LastName",
				unexportedName: "Name",
			},
			b: exampleStruct{
				FirstName:      "FirstName",
				LastName:       "LastName",
				unexportedName: "Name",
			},
			expectedResult: ``,
		},
		"Different object types": {
			a: exampleStruct{
				FirstName:      "FirstName",
				LastName:       "LastName",
				unexportedName: "Name",
			},
			b:              time.Time{},
			expectedResult: `difference(s) found between actual and expected: testutils.exampleStruct != time.Time`,
		},
		"Objects with one difference": {
			a: exampleStruct{
				FirstName:      "FirstName",
				LastName:       "LastName",
				unexportedName: "Name",
			},
			b: exampleStruct{
				FirstName:      "Other FirstName",
				LastName:       "LastName",
				unexportedName: "Name",
			},
			expectedResult: `difference(s) found between actual and expected: FirstName: FirstName != Other FirstName`,
		},
		"Objects with several differences": {
			a: exampleStruct{
				FirstName:      "FirstName",
				LastName:       "LastName",
				unexportedName: "Name",
			},
			b: exampleStruct{
				FirstName:      "Other FirstName",
				LastName:       "Other LastName",
				unexportedName: "Name",
			},
			expectedResult: `difference(s) found between actual and expected:
- FirstName: FirstName != Other FirstName
- LastName: LastName != Other LastName`,
		},
		"Objects with difference but it is in an unexported field": {
			a: exampleStruct{
				FirstName:      "FirstName",
				LastName:       "LastName",
				unexportedName: "Name",
			},
			b: exampleStruct{
				FirstName:      "FirstName",
				LastName:       "LastName",
				unexportedName: "Other Secret Name",
			},
			expectedResult: ``,
		},
		"Nil with empty slice": {
			a:              nil,
			b:              []string{},
			expectedResult: `difference(s) found between actual and expected: <nil pointer> != []`,
		},
		"Int with solid float": {
			a:              int(1),
			b:              float64(1.0),
			expectedResult: `difference(s) found between actual and expected: int != float64`,
		},
		"Slices containing different types": {
			a:              []string{},
			b:              []interface{}{},
			expectedResult: `difference(s) found between actual and expected: []string != []interface {}`,
		},
		"Errors with same messages but different underlying types": { // in those cases, avoid testutils.DeepEqual and use testutils.ErrorEqual
			a:              fmt.Errorf("wrapping error: %w", errors.New("wrapped error")),
			b:              errors.New("wrapping error: wrapped error"),
			expectedResult: `difference(s) found between actual and expected: *fmt.wrapError != *errors.errorString`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualResult := DeepEqual(test.a, test.b)

			if actualResult != test.expectedResult {
				t.Errorf("got result [%v], wanted [%v]", actualResult, test.expectedResult)
			}
		})
	}
}
