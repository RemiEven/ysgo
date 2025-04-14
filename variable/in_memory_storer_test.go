package variable

import (
	"testing"

	"github.com/remieven/ysgo/internal/testutils"
)

func TestInMemoryStorer(t *testing.T) {
	storer := NewInMemoryStorer()

	t.Run("number storage", func(t *testing.T) {
		const varName = "number"
		if storer.Contains(varName) {
			t.Errorf("Contains returned that variable was stored when it was not")
			return
		}
		if _, ok := storer.GetValue(varName); ok {
			t.Errorf("a value was returned by GetValue when it should not be")
		}

		firstValue := 6.4
		storer.SetNumberValue(varName, firstValue)
		if !storer.Contains(varName) {
			t.Errorf("Contains returned that variable was not stored when it should be")
			return
		}
		if value, ok := storer.GetValue(varName); !ok {
			t.Errorf("no value was returned by GetValue when one should have been")
			return
		} else if value.Number == nil {
			t.Errorf("expected the returned value %+v to be a number but was not", value)
			return
		} else if *value.Number != firstValue {
			t.Errorf("unexpected returned number: got [%f], wanted [%f]", *value.Number, firstValue)
			return
		}

		secondValue := 9.3
		storer.SetNumberValue(varName, secondValue)
		if value, ok := storer.GetValue(varName); !ok {
			t.Errorf("no value was returned by GetValue when one should have been")
			return
		} else if value.Number == nil {
			t.Errorf("expected the returned value %+v to be a number but was not", value)
			return
		} else if *value.Number != secondValue {
			t.Errorf("unexpected returned number: got [%f], wanted [%f]", *value.Number, firstValue)
			return
		}

		storer.Clear()
		if storer.Contains(varName) {
			t.Errorf("Contains returned that variable was stored when it should have been cleared")
			return
		}
	})

	t.Run("boolean storage", func(t *testing.T) {
		const varName = "boolean"
		if storer.Contains(varName) {
			t.Errorf("Contains returned that variable was stored when it was not")
			return
		}
		if _, ok := storer.GetValue(varName); ok {
			t.Errorf("a value was returned by GetValue when it should not be")
		}

		firstValue := true
		storer.SetBooleanValue(varName, firstValue)
		if !storer.Contains(varName) {
			t.Errorf("Contains returned that variable was not stored when it should be")
			return
		}
		if value, ok := storer.GetValue(varName); !ok {
			t.Errorf("no value was returned by GetValue when one should have been")
			return
		} else if value.Boolean == nil {
			t.Errorf("expected the returned value %+v to be a boolean but was not", value)
			return
		} else if *value.Boolean != firstValue {
			t.Errorf("unexpected returned boolean: got [%t], wanted [%t]", *value.Boolean, firstValue)
			return
		}

		secondValue := false
		storer.SetBooleanValue(varName, secondValue)
		if value, ok := storer.GetValue(varName); !ok {
			t.Errorf("no value was returned by GetValue when one should have been")
			return
		} else if value.Boolean == nil {
			t.Errorf("expected the returned value %+v to be a boolean but was not", value)
			return
		} else if *value.Boolean != secondValue {
			t.Errorf("unexpected returned boolean: got [%t], wanted [%t]", *value.Boolean, firstValue)
			return
		}

		storer.Clear()
		if storer.Contains(varName) {
			t.Errorf("Contains returned that variable was stored when it should have been cleared")
			return
		}
	})

	t.Run("string storage", func(t *testing.T) {
		const varName = "string"
		if storer.Contains(varName) {
			t.Errorf("Contains returned that variable was stored when it was not")
			return
		}
		if _, ok := storer.GetValue(varName); ok {
			t.Errorf("a value was returned by GetValue when it should not be")
		}

		firstValue := "a first string value"
		storer.SetStringValue(varName, firstValue)
		if !storer.Contains(varName) {
			t.Errorf("Contains returned that variable was not stored when it should be")
			return
		}
		if value, ok := storer.GetValue(varName); !ok {
			t.Errorf("no value was returned by GetValue when one should have been")
			return
		} else if value.String == nil {
			t.Errorf("expected the returned value %+v to be a string but was not", value)
			return
		} else if *value.String != firstValue {
			t.Errorf("unexpected returned string: got [%s], wanted [%s]", *value.String, firstValue)
			return
		}

		secondValue := "a second string value"
		storer.SetStringValue(varName, secondValue)
		if value, ok := storer.GetValue(varName); !ok {
			t.Errorf("no value was returned by GetValue when one should have been")
			return
		} else if value.String == nil {
			t.Errorf("expected the returned value %+v to be a string but was not", value)
			return
		} else if *value.String != secondValue {
			t.Errorf("unexpected returned string: got [%s], wanted [%s]", *value.String, firstValue)
			return
		}

		storer.Clear()
		if storer.Contains(varName) {
			t.Errorf("Contains returned that variable was stored when it should have been cleared")
			return
		}
	})

	t.Run("Export of empty store", func(t *testing.T) {
		storer := NewInMemoryStorer()

		if len(storer.GetValues()) != 0 {
			t.Errorf("expected GetValues of empty storer to return 0 values but got %v", len(storer.GetValues()))
		}
	})

	t.Run("Export of filled store", func(t *testing.T) {
		storer := NewInMemoryStorer()
		storer.SetBooleanValue("bool1", true)
		storer.SetBooleanValue("bool2", false)
		storer.SetNumberValue("number", 70400)
		storer.SetStringValue("string", "some string value")

		expected := map[string]Value{
			"bool1":  *NewBoolean(true),
			"bool2":  *NewBoolean(false),
			"number": *NewNumber(70400),
			"string": *NewString("some string value"),
		}

		if diff := testutils.DeepEqual(storer.GetValues(), expected); diff != "" {
			t.Error("unexpected result from GetValues: " + diff)
		}

	})
}
