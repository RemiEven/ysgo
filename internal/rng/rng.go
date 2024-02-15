package rng

import (
	"fmt"
	"math/rand"
)

// RNG can be used to generate random values.
type RNG struct {
	seed   string
	source *rand.Rand
}

// NewRNG creates a new RNG based on either a given seed or a randomly generated one.
func NewRNG(seed string) (*RNG, error) {
	seedValue := int64(0)
	if seed == "" {
		seedValue = rand.Int63()
		seed = int64ToSeed(seedValue)
	} else {
		var err error
		seedValue, err = seedToInt64(seed)
		if err != nil {
			return nil, fmt.Errorf("invalid seed: %w", err)
		}
	}
	return &RNG{
		seed:   seed,
		source: rand.New(rand.NewSource(seedValue)),
	}, nil
}

// Seed returns the RNG seed.
func (rng *RNG) Seed() string {
	return rng.seed
}

// IntBetween returns a random integer value between lower and upper bounds, inclusive.
func (rng *RNG) IntBetween(lowerBound, upperBound int) int {
	return lowerBound + rng.source.Intn(upperBound-lowerBound+1)
}

// Float returns a random float between 0 (inclusive) and 1 (exclusive).
func (rng *RNG) Float() float64 {
	return rng.source.Float64()
}
