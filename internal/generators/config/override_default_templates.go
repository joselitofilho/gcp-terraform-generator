package config

type OverrideDefaultTemplates struct {
	AppEngine []FilenameTemplateMap `yaml:"bigquery,omitempty"`
	BigQuery  []FilenameTemplateMap `yaml:"bigquery,omitempty"`
	BigTable  []FilenameTemplateMap `yaml:"bigtable,omitempty"`
	DataFlow  []FilenameTemplateMap `yaml:"dataflow,omitempty"`
	Function  []FilenameTemplateMap `yaml:"function,omitempty"`
	IoTCore   []FilenameTemplateMap `yaml:"iotcore,omitempty"`
	PubSub    []FilenameTemplateMap `yaml:"pubsub,omitempty"`
	Storage   []FilenameTemplateMap `yaml:"storage,omitempty"`
}
