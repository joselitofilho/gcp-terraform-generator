package resourcestoyaml

import (
	"fmt"

	"github.com/ettle/strcase"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

const (
	defaultSource                    = "."
	defaultRuntime                   = "go116"
	defaultSourceArchiveBucket       = "${google_storage_bucket.archive_funcs_bucket.name}"
	defaultSourceArchiveObjectFormat = "${google_storage_bucket_object.%s_archive_funcs_object.name}"
	defaultTriggerHTTP               = "true"
)

func (t *Transformer) buildFunctions() (result []*config.Function) {
	for _, fn := range t.resourcesByTypeMap[gcpresources.Function] {
		name := fn.Value()

		result = append(result, &config.Function{
			Name:                name,
			Source:              defaultSource,
			Runtime:             defaultRuntime,
			SourceArchiveBucket: defaultSourceArchiveBucket,
			SourceArchiveObject: fmt.Sprintf(defaultSourceArchiveObjectFormat, strcase.ToSnake(name)),
			TriggerHTTP:         defaultTriggerHTTP,
			EntryPoint:          fmt.Sprintf("%sFunction", name),
		})
	}

	return result
}
