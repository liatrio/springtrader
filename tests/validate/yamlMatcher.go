package validate

import (
	"fmt"
	"io/ioutil"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"gopkg.in/yaml.v2"
)

func ValidateYamlObject(expected interface{}, failureMessage *string) types.GomegaMatcher {
	return &validateYaml{
		expected:       expected,
		failureMessage: failureMessage,
	}
}

type validateYaml struct {
	expected       interface{}
	failureMessage *string
}

func (matcher *validateYaml) Match(actual interface{}) (success bool, err error) {
	switch expectedType := matcher.expected.(type) {
	case map[interface{}]interface{}:
		actualMap, ok := actual.(map[interface{}]interface{})
		if !ok {
			return false, typeMismatchError(matcher, actual, expectedType)
		}
		for key, value := range actualMap {
			if expectedTypeValue, ok := expectedType[key.(string)]; ok {
				nestedExpectedObject := validateYaml{expectedTypeValue, matcher.failureMessage}
				_, err := nestedExpectedObject.Match(actualMap[key.(string)])
				if err != nil {
					return false, recursiveCallError(matcher, nestedExpectedObject, actualMap[key.(string)], err)
				}
			} else {
				return false, valueComparisonError(matcher, actual, value, expectedType, expectedTypeValue)
			}
		}
		return true, nil
	case []interface{}:
		actualSlice, ok := actual.([]interface{})
		if !ok {
			return false, typeMismatchError(matcher, actual, expectedType)
		}
		for i, value := range actualSlice {
			if expectedTypeValue := expectedType[i]; ok {
				nestedExpectedObject := validateYaml{expectedTypeValue, matcher.failureMessage}
				_, err := nestedExpectedObject.Match(actualSlice[i])
				if err != nil {
					return false, recursiveCallError(matcher, nestedExpectedObject, actualSlice[i], err)
				}
			} else {
				return false, valueComparisonError(matcher, actual, value, expectedType, expectedTypeValue)
			}
		}
		return true, nil
	case string:
		actualString, ok := actual.(string)
		if !ok {
			return false, typeMismatchError(matcher, actual, expectedType)
		}
		if actualString != expectedType {
			return false, valueComparisonError(matcher, actualString, nil, expectedType, nil)
		}
		return true, nil
	case int:
		actualInt, ok := actual.(int)
		if !ok {
			return false, typeMismatchError(matcher, actual, expectedType)
		}
		if actualInt != expectedType {
			return false, valueComparisonError(matcher, actualInt, nil, expectedType, nil)
		}
		return true, nil
	case bool:
		actualBool, ok := actual.(bool)
		if !ok {
			return false, typeMismatchError(matcher, actual, expectedType)
		}
		if actualBool != expectedType {
			return false, valueComparisonError(matcher, actualBool, nil, expectedType, nil)
		}
		return true, nil
	default:
		return false, fmt.Errorf("expectedType of %T did not match any expected types", expectedType)
	}
}

func (matcher *validateYaml) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto contain the JSON representation of\n\t%#v", actual, matcher.expected)
}

func (matcher *validateYaml) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nnot to contain the JSON representation of\n\t%#v", actual, matcher.expected)
}

func ExpectYamlToParse(path string) interface{} {
	var output interface{}
	file, err := ioutil.ReadFile(path)
	failMessage := fmt.Sprintf("File at the path, %s, cannot be found. File may be in wrong location or misnamed.\n", path)
	Expect(err).To(BeNil(), failMessage)
	err = yaml.Unmarshal([]byte(file), &output)
	failMessage = fmt.Sprintf("File at the path, %s, could not be parsed as YAML. Error: %s\n", path, err)
	Expect(err).To(BeNil(), failMessage)
	return output
}

func typeMismatchError(matcher *validateYaml, actual interface{}, expected interface{}) error {
	errStr := fmt.Sprintf("%s; Your value type %T, is not the same as the correct type, %T", *matcher.failureMessage, expected, actual)
	matcher.failureMessage = &errStr
	return fmt.Errorf(errStr)
}

func valueComparisonError(matcher *validateYaml, actual interface{}, actualValue interface{}, expected interface{}, expectedValue interface{}) error {
	var errStr string
	if actualValue == nil && expectedValue == nil {
		switch actualType := actual.(type) {
		case string:
			errStr = fmt.Sprintf("%s; Your value, %v, did not have the correct value, %v", *matcher.failureMessage, actualType, expected.(string))
		case int:
			errStr = fmt.Sprintf("%s; Your value, %d, did not have the correct value, %d", *matcher.failureMessage, actualType, expected.(int))
		case bool:
		default:
			errStr = fmt.Sprintf("%s; Your value, %t, did not have the correct value, %t", *matcher.failureMessage, actualType, expected.(bool))
		}
	} else {
		errStr = fmt.Sprintf("%s; Your %T with value, %T, did not have the correct value, %T , of field,  %T", *matcher.failureMessage, actual, actualValue, expected, expectedValue)
	}
	matcher.failureMessage = &errStr
	return fmt.Errorf(errStr)
}

func recursiveCallError(matcher *validateYaml, expected interface{}, actual interface{}, err error) error {
	return err
}
