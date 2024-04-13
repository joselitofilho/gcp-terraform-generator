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

	attributeName  = "name"
	attributeTopic = "topic"
)

type FNProcess func(key string) (hasRelationship bool, suggestionLabels []string)

type Transformer struct {
	yamlConfig *config.Config
	tfConfig   *hcl.Config

	resources     []resources.Resource
	relationships []resources.Relationship

	// ByName.
	appEngineByName map[string]resources.Resource
	bqTableByName   map[string]resources.Resource
	bigTableByName  map[string]resources.Resource
	dataFlowByName  map[string]resources.Resource
	functionByName  map[string]resources.Resource
	iotCoreByName   map[string]resources.Resource
	pubSubByName    map[string]resources.Resource
	storageByName   map[string]resources.Resource

	// GSCByName.
	pbSubscriptionGCSByName map[string]*gcpresources.ResourceGCS

	// ByLabel.
	appEngineByLabel map[string]resources.Resource
	bqTableByLabel   map[string]resources.Resource
	bigTableByLabel  map[string]resources.Resource
	dataFlowByLabel  map[string]resources.Resource
	functionByLabel  map[string]resources.Resource
	iotCoreByLabel   map[string]resources.Resource
	pubSubByLabel    map[string]resources.Resource
	storageByLabel   map[string]resources.Resource

	// GCSByLabel.
	bqDatasetGCSByLabel map[string]*gcpresources.ResourceGCS

	// Reletionship.
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

		appEngineByName: map[string]resources.Resource{},
		bqTableByName:   map[string]resources.Resource{},
		bigTableByName:  map[string]resources.Resource{},
		dataFlowByName:  map[string]resources.Resource{},
		functionByName:  map[string]resources.Resource{},
		iotCoreByName:   map[string]resources.Resource{},
		pubSubByName:    map[string]resources.Resource{},
		storageByName:   map[string]resources.Resource{},

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
		resource = t.bigTableByName[gcs.Name]
	case gcpresources.LabelDataFlow:
		resource = t.dataFlowByName[gcs.Name]
	case gcpresources.LabelFunction:
		resource = t.functionByName[gcs.Name]
	case gcpresources.LabelIoTCore:
		resource = t.iotCoreByName[gcs.Name]
	case gcpresources.LabelPubSub:
		resource = t.pubSubByName[gcs.Name]
	case gcpresources.LabelPubSubSubscription:
		pbSubLabel := t.pbSubscriptionGCSByName[gcs.Name].Label
		resource = t.pubSubByLabel[t.pubSubByPubSubSubscriptionLabel[pbSubLabel].Label]
	case gcpresources.LabelStorage:
		resource = t.storageByName[gcs.Name]
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
	conf *hcl.Resource, resourceType gcpresources.ResourceType, attribute, labelSuffix string,
	resourcesByName, resourcesByLabel map[string]resources.Resource,
) {
	label := conf.Labels[1]
	name := getResourceNameFromLabel(label, "_"+labelSuffix)

	if attribute != EmptyAttributeName {
		name = replaceVars(conf.Attributes[attribute].(string), t.tfConfig.Variables, t.tfConfig.Locals,
			t.yamlConfig.Draw.ReplaceableTexts)
	}

	if _, ok := resourcesByLabel[label]; !ok {
		resource := resources.NewGenericResource(fmt.Sprintf("%d", t.id), name, resourceType.String())
		t.id++

		t.resources = append(t.resources, resource)
		resourcesByName[name] = resource
		resourcesByLabel[label] = resource
	}
}

func (t *Transformer) processRelationshipByAttrsMap(
	conf *hcl.Resource, sourceAttr, targetAttr string, fnProcess FNProcess,
) {
	attrMap, ok := conf.Attributes[targetAttr]
	if !ok {
		return
	}

	sourceValue := replaceVars(conf.Attributes[sourceAttr].(string), t.tfConfig.Variables, t.tfConfig.Locals,
		t.yamlConfig.Draw.ReplaceableTexts)
	sourceGCS := gcpresources.ParseResourceGCS(sourceValue, conf.Labels)

	if attrMap, ok := attrMap.(map[string]any); ok {
		for k, v := range attrMap {
			value, ok := v.(string)
			if !ok {
				continue
			}

			value = extractTextFromTFVar(value)

			hasRelationship, suggestionLabels := fnProcess(k)

			if hasRelationship {
				targetValue := replaceVars(value, t.tfConfig.Variables, t.tfConfig.Locals,
					t.yamlConfig.Draw.ReplaceableTexts)
				targetGCS := gcpresources.ParseResourceGCS(targetValue, suggestionLabels)

				t.relationshipsMap[sourceGCS] = append(t.relationshipsMap[sourceGCS], targetGCS)
			}
		}
	}
}

func (t *Transformer) processAppEngine(conf *hcl.Resource) {
	t.processResource(conf,
		gcpresources.AppEngine, EmptyAttributeName, suffixAppEngine, t.appEngineByName, t.appEngineByLabel)
}

func (t *Transformer) processBigQueryDataset(conf *hcl.Resource) {
	label := conf.Labels[1]

	datasetIDValue := replaceVars(conf.Attributes["dataset_id"].(string), t.tfConfig.Variables, t.tfConfig.Locals,
		t.yamlConfig.Draw.ReplaceableTexts)

	t.bqDatasetGCSByLabel[label] = gcpresources.ParseResourceGCS(datasetIDValue, conf.Labels)
}

