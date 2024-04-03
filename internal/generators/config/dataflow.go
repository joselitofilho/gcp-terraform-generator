package config

type DataFlow struct {
	Name              string   `yaml:"name"`
	TemplateGcsPath   string   `yaml:template_gcs_path`
	TempGcsLocation   string   `yaml:temp_gcs_location`
	InputTopics       []string `yaml:"input_topics,omitempty"`
	OutputDirectories []string `yaml:"output_directories,omitempty"`
	OutputTables      []string `yaml:"output_tables,omitempty"`
	Files             []File   `yaml:"files,omitempty"`
}

func (r *DataFlow) GetName() string { return r.Name }
