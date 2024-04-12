package terraformtoresources

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/diagram-code-generator/resources/pkg/resources"

	hcl "github.com/joselitofilho/hcl-parser-go/pkg/parser/hcl"

	"github.com/joselitofilho/gcp-terraform-generator/internal/fmtcolor"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"
)

const (
	suffixAppEngine     = "app"
	suffixBigQueryTable = "table"
	suffixBigTable      = "instance"
	suffixDataFlow      = "df_job"
	suffixFunction      = "function"
	suffixIoTCore       = "registry"
	suffixPubSub        = "topic"
	suffixStorage       = "bucket"
)

type Transformer struct {
	yamlConfig *config.Config
	tfConfig   *hcl.Config

	resources     []resources.Resource
	relationships []resources.Relationship

	// ByName
	bqTableByName map[string]resources.Resource

	// GSCByName
	pbSubscriptionGCSByName map[string]*gcpresources.ResourceGCS

	// ByLabel
	appEngineByLabel map[string]resources.Resource
	bqTableByLabel   map[string]resources.Resource
	bigTableByLabel  map[string]resources.Resource
	dataFlowByLabel  map[string]resources.Resource
	functionByLabel  map[string]resources.Resource
	iotCoreByLabel   map[string]resources.Resource
	pubSubByLabel    map[string]resources.Resource
	storageByLabel   map[string]resources.Resource

	// GCSByLabel
	bqDatasetGCSByLabel map[string]*gcpresources.ResourceGCS

	// Reletionship
	pubSubByPubSubSubscriptionLabel map[string]*gcpresources.ResourceGCS

	relationshipsMap map[*gcpresources.ResourceGCS][]*gcpresources.ResourceGCS

	id int
}

func NewTransformer(yamlConfig *config.Config, tfConfig *hcl.Config) *Transformer {
	return &Transformer{
		yamlConfig: yamlConfig,
		tfConfig:   tfConfig,

		resources:     []resources.Resource{},
		relationships: []resources.Relationship{},

		bqTableByName: map[string]resources.Resource{},

		pbSubscriptionGCSByName: map[string]*gcpresources.ResourceGCS{},

		appEngineByLabel: map[string]resources.Resource{},
		bqTableByLabel:   map[string]resources.Resource{},
		bigTableByLabel:  map[string]resources.Resource{},
		dataFlowByLabel:  map[string]resources.Resource{},
		functionByLabel:  map[string]resources.Resource{},
		iotCoreByLabel:   map[string]resources.Resource{},
		pubSubByLabel:    map[string]resources.Resource{},
		storageByLabel:   map[string]resources.Resource{},

		bqDatasetGCSByLabel: map[string]*gcpresources.ResourceGCS{},

		pubSubByPubSubSubscriptionLabel: map[string]*gcpresources.ResourceGCS{},

		relationshipsMap: map[*gcpresources.ResourceGCS][]*gcpresources.ResourceGCS{},

		id: 1,
	}
}

func (t *Transformer) Transform() *resources.ResourceCollection {
	t.processTerraformResources()

	t.buildRelationships()

	t.applyFiltersInResources()
	t.applyFiltersInRelationships()

	return &resources.ResourceCollection{Resources: t.resources, Relationships: t.relationships}
}

func (t *Transformer) applyFiltersInResources() {
	filtered := make([]resources.Resource, 0, len(t.resources))

	for _, res := range t.resources {
		if t.hasResourceMatched(res, t.yamlConfig.Draw.Filters) {
			filtered = append(filtered, res)
		}
	}

	t.resources = filtered
}

func (t *Transformer) applyFiltersInRelationships() {
	filtered := make([]resources.Relationship, 0, len(t.relationships))

	for _, rel := range t.relationships {
		sourceMatch := t.hasResourceMatched(rel.Source, t.yamlConfig.Draw.Filters)
		targetMatch := t.hasResourceMatched(rel.Target, t.yamlConfig.Draw.Filters)

		if sourceMatch && targetMatch {
			filtered = append(filtered, rel)
		}
	}

	t.relationships = filtered
}

func (t *Transformer) buildRelationships() {
	for sourceGCS, rel := range t.relationshipsMap {
		source := t.getResourceByGCS(sourceGCS)

		for i := range rel {
			targetGCS := rel[i]
			target := t.getResourceByGCS(targetGCS)

			if source != nil && target != nil {
				t.relationships = append(t.relationships, resources.Relationship{Source: source, Target: target})
			}
		}
	}
}

func (t *Transformer) getResourceByGCS(gcs *gcpresources.ResourceGCS) resources.Resource {
	if gcs.Label == "" {
		return t.getResourceByGCSName(gcs)
	}

	return t.getResourceByGCSLabel(gcs)
}

