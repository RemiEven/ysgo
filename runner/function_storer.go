package runner

import (
	"fmt"
	"reflect"

	"github.com/remieven/ysgo/internal/rng"
	"github.com/remieven/ysgo/variable"
)

type functionCaller interface {
	call(functionID string, args []*variable.Value) (*variable.Value, error)
}

// YarnSpinnerFunction is a Go function than can be called as a function from a YarnSpinner script.
// It can take zero, one or more input values, and returns possibly a value and/or an error.
type YarnSpinnerFunction func([]*variable.Value) (*variable.Value, error)

type functionStorer struct {
	functionsByID map[string]YarnSpinnerFunction
}

func newFunctionStorer(rng *rng.RNG) *functionStorer {
	storer := &functionStorer{
		functionsByID: map[string]YarnSpinnerFunction{
			"string": toString,
			"bool":   toBoolean,
			"number": toFloat,
		},
	}

	for functionID, f := range map[string]any{
		"dice":         dice(rng),
		"random":       random(rng),
		"random_range": randomRange(rng),
		"round":        round,
		"round_places": roundPlaces,
		"floor":        floor,
		"ceil":         ceil,
		"inc":          inc,
		"dec":          dec,
		"decimal":      decimal,
		"integer":      integer,
	} {
		if err := storer.convertAndAddFunction(functionID, f); err != nil {
			panic(fmt.Errorf("failed to convert base function %s: %w", functionID, err))
		}
	}

	return storer
}

func (storer *functionStorer) addFunction(functionID string, function YarnSpinnerFunction) {
	storer.functionsByID[functionID] = function
}

func (storer *functionStorer) convertAndAddFunction(functionID string, function any) error {
	yarnSpinnerFunction, err := newYarnSpinnerFunction(function)
	if err != nil {
		return fmt.Errorf("failed to convert function to use yarn spinner values: %w", err)
	}
	storer.addFunction(functionID, yarnSpinnerFunction)
	return nil
}

func (storer *functionStorer) call(functionID string, args []*variable.Value) (*variable.Value, error) {
	function, ok := storer.functionsByID[functionID]
	if !ok {
		return nil, fmt.Errorf("unknown function")
	}
	value, err := function(args)
	if err != nil {
		return nil, fmt.Errorf("execution failed: %w", err)
	}
	return value, nil
}

var typeError = reflect.TypeOf((*error)(nil)).Elem()

func newYarnSpinnerFunction(function any) (YarnSpinnerFunction, error) {
	functionType := reflect.TypeOf(function)
	if functionType.Kind() != reflect.Func {
		return nil, fmt.Errorf("newYarnSpinnerFunction expects an argument which is a function")
	}

	returnSignature, err := checkFunctionOutputParameters(functionType)
	if err != nil {
		return nil, fmt.Errorf("received function has unsupported output parameters: %w", err)
	}

	inputConverter, err := createInputConverter(functionType)
	if err != nil {
		return nil, fmt.Errorf("failed to create input converter: %w", err)
	}

	return func(args []*variable.Value) (*variable.Value, error) {
		inputParameters, err := inputConverter(args)
		if err != nil {
			return nil, fmt.Errorf("failed to convert input parameters: %w", err)
		}

		outputParameters := reflect.ValueOf(function).Call(inputParameters)
		switch returnSignature {
		case noReturn:
			return nil, nil
		case valueReturn:
			value, err := getTreeValue(outputParameters[0])
			if err != nil {
				return nil, fmt.Errorf("failed to convert return value of function: %w", err)
			}
			return value, nil
		case errorReturn:
			if outputParameters[0].Interface() == nil {
				return nil, nil
			}
			err, ok := outputParameters[0].Interface().(error)
			if !ok {
				return nil, fmt.Errorf("function did not return an error value like expected")
			}
			return nil, err
		case valueErrorReturn:
			returnedErr := outputParameters[1].Interface()
			if returnedErr != nil {
				if _, ok := returnedErr.(error); !ok {
					return nil, fmt.Errorf("function did not return an error value like expected")
				} else {
					return nil, returnedErr.(error)
				}
			}
			value, err := getTreeValue(outputParameters[0])
			if err != nil {
				return nil, fmt.Errorf("failed to convert return value of function: %w", err)
			}
			return value, nil
		}
		return nil, nil
	}, nil
}

