package resourcestoyaml

import (
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

const defaultProjectID = "us-central"

func (t *Transformer) buildAppEngineRelationship(source, appEngine resources.Resource) {
	if source.ResourceType() == resources.PubSub {
		t.buildPubSubToAppEngine(source, appEngine)
	}
}

func (t *Transformer) buildAppEngines() (result []*config.AppEngine) {
	for _, bq := range t.resourcesByTypeMap[resources.AppEngine] {
		result = append(result, &config.AppEngine{Name: bq.Value(), LocationID: defaultProjectID})
	}

	return result
}
