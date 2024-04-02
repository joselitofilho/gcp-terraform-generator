package config

type DataFlow struct {
	Name         string   `yaml:"name"`
	InputTopics  []string `yaml:"input_topics,omitempty"`
	OutputTables []string `yaml:"output_tables,omitempty"`
	Files        []*File  `yaml:"files,omitempty"`
}

func (r *DataFlow) GetName() string { return r.Name }
