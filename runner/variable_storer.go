package runner

import "github.com/RemiEven/ysgo/tree"

type variableRetriever interface {
	GetValue(variableName string) (*tree.Value, bool)

	Contains(variableName string) bool
}

type VariableStorer interface {
	variableRetriever

	SetNumberValue(variableName string, value float64)
	SetBooleanValue(variableName string, value bool)
	SetStringValue(variableName string, value string)

	Clear()
}
