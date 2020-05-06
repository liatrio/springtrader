package validate

import (
	"fmt"
	"os"
)

func fileExists(filename string) (err error) {
	_, err = os.Stat(filename)
	return err
}

func treeValue(values interface{}, path []interface{}) (string, error) {
	if len(path) == 0 {
		return values.(string), nil
	}
	switch step := path[0].(type) {
	case string:
		v, ok := values.(map[interface{}]interface{})
		if !ok {
			return "", fmt.Errorf("%v is not a map in %v", step, v)
		}
		return treeValue(v[step], path[1:])
	case int:
		v, ok := values.([]interface{})
		if !ok {
			return "", fmt.Errorf("%v is not a slice in %v", step, v)
		}
		return treeValue(v[step], path[1:])
	default:
		return "", fmt.Errorf("cannot navigate path step %v of type %t", step, step)
	}
}

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

func treeCompare(actual interface{}, expected interface{}) error {
	switch expectedType := expected.(type) {
	case map[string]interface{}:
		actualMap, ok := actual.(map[string]interface{})
		if !ok {
			/*
				err := new *treeCompareError
				return new *treeCompareError{ err.Errorf(fmt.Sprintf("actual value is of type %T, expected %T", actualMap, expectedType))}
			*/
			return fmt.Errorf("actual value is of type %T, expectedType %T", actual, expectedType)
		}
		for key := range expectedType {
			if actualMapValue, ok := actualMap[key]; ok {
				err := treeCompare(actualMapValue, expectedType[key])
				if err != nil {
					//err.Add(key)
					return err
				}
			} else {
				//return fmt.Errorf("actual map did not contain key %s", key)
				return fmt.Errorf("actual value is of type %T, expectedType %T", actual, expectedType)
			}
		}
		return nil
	case []interface{}:
		actualSlice, ok := actual.([]interface{})
		if !ok {
			return fmt.Errorf("actual value is of type %T, expectedType %T", actual, expectedType)
		}
		for key := range expectedType {
			if actualSliceValue := actualSlice[key]; ok {
				err := treeCompare(actualSliceValue, expectedType[key])
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("actual slice did not contain key %d", key)
			}
		}
		return nil
	case string:
		actualString, ok := actual.(string)
		if !ok {
			return fmt.Errorf("actual value is of type %T, expectedType %T", actual, expectedType)
		}
		if actualString != expectedType {
			return fmt.Errorf("actual value of %s does not match expectedType string of %s", actualString, expectedType)
		}
		return nil
	case int:
		actualInt, ok := actual.(int)
		if !ok {
			return fmt.Errorf("actual value is of type %T, expectedType %T", actual, expectedType)
		}
		if actualInt != expectedType {
			return fmt.Errorf("actual value of %d does not match expectedType integer of %d", actualInt, expectedType)
		}
		return nil
	case bool:
		actualBool, ok := actual.(bool)
		if !ok {
			return fmt.Errorf("actual value is of type %T, expectedType %T", actual, expectedType)
		}
		if actualBool != expectedType {
			return fmt.Errorf("actual value of %v does not match expectedType boolean of %v", actualBool, expectedType)
		}
		return nil
	}
	return nil
}
