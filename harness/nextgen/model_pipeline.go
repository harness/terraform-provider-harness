package nextgen

type Pipeline struct {
	Name              string `yaml:"name"`
	Identifier        string `yaml:"identifier"`
	ProjectIdentifier string `yaml:"projectIdentifier"`
	OrgIdentifier     string `yaml:"orgIdentifier"`
	Yaml              string `yaml:"-"`
}
