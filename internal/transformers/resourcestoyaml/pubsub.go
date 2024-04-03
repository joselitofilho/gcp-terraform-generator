package resourcestoyaml

import (
	"fmt"

	"github.com/ettle/strcase"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

func (t *Transformer) buildPubSubRelationship(source, pubSub resources.Resource) {
	switch source.ResourceType() {
	case resources.Dataflow:
		t.buildDataFlowToPubSub(source, pubSub)
	case resources.IoTCore:
		t.buildIoTCoreToPubSub(source, pubSub)
	}
}

func (t *Transformer) buildPubSubs() (result []*config.PubSub) {
	for _, ps := range t.resourcesByTypeMap[resources.PubSub] {
		var labels map[string]string
		if len(t.appEngineByPubSubID[ps.ID()]) > 0 {
			labels = map[string]string{}

			for _, a := range t.appEngineByPubSubID[ps.ID()] {
				k := fmt.Sprintf("%s-subscriber", a.Value())
				v := fmt.Sprintf(`"${google_app_engine_application.%s_app.name}"`, strcase.ToSnake(a.Value()))
				labels[k] = v
			}
		}

		result = append(result, &config.PubSub{Name: ps.Value(), Labels: labels})
	}

	return result
}