func isTypeConvertibleToValue(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		return true
	case reflect.Bool:
		return true
	case reflect.String:
		return true
	}
	return false
}

func getTreeValue(v reflect.Value) (*variable.Value, error) {
	switch {
	case v.CanFloat():
		return variable.NewNumber(v.Float()), nil
	case v.CanInt():
		return variable.NewNumber(float64(v.Int())), nil
	case v.Kind() == reflect.Bool:
		return variable.NewBoolean(v.Bool()), nil
	case v.Kind() == reflect.String:
		return variable.NewString(v.String()), nil
	}
	return nil, fmt.Errorf("unsupported value type received when trying to create a tree value")
}

type returnSignature int

const (
	noReturn returnSignature = iota
	valueReturn
	errorReturn
	errorChanReturn
	valueErrorReturn
)

func checkFunctionOutputParameters(functionType reflect.Type) (returnSignature, error) {
	switch functionType.NumOut() {
	case 0:
		return noReturn, nil
	case 1:
		outputType := functionType.Out(0)
		switch {
		case isTypeConvertibleToValue(outputType):
			return valueReturn, nil
		case outputType.ConvertibleTo(typeError):
			return errorReturn, nil
		}
		return 0, fmt.Errorf("a single output parameter must be either convertible to a yarn spinner value or an error")
	case 2:
		firstOutputType, secondOutputType := functionType.Out(0), functionType.Out(1)
		if !isTypeConvertibleToValue(firstOutputType) {
			return 0, fmt.Errorf("the first output parameter of a function with two must be convertible to a yarn spinner value")
		}
		if !secondOutputType.ConvertibleTo(typeError) {
			return 0, fmt.Errorf("the second output parameter of a function with two must be an error")
		}
		return valueErrorReturn, nil
	default:
		return 0, fmt.Errorf("functions can have at most two output parameters")
	}
}

func createInputConverter(functionType reflect.Type) (func([]*variable.Value) ([]reflect.Value, error), error) {
	isVariadic := functionType.IsVariadic()
	if isVariadic {
		return createVariadicInputConverter(functionType)
	}

	numIn := functionType.NumIn()
	argConverters := make([]func(*variable.Value) (reflect.Value, error), 0, numIn)
	for i := 0; i < numIn; i++ {
		argConverter, ok := argConverterByGoalKind[functionType.In(i).Kind()]
		if !ok {
			return nil, fmt.Errorf("argument number %d has an unsupported type %v", i, functionType.In(i).Kind())
		}
		argConverters = append(argConverters, argConverter)
	}
	return func(args []*variable.Value) ([]reflect.Value, error) {
		if len(args) < numIn {
			return nil, fmt.Errorf("received too few arguments")
		} else if len(args) > numIn {
			return nil, fmt.Errorf("received too many arguments")
		}

		inputParameters := make([]reflect.Value, 0, numIn)
		for i := 0; i < numIn; i++ {
			inputParameter, err := argConverters[i](args[i])
			if err != nil {
				return nil, fmt.Errorf("failed to convert argument number %d: %w", i, err)
			}
			inputParameters = append(inputParameters, inputParameter)
		}

		return inputParameters, nil
	}, nil
}

