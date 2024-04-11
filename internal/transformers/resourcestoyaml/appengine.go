package resourcestoyaml

import (
	"github.com/diagram-code-generator/resources/pkg/resources"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

const (
	defaultProjectID  = "${var.project_id}"
	defaultLocationID = "US"
)

func (t *Transformer) buildAppEngineRelationship(source, appEngine resources.Resource) {
	if gcpresources.ParseResourceType(source.ResourceType()) == gcpresources.PubSub {
		t.buildPubSubToAppEngine(source, appEngine)
	}
}

func (t *Transformer) buildAppEngines() (result []*config.AppEngine) {
	var locationID string
	if t.yamlConfig.Diagram != nil {
		locationID = t.yamlConfig.Diagram.DefaultLocation
	}

	if locationID == "" {
		locationID = defaultLocation
	}

	for _, ae := range t.resourcesByTypeMap[gcpresources.AppEngine] {
		result = append(result, &config.AppEngine{Name: ae.Value(), Project: defaultProjectID, LocationID: locationID})
	}

	return result
}
