package variable

type Retriever interface {
	GetValue(variableName string) (*Value, bool)

	Contains(variableName string) bool
}

type Storer interface {
	Retriever

	SetNumberValue(variableName string, value float64)
	SetBooleanValue(variableName string, value bool)
	SetStringValue(variableName string, value string)

	Clear()
}
