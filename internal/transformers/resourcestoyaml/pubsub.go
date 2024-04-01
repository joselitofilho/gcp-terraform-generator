package resourcestoyaml

import (
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

func (t *Transformer) buildPubSubRelationship(source, pubsub resources.Resource) {
	if source.ResourceType() == resources.IoTCore {
		t.buildIoTCoreToPubSub(source, pubsub)
	}
}

func (t *Transformer) buildPubSubs() (pubsubs []*config.PubSub) {
	for _, pubsub := range t.resourcesByTypeMap[resources.PubSub] {
		pubsubs = append(pubsubs, &config.PubSub{Name: pubsub.Value()})
	}

	return pubsubs
}
