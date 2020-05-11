package validate

import (
	"fmt"
	"io/ioutil"

	"github.com/onsi/gomega/types"
	"gopkg.in/yaml.v2"
)

/*
type treeCompareError struct {
	Path []string
}

func (err *treeCompareError) Error() string {
	var output = "Error found in %s"
	for i := range err.Path {
		output += err.Path[i]
	}
	return output
}

func (err *treeCompareError) Errorf(message string) *treeCompareError {
	err.Path = append(err.Path, message)
	return err
}

func (err *treeCompareError) Add(path string) {
	err.Path = append(err.Path, path)
}
*/

func ValidateYamlObject(expected interface{}) types.GomegaMatcher {
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
			return false, fmt.Errorf("actual value is of type %T, expectedType %T", actual, expectedType)
		}
		for key := range actualMap {
			if expectedTypeValue, ok := expectedType[key.(string)]; ok {
				nestedExpectedObject := validateYaml{expectedTypeValue}
				_, err := nestedExpectedObject.Match(actualMap[key.(string)])
				if err != nil {
					return false, err
				}
			} else {
				return false, fmt.Errorf("actual value %s of type %T does not match up with expected value %s of type %T", actualMap[key.(string)], actualMap, expectedTypeValue, expectedType)
			}
		}
		return true, nil
	case []interface{}:
		actualSlice, ok := actual.([]interface{})
		if !ok {
			return false, fmt.Errorf("actual value is of type %T, expectedType %T", actual, expectedType)
		}
		for key := range actualSlice {
			if expectedTypeValue := expectedType[key]; ok {
				nestedExpectedObject := validateYaml{expectedTypeValue}
				_, err := nestedExpectedObject.Match(actualSlice[key])
				if err != nil {
					return false, err
				}
			} else {
				return false, fmt.Errorf("actual value %s of type %T does not match up with expected value %s of type %T", actualSlice[key], actualSlice, expectedTypeValue, expectedType)
			}
		}
		return true, nil
	case string:
		actualString, ok := actual.(string)
		if !ok {
			return false, fmt.Errorf("actual value is of type %T, expectedType %T", actual, expectedType)
		}
		if actualString != expectedType {
			return false, fmt.Errorf("actual value of %s does not match expectedType string of %s", actualString, expectedType)
		}
		return true, nil
	case int:
		actualInt, ok := actual.(int)
		if !ok {
			return false, fmt.Errorf("actual value is of type %T, expectedType %T", actual, expectedType)
		}
		if actualInt != expectedType {
			return false, fmt.Errorf("actual value of %d does not match expectedType integer of %d", actualInt, expectedType)
		}
		return true, nil
	case bool:
		actualBool, ok := actual.(bool)
		if !ok {
			return false, fmt.Errorf("actual value is of type %T, expectedType %T", actual, expectedType)
		}
		if actualBool != expectedType {
			return false, fmt.Errorf("actual value of %v does not match expectedType boolean of %v", actualBool, expectedType)
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
	failMessage = fmt.Sprintf("File at the path, %s, could not be parsed as YAML.\n Error: %s\n", path, err)
	Expect(err).To(BeNil(), failMessage)
	return output
}
