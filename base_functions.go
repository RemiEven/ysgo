package ysgo

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/remieven/ysgo/internal/rng"
	"github.com/remieven/ysgo/variable"
)

func toString(args []*variable.Value) (*variable.Value, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected exactly one argument")
	}
	toConvert := args[0]
	if toConvert.Number == nil && toConvert.Boolean == nil && toConvert.String == nil {
		return nil, errors.New("received a value which was not a number, a boolean or a string")
	}
	return variable.NewString(toConvert.ToString()), nil
}

var _ = (YarnSpinnerFunction)(toString)

func toBoolean(args []*variable.Value) (*variable.Value, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected exactly one argument")
	}
	toConvert := args[0]
	switch {
	case toConvert.Number != nil:
		return variable.NewBoolean(*toConvert.Number != 0), nil
	case toConvert.Boolean != nil:
		return toConvert, nil
	case toConvert.String != nil:
		b, err := strconv.ParseBool(*toConvert.String)
		if err != nil {
			return nil, fmt.Errorf("failed to parse boolean from string: %w", err)
		}
		return variable.NewBoolean(b), nil
	}
	return nil, fmt.Errorf("received a value which was not a number, a boolean or a string")
}

var _ = (YarnSpinnerFunction)(toBoolean)

func toFloat(args []*variable.Value) (*variable.Value, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expected exactly one argument")
	}
	toConvert := args[0]
	switch {
	case toConvert.Number != nil:
		return toConvert, nil
	case toConvert.Boolean != nil:
		value := 0.0
		if *toConvert.Boolean {
			value = 1
		}
		return variable.NewNumber(value), nil
	case toConvert.String != nil:
		f, err := strconv.ParseFloat(*toConvert.String, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse number from string: %w", err)
		}
		return variable.NewNumber(f), nil
	}
	return nil, fmt.Errorf("received a value which was not a number, a boolean or a string")
}

var _ = (YarnSpinnerFunction)(toFloat)

// random returns a random number between 0 and 1
func random(rng *rng.RNG) func() float64 {
	return func() float64 {
		return rng.Float()
	}
}

// randomeRange returns a random integer between a and b, inclusive
func randomRange(rng *rng.RNG) func(int, int) int {
	return func(lowerBound, upperBound int) int {
		return rng.IntBetween(lowerBound, upperBound)
	}
}

// dice returns a random integer between 1 and sides, inclusive
func dice(rng *rng.RNG) func(int) int {
	return func(sides int) int {
		return rng.IntBetween(1, sides)
	}
}

// round rounds f to the nearest integer
func round(f float64) float64 {
	return math.Round(f)
}

// roundPlaces rounds f to the nearest number with places decimal points
func roundPlaces(f float64, places int) float64 {
	tenPowerPlaces := math.Pow10(places)
	return math.Round(f*tenPowerPlaces) / tenPowerPlaces
}

// floor rounds f down to the nearest integer, towards negative infinity
func floor(f float64) float64 {
	return math.Floor(f)
}

// ceil rounds f up to the nearest integer, towards positive infinity
func ceil(f float64) float64 {
	return math.Ceil(f)
}

// inc rounds f up to the nearest integer. If f is already an integer, inc returns f + 1
func inc(f float64) float64 {
	return floor(f) + 1
}

// dec rounds f down to the nearest integer. If f is already an integer, dec returns f - 1
func dec(f float64) float64 {
	return ceil(f) - 1
}

// decimal returns the decimal portion of f. This will always be a number between 0 and 1
func decimal(f float64) float64 {
	return f - integer(f)
}

// integer rounds down f down to the nearest integer, towards zero
func integer(f float64) float64 {
	return math.Trunc(f)
}
