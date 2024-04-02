package resourcestoyaml

import (
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

func (t *Transformer) buildDataFlowRelationship(source, dataFlow resources.Resource) {
	switch source.ResourceType() {
	case resources.PubSub:
		t.buildPubSubToDataFlow(source, dataFlow)
	}
}

func (t *Transformer) buildDataFlows() (result []*config.DataFlow) {
	for _, df := range t.resourcesByTypeMap[resources.Dataflow] {
		inputTopics := []string{}

		for _, ps := range t.pubSubByDataFlowID[df.ID()] {
			inputTopics = append(inputTopics, ps.Value())
		}

		result = append(result, &config.DataFlow{Name: df.Value(), InputTopics: inputTopics})
	}

	return result
}
