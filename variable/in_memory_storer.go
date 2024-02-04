package variable

// InMemoryStorer stores variables from a dialogue using in-memory maps.
type InMemoryStorer struct {
	numbers  map[string]float64
	booleans map[string]bool
	strings  map[string]string
}

// NewInMemoryStorer creates a new empty InMemoryStorer.
func NewInMemoryStorer() *InMemoryStorer {
	storer := &InMemoryStorer{}
	storer.Clear()
	return storer
}

// GetValue returns the value of a variable, if it is stored.
func (storer *InMemoryStorer) GetValue(variableName string) (*Value, bool) {
	if value, ok := storer.numbers[variableName]; ok {
		return NewNumber(value), true
	} else if value, ok := storer.booleans[variableName]; ok {
		return NewBoolean(value), true
	} else if value, ok := storer.strings[variableName]; ok {
		return NewString(value), true
	}
	return nil, false
}

// GetValues returns a map of all stored values.
func (storer *InMemoryStorer) GetValues() map[string]Value {
	values := make(map[string]Value, len(storer.booleans)+len(storer.numbers)+len(storer.strings))

	for variable, value := range storer.booleans {
		values[variable] = *NewBoolean(value)
	}
	for variable, value := range storer.numbers {
		values[variable] = *NewNumber(value)
	}
	for variable, value := range storer.strings {
		values[variable] = *NewString(value)
	}

	return values
}

// SetNumberValue stores a number.
func (storer *InMemoryStorer) SetNumberValue(variableName string, value float64) {
	storer.numbers[variableName] = value
}

// SetBooleanValue stores a boolean.
func (storer *InMemoryStorer) SetBooleanValue(variableName string, value bool) {
	storer.booleans[variableName] = value
}

// SetStringValue stores a string.
func (storer *InMemoryStorer) SetStringValue(variableName string, value string) {
	storer.strings[variableName] = value
}

// Clear removes all variables from the store.
func (storer *InMemoryStorer) Clear() {
	storer.numbers = map[string]float64{}
	storer.booleans = map[string]bool{}
	storer.strings = map[string]string{}
}

// Contains returns whether a variable with the given name is currently stored.
func (storer *InMemoryStorer) Contains(variableName string) bool {
	if _, ok := storer.numbers[variableName]; ok {
		return true
	} else if _, ok := storer.booleans[variableName]; ok {
		return true
	} else if _, ok := storer.strings[variableName]; ok {
		return true
	}
	return false
}

var _ Storer = (*InMemoryStorer)(nil)