func (t *Transformer) processBigQueryTable(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.BigQuery, "table_id", suffixBigQueryTable, t.bqTableByName, t.bqTableByLabel)

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
		oldName := target.Value()
		newName := t.bqDatasetGCSByLabel[bqDatasetGCS.Label].Name + "." + oldName
		target = resources.NewGenericResource(target.ID(), newName, target.ResourceType())

		t.resources[len(t.resources)-1] = target
		t.bqTableByLabel[label] = target

		delete(t.bqTableByName, oldName)
		t.bqTableByName[newName] = target
	}
}

func (t *Transformer) processBigTable(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.BigTable, "name", suffixBigTable, t.bigTableByName, t.bigTableByLabel)
}

func (t *Transformer) processDataFlow(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.Dataflow, attributeName, suffixDataFlow, t.dataFlowByName, t.dataFlowByLabel)

	parameters, ok := conf.Attributes["parameters"]
	if !ok {
		return
	}

	if parameters, ok := parameters.(map[string]any); ok {
		sourceValue := replaceVars(conf.Attributes[attributeName].(string), t.tfConfig.Variables, t.tfConfig.Locals,
			t.yamlConfig.Draw.ReplaceableTexts)
		sourceGCS := gcpresources.ParseResourceGCS(sourceValue, conf.Labels)

		t.processDataFlowParameters(parameters, sourceGCS)
	}
}

func (t *Transformer) processDataFlowParameters(parameters map[string]any, sourceGCS *gcpresources.ResourceGCS) {
	for targetAttr, v := range parameters {
		value, ok := v.(string)
		if !ok {
			continue
		}

		value = extractTextFromTFVar(value)

		var suggestionLabels []string

		hasRelationship := true

		switch {
		case strings.HasPrefix(targetAttr, "outputTopic"):
			suggestionLabels = []string{gcpresources.LabelPubSubSubscription}
		case strings.HasPrefix(targetAttr, "outputTable"):
			parts := strings.Split(value, ":")
			if len(parts) > 1 {
				value = parts[1]
			}

			suggestionLabels = []string{gcpresources.LabelBigQueryTable}
		case strings.HasPrefix(targetAttr, "outputDirectory"):
			suggestionLabels = []string{gcpresources.LabelStorage}
		default:
			hasRelationship = false
		}

		if hasRelationship {
			targetValue := replaceVars(value, t.tfConfig.Variables, t.tfConfig.Locals,
				t.yamlConfig.Draw.ReplaceableTexts)
			targetGCS := gcpresources.ParseResourceGCS(targetValue, suggestionLabels)

			t.relationshipsMap[sourceGCS] = append(t.relationshipsMap[sourceGCS], targetGCS)
		}
	}
}

func (t *Transformer) processFunction(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.Function, attributeName, suffixFunction, t.functionByName, t.functionByLabel)

	t.processRelationshipByAttrsMap(conf, attributeName, "environment_variables", t.processFunctionEnvarsToResource)
}

func (t *Transformer) processFunctionEnvarsToResource(key string) (hasRelationship bool, suggestionLabels []string) {
	hasRelationship = true

	switch {
	case strings.HasSuffix(key, "TO_TOPIC_NAME"):
		suggestionLabels = append(suggestionLabels, gcpresources.LabelPubSub)
	default:
		hasRelationship = false
	}

	return
}

func (t *Transformer) processIoTCore(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.IoTCore, attributeName, suffixIoTCore, t.iotCoreByName, t.iotCoreByLabel)

	t.processRelationshipByAttrsMap(conf, attributeName, "event_notification_configs",
		t.processIoTCoreEventNotificationConfigs)
}

func (t *Transformer) processIoTCoreEventNotificationConfigs(
	value string,
) (hasRelationship bool, suggestionLabels []string) {
	hasRelationship = true

	switch {
	case strings.HasSuffix(value, "topic_name"):
		suggestionLabels = append(suggestionLabels, gcpresources.LabelPubSub)
	default:
		hasRelationship = false
	}

	return
}

func (t *Transformer) processPubSub(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.PubSub, "name", suffixPubSub, t.pubSubByName, t.pubSubByLabel)
}

func (t *Transformer) processPubSubSubscription(conf *hcl.Resource) {
	label := conf.Labels[1]

	name := replaceVars(conf.Attributes[attributeName].(string), t.tfConfig.Variables, t.tfConfig.Locals,
		t.yamlConfig.Draw.ReplaceableTexts)
	subsGCS := gcpresources.ParseResourceGCS(name, conf.Labels)

	topic := replaceVars(conf.Attributes[attributeTopic].(string), t.tfConfig.Variables, t.tfConfig.Locals,
		t.yamlConfig.Draw.ReplaceableTexts)
	topicGCS := gcpresources.ParseResourceGCS(topic, conf.Labels)

	t.pbSubscriptionGCSByName[name] = subsGCS
	t.pubSubByPubSubSubscriptionLabel[label] = topicGCS

	t.processRelationshipByAttrsMap(conf, attributeTopic, "push_config", t.processPubSubSubsPushConfig)
}

func (t *Transformer) processPubSubSubsPushConfig(key string) (hasRelationship bool, suggestionLabels []string) {
	if key == "push_endpoint" {
		suggestionLabels = append(suggestionLabels, gcpresources.LabelFunction)
		hasRelationship = true
	}

	return
}

func (t *Transformer) processStorage(conf *hcl.Resource) {
	t.processResource(conf, gcpresources.Storage, "name", suffixStorage, t.storageByName, t.storageByLabel)
}