func (t *Transformer) getResourceByGCSName(gcs *gcpresources.ResourceGCS) (resource resources.Resource) {
	switch gcs.Type {
	case gcpresources.LabelAppEngine:
	case gcpresources.LabelBigQueryTable:
		resource = t.bqTableByName[gcs.Name]
	case gcpresources.LabelBigTable:
	case gcpresources.LabelDataFlow:
	case gcpresources.LabelFunction:
	case gcpresources.LabelIoTCore:
	case gcpresources.LabelPubSub:
	case gcpresources.LabelPubSubSubscription:
		pbSubLabel := t.pbSubscriptionGCSByName[gcs.Name].Label
		resource = t.pubSubByLabel[t.pubSubByPubSubSubscriptionLabel[pbSubLabel].Label]
	case gcpresources.LabelStorage:
	}

	return resource
}

func (t *Transformer) getResourceByGCSLabel(gcs *gcpresources.ResourceGCS) (resource resources.Resource) {
	switch gcs.Type {
	case gcpresources.LabelAppEngine:
		resource = t.appEngineByLabel[gcs.Label]
	case gcpresources.LabelBigQueryTable:
		resource = t.bqTableByLabel[gcs.Label]
	case gcpresources.LabelBigTable:
		resource = t.bigTableByLabel[gcs.Label]
	case gcpresources.LabelDataFlow:
		resource = t.dataFlowByLabel[gcs.Label]
	case gcpresources.LabelFunction:
		resource = t.functionByLabel[gcs.Label]
	case gcpresources.LabelIoTCore:
		resource = t.iotCoreByLabel[gcs.Label]
	case gcpresources.LabelPubSub:
		resource = t.pubSubByLabel[gcs.Label]
	case gcpresources.LabelPubSubSubscription:
		resource = t.pubSubByLabel[t.pubSubByPubSubSubscriptionLabel[gcs.Label].Label]
	case gcpresources.LabelStorage:
		resource = t.storageByLabel[gcs.Label]
	}

	return resource
}

func (t *Transformer) hasResourceMatched(res resources.Resource, filters config.Filters) bool {
	if res == nil {
		return false
	}

	filter, hasFilter := filters[gcpresources.ParseResourceType(res.ResourceType())]
	if !hasFilter {
		return true
	}

	match := len(filter.Match) == 0

	for _, pattern := range filter.Match {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			fmtcolor.Yellow.Println("error compiling match regex:", err)
			continue
		}

		if regex.MatchString(res.Value()) {
			match = true
			break
		}
	}

	for _, pattern := range filter.NotMatch {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			fmtcolor.Yellow.Println("error compiling not_match regex:", err)
			continue
		}

		if regex.MatchString(res.Value()) {
			match = false
			break
		}
	}

	return match
}

func (t *Transformer) processTerraformResources() {
	for _, tfResourceConf := range t.tfConfig.Resources {
		if len(tfResourceConf.Labels) == 2 {
			switch tfResourceConf.Labels[0] {
			case gcpresources.LabelBigQueryDataset:
				t.processBigQueryDataset(tfResourceConf)
			case gcpresources.LabelPubSubSubscription:
				t.processPubSubSubscription(tfResourceConf)
			}
		}
	}

	for _, tfResourceConf := range t.tfConfig.Resources {
		if len(tfResourceConf.Labels) == 2 {
			switch tfResourceConf.Labels[0] {
			case gcpresources.LabelAppEngine:
				t.processAppEngine(tfResourceConf)
			case gcpresources.LabelBigQueryTable:
				t.processBigQueryTable(tfResourceConf)
			case gcpresources.LabelBigTable:
				t.processBigTable(tfResourceConf)
			case gcpresources.LabelDataFlow:
				t.processDataFlow(tfResourceConf)
			case gcpresources.LabelFunction:
				t.processFunction(tfResourceConf)
			case gcpresources.LabelIoTCore:
				t.processIoTCore(tfResourceConf)
			case gcpresources.LabelPubSub:
				t.processPubSub(tfResourceConf)
			case gcpresources.LabelStorage:
				t.processStorage(tfResourceConf)
			}
		}
	}
}

func (t *Transformer) processResource(
	conf *hcl.Resource, resourceType gcpresources.ResourceType, attributeName, labelSuffix string,
	resourcesByLabel map[string]resources.Resource,
) {
	label := conf.Labels[1]
	name := getResourceNameFromLabel(label, "_"+labelSuffix)

	if attributeName != EmptyAttributeName {
		name = replaceVars(conf.Attributes[attributeName].(string), t.tfConfig.Variables, t.tfConfig.Locals,
			t.yamlConfig.Draw.ReplaceableTexts)
	}

	if _, ok := resourcesByLabel[label]; !ok {
		resource := resources.NewGenericResource(fmt.Sprintf("%d", t.id), name, resourceType.String())
		t.id++

		t.resources = append(t.resources, resource)
		resourcesByLabel[label] = resource
	}
}

func (t *Transformer) processAppEngine(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.AppEngine, EmptyAttributeName, suffixAppEngine, t.appEngineByLabel)
}

