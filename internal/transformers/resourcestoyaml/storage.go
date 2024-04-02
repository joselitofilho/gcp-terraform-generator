package resourcestoyaml

import (
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

const defaultLocation = "US"

func (t *Transformer) buildStorageRelationship(source, storage resources.Resource) {
	if source.ResourceType() == resources.Dataflow {
		t.buildDataFlowToStorage(source, storage)
	}
}

func (t *Transformer) buildStorages() (result []*config.Storage) {
	for _, s := range t.resourcesByTypeMap[resources.Storage] {
		result = append(result, &config.Storage{Name: s.Value(), Location: defaultLocation})
	}

	return result
}
