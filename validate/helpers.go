package validate

import (
	"fmt"
	"internal/testlog"
	"os"
	"reflect"
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

func treeCompare(actual interface{}, expected interface{}) error {
	switch expectedType := expected.(type) {
	case map[string]interface{}:
		if actualMap, ok := actual.(map[string]interface{}); !ok {
			return fmt.Errorf("actual value is of type %T, expected %T", actual, expected)
		}
		for key, value := range expected {
			if v, ok := actualMap[key]; ok {
				err := treeCompare(v, expected[key])
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("actual map did not contain key %s", key)
			}
		}
		return nil
	case []interface{}:
		if actualSlice, ok := actual.([]interface{}); !ok {
			return fmt.Errorf("actual value is of type %T, expected %T", actual, expected)
		}
		for key, value := range expected {
			if v, ok := actualSlice[key]; ok {
				err := treeCompare(v, expected[key])
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("actual slice did not contain key %d", key)
			}
		}
		return nil
	case string:
		if actualString, ok := actual.(string); !ok {
			return fmt.Errorf("actual value is of type %T, expected %T", actual, expected)
		}
		if actualString != string(expected) {
			return fmt.Errorf("actual value of %s does not match expected string of %s", actualString, expected)
		}
		return nil
	case int:
		if actualInt, ok := actual.(int); !ok {
			return fmt.Errorf("actual value is of type %T, expected %T", actual, expected)
		}
		if actualInt != bool(expected) {
			return fmt.Errorf("actual value of %d does not match expected integer of %d", actualInt, expected)
		}
		return nil
	case bool:
		if actualBool, ok := actual.(bool); !ok {
			return fmt.Errorf("actual value is of type %T, expected %T", actual, expected)
		}
		if actualBool != bool(expected) {
			return fmt.Errorf("actual value of %v does not match expected boolean of %v", actualString, expected)
		}
		return nil
	}
}
