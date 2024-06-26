package resourcestoyaml

import (
	"fmt"

	"github.com/ettle/strcase"

	"github.com/diagram-code-generator/resources/pkg/resources"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

func (t *Transformer) buildPubSubRelationship(source, pubSub resources.Resource) {
	switch gcpresources.ParseResourceType(source.ResourceType()) {
	case gcpresources.Dataflow:
		t.buildDataFlowToPubSub(source, pubSub)
	case gcpresources.Function:
		t.buildFunctionToPubSub(source, pubSub)
	case gcpresources.IoTCore:
		t.buildIoTCoreToPubSub(source, pubSub)
	}
}

func (t *Transformer) buildPubSubs() (result []*config.PubSub) {
	for _, ps := range t.resourcesByTypeMap[gcpresources.PubSub] {
		var labels map[string]string
		if len(t.appEngineByPubSubID[ps.ID()]) > 0 {
			labels = map[string]string{}

			for _, fn := range t.functionPublisherByPubSubID[ps.ID()] {
				k := fmt.Sprintf("%s-publisher", fn.Value())
				v := fmt.Sprintf("google_cloudfunctions_function.%s_function.name", strcase.ToSnake(fn.Value()))
				labels[k] = v
			}

			for _, a := range t.appEngineByPubSubID[ps.ID()] {
				k := fmt.Sprintf("%s-subscriber", a.Value())
				v := fmt.Sprintf("google_app_engine_application.%s_app.name", strcase.ToSnake(a.Value()))
				labels[k] = v
			}
		}

		var pushEndpoint string
		if fn, ok := t.functionSubscriberByPubSubID[ps.ID()]; ok {
			pushEndpoint = fn.Value()
		}

		result = append(result, &config.PubSub{Name: ps.Value(), Labels: labels, PushEnpoint: pushEndpoint})
	}

	return result
}
