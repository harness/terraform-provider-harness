package cac

import (
	"gopkg.in/yaml.v3"
)

type ServiceVariable struct {
	Name      string            `yaml:"name,omitempty"`
	Value     string            `yaml:"value,omitempty"`
	ValueType VariableValueType `yaml:"valueType,omitempty"`
}

// We need to customize the marshaling of this object because of the way the `Value`
// field works. When the ValueType is ENCRYPTED_TEXT the value needs to be secretName:<secretName>.

func (s *ServiceVariable) UnmarshalYAML(unmarshal func(interface{}) error) error {

	type Alias ServiceVariable

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := unmarshal(&aux.Alias); err != nil {
		return err
	}

	if s.ValueType == VariableOverrideValueTypes.Text {
		s.Value = aux.Value
	} else if s.ValueType == VariableOverrideValueTypes.EncryptedText {
		ref := &SecretRef{}
		if err := yaml.Unmarshal([]byte(aux.Value), ref); err != nil {
			return err
		}

		s.Value = ref.Name
	}

	return nil
}

func (r *ServiceVariable) MarshalYAML() (interface{}, error) {
	var aux ServiceVariable = *r

	if r.ValueType == VariableOverrideValueTypes.Text {
		aux.Value = r.Value
	} else if r.ValueType == VariableOverrideValueTypes.EncryptedText {
		ref := &SecretRef{
			Name: r.Value,
		}
		aux.Value = ref.String()
	}

	return aux, nil
}
