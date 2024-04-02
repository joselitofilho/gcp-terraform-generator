package config

type BigQueryTable struct {
	Name   string `yaml:"name"`
	Schema string `yaml:"schema,omitempty"`
	Files  []File `yaml:"files,omitempty"`
}

func (r *BigQueryTable) GetName() string { return r.Name }
