package utils

import (
	"fmt"
	"reflect"
)

func MustSetField(i interface{}, fieldName string, value interface{}) {
	if ok, err := TrySetField(i, fieldName, value); !ok {
		panic(err)
	}
}

func TrySetField(i interface{}, fieldName string, value interface{}) (bool, error) {
	valueI := reflect.ValueOf(i)

	// Check if the passed interface is a pointer
	if valueI.Type().Kind() != reflect.Ptr {
		// Create a new type of i's Type, so we have a pointer to work with
		valueI = reflect.New(reflect.TypeOf(i))
	}

	// 'dereference' with Elem() and get the field by name
	field := valueI.Elem().FieldByName(fieldName)
	if !field.IsValid() {
		return false, fmt.Errorf("interface `%s` does not have the field `%s`", valueI.Type(), fieldName)
	}

	if !field.CanSet() {
		return false, fmt.Errorf("unable to set field `%s` with value `%s`", fieldName, value)
	}

	field.Set(reflect.ValueOf(value))

	return true, nil
}

func HasField(i interface{}, fieldName string) bool {
	valueI := reflect.ValueOf(i)

	// Check if the passed interface is a pointer
	if valueI.Type().Kind() != reflect.Ptr {
		// Create a new type of i's Type, so we have a pointer to work with
		valueI = reflect.New(reflect.TypeOf(i))
	}

	// 'dereference' with Elem() and get the field by name
	return valueI.Elem().FieldByName(fieldName).IsValid()
}

func TryGetFieldValue(i interface{}, fieldName string) (interface{}, bool) {
	valueI := reflect.ValueOf(i)

	// Check if the passed interface is a pointer
	if valueI.Type().Kind() != reflect.Ptr {
		// Create a new type of i's Type, so we have a pointer to work with
		valueI = reflect.New(reflect.TypeOf(i))
	}

	// 'dereference' with Elem() and get the field by name
	field := valueI.Elem().FieldByName(fieldName)
	if !field.IsValid() {
		return nil, false
	}

	return field.Interface(), true
}

func RequiredStringFieldsSet(obj interface{}, fields []string) (bool, error) {
	for _, fieldName := range fields {
		val, _ := TryGetFieldValue(obj, fieldName)
		if val == nil || val == "" {
			return false, fmt.Errorf("expected %s to be set", fieldName)
		}
	}

	return true, nil
}

// Takes an object with a list of field names and their default values.
// Returns true if the all of the fields are set with a non-default value.
// Example usage:
// 		if ok, err := utils.RequiredFieldsSetWithDefaultValues(obj, map[string]interface{}{
// 			"field1": "default1",
// 			"field2": "default2",
// 		})
// 		if !ok {
// 			panic(err)
// 		}
func RequiredFieldsSetWithDefaultValues(obj interface{}, fieldValues map[string]interface{}) (bool, error) {
	for fieldName, defaultValue := range fieldValues {
		if val, ok := TryGetFieldValue(obj, fieldName); !ok || val == defaultValue {
			return false, fmt.Errorf("expected %s to be set", fieldName)
		}
	}

	return true, nil
}
func RequiredFieldsSet(obj interface{}, defaultValues map[string]interface{}) (bool, error) {
	for fieldName, value := range defaultValues {
		if !HasField(obj, fieldName) {
			return false, fmt.Errorf("expected object to have field '%s'", fieldName)
		}

		if v, ok := TryGetFieldValue(obj, fieldName); !ok || v == value {
			return false, fmt.Errorf("expected object to have field '%s' set with non-default value", fieldName)
		}

		return true, nil
	}

	return true, nil
}

func RequiredValuesSet(obj interface{}, fieldValues map[string]interface{}) (bool, error) {
	for fieldName, expectedValue := range fieldValues {
		if val, ok := TryGetFieldValue(obj, fieldName); !ok || val != expectedValue {
			return false, fmt.Errorf("expected %s to be set to %s", fieldName, expectedValue)
		}
	}

	return true, nil
}

func RequiredValueOptionsSet(obj interface{}, fieldValuesMap map[string][]interface{}) (bool, error) {
	for fieldName, expectedValue := range fieldValuesMap {
		if val, ok := TryGetFieldValue(obj, fieldName); !ok {
			valueMatch := false
			for v := range expectedValue {
				if v == val {
					valueMatch = true
					break
				}
			}

			if !valueMatch {
				return false, fmt.Errorf("expected %s to be one of %s", fieldName, expectedValue)
			}
		}
	}

	return true, nil
}

// func RequiredFieldsSet(obj interface{}, fields []string) (bool, error) {
// 	for _, fieldName := range fields {
// 		if val, ok := TryGetFieldValue(obj, fieldName); !ok || val == nil {
// 			return false, fmt.Errorf("expected %s to be set", fieldName)
// 		}
// 	}

// 	return true, nil
// }

// func RequiredSliceFieldsSet(obj interface{}, fields []string) (bool, error) {
// 	for _, fieldName := range fields {
// 		if val, ok := TryGetFieldValue(obj, fieldName); !ok || val == nil {
// 			return false, fmt.Errorf("expected %s to be set", fieldName)
// 		}
// 	}

// 	return true, nil
// }

func RequiredFieldsCheck(obj interface{}, fields []string) (bool, error) {
	for _, fieldName := range fields {
		if !HasField(obj, fieldName) {
			return false, fmt.Errorf("expected object to have field '%s'", fieldName)
		}
	}

	return true, nil
}
