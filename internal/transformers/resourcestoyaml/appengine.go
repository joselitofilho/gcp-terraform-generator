package resourcestoyaml

import (
	"github.com/diagram-code-generator/resources/pkg/resources"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

const defaultProjectID = "us-central"

func (t *Transformer) buildAppEngineRelationship(source, appEngine resources.Resource) {
	if gcpresources.ParseResourceType(source.ResourceType()) == gcpresources.PubSub {
		t.buildPubSubToAppEngine(source, appEngine)
	}
}

func (t *Transformer) buildAppEngines() (result []*config.AppEngine) {
	for _, bq := range t.resourcesByTypeMap[gcpresources.AppEngine] {
		result = append(result, &config.AppEngine{Name: bq.Value(), LocationID: defaultProjectID})
	}

	return result
}