func createVariadicInputConverter(functionType reflect.Type) (func([]*variable.Value) ([]reflect.Value, error), error) {
	numIn := functionType.NumIn()
	argConverters := make([]func(*variable.Value) (reflect.Value, error), 0, numIn)
	for i := 0; i < numIn-1; i++ {
		argConverter, ok := argConverterByGoalKind[functionType.In(i).Kind()]
		if !ok {
			return nil, fmt.Errorf("argument number %d has an unsupported type %v", i, functionType.In(i).Kind())
		}
		argConverters = append(argConverters, argConverter)
	}

	variadicArgsConverter, ok := argConverterByGoalKind[functionType.In(numIn-1).Elem().Kind()]
	if !ok {
		return nil, fmt.Errorf("argument number %d has an unsupported type %v", numIn-1, functionType.In(numIn-1).Kind())
	}

	return func(args []*variable.Value) ([]reflect.Value, error) {
		if len(args) < numIn-1 {
			return nil, fmt.Errorf("received too few arguments")
		}

		inputParameters := make([]reflect.Value, 0, numIn)
		for i := 0; i < numIn-1; i++ {
			inputParameter, err := argConverters[i](args[i])
			if err != nil {
				return nil, fmt.Errorf("failed to convert argument number %d: %w", i, err)
			}
			inputParameters = append(inputParameters, inputParameter)
		}

		for i := numIn - 1; i < len(args); i++ {
			inputParameter, err := variadicArgsConverter(args[i])
			if err != nil {
				return nil, fmt.Errorf("failed to convert argument number %d: %w", i, err)
			}
			inputParameters = append(inputParameters, inputParameter)
		}

		return inputParameters, nil
	}, nil
}

var argConverterByGoalKind map[reflect.Kind]func(*variable.Value) (reflect.Value, error) = map[reflect.Kind]func(*variable.Value) (reflect.Value, error){
	reflect.Int: func(value *variable.Value) (reflect.Value, error) {
		if value.Number == nil {
			return reflect.Value{}, fmt.Errorf("expected a number value but got something else")
		}
		return reflect.ValueOf(int(*value.Number)), nil
	},
	reflect.Int8: func(value *variable.Value) (reflect.Value, error) {
		if value.Number == nil {
			return reflect.Value{}, fmt.Errorf("expected a number value but got something else")
		}
		return reflect.ValueOf(int8(*value.Number)), nil
	},
	reflect.Int16: func(value *variable.Value) (reflect.Value, error) {
		if value.Number == nil {
			return reflect.Value{}, fmt.Errorf("expected a number value but got something else")
		}
		return reflect.ValueOf(int16(*value.Number)), nil
	},
	reflect.Int32: func(value *variable.Value) (reflect.Value, error) {
		if value.Number == nil {
			return reflect.Value{}, fmt.Errorf("expected a number value but got something else")
		}
		return reflect.ValueOf(int32(*value.Number)), nil
	},
	reflect.Int64: func(value *variable.Value) (reflect.Value, error) {
		if value.Number == nil {
			return reflect.Value{}, fmt.Errorf("expected a number value but got something else")
		}
		return reflect.ValueOf(int64(*value.Number)), nil
	},
	reflect.Float32: func(value *variable.Value) (reflect.Value, error) {
		if value.Number == nil {
			return reflect.Value{}, fmt.Errorf("expected a number value but got something else")
		}
		return reflect.ValueOf(float32(*value.Number)), nil
	},
	reflect.Float64: func(value *variable.Value) (reflect.Value, error) {
		if value.Number == nil {
			return reflect.Value{}, fmt.Errorf("expected a number value but got something else")
		}
		return reflect.ValueOf(float64(*value.Number)), nil
	},
	reflect.Bool: func(value *variable.Value) (reflect.Value, error) {
		if value.Boolean == nil {
			return reflect.Value{}, fmt.Errorf("expected a boolean value but got something else")
		}
		return reflect.ValueOf(*value.Boolean), nil
	},
	reflect.String: func(value *variable.Value) (reflect.Value, error) {
		if value.String == nil {
			return reflect.Value{}, fmt.Errorf("expected a string value but got something else")
		}
		return reflect.ValueOf(*value.String), nil
	},
}

var _ functionCaller = (*functionStorer)(nil)
