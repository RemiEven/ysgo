package ysgo

// InitOptions contains options used to initialize a dialogue runner.
// The zero value is valid.
type InitOptions struct {
	// RNGSeed is used to create deterministic random values. It will be
	// used when eg. the dice(6) function is called in a dialogue. If left
	// empty, a random value will be used.
	RNGSeed string
	// FirstNode is the name of the node where the dialogue should begin.
	// If left empty, the first parsed node will be used.
	FirstNode string
}
