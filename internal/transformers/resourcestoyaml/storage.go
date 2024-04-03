package resourcestoyaml

import (
	"github.com/diagram-code-generator/resources/pkg/resources"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

const defaultLocation = "US"

func (t *Transformer) buildStorageRelationship(source, storage resources.Resource) {
	if gcpresources.ParseResourceType(source.ResourceType()) == gcpresources.Dataflow {
		t.buildDataFlowToStorage(source, storage)
	}
}

func (t *Transformer) buildStorages() (result []*config.Storage) {
	for _, s := range t.resourcesByTypeMap[gcpresources.Storage] {
		result = append(result, &config.Storage{Name: s.Value(), Location: defaultLocation})
	}

	return result
}
