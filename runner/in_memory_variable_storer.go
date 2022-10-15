package runner

import "github.com/RemiEven/ysgo/tree"

type InMemoryVariableStorer struct {
	numbers  map[string]float64
	booleans map[string]bool
	strings  map[string]string
}

func NewInMemoryVariableStorer() *InMemoryVariableStorer {
	storer := &InMemoryVariableStorer{}
	storer.Clear()
	return storer
}

func (storer *InMemoryVariableStorer) GetValue(variableName string) (*tree.Value, bool) {
	if value, ok := storer.numbers[variableName]; ok {
		return tree.NewNumberValue(value), true
	} else if value, ok := storer.booleans[variableName]; ok {
		return tree.NewBooleanValue(value), true
	} else if value, ok := storer.strings[variableName]; ok {
		return tree.NewStringValue(value), true
	}
	return nil, false
}

func (storer *InMemoryVariableStorer) SetNumberValue(variableName string, value float64) {
	storer.numbers[variableName] = value
}

func (storer *InMemoryVariableStorer) SetBooleanValue(variableName string, value bool) {
	storer.booleans[variableName] = value
}

func (storer *InMemoryVariableStorer) SetStringValue(variableName string, value string) {
	storer.strings[variableName] = value
}

func (storer *InMemoryVariableStorer) Clear() {
	storer.numbers = map[string]float64{}
	storer.booleans = map[string]bool{}
	storer.strings = map[string]string{}
}

func (storer *InMemoryVariableStorer) Contains(variableName string) bool {
	if _, ok := storer.numbers[variableName]; ok {
		return true
	} else if _, ok := storer.booleans[variableName]; ok {
		return true
	} else if _, ok := storer.strings[variableName]; ok {
		return true
	}
	return false
}

var _ VariableStorer = (*InMemoryVariableStorer)(nil)
