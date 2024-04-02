package config

type Resource interface {
	GetName() string
}

// Config represents a configuration object that can be populated from a YAML file.
type Config struct {
	Diagram   *Diagram    `yaml:"diagram,omitempty"`
	Structure *Structure  `yaml:"structure,omitempty"`
	IoTCores  []*IoTCore  `yaml:"iot_cores,omitempty"`
	DataFlows []*DataFlow `yaml:"dataflows,omitempty"`
	PubSubs   []*PubSub   `yaml:"pubsubs,omitempty"`
}
