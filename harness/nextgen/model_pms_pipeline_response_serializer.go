package nextgen

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func (a *PmsPipelineResponse) UnmarshalJSON(data []byte) error {

	type Alias PmsPipelineResponse

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(aux.YamlPipeline), &a.PipelineData)
	// a.PipelineData.Pipeline.Yaml = string(data)

	return err
}

// func (a *PmsPipelineResponse) MarshalJSON() ([]byte, error) {
// 	type Alias PmsPipelineResponse

// 	a.YamlPipeline = a.PipelineData.Pipeline.Yaml

// 	return json.Marshal((*Alias)(a))
// }
