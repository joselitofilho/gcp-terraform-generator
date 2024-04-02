package config

type OverrideDefaultTemplates struct {
	DataFlow []FilenameTemplateMap `yaml:"dataflow,omitempty"`
	IoTCore  []FilenameTemplateMap `yaml:"iotcore,omitempty"`
	PubSub   []FilenameTemplateMap `yaml:"pubsub,omitempty"`
	Storage  []FilenameTemplateMap `yaml:"storage,omitempty"`
}
