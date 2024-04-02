package resourcestoyaml

import (
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

func (t *Transformer) buildPubSubRelationship(source, pubSub resources.Resource) {
	if source.ResourceType() == resources.IoTCore {
		t.buildIoTCoreToPubSub(source, pubSub)
	}
}

func (t *Transformer) buildPubSubs() (result []*config.PubSub) {
	for _, ps := range t.resourcesByTypeMap[resources.PubSub] {
		result = append(result, &config.PubSub{Name: ps.Value()})
	}

	return result
}
