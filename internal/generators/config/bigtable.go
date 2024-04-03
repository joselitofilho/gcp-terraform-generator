package config

type BigTable struct {
	Name   string            `yaml:"name"`
	Labels map[string]string `yaml:"labels,omitempty"`
	Files  []File            `yaml:"files,omitempty"`
}

func (r *BigTable) GetName() string { return r.Name }
