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
		if val, ok := TryGetFieldValue(obj, fieldName); !ok || val == "" {
			return false, fmt.Errorf("expected %s to be set", fieldName)
		}
	}

	return true, nil
}

func RequiredFieldsCheck(obj interface{}, fields []string) (bool, error) {
	for _, fieldName := range fields {
		if !HasField(obj, fieldName) {
			return false, fmt.Errorf("expected object to have field '%s'", fieldName)
		}
	}

	return true, nil
}
