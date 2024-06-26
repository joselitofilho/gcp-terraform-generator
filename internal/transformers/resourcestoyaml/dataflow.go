package resourcestoyaml

import (
	"github.com/diagram-code-generator/resources/pkg/resources"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

const (
	defaultTemplateGcsPath = "gs://my-bucket/templates/template_file"
	defaultTempGcsLocation = "gs://my-bucket/tmp_dir"
)

func (t *Transformer) buildDataFlowRelationship(source, dataFlow resources.Resource) {
	if gcpresources.ParseResourceType(source.ResourceType()) == gcpresources.PubSub {
		t.buildPubSubToDataFlow(source, dataFlow)
	}
}

func (t *Transformer) buildDataFlows() (result []*config.DataFlow) {
	for _, df := range t.resourcesByTypeMap[gcpresources.Dataflow] {
		inputTopics := []string{}
		for _, ps := range t.inputPubSubByDataFlowID[df.ID()] {
			inputTopics = append(inputTopics, ps.Value())
		}

		outputTopics := []string{}
		for _, ps := range t.outputPubSubByDataFlowID[df.ID()] {
			outputTopics = append(outputTopics, ps.Value())
		}

		outputDirectories := []string{}
		for _, s := range t.storageByDataFlowID[df.ID()] {
			outputDirectories = append(outputDirectories, s.Value())
		}

		outputTables := []string{}
		for _, bq := range t.bqTablesByDataFlowID[df.ID()] {
			outputTables = append(outputTables, bq.Value())
		}

		result = append(result, &config.DataFlow{
			Name:              df.Value(),
			TemplateGcsPath:   defaultTemplateGcsPath,
			TempGcsLocation:   defaultTemplateGcsPath,
			InputTopics:       inputTopics,
			OutputTopics:      outputTopics,
			OutputDirectories: outputDirectories,
			OutputTables:      outputTables,
		})
	}

	return result
}
