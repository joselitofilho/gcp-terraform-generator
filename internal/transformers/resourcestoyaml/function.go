package resourcestoyaml

import (
	"fmt"

	"github.com/diagram-code-generator/resources/pkg/resources"
	"github.com/ettle/strcase"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

const (
	defaultSource                    = "."
	defaultRuntime                   = "go1.x"
	defaultSourceArchiveBucket       = "${google_storage_bucket.archive_funcs_bucket.name}"
	defaultSourceArchiveObjectFormat = "${google_storage_bucket_object.%s_archive_funcs_object.name}"
	defaultTriggerHTTP               = "true"
)

func (t *Transformer) buildFunctionRelationship(source, function resources.Resource) {
	if gcpresources.ParseResourceType(source.ResourceType()) == gcpresources.PubSub {
		t.buildPubSubToFunction(source, function)
	}
}

func (t *Transformer) buildFunctions() (result []*config.Function) {
	for _, fn := range t.resourcesByTypeMap[gcpresources.Function] {
		fnName := fn.Value()
		envars := map[string]string{}

		for _, ps := range t.pubSubByFunctionID[fn.ID()] {
			k := fmt.Sprintf("%s_TOPIC_NAME", strcase.ToSNAKE(ps.Value()))
			v := fmt.Sprintf("google_pubsub_topic.%s_topic.name", strcase.ToSnake(ps.Value()))
			envars[k] = v
		}

		result = append(result, &config.Function{
			Name:                fnName,
			Source:              defaultSource,
			Runtime:             defaultRuntime,
			SourceArchiveBucket: defaultSourceArchiveBucket,
			SourceArchiveObject: fmt.Sprintf(defaultSourceArchiveObjectFormat, strcase.ToSnake(fnName)),
			TriggerHTTP:         defaultTriggerHTTP,
			EntryPoint:          fmt.Sprintf("%sFunction", fnName),
			Envars:              envars,
		})
	}

	return result
}
