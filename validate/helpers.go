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

type treeCompareError struct {
  Path []string
}

func (err *treeCompareError) Error() (string) {

}

func (err *treeCompareError) Add(path string) {
	err.Path.append(path)
}

func treeCompare(actual interface{}, expected interface{}) (treeCompareError) {
	// if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
	// 	return fmt.Errorf("actual not same type as expected")
	// }

	switch expectedType := expected.(type) {
	case map[string]interface{}:
		if actualMap, ok := actual.(map[string]interface{}); !ok {
			return fmt.Errorf("actual not object(map)")
		}
		for key, value := range expectedType {
			if v, ok := actualMap[key]; ok {
				err := treeCompare(actualMap[key], expectedMap[key])
				if err != nil {
					err.Add(key)
					return err;
				}
			} else {

			}
		}
	}

	switch reflect.TypeOf(actual).Kind() {
	case reflect.Map:
		expected.(map[string]interface{}) range
	case reflect.Slice:
	case reflect.String:
		if actual.(string) == expected.(string) {
			return nil
		} else {
			return fmt.Errorf("%s does not equal %s", actual, expected)
		}
	case reflect.Bool:
	}
}