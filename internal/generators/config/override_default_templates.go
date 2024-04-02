package config

type OverrideDefaultTemplates struct {
	IoTCore []FilenameTemplateMap `yaml:"iotcore,omitempty"`
	PubSub  []FilenameTemplateMap `yaml:"pubsub,omitempty"`
	Storage []FilenameTemplateMap `yaml:"storage,omitempty"`
}
