package config

type IoTCore struct {
	Name                     string                     `yaml:"name"`
	EventNotificationConfigs []*EventNotificationConfig `yaml:"event_notification_configs,omitempty"`
	Files                    []*File                    `yaml:"files,omitempty"`
}

func (r *IoTCore) GetName() string { return r.Name }

type EventNotificationConfig struct {
	TopicName string `yaml:"pubsub_topic_name"`
}
