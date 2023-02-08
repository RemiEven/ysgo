package runner

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/remieven/ysgo/internal/rng"
	"github.com/remieven/ysgo/internal/testutils"
	"github.com/remieven/ysgo/variable"
)

func TestRound(t *testing.T) {
	testRoundingFunc(t, round, map[float64]float64{
		3.522: 4,
		1.222: 1,
		-5.66: -6,
		-8.33: -8,
	})
}

func TestFloor(t *testing.T) {
	testRoundingFunc(t, floor, map[float64]float64{
		3.522: 3,
		-3.52: -4,
		1.222: 1,
		-1.22: -2,
	})
}

func TestCeil(t *testing.T) {
	testRoundingFunc(t, ceil, map[float64]float64{
		3.522: 4,
		-3.52: -3,
		1.444: 2,
		-1.44: -1,
	})
}

func TestInc(t *testing.T) {
	testRoundingFunc(t, inc, map[float64]float64{
		3.522: 4,
		4:     5,
		1.444: 2,
		-3.52: -3,
		-4:    -3,
		-1.44: -1,
	})
}

func TestDec(t *testing.T) {
	testRoundingFunc(t, dec, map[float64]float64{
		3.522: 3,
		4:     3,
		1.444: 1,
		-3.52: -4,
		-4:    -5,
		-1.44: -2,
	})
}

func TestDecimal(t *testing.T) {
	testRoundingFunc(t, decimal, map[float64]float64{
		3.522: 0.522,
		4:     0,
		-3.52: -0.52,
		-4:    0,
		-1.44: -0.44,
	})
}

func TestInteger(t *testing.T) {
	testRoundingFunc(t, integer, map[float64]float64{
		3.522: 3,
		4:     4,
		-3.52: -3,
		-4:    -4,
		-1.44: -1,
	})
}

func testRoundingFunc(t *testing.T, roundingFunc func(float64) float64, testCases map[float64]float64) {
	for value, expectedResult := range testCases {
		t.Run(fmt.Sprintf("%f", value), func(t *testing.T) {
			actualResult := roundingFunc(value)
			if !closeEnough(actualResult, expectedResult) {
				t.Errorf("unexpected result: got [%f], wanted [%f]", actualResult, expectedResult)
			}
		})
	}
}

func closeEnough(x, y float64) bool {
	return math.Abs(x-y) < math.Pow10(-9)
}

func TestRoundPlaces(t *testing.T) {
	testCases := map[struct {
		f      float64
		places int
	}]float64{
		{10.234567, 0}: 10,
		{10.234567, 1}: 10.2,
		{10.234567, 2}: 10.23,
		{10.234567, 3}: 10.235,
		{10.234567, 4}: 10.2346,
		{10.234567, 5}: 10.23457,
		{10.234567, 6}: 10.234567,
		{10.234567, 7}: 10.234567,
	}

	for args, expectedResult := range testCases {
		t.Run(fmt.Sprintf("%f_%v", args.f, args.places), func(t *testing.T) {
			actualResult := roundPlaces(args.f, args.places)
			if !closeEnough(actualResult, expectedResult) {
				t.Errorf("unexpected result: got [%f], wanted [%f]", actualResult, expectedResult)
			}
		})
	}
}

func TestDice(t *testing.T) {
	rng, err := rng.NewRNG("")
	if err != nil {
		t.Errorf("failed to initialize rng which is needed to create the dice function: %v", err)
		return
	}
	var (
		dice                        = dice(rng)
		upperBound                  = 6
		rolls                       = 100
		previousResult              int
		resultHasChangedAtLeastOnce = false
	)

	for i := 0; i < rolls; i++ {
		result := dice(upperBound)
		if result < 1 || result > upperBound {
			t.Errorf("got out of bound result from dice: %v", result)
			return
		}
		if i != 0 {
			resultHasChangedAtLeastOnce = resultHasChangedAtLeastOnce || previousResult != result
			previousResult = result
		}
	}

	if !resultHasChangedAtLeastOnce {
		t.Errorf("always got the same result, expected some change")
	}
}

func TestRandomRange(t *testing.T) {
	rng, err := rng.NewRNG("")
	if err != nil {
		t.Errorf("failed to initialize rng which is needed to create the randomRange function: %v", err)
		return
	}
	var (
		randomRange                 = randomRange(rng)
		lowerBound, upperBound      = 10, 100
		rolls                       = 100
		previousResult              int
		resultHasChangedAtLeastOnce = false
	)

	for i := 0; i < rolls; i++ {
		result := randomRange(lowerBound, upperBound)
		if result < lowerBound || result > upperBound {
			t.Errorf("got out of bound result from randomRange: %v", result)
			return
		}
		if i != 0 {
			resultHasChangedAtLeastOnce = resultHasChangedAtLeastOnce || previousResult != result
			previousResult = result
		}
	}

	if !resultHasChangedAtLeastOnce {
		t.Errorf("always got the same result, expected some change")
	}
}

func TestRandom(t *testing.T) {
	rng, err := rng.NewRNG("")
	if err != nil {
		t.Errorf("failed to initialize rng which is needed to create the random function: %v", err)
		return
	}
	var (
		random                      = random(rng)
		rolls                       = 100
		previousResult              float64
		resultHasChangedAtLeastOnce = false
	)

	for i := 0; i < rolls; i++ {
		result := random()
		if result < 0 || result > 1 {
			t.Errorf("got out of bound result from randomRange: %v", result)
			return
		}
		if i != 0 {
			resultHasChangedAtLeastOnce = resultHasChangedAtLeastOnce || previousResult != result
			previousResult = result
		}
	}

	if !resultHasChangedAtLeastOnce {
		t.Errorf("always got the same result, expected some change")
	}
}

