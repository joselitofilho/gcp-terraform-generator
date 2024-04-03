package resourcestoyaml

import (
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

const (
	defaultTemplateGcsPath = "gs://my-bucket/templates/template_file"
	defaultTempGcsLocation = "gs://my-bucket/tmp_dir"
)

func (t *Transformer) buildDataFlowRelationship(source, dataFlow resources.Resource) {
	if source.ResourceType() == resources.PubSub {
		t.buildPubSubToDataFlow(source, dataFlow)
	}
}

func (t *Transformer) buildDataFlows() (result []*config.DataFlow) {
	for _, df := range t.resourcesByTypeMap[resources.Dataflow] {
		inputTopics := []string{}
		for _, ps := range t.pubSubByDataFlowID[df.ID()] {
			inputTopics = append(inputTopics, ps.Value())
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
			OutputDirectories: outputDirectories,
			OutputTables:      outputTables,
		})
	}

	return result
}
