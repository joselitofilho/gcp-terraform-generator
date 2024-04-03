package resourcestoyaml

import (
	"fmt"

	"github.com/diagram-code-generator/resources/pkg/resources"
	"github.com/ettle/strcase"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

func (t *Transformer) buildBigTableRelationship(source, bigTable resources.Resource) {
	if gcpresources.ParseResourceType(source.ResourceType()) == gcpresources.AppEngine {
		t.buildAppEngineToBigTable(source, bigTable)
	}
}

func (t *Transformer) buildBigTables() (result []*config.BigTable) {
	for _, bt := range t.resourcesByTypeMap[gcpresources.BigTable] {
		var labels map[string]string
		if len(t.appEngineByBigTableID[bt.ID()]) > 0 {
			labels = map[string]string{}

			for _, a := range t.appEngineByBigTableID[bt.ID()] {
				k := fmt.Sprintf("%s-sender", a.Value())
				v := fmt.Sprintf("google_app_engine_application.%s_app.name", strcase.ToSnake(a.Value()))
				labels[k] = v
			}
		}

		result = append(result, &config.BigTable{Name: bt.Value(), Labels: labels})
	}

	return result
}
