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