func (t *Transformer) processBigQueryDataset(conf *hcl.Resource) {
	label := conf.Labels[1]

	datasetIDValue := replaceVars(conf.Attributes["dataset_id"].(string), t.tfConfig.Variables, t.tfConfig.Locals,
		t.yamlConfig.Draw.ReplaceableTexts)

	t.bqDatasetGCSByLabel[label] = gcpresources.ParseResourceGCS(datasetIDValue, conf.Labels)
}

func (t *Transformer) processBigQueryTable(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.BigQuery, "table_id", suffixBigQueryTable, t.bqTableByLabel)

	datasetID, ok := conf.Attributes["dataset_id"]
	if !ok {
		return
	}

	label := conf.Labels[1]
	target := t.bqTableByLabel[label]

	bqDatasetValue := replaceVars(datasetID.(string), t.tfConfig.Variables, t.tfConfig.Locals,
		t.yamlConfig.Draw.ReplaceableTexts)
	bqDatasetGCS := gcpresources.ParseResourceGCS(bqDatasetValue, conf.Labels)

	if gcpresources.ParseResourceType(target.ResourceType()) == gcpresources.BigQuery {
		value := t.bqDatasetGCSByLabel[bqDatasetGCS.Label].Name + "." + target.Value()
		target = resources.NewGenericResource(target.ID(), value, target.ResourceType())

		t.resources[len(t.resources)-1] = target
		t.bqTableByName[value] = target
		t.bqTableByLabel[label] = target
	}
}

func (t *Transformer) processBigTable(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.BigTable, "name", suffixBigTable, t.bigTableByLabel)
}

func (t *Transformer) processDataFlow(conf *hcl.Resource) {
	sourceAttribute := "name"

	t.processResource(conf, gcpresources.Dataflow, sourceAttribute, suffixDataFlow, t.dataFlowByLabel)

	sourceValue := replaceVars(conf.Attributes[sourceAttribute].(string), t.tfConfig.Variables, t.tfConfig.Locals,
		t.yamlConfig.Draw.ReplaceableTexts)
	sourceGCS := gcpresources.ParseResourceGCS(sourceValue, conf.Labels)

	parameters, ok := conf.Attributes["parameters"]
	if !ok {
		return
	}

	if parameters, ok := parameters.(map[string]any); ok {
		for targetAttr, v := range parameters {
			value, ok := v.(string)
			if !ok {
				continue
			}

			var labels []string

			hasRelationship := true
			switch {
			case strings.HasPrefix(targetAttr, "outputTopic"):
				labels = []string{gcpresources.LabelPubSubSubscription}
			case strings.HasPrefix(targetAttr, "outputTable"):
				overrideValue := false

				parts := strings.Split(value, ":")
				if len(parts) > 1 {
					parts = strings.Split(parts[1], ".")
					for i := range parts {
						if strings.Contains(parts[i], gcpresources.LabelBigQueryTable) {
							value = fmt.Sprintf("%s.%s", strings.ReplaceAll(parts[i], "${", ""), parts[i+1])
							overrideValue = true
							break
						}
					}

					if !overrideValue {
						value = strings.Join(parts, ".")
						labels = []string{gcpresources.LabelBigQueryTable}
					}
				}
			default:
				hasRelationship = false
			}

			if hasRelationship {
				targetValue := replaceVars(value, t.tfConfig.Variables, t.tfConfig.Locals,
					t.yamlConfig.Draw.ReplaceableTexts)
				targetGCS := gcpresources.ParseResourceGCS(targetValue, labels)

				t.relationshipsMap[sourceGCS] = append(t.relationshipsMap[sourceGCS], targetGCS)
			}
		}
	}
}

func (t *Transformer) processFunction(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.Function, "name", suffixFunction, t.functionByLabel)
}

func (t *Transformer) processIoTCore(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.IoTCore, "name", suffixIoTCore, t.iotCoreByLabel)
}

func (t *Transformer) processPubSub(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.PubSub, "name", suffixPubSub, t.pubSubByLabel)
}

func (t *Transformer) processPubSubSubscription(conf *hcl.Resource) {
	label := conf.Labels[1]

	name := replaceVars(conf.Attributes["name"].(string), t.tfConfig.Variables, t.tfConfig.Locals,
		t.yamlConfig.Draw.ReplaceableTexts)
	subsGCS := gcpresources.ParseResourceGCS(name, conf.Labels)

	topic := replaceVars(conf.Attributes["topic"].(string), t.tfConfig.Variables, t.tfConfig.Locals,
		t.yamlConfig.Draw.ReplaceableTexts)
	topicGCS := gcpresources.ParseResourceGCS(topic, conf.Labels)

	t.pbSubscriptionGCSByName[name] = subsGCS
	t.pubSubByPubSubSubscriptionLabel[label] = topicGCS
}

func (t *Transformer) processStorage(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.Storage, "name", suffixStorage, t.storageByLabel)
}
