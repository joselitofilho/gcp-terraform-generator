package resourcestoyaml

import (
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

func (t *Transformer) buildIoTCores() (result []*config.IoTCore) {
	for _, c := range t.resourcesByTypeMap[resources.IoTCore] {
		coreID := c.ID()
		eventNotificationConfigs := make([]config.EventNotificationConfig, 0, len(t.pubSubByIoTCoreID[coreID]))

		for _, r := range t.pubSubByIoTCoreID[coreID] {
			eventNotificationConfigs = append(eventNotificationConfigs,
				config.EventNotificationConfig{TopicName: r.Value()})
		}

		result = append(result, &config.IoTCore{Name: c.Value(), EventNotificationConfigs: eventNotificationConfigs})
	}

	return result
}
