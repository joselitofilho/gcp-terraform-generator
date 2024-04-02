package config

type DataFlow struct {
	Name        string   `yaml:"name"`
	InputTopics []string `yaml:"input_topics,omitempty"`
	Files       []*File  `yaml:"files,omitempty"`
}

func (r *DataFlow) GetName() string { return r.Name }
