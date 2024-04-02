package config

type OverrideDefaultTemplates struct {
	BigQuery []FilenameTemplateMap `yaml:"bigquery,omitempty"`
	DataFlow []FilenameTemplateMap `yaml:"dataflow,omitempty"`
	IoTCore  []FilenameTemplateMap `yaml:"iotcore,omitempty"`
	PubSub   []FilenameTemplateMap `yaml:"pubsub,omitempty"`
	Storage  []FilenameTemplateMap `yaml:"storage,omitempty"`
}
