package rng

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const radix = '9' - '0' + 'z' - 'a' + 2

func toRadix36(r rune) (int64, error) {
	if '0' <= r && r <= '9' {
		return int64(r - 48), nil
	}
	if 'a' <= r && r <= 'z' {
		return int64(r - 'a' + ('9' - '0' + 1)), nil
	}
	return 0, errors.New("only number and lowercase characters are supported")
}

func isValidSeed(seed string) bool {
	_, err := seedToInt64(seed)
	return err == nil
}

func seedToInt64(seed string) (int64, error) {
	result := int64(0)
	for i, r := range []rune(seed) {
		v, err := toRadix36(r)
		if err != nil {
			return 0, fmt.Errorf("invalid rune at position %v: %w", i, err)
		}
		result = radix*result + v
	}
	return result, nil
}

func toRune(value int64) rune {
	if 0 <= value && value <= '9'-'0' {
		return '0' + rune(value)
	}
	if '9'-'0'+1 <= value && value <= radix-1 {
		return 'a' - ('9' - '0' + 1) + rune(value)
	}
	return -1
}

func int64ToSeed(value int64) string {
	e := int(math.Floor(math.Log(float64(value)) / math.Log(radix)))
	seed := make([]rune, 0, e)
	posValue := int64(math.Pow(radix, float64(e)))
	for e >= 0 {
		digit := value / posValue
		seed = append(seed, toRune(digit))
		value %= posValue
		posValue /= radix
		e--
	}
	return string(seed)
}

// RandomSeed generates a random seed
func RandomSeed() string {
	return int64ToSeed(rand.NewSource(time.Now().Unix()).Int63())
}
