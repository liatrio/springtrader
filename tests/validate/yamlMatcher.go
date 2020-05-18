package validate

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/onsi/gomega/types"
	"gopkg.in/yaml.v2"
)

func ValidateYamlObject(expected interface{}, failureMessage *string) types.GomegaMatcher {
	return &validateYaml{
		expected: expected,
	}
}

type validateYaml struct {
	expected interface{}
}

func (matcher *validateYaml) Match(actual interface{}) (success bool, err error) {
	switch expectedType := matcher.expected.(type) {
	case map[interface{}]interface{}:
		actualMap, ok := actual.(map[interface{}]interface{})
		if !ok {
			return false, typeMismatchError(actual, expectedType)
		}
		for key := range actualMap {
			if expectedTypeValue, ok := expectedType[key.(string)]; ok {
				nestedExpectedObject := validateYaml{expectedTypeValue}
				_, err := nestedExpectedObject.Match(actualMap[key.(string)])
				if err != nil {
					return false, recursiveCallError(nestedExpectedObject, actualMap[key.(string)], err)
				}
			} else {
				return false, keyComparisonError(key, actual, matcher.expected)
			}
		}
		return true, nil
	case []interface{}:
		actualSlice, ok := actual.([]interface{})
		if !ok {
			return false, typeMismatchError(actual, expectedType)
		}
		for i := range actualSlice {
			if expectedTypeValue := expectedType[i]; ok {
				nestedExpectedObject := validateYaml{expectedTypeValue}
				_, err := nestedExpectedObject.Match(actualSlice[i])
				if err != nil {
					return false, recursiveCallError(nestedExpectedObject, actualSlice[i], err)
				}
			} else {
				return false, keyComparisonError(i, actual, matcher.expected)
			}
		}
		return true, nil
	case string:
		actualString, ok := actual.(string)
		if !ok {
			return false, typeMismatchError(actual, expectedType)
		}
		if actualString != expectedType {
			return false, valueComparisonError(actualString, nil, expectedType, nil)
		}
		return true, nil
	case int:
		actualInt, ok := actual.(int)
		if !ok {
			return false, typeMismatchError(actual, expectedType)
		}
		if actualInt != expectedType {
			return false, valueComparisonError(actualInt, nil, expectedType, nil)
		}
		return true, nil
	case float64:
		actualFloat, ok := actual.(float64)
		if !ok {
			return false, typeMismatchError(actual, expectedType)
		}
		if actualFloat != expectedType {
			return false, valueComparisonError(actualFloat, nil, expectedType, nil)
		}
		return true, nil
	case bool:
		actualBool, ok := actual.(bool)
		if !ok {
			return false, typeMismatchError(actual, expectedType)
		}
		if actualBool != expectedType {
			return false, valueComparisonError(actualBool, nil, expectedType, nil)
		}
		return true, nil
	case nil:
		if actual != nil {
			return false, typeMismatchError(actual, expectedType)
		}
		return true, nil
	default:
		return false, fmt.Errorf("Type of %T did not match any expected types", expectedType)
	}
}

func (matcher *validateYaml) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %v to be the same value as %v", actual, matcher.expected)
}

func (matcher *validateYaml) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %v to be the same value as %v", actual, matcher.expected)
}

func ExpectYamlToParse(path string) (interface{}, string) {
	var output interface{}
	file, err := ioutil.ReadFile(path)
	failMessage := fmt.Sprintf("Your %s file cannot be found. File may be in wrong location or misnamed.\n", filepath.Base(path))
	if err != nil {
		return nil, failMessage
	}
	err = yaml.Unmarshal([]byte(file), &output)
	failMessage = fmt.Sprintf("Your %s file could not be read as YAML. Possible issues with formatting: %s\n", filepath.Base(path), err)
	if err != nil {
		return nil, failMessage
	}
	return output, ""
}

func typeMismatchError(actual interface{}, expected interface{}) error {
	return fmt.Errorf("Your value type %T, is not the same as the correct type, %T", expected, actual)
}

func valueComparisonError(actual interface{}, actualValue interface{}, expected interface{}, expectedValue interface{}) error {
	var errStr string
	switch actualType := actual.(type) {
	case string:
		errStr = fmt.Sprintf("Your value, %v, did not have the correct value, %v", expected.(string), actualType)
	case int:
		errStr = fmt.Sprintf("Your value, %d, did not have the correct value, %d", expected.(int), actualType)
	case float64:
		errStr = fmt.Sprintf("Your value, %f, did not have the correct value, %f", expected.(float64), actualType)
	case bool:
		errStr = fmt.Sprintf("Your value, %t, did not have the correct value, %t", expected.(bool), actualType)
	case nil:
		errStr = fmt.Sprintf("Your value should have been empty")
	default:
		errStr = fmt.Sprintf("Your %T with value, %T, did not have the correct value, %T, of field,  %T", actual, actualValue, expected, expectedValue)
	}
	return fmt.Errorf(errStr)
}

func keyComparisonError(key interface{}, actual interface{}, expected interface{}) error {
	var errStr string
	switch actualType := actual.(type) {
	case map[interface{}]interface{}:
		expectedMap := expected.(map[interface{}]interface{})
		errStr = fmt.Sprintf("There was a mismatch between the correct field, %s, and one of your fields:", key.(string))
		for expectedKey := range expectedMap {
			errStr += fmt.Sprintf(" %s", expectedKey)
		}
	case []interface{}:
		expectedList := expected.([]interface{})
		errStr = fmt.Sprintf("Your file has %d fields, which does not match the correct number amount of fields, %d", len(expectedList), len(actualType))
	default:
	}
	return fmt.Errorf(errStr)
}

func recursiveCallError(expected interface{}, actual interface{}, err error) error {
	return err
}
