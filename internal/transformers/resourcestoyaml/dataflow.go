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
		outputTables := []string{}

		for _, ps := range t.pubSubByDataFlowID[df.ID()] {
			inputTopics = append(inputTopics, ps.Value())
		}

		for _, bq := range t.bqTablesByDataFlowID[df.ID()] {
			outputTables = append(outputTables, bq.Value())
		}

		result = append(result, &config.DataFlow{
			Name:         df.Value(),
			InputTopics:  inputTopics,
			OutputTables: outputTables,
		})
	}

	return result
}
