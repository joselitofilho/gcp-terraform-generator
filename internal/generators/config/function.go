package config

type Function struct {
	Name                string            `yaml:"name"`
	Source              string            `yaml:"source"`
	Runtime             string            `yaml:"runtime"`
	SourceArchiveBucket string            `yaml:"source_archive_bucket,omitempty"`
	SourceArchiveObject string            `yaml:"source_archive_object,omitempty"`
	TriggerHTTP         string            `yaml:"trigger_http,omitempty"`
	EntryPoint          string            `yaml:"entry_point,omitempty"`
	Envars              map[string]string `yaml:"envars,omitempty"`
	Files               []File            `yaml:"files,omitempty"`
}

func (r *Function) GetName() string { return r.Name }
