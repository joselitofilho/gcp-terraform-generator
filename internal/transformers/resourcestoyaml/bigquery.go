package resourcestoyaml

import (
	"github.com/diagram-code-generator/resources/pkg/resources"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

const defaultSchema = `<<EOF
# Define your BigQuery schema here
EOF`

func (t *Transformer) buildBigQueryRelationship(source, bq resources.Resource) {
	if gcpresources.ParseResourceType(source.ResourceType()) == gcpresources.Dataflow {
		t.buildDataFlowToBigQuery(source, bq)
	}
}

func (t *Transformer) buildBigQueryTables() (result []*config.BigQueryTable) {
	for _, bq := range t.resourcesByTypeMap[gcpresources.BigQuery] {
		result = append(result, &config.BigQueryTable{Name: bq.Value(), Schema: defaultSchema})
	}

	return result
}
