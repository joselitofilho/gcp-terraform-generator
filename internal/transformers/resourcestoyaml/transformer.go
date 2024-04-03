package resourcestoyaml

import (
	"github.com/diagram-code-generator/resources/pkg/resources"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

type Transformer struct {
	yamlConfig *config.Config
	resc       *resources.ResourceCollection

	appEngineByPubSubID      map[string][]resources.Resource
	bqTablesByDataFlowID     map[string][]resources.Resource
	pubSubByIoTCoreID        map[string][]resources.Resource
	inputPubSubByDataFlowID  map[string][]resources.Resource
	outputPubSubByDataFlowID map[string][]resources.Resource
	storageByDataFlowID      map[string][]resources.Resource

	resourcesByTypeMap map[gcpresources.ResourceType][]resources.Resource
}

func NewTransformer(yamlConfig *config.Config, resc *resources.ResourceCollection) *Transformer {
	return &Transformer{
		yamlConfig: yamlConfig,
		resc:       resc,

		appEngineByPubSubID:      map[string][]resources.Resource{},
		bqTablesByDataFlowID:     map[string][]resources.Resource{},
		pubSubByIoTCoreID:        map[string][]resources.Resource{},
		inputPubSubByDataFlowID:  map[string][]resources.Resource{},
		outputPubSubByDataFlowID: map[string][]resources.Resource{},
		storageByDataFlowID:      map[string][]resources.Resource{},

		resourcesByTypeMap: map[gcpresources.ResourceType][]resources.Resource{},
	}
}

func (t *Transformer) Transform() (*config.Config, error) {
	t.buildResourcesByTypeMap()
	t.buildResourceRelationships()

	appEngines := t.buildAppEngines()
	bigQueryTables := t.buildBigQueryTables()
	dataFlows := t.buildDataFlows()
	iotCores := t.buildIoTCores()
	pubSubs := t.buildPubSubs()
	storages := t.buildStorages()

	return &config.Config{
		AppEngines:     appEngines,
		BigQueryTables: bigQueryTables,
		DataFlows:      dataFlows,
		IoTCores:       iotCores,
		PubSubs:        pubSubs,
		Storages:       storages,
	}, nil
}

func (t *Transformer) buildResourcesByTypeMap() {
	for _, resource := range t.resc.Resources {
		resType := gcpresources.ParseResourceType(resource.ResourceType())
		t.resourcesByTypeMap[resType] = append(t.resourcesByTypeMap[resType], resource)
	}
}

func (t *Transformer) buildResourceRelationships() {
	for _, rel := range t.resc.Relationships {
		target := rel.Target
		source := rel.Source

		switch gcpresources.ParseResourceType(target.ResourceType()) {
		case gcpresources.AppEngine:
			t.buildAppEngineRelationship(source, target)
		case gcpresources.BigQuery:
			t.buildBigQueryRelationship(source, target)
		case gcpresources.Dataflow:
			t.buildDataFlowRelationship(source, target)
		case gcpresources.PubSub:
			t.buildPubSubRelationship(source, target)
		case gcpresources.Storage:
			t.buildStorageRelationship(source, target)
		}
	}
}