func TestToString(t *testing.T) {
	tests := map[string]struct {
		input          []*variable.Value
		expectedResult *variable.Value
		expectedErr    error
	}{
		"already a string": {
			input:          []*variable.Value{variable.NewString("hello")},
			expectedResult: variable.NewString("hello"),
		},
		"number value (integer)": {
			input:          []*variable.Value{variable.NewNumber(8)},
			expectedResult: variable.NewString("8"),
		},
		"number value (with decimals)": {
			input:          []*variable.Value{variable.NewNumber(5.22)},
			expectedResult: variable.NewString("5.22"),
		},
		"boolean value": {
			input:          []*variable.Value{variable.NewBoolean(true)},
			expectedResult: variable.NewString("True"),
		},
		"empty value": {
			input:       []*variable.Value{{}},
			expectedErr: errors.New("received a value which was not a number, a boolean or a string"),
		},
		"no arguments": {
			input:       []*variable.Value{},
			expectedErr: errors.New("expected exactly one argument"),
		},
		"too many arguments": {
			input:       []*variable.Value{variable.NewBoolean(true), variable.NewBoolean(true)},
			expectedErr: errors.New("expected exactly one argument"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualResult, err := toString(test.input)
			if !testutils.ErrorEqual(err, test.expectedErr) {
				t.Errorf("unexpected error: got [%v], wanted [%v]", err, test.expectedErr)
				return
			}
			if diff := testutils.DeepEqual(actualResult, test.expectedResult); diff != "" {
				t.Errorf("unexpected result: got [%v], wanted [%v]", actualResult, test.expectedResult)
				return
			}
		})
	}
}

func TestToBoolean(t *testing.T) {
	tests := map[string]struct {
		input          []*variable.Value
		expectedResult *variable.Value
		expectedErr    error
	}{
		"already a boolean": {
			input:          []*variable.Value{variable.NewBoolean(true)},
			expectedResult: variable.NewBoolean(true),
		},
		"number value (non zero)": {
			input:          []*variable.Value{variable.NewNumber(8)},
			expectedResult: variable.NewBoolean(true),
		},
		"number value (zero)": {
			input:          []*variable.Value{variable.NewNumber(0)},
			expectedResult: variable.NewBoolean(false),
		},
		"string value (false)": {
			input:          []*variable.Value{variable.NewString("false")},
			expectedResult: variable.NewBoolean(false),
		},
		"string value (true)": {
			input:          []*variable.Value{variable.NewString("true")},
			expectedResult: variable.NewBoolean(true),
		},
		"string value (not a boolean)": {
			input:       []*variable.Value{variable.NewString("not a boolean")},
			expectedErr: errors.New(`failed to parse boolean from string: strconv.ParseBool: parsing "not a boolean": invalid syntax`),
		},
		"empty value": {
			input:       []*variable.Value{{}},
			expectedErr: errors.New("received a value which was not a number, a boolean or a string"),
		},
		"no arguments": {
			input:       []*variable.Value{},
			expectedErr: errors.New("expected exactly one argument"),
		},
		"too many arguments": {
			input:       []*variable.Value{variable.NewBoolean(true), variable.NewBoolean(true)},
			expectedErr: errors.New("expected exactly one argument"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualResult, err := toBoolean(test.input)
			if !testutils.ErrorEqual(err, test.expectedErr) {
				t.Errorf("unexpected error: got [%v], wanted [%v]", err, test.expectedErr)
				return
			}
			if diff := testutils.DeepEqual(actualResult, test.expectedResult); diff != "" {
				t.Errorf("unexpected result: got [%v], wanted [%v]", actualResult, test.expectedResult)
				return
			}
		})
	}
}

func TestToFloat(t *testing.T) {
	tests := map[string]struct {
		input          []*variable.Value
		expectedResult *variable.Value
		expectedErr    error
	}{
		"already a number": {
			input:          []*variable.Value{variable.NewNumber(1.23)},
			expectedResult: variable.NewNumber(1.23),
		},
		"boolean value (true)": {
			input:          []*variable.Value{variable.NewBoolean(true)},
			expectedResult: variable.NewNumber(1),
		},
		"boolean value (false)": {
			input:          []*variable.Value{variable.NewBoolean(false)},
			expectedResult: variable.NewNumber(0),
		},
		"string value (integer)": {
			input:          []*variable.Value{variable.NewString("1")},
			expectedResult: variable.NewNumber(1),
		},
		"string value (negative, with decimals)": {
			input:          []*variable.Value{variable.NewString("-1.74")},
			expectedResult: variable.NewNumber(-1.74),
		},
		"string value (not a number)": {
			input:       []*variable.Value{variable.NewString("not a number")},
			expectedErr: errors.New(`failed to parse number from string: strconv.ParseFloat: parsing "not a number": invalid syntax`),
		},
		"empty value": {
			input:       []*variable.Value{{}},
			expectedErr: errors.New("received a value which was not a number, a boolean or a string"),
		},
		"no arguments": {
			input:       []*variable.Value{},
			expectedErr: errors.New("expected exactly one argument"),
		},
		"too many arguments": {
			input:       []*variable.Value{variable.NewBoolean(true), variable.NewBoolean(true)},
			expectedErr: errors.New("expected exactly one argument"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualResult, err := toFloat(test.input)
			if !testutils.ErrorEqual(err, test.expectedErr) {
				t.Errorf("unexpected error: got [%v], wanted [%v]", err, test.expectedErr)
				return
			}
			if diff := testutils.DeepEqual(actualResult, test.expectedResult); diff != "" {
				t.Errorf("unexpected result: got [%v], wanted [%v]", actualResult, test.expectedResult)
				return
			}
		})
	}
}
