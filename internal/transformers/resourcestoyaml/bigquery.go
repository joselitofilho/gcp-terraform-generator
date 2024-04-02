package resourcestoyaml

import "github.com/joselitofilho/gcp-terraform-generator/internal/resources"

func (t *Transformer) buildBigQueryRelationship(source, bq resources.Resource) {
	if source.ResourceType() == resources.Dataflow {
		t.buildDataFlowToBigQuery(source, bq)
	}
}
