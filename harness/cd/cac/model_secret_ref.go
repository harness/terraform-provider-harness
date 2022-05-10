package cac

import (
	"errors"
	"fmt"
	"strings"
)

type SecretRef struct {
	Name string
}

func (r *SecretRef) MarshalYAML() (interface{}, error) {
	if (r == &SecretRef{}) {
		return []byte{}, nil
	}

	if r.Name == "" {
		return nil, errors.New("name must be set")
	}

	return fmt.Sprintf("secretName:%s", r.Name), nil
}

func (r *SecretRef) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var val interface{}
	err := unmarshal(&val)
	if err != nil {
		return err
	}

	value := val.(string)

	parts := strings.Split(value, ":")

	if len(parts) == 1 {
		r.Name = parts[0]
	} else if len(parts) == 2 {
		r.Name = parts[1]
	}

	return nil
}

func (r *SecretRef) String() string {
	return fmt.Sprintf("secretName:%s", r.Name)
}
