package resourcestoyaml

import (
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

const defaultSchema = `<<EOF
# Define your BigQuery schema here
EOF`

func (t *Transformer) buildBigQueryRelationship(source, bq resources.Resource) {
	if source.ResourceType() == resources.Dataflow {
		t.buildDataFlowToBigQuery(source, bq)
	}
}

func (t *Transformer) buildBigQueryTables() (result []*config.BigQueryTable) {
	for _, bq := range t.resourcesByTypeMap[resources.BigQuery] {
		result = append(result, &config.BigQueryTable{Name: bq.Value(), Schema: defaultSchema})
	}

	return result
}
