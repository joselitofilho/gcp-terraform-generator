package resourcestoyaml

import (
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

func (t *Transformer) buildIoTCores() (cores []*config.IoTCore) {
	for _, core := range t.resourcesByTypeMap[resources.IoTCore] {
		coreID := core.ID()
		eventNotificationConfigs := make([]*config.EventNotificationConfig, 0, len(t.pubSubByIoTCoreID[coreID]))

		for i := range t.pubSubByIoTCoreID[coreID] {
			eventNotificationConfigs = append(eventNotificationConfigs,
				&config.EventNotificationConfig{TopicName: t.pubSubByIoTCoreID[coreID][i].Value()})
		}

		cores = append(cores, &config.IoTCore{Name: core.Value(), EventNotificationConfigs: eventNotificationConfigs})
	}

	return cores
}
