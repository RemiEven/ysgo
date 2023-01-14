package markup

import (
	"errors"
	"fmt"
	"strings"
)

type markerProcessor func(*attributeMarker) (string, error)

func getProcessor(name string) markerProcessor {
	switch name {
	case "nomarkup":
		return processNoMarkup
	case "select":
		return processSelect
	case "plural":
		return processPlural
	case "ordinal":
		return processOrdinal
	}
	return nil
}

func processNoMarkup(marker *attributeMarker) (string, error) {
	content, ok := marker.GetProperty(replacementMarkerContents)
	if !ok {
		return "", nil
	}
	return content.toString(), nil
}

func processSelect(marker *attributeMarker) (string, error) {
	value, ok := marker.GetProperty("value")
	if !ok {
		return "", errors.New(`did not find a "value" property`)
	}
	valueString := value.toString()

	replacement, ok := marker.GetProperty(valueString)
	if !ok {
		return "", fmt.Errorf(`no replacement found for "` + valueString + `"`)
	}
	replacementString := replacement.toString()

	replacementString = replacePlaceholders(replacementString, valueString)

	return replacementString, nil
}

func processPlural(marker *attributeMarker) (string, error) {
	value, ok := marker.GetProperty("value")
	if !ok {
		return "", errors.New(`did not find a "value" property`)
	}
	pluralCase := pluralCaseOther
	switch value.ValueType {
	case ValueTypeFloat:
		// nothing to do
	case ValueTypeInteger:
		if value.IntegerValue == 1 {
			pluralCase = pluralCaseOne
		}
	default:
		return "", errors.New(`"value" property wasn't a number`)
	}

	replacement, ok := marker.GetProperty(pluralCase)
	if !ok {
		return "", fmt.Errorf(`no replacement found for case "` + pluralCase + `"`)
	}
	replacementString := replacement.toString()

	replacementString = replacePlaceholders(replacementString, value.toString())

	return replacementString, nil
}

func processOrdinal(marker *attributeMarker) (string, error) {
	value, ok := marker.GetProperty("value")
	if !ok {
		return "", errors.New(`did not find a "value" property`)
	}
	if value.ValueType != ValueTypeInteger {
		return "", errors.New(`"value" property wasn't an integer`)
	}
	pluralCase := pluralCaseOther
	n := value.IntegerValue
	switch {
	case (n%10 == 1) && (n%100 != 11):
		pluralCase = pluralCaseOne
	case (n%10 == 2) && (n%100 != 12):
		pluralCase = pluralCaseTwo
	case (n%10 == 3) && (n%100 != 13):
		pluralCase = pluralCaseFew
	}

	replacement, ok := marker.GetProperty(pluralCase)
	if !ok {
		return "", fmt.Errorf(`no replacement found for case "` + pluralCase + `"`)
	}
	replacementString := replacement.toString()

	replacementString = replacePlaceholders(replacementString, value.toString())

	return replacementString, nil
}

func replacePlaceholders(replacement, value string) string {
	if strings.Count(replacement, "%") == 0 {
		return replacement
	}
	r := strings.ReplaceAll(replacement, "%", value)
	return strings.ReplaceAll(r, "\\"+value, "%")
}

const (
	pluralCaseOne   = "one"
	pluralCaseTwo   = "two"
	pluralCaseFew   = "few"
	pluralCaseMany  = "many"
	pluralCaseOther = "other"
)
