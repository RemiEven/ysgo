package runner

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/RemiEven/ysgo/internal/testutils"
	"github.com/RemiEven/ysgo/variable"
)

func TestNewYarnSpinnerCommand(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		command                 any
		expectedConversionError error
		inputParameters         []*variable.Value
		expectedError           error
	}{
		"input is not a function": {
			command:                 "not a function",
			expectedConversionError: errors.New("newYarnSpinnerCommand expects an argument which is a function"),
		},
		"no arg, no return": {
			command: func() {},
		},
		"no arg, error value (nil) return": {
			command: func() error { return nil },
		},
		"no arg, error value (non-nil) return": {
			command:       func() error { return errors.New("some error") },
			expectedError: errors.New("some error"),
		},
		"no arg, buffered error chan (sending nil) return": {
			command: func() <-chan error {
				return chanWithImmediateValue(error(nil))
			},
		},
		"no arg, buffered error chan (sending non-nil) return": {
			command: func() <-chan error {
				return chanWithImmediateValue(errors.New("some error"))
			},
			expectedError: errors.New("some error"),
		},
		"no arg, unbuffered error chan (sending non-nil) return": {
			command: func() chan error {
				errChan := make(chan error)
				go func() {
					errChan <- errors.New("some error")
				}()
				return errChan
			},
			expectedError: errors.New("some error"),
		},
		"no arg, unbuffered error chan (sending nil) taking longish to return": {
			command: func() chan error {
				errChan := make(chan error)
				go func() {
					time.Sleep(10 * time.Millisecond)
					errChan <- nil
				}()
				return errChan
			},
		},
		"no arg, unbuffered error chan (sending non-nil) taking longish to return": {
			command: func() chan error {
				errChan := make(chan error)
				go func() {
					time.Sleep(10 * time.Millisecond)
					errChan <- errors.New("some error")
				}()
				return errChan
			},
			expectedError: errors.New("some error"),
		},
		"no arg, command returns a nil chan": {
			command:       func() chan error { return nil },
			expectedError: errors.New("command returned a nil chan"),
		},
		"no arg, unsupported return type": {
			command:                 func() complex128 { return 1 + 2i },
			expectedConversionError: errors.New("received command has unsupported output parameters: a single output parameter must be either convertible to an error or to a chan error"),
		},
		"too many returns": {
			command:                 func() (a, b float64) { return },
			expectedConversionError: errors.New("received command has unsupported output parameters: commands can have at most one output parameter"),
		},
		"one arg of unsupported type": {
			command:                 func(c complex128) {},
			expectedConversionError: errors.New("failed to create input converter: argument number 0 has an unsupported type complex128"),
		},
		"several args, all of supported type": {
			command: func(i int, f1 float64, f2 float64, b bool, s string) error {
				value := fmt.Sprintf("%v %v %v %v %v", i, f1, f2, b, s)
				if value != "1 -3.4 999999.999 false hello" {
					return errors.New("unexpected value")
				}
				return nil
			},
			inputParameters: []*variable.Value{
				variable.NewNumber(1),
				variable.NewNumber(-3.4),
				variable.NewNumber(999_999.999),
				variable.NewBoolean(false),
				variable.NewString("hello"),
			},
		},
		"non variadic function, passing too many arguments": {
			command:         func() {},
			inputParameters: []*variable.Value{variable.NewNumber(1)},
			expectedError:   errors.New("failed to convert input parameters: received too many arguments"),
		},
		"non variadic function, passing too few arguments": {
			command:         func(i int) {},
			inputParameters: []*variable.Value{},
			expectedError:   errors.New("failed to convert input parameters: received too few arguments"),
		},
		"purely variadic function, one return, no variadic argument": {
			command:         sumIntAndExpect(0),
			inputParameters: []*variable.Value{},
		},
		"purely variadic function, one return, single variadic argument": {
			command:         sumIntAndExpect(4),
			inputParameters: []*variable.Value{variable.NewNumber(4)},
		},
		"purely variadic function, one return, several variadic arguments": {
			command:         sumIntAndExpect(7),
			inputParameters: []*variable.Value{variable.NewNumber(4), variable.NewNumber(2), variable.NewNumber(1)},
		},
		"mixed variadic function, too few arguments": {
			command:         concatenateNTimesAndExpect(""),
			inputParameters: []*variable.Value{},
			expectedError:   errors.New("failed to convert input parameters: received too few arguments"),
		},
		"mixed variadic function, no variadic arguments": {
			command:         concatenateNTimesAndExpect(""),
			inputParameters: []*variable.Value{variable.NewNumber(2)},
		},
		"mixed variadic function, single variadic argument": {
			command:         concatenateNTimesAndExpect("aa"),
			inputParameters: []*variable.Value{variable.NewNumber(2), variable.NewString("a")},
		},
		"mixed variadic function, several variadic arguments": {
			command:         concatenateNTimesAndExpect("abcabc"),
			inputParameters: []*variable.Value{variable.NewNumber(2), variable.NewString("a"), variable.NewString("b"), variable.NewString("c")},
		},
		"variadic function, variadic argument of unsupported type": {
			command:                 func(...complex64) {},
			expectedConversionError: errors.New("failed to create input converter: argument number 0 has an unsupported type slice"),
		},
		"variadic function, non-variadic argument of unsupported type": {
			command:                 func(complex64, ...int) {},
			expectedConversionError: errors.New("failed to create input converter: argument number 0 has an unsupported type complex64"),
		},
		"variadic function, wrong non-variadic argument type when calling": {
			command:         concatenateNTimesAndExpect(""),
			inputParameters: []*variable.Value{variable.NewBoolean(true)},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 0: expected a number value but got something else"),
		},
		"variadic function, wrong variadic argument type when calling": {
			command:         concatenateNTimesAndExpect(""),
			inputParameters: []*variable.Value{variable.NewNumber(3), variable.NewString("a"), variable.NewNumber(2)},
			expectedError:   errors.New("failed to convert input parameters: failed to convert argument number 2: expected a string value but got something else"),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c, err := newYarnSpinnerCommand(test.command)
			if !testutils.ErrorEqual(err, test.expectedConversionError) {
				t.Errorf("unexpected error while converting command: got [%v], wanted [%v]", err, test.expectedConversionError)
				return
			} else if test.expectedConversionError != nil {
				return
			}

			errChan := c(test.inputParameters)
			select {
			case returnedErr := <-errChan:
				if !testutils.ErrorEqual(returnedErr, test.expectedError) {
					t.Errorf("unexpected error while calling command: got [%v], wanted [%v]", returnedErr, test.expectedError)
				}
			case <-time.After(40 * time.Millisecond):
				t.Errorf("timed out waiting for command to complete")
				return
			}
		})
	}
}

