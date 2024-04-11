package config

type Diagram struct {
	StackName       string `yaml:"stack_name"`
	DefaultLocation string `yaml:"default_location,omitempty"`
}
