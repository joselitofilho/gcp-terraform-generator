package config

type Resource interface {
	GetName() string
}

// Config represents a configuration object that can be populated from a YAML file.
type Config struct {
	Diagram Diagram `yaml:"diagram,omitempty"`
}