func concatenateNTimesAndExpect(expected string) func(int, ...string) error {
	return func(n int, strings ...string) error {
		value := concatenateNTimes(n, strings...)
		if value != expected {
			return errors.New("unexpected value")
		}
		return nil
	}
}

func sumIntAndExpect(expected int) func(...int) error {
	return func(numbers ...int) error {
		value := sumInt(numbers...)
		if value != expected {
			return errors.New("unexpected value")
		}
		return nil
	}
}

func TestCommandStorer(t *testing.T) {
	t.Parallel()
	storer := newCommandStorer()

	errChan := storer.call("unknown_command", nil)
	select {
	case err := <-errChan:
		expectedError := errors.New("unknown command")
		if !testutils.ErrorEqual(err, expectedError) {
			t.Errorf("unexpected error while calling an unknown command: got [%v], wanted [%v]", err, expectedError)
		}
	default:
		t.Errorf("calling an unknown command should result in an immediate error")
	}

	errChan = storer.call("wait", []*variable.Value{variable.NewNumber(0.02)})
	select {
	case err := <-errChan:
		t.Errorf("wait returned too soon with error [%v]", err)
		return
	default:
		// go on with the test
	}
	time.Sleep(25 * time.Millisecond)
	select {
	case err := <-errChan:
		if err != nil {
			t.Errorf("wait has returned with an unexpected error: [%v]", err)
		}
	default:
		t.Errorf("wait still has not returned when it should already have")
	}
}
