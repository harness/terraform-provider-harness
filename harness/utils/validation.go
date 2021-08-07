package utils

import "fmt"

func CheckRequiredParameters(param string, badValue interface{}) (bool, error) {
	if param == badValue {
		return false, fmt.Errorf("expected '%s' to not be '%s'", param, badValue)
	}

	return true, nil
}
