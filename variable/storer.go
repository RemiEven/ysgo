package variable

// Retriever allows to access stored variable values.
type Retriever interface {
	GetValue(variableName string) (*Value, bool)

	Contains(variableName string) bool
}

// Storer allows to access and modify stored variable values.
type Storer interface {
	Retriever

	SetNumberValue(variableName string, value float64)
	SetBooleanValue(variableName string, value bool)
	SetStringValue(variableName string, value string)

	Clear()
}
