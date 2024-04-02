package iotcore

type EventNotificationConfig struct {
	TopicName string
}

type Data struct {
	Name                     string
	EventNotificationConfigs []EventNotificationConfig
}
