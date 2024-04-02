package resourcestoyaml

import "github.com/joselitofilho/gcp-terraform-generator/internal/resources"

func (t *Transformer) buildBigQueryRelationship(source, bq resources.Resource) {
	switch source.ResourceType() {
	case resources.Dataflow:
		t.buildDataFlowToBigQuery(source, bq)
	}
}
