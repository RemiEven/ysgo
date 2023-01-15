package variable

type InMemoryStorer struct {
	numbers  map[string]float64
	booleans map[string]bool
	strings  map[string]string
}

func NewInMemoryStorer() *InMemoryStorer {
	storer := &InMemoryStorer{}
	storer.Clear()
	return storer
}

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

func (storer *InMemoryStorer) SetNumberValue(variableName string, value float64) {
	storer.numbers[variableName] = value
}

func (storer *InMemoryStorer) SetBooleanValue(variableName string, value bool) {
	storer.booleans[variableName] = value
}

func (storer *InMemoryStorer) SetStringValue(variableName string, value string) {
	storer.strings[variableName] = value
}

func (storer *InMemoryStorer) Clear() {
	storer.numbers = map[string]float64{}
	storer.booleans = map[string]bool{}
	storer.strings = map[string]string{}
}

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
