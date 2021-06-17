package caac

import (
	"errors"
	"fmt"

	"github.com/micahlmartin/terraform-provider-harness/harness/utils"
	"gopkg.in/yaml.v3"
)

func (i *ConfigAsCodeItem) ParseYamlContent() (interface{}, error) {
	if i.Yaml == "" {
		return nil, nil
	}

	fmt.Println(i.Yaml)

	tmp := map[string]interface{}{}
	data := []byte(i.Yaml)

	err := yaml.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	val, ok := tmp["type"]
	if !ok {
		return nil, errors.New("could not find field 'type' in yaml object")
	}

	switch val {
	case ObjectTypes.Service:
		obj := &Service{}
		if err := yaml.Unmarshal(data, &obj); err != nil {
			return nil, err
		}
		obj.Name = utils.TrimFileExtension(i.Name)
		return obj, err
	default:
		return nil, fmt.Errorf("could not parse object type of '%s'", val)
	}

}

func (s *Service) Validate() (bool, error) {
	if s.ApplicationId == "" {
		return false, errors.New("service is invalid. missing field `ApplicationId`")
	}

	return true, nil
}

func (i *ConfigAsCodeItem) IsEmpty() bool {
	return i == &ConfigAsCodeItem{}
}

// Indicates an error condition
func (r *Response) IsEmpty() bool {
	// return true
	return r.Metadata == ResponseMetadata{} && r.Resource.IsEmpty() && len(r.ResponseMessages) == 0
}

func (m *ResponseMessage) ToError() error {
	return fmt.Errorf("%s: %s", m.Code, m.Message)
}
