package runner

import (
	"fmt"
	"reflect"
	"time"

	"github.com/remieven/ysgo/variable"
)

type commandCaller interface {
	call(commandID string, args []*variable.Value) <-chan error
}

// YarnSpinnerCommand is a Go function than can be called as a command from a YarnSpinner script.
// It can take zero, one or more input values, and returns a chan that must return either nil or
// an error when the command has finished its job.
type YarnSpinnerCommand func([]*variable.Value) <-chan error

type commandStorer struct {
	commandsByID map[string]YarnSpinnerCommand
}

var _ commandCaller = (*commandStorer)(nil)

func newCommandStorer() *commandStorer {
	storer := &commandStorer{
		commandsByID: map[string]YarnSpinnerCommand{
			"wait": waitCommand,
		},
	}

	return storer
}

func (storer *commandStorer) call(commandID string, args []*variable.Value) <-chan error {
	command, ok := storer.commandsByID[commandID]
	if !ok {
		errChan := make(chan error, 1)
		errChan <- fmt.Errorf("unknown command")
		return errChan
	}
	return command(args)
}

func chanWithImmediateValue[V any](value V) <-chan V {
	ch := make(chan V, 1)
	ch <- value
	return ch
}

func waitCommand(args []*variable.Value) <-chan error {
	if len(args) != 1 {
		return chanWithImmediateValue(fmt.Errorf("expected exactly one argument"))
	}
	duration := args[0]
	if duration.Number == nil {
		return chanWithImmediateValue(fmt.Errorf("received a duration which was not a number"))
	}
	ch := make(chan error, 1)
	go func() {
		time.Sleep(time.Duration(*duration.Number) * time.Second)
		ch <- nil
	}()
	return ch
}

func (storer *commandStorer) addCommand(commandID string, command YarnSpinnerCommand) {
	storer.commandsByID[commandID] = command
}

func (storer *commandStorer) convertAndAddCommand(commandID string, command any) error {
	yarnSpinnerCommand, err := newYarnSpinnerCommand(command)
	if err != nil {
		return fmt.Errorf("failed to convert command to use yarn spinner values: %w", err)
	}
	storer.addCommand(commandID, yarnSpinnerCommand)
	return nil
}

func newYarnSpinnerCommand(command any) (YarnSpinnerCommand, error) {
	commandType := reflect.TypeOf(command)
	if commandType.Kind() != reflect.Func {
		return nil, fmt.Errorf("newYarnSpinnerCommand expects an argument which is a function")
	}

	returnSignature, err := checkCommandOutputParameters(commandType)
	if err != nil {
		return nil, fmt.Errorf("received command has unsupported output parameters: %w", err)
	}

	inputConverter, err := createInputConverter(commandType)
	if err != nil {
		return nil, fmt.Errorf("failed to create input converter: %w", err)
	}

	return func(args []*variable.Value) <-chan error {
		errChan := make(chan error, 1)
		inputParameters, err := inputConverter(args)
		if err != nil {
			errChan <- fmt.Errorf("failed to convert input parameters: %w", err)
			return errChan
		}

		if returnSignature == errorChanReturn {
			outputParameters := reflect.ValueOf(command).Call(inputParameters)
			if outputParameters[0].IsNil() {
				errChan <- fmt.Errorf("command returned a nil chan")
				return errChan
			}
			switch returnChan := outputParameters[0].Interface().(type) {
			case <-chan error:
				return returnChan
			case chan error:
				return returnChan
			}
			errChan <- fmt.Errorf("command did not return a chan error like expected")
			return errChan
		}

		go func() {
			outputParameters := reflect.ValueOf(command).Call(inputParameters)
			switch returnSignature {
			case noReturn:
				errChan <- nil
			case errorReturn:
				if outputParameters[0].Interface() == nil {
					errChan <- nil
					return
				}
				err, ok := outputParameters[0].Interface().(error)
				if !ok {
					errChan <- fmt.Errorf("command did not return an error value like expected")
					return
				}
				errChan <- err
				return
			}
		}()

		return errChan
	}, nil
}

func checkCommandOutputParameters(commandType reflect.Type) (returnSignature, error) {
	switch commandType.NumOut() {
	case 0:
		return noReturn, nil
	case 1:
		outputType := commandType.Out(0)
		switch {
		case outputType.ConvertibleTo(typeError):
			return errorReturn, nil
		case isTypeErrChan(outputType):
			return errorChanReturn, nil
		}
		return 0, fmt.Errorf("a single output parameter must be either convertible to an error or to a chan error")
	default:
		return 0, fmt.Errorf("commands can have at most one output parameter")
	}
}

var _ commandCaller = (*commandStorer)(nil)

func isTypeErrChan(t reflect.Type) bool {
	if t.Kind() != reflect.Chan {
		return false
	}
	return t.Elem().ConvertibleTo(typeError)
}
