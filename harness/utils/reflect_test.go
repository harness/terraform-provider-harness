package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type CustomString string

var CustomStrings = struct {
	Foo CustomString
}{
	Foo: "test",
}

type TestObj struct {
	Custom   CustomString
	Foo      string
	ObjField interface{}
}

func TestRequiredStringFieldsSet_FieldDoesNotExist(t *testing.T) {
	obj := &TestObj{}
	ok, err := RequiredStringFieldsSet(obj, []string{"Bar"})
	require.False(t, ok)
	require.NotNil(t, err)
}

func TestRequiredStringFieldsSet_FieldExistsButIsNotset(t *testing.T) {
	obj := &TestObj{}
	ok, err := RequiredStringFieldsSet(obj, []string{"Foo"})
	require.False(t, ok)
	require.NotNil(t, err)
}

func TestRequiredStringFieldsSet_FieldExistsButIsSetToEmpty(t *testing.T) {
	obj := &TestObj{
		Foo: "",
	}
	ok, err := RequiredStringFieldsSet(obj, []string{"Foo"})
	require.False(t, ok)
	require.NotNil(t, err)
}

func TestRequiredStringFieldsSet_FieldIsSet(t *testing.T) {
	obj := &TestObj{
		Foo: "test",
	}
	ok, err := RequiredStringFieldsSet(obj, []string{"Foo"})
	require.True(t, ok)
	require.Nil(t, err)
}

func TestRequiredStringFieldsSet_FieldIsObject_NotSet(t *testing.T) {
	obj := &TestObj{}

	ok, err := RequiredStringFieldsSet(obj, []string{"ObjField"})
	require.False(t, ok)
	require.NotNil(t, err)
}

func TestRequiredStringFieldsSet_FieldIsObject_Set(t *testing.T) {
	obj := &TestObj{
		ObjField: []string{},
	}

	ok, err := RequiredStringFieldsSet(obj, []string{"ObjField"})
	require.True(t, ok)
	require.Nil(t, err)
}

func TestRequiredStringFieldsSet_CustomType_IsSet(t *testing.T) {
	obj := &TestObj{
		Custom: CustomStrings.Foo,
	}

	ok, err := RequiredStringFieldsSet(obj, []string{"Custom"})
	require.True(t, ok)
	require.Nil(t, err)
}

func TestRequiredFieldsSet_IsSet(t *testing.T) {
	obj := &TestObj{
		Custom: CustomStrings.Foo,
	}

	ok, err := RequiredFieldsSet(obj, map[string]interface{}{"Custom": CustomString("")})
	require.True(t, ok)
	require.Nil(t, err)
}

func TestRequiredFieldsSet_IsNotSet(t *testing.T) {
	obj := &TestObj{}

	ok, err := RequiredFieldsSet(obj, map[string]interface{}{"Custom": CustomString("")})
	require.False(t, ok)
	require.NotNil(t, err)
}
