package runner

import (
	"errors"
	"fmt"
	"testing"

	"github.com/RemiEven/ysgo/internal/testutils"
	"github.com/RemiEven/ysgo/tree"
)

func TestNewYarnSpinnerFunction(t *testing.T) {
	tests := map[string]struct {
		function                any
		expectedConversionError error
		inputParameters         []*tree.Value
		expectedValue           *tree.Value
		expectedError           error
	}{
		"input is not a function": {
			function:                "not a function",
			expectedConversionError: errors.New("newYarnSpinnerFunction expects an argument which is a function"),
		},
		"no arg, no return": {
			function: func() {},
		},
		"no arg, int value return": {
			function:      func() int { return 1 },
			expectedValue: tree.NewNumberValue(1),
		},
		"no arg, float value return": {
			function:      func() float64 { return 1 },
			expectedValue: tree.NewNumberValue(1),
		},
		"no arg, bool value return": {
			function:      func() bool { return true },
			expectedValue: tree.NewBooleanValue(true),
		},
		"no arg, string value return": {
			function:      func() string { return "hello world" },
			expectedValue: tree.NewStringValue("hello world"),
		},
		"no arg, error value (nil) return": {
			function: func() error { return nil },
		},
		"no arg, error value (non-nil) return": {
			function:      func() error { return errors.New("some error") },
			expectedError: errors.New("some error"),
		},
		"no arg, value and error (nil) return": {
			function:      func() (string, error) { return "hello world", nil },
			expectedValue: tree.NewStringValue("hello world"),
		},
		"no arg, value and error (non-nil) return": {
			function:      func() (string, error) { return "hello world", errors.New("some error") },
			expectedError: errors.New("some error"),
		},
		"no arg, unsupported return type": {
			function:                func() complex128 { return 1 + 2i },
			expectedConversionError: errors.New("received function has unsupported output parameters: a single output parameter must be either convertible to a yarn spinner value or an error"),
		},
		"no arg, unsupported first return type": {
			function:                func() (complex128, error) { return 0, nil },
			expectedConversionError: errors.New("received function has unsupported output parameters: the first output parameter of a function with two must be convertible to a yarn spinner value"),
		},
		"no arg, unsupported second return type": {
			function:                func() (float64, complex64) { return 0, 0 },
			expectedConversionError: errors.New("received function has unsupported output parameters: the second output parameter of a function with two must be an error"),
		},
		"too many returns": {
			function:                func() (a, b, c float64) { return },
			expectedConversionError: errors.New("received function has unsupported output parameters: functions can have at most two output parameters"),
		},
		"one int arg": {
			function:        func(i int) {},
			inputParameters: []*tree.Value{tree.NewNumberValue(3)},
		},
		"one int arg, wrong argument when calling function": {
			function:        func(i int) {},
			inputParameters: []*tree.Value{tree.NewStringValue("not a number")},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 0: expected a number value but got something else"),
		},
		"one int8 arg": {
			function:        func(i int8) {},
			inputParameters: []*tree.Value{tree.NewNumberValue(3)},
		},
		"one int8 arg, wrong argument when calling function": {
			function:        func(i int8) {},
			inputParameters: []*tree.Value{tree.NewStringValue("not a number")},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 0: expected a number value but got something else"),
		},
		"one int16 arg": {
			function:        func(i int16) {},
			inputParameters: []*tree.Value{tree.NewNumberValue(3)},
		},
		"one int16 arg, wrong argument when calling function": {
			function:        func(i int16) {},
			inputParameters: []*tree.Value{tree.NewStringValue("not a number")},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 0: expected a number value but got something else"),
		},
		"one int32 arg": {
			function:        func(i int32) {},
			inputParameters: []*tree.Value{tree.NewNumberValue(3)},
		},
		"one int32 arg, wrong argument when calling function": {
			function:        func(i int32) {},
			inputParameters: []*tree.Value{tree.NewStringValue("not a number")},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 0: expected a number value but got something else"),
		},
		"one int64 arg": {
			function:        func(i int64) {},
			inputParameters: []*tree.Value{tree.NewNumberValue(3)},
		},
		"one int64 arg, wrong argument when calling function": {
			function:        func(i int64) {},
			inputParameters: []*tree.Value{tree.NewStringValue("not a number")},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 0: expected a number value but got something else"),
		},
		"one float32 arg": {
			function:        func(f float32) {},
			inputParameters: []*tree.Value{tree.NewNumberValue(4.6)},
		},
		"one float32 arg, wrong argument when calling function": {
			function:        func(f float32) {},
			inputParameters: []*tree.Value{tree.NewStringValue("not a number")},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 0: expected a number value but got something else"),
		},
		"one float64 arg": {
			function:        func(f float64) {},
			inputParameters: []*tree.Value{tree.NewNumberValue(4.6)},
		},
		"one float64 arg, wrong argument when calling function": {
			function:        func(f float64) {},
			inputParameters: []*tree.Value{tree.NewStringValue("not a number")},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 0: expected a number value but got something else"),
		},
		"one bool arg": {
			function:        func(b bool) {},
			inputParameters: []*tree.Value{tree.NewBooleanValue(true)},
		},
		"one bool arg, wrong argument when calling function": {
			function:        func(b bool) {},
			inputParameters: []*tree.Value{tree.NewStringValue("not a boolean")},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 0: expected a boolean value but got something else"),
		},
		"one string arg": {
			function:        func(s string) {},
			inputParameters: []*tree.Value{tree.NewStringValue("hello")},
		},
		"one string arg, wrong argument when calling function": {
			function:        func(s string) {},
			inputParameters: []*tree.Value{tree.NewBooleanValue(false)},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 0: expected a string value but got something else"),
		},
		"one arg of unsupported type": {
			function:                func(c complex128) {},
			expectedConversionError: errors.New("failed to create input converter: argument number 0 has an unsupported type complex128"),
		},
		"several args, all of supported type": {
			function: func(i int, f1 float64, f2 float64, b bool, s string) string {
				return fmt.Sprintf("%v %v %v %v %v", i, f1, f2, b, s)
			},
			inputParameters: []*tree.Value{
				tree.NewNumberValue(1),
				tree.NewNumberValue(-3.4),
				tree.NewNumberValue(999_999.999),
				tree.NewBooleanValue(false),
				tree.NewStringValue("hello"),
			},
			expectedValue: tree.NewStringValue("1 -3.4 999999.999 false hello"),
		},
		"non variadic function, passing too many arguments": {
			function:        func() {},
			inputParameters: []*tree.Value{tree.NewNumberValue(1)},
			expectedError:   errors.New("failed to convert input parameters: received too many arguments"),
		},
		"non variadic function, passing too few arguments": {
			function:        func(i int) {},
			inputParameters: []*tree.Value{},
			expectedError:   errors.New("failed to convert input parameters: received too few arguments"),
		},
		"purely variadic function, one return, no variadic argument": {
			function:        sumInt,
			inputParameters: []*tree.Value{},
			expectedValue:   tree.NewNumberValue(0),
		},
		"purely variadic function, one return, single variadic argument": {
			function:        sumInt,
			inputParameters: []*tree.Value{tree.NewNumberValue(4)},
			expectedValue:   tree.NewNumberValue(4),
		},
		"purely variadic function, one return, several variadic arguments": {
			function:        sumInt,
			inputParameters: []*tree.Value{tree.NewNumberValue(4), tree.NewNumberValue(2), tree.NewNumberValue(1)},
			expectedValue:   tree.NewNumberValue(7),
		},
		"mixed variadic function, too few arguments": {
			function:        concatenateNTimes,
			inputParameters: []*tree.Value{},
			expectedError:   errors.New("failed to convert input parameters: received too few arguments"),
		},
		"mixed variadic function, no variadic arguments": {
			function:        concatenateNTimes,
			inputParameters: []*tree.Value{tree.NewNumberValue(2)},
			expectedValue:   tree.NewStringValue(""),
		},
		"mixed variadic function, single variadic argument": {
			function:        concatenateNTimes,
			inputParameters: []*tree.Value{tree.NewNumberValue(2), tree.NewStringValue("a")},
			expectedValue:   tree.NewStringValue("aa"),
		},
		"mixed variadic function, several variadic arguments": {
			function:        concatenateNTimes,
			inputParameters: []*tree.Value{tree.NewNumberValue(2), tree.NewStringValue("a"), tree.NewStringValue("b"), tree.NewStringValue("c")},
			expectedValue:   tree.NewStringValue("abcabc"),
		},
		"variadic function, variadic argument of unsupported type": {
			function:                func(...complex64) {},
			expectedConversionError: errors.New("failed to create input converter: argument number 0 has an unsupported type slice"),
		},
		"variadic function, non-variadic argument of unsupported type": {
			function:                func(complex64, ...int) {},
			expectedConversionError: errors.New("failed to create input converter: argument number 0 has an unsupported type complex64"),
		},
		"variadic function, wrong non-variadic argument type when calling": {
			function:        concatenateNTimes,
			inputParameters: []*tree.Value{tree.NewBooleanValue(true)},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 0: expected a number value but got something else"),
		},
		"variadic function, wrong variadic argument type when calling": {
			function:        concatenateNTimes,
			inputParameters: []*tree.Value{tree.NewNumberValue(3), tree.NewStringValue("a"), tree.NewNumberValue(2)},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 2: expected a string value but got something else"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			f, err := newYarnSpinnerFunction(test.function)
			if !testutils.ErrorEqual(err, test.expectedConversionError) {
				t.Errorf("unexpected error while converting function: got [%v], wanted [%v]", err, test.expectedConversionError)
				return
			} else if test.expectedConversionError != nil {
				return
			}

			value, err := f(test.inputParameters)
			if !testutils.ErrorEqual(err, test.expectedError) {
				t.Errorf("unexpected error while calling function: got [%v], wanted [%v]", err, test.expectedError)
			}
			if diff := testutils.DeepEqual(value, test.expectedValue); diff != "" {
				t.Errorf("unexpected returned value: " + diff)
			}
		})
	}
}

func sumInt(numbers ...int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func concatenateNTimes(n int, strings ...string) string {
	if n == 0 || len(strings) == 0 {
		return ""
	}
	concatenated := ""
	for _, s := range strings {
		concatenated += s
	}

	concatenatedNTimes := ""
	for i := 0; i < n; i++ {
		concatenatedNTimes += concatenated
	}

	return concatenatedNTimes
}

func TestFunctionStorer(t *testing.T) {
	storer := newFunctionStorer(nil)

	if _, err := storer.Call("unknown_function", nil); err == nil {
		t.Errorf("expected an error while calling an unknown function but did not get any")
	} else if expectedErr := errors.New("unknown function"); !testutils.ErrorEqual(err, expectedErr) {
		t.Errorf("unexpected error: got [%v], wanted [%v]", err, expectedErr)
	}

	actual, err := storer.Call("round", []*tree.Value{tree.NewNumberValue(3.7)})
	expected := tree.NewNumberValue(4)
	if err != nil {
		t.Errorf("expected no error while calling a function but got [%v]", err)
	} else if diff := testutils.DeepEqual(actual, expected); diff != "" {
		t.Errorf("unexpected result: got [%+v], wanted [%+v]", actual, expected)
	}
}
