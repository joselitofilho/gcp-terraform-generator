package resourcestoyaml

import (
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

type Transformer struct {
	yamlConfig *config.Config
	resc       *resources.ResourceCollection

	pubSubByIoTCoreID  map[string][]resources.Resource
	pubSubByDataFlowID map[string][]resources.Resource

	resourcesByTypeMap map[resources.ResourceType][]resources.Resource
}

func NewTransformer(yamlConfig *config.Config, resc *resources.ResourceCollection) *Transformer {
	return &Transformer{
		yamlConfig: yamlConfig,
		resc:       resc,

		pubSubByIoTCoreID:  map[string][]resources.Resource{},
		pubSubByDataFlowID: map[string][]resources.Resource{},

		resourcesByTypeMap: map[resources.ResourceType][]resources.Resource{},
	}
}

func (t *Transformer) Transform() (*config.Config, error) {
	t.buildResourcesByTypeMap()
	t.buildResourceRelationships()

	iotCores := t.buildIoTCores()
	pubSubs := t.buildPubSubs()
	dataFlows := t.buildDataFlows()

	return &config.Config{
		IoTCores:  iotCores,
		PubSubs:   pubSubs,
		DataFlows: dataFlows,
	}, nil
}

func (t *Transformer) buildResourcesByTypeMap() {
	for _, resource := range t.resc.Resources {
		t.resourcesByTypeMap[resource.ResourceType()] = append(t.resourcesByTypeMap[resource.ResourceType()], resource)
	}
}

func (t *Transformer) buildResourceRelationships() {
	for _, rel := range t.resc.Relationships {
		target := rel.Target
		source := rel.Source

		switch target.ResourceType() {
		case resources.PubSub:
			t.buildPubSubRelationship(source, target)
		case resources.Dataflow:
			t.buildDataFlowRelationship(source, target)
		}
	}
}
