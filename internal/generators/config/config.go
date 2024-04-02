package config

type Resource interface {
	GetName() string
}

// Config represents a configuration object that can be populated from a YAML file.
type Config struct {
	OverrideDefaultTemplates OverrideDefaultTemplates `yaml:"override_default_templates,omitempty"`
	Diagram                  *Diagram                 `yaml:"diagram,omitempty"`
	Structure                *Structure               `yaml:"structure,omitempty"`
	DataFlows                []*DataFlow              `yaml:"dataflows,omitempty"`
	IoTCores                 []*IoTCore               `yaml:"iotcores,omitempty"`
	PubSubs                  []*PubSub                `yaml:"pubsubs,omitempty"`
	Storages                 []*Storage               `yaml:"storages,omitempty"`
}
