package resources

import "strings"

type ResourceGCS struct {
	Type  string
	Name  string
	Label string
}

var labelByResourceType = map[ResourceType]string{
	AppEngine:   LabelAppEngine,
	BigQuery:    LabelBigQueryTable,
	BigTable:    LabelBigTable,
	Dataflow:    LabelDataFlow,
	Function:    LabelFunction,
	IoTCore:     LabelIoTCore,
	PubSub:      LabelPubSub,
	Storage:     LabelStorage,
	UnknownType: "",
}

func ParseResourceGCS(gcs string, labels []string) *ResourceGCS {
	var gcsType, name, label string

	switch {
	case strings.HasPrefix(gcs, "gs:"):
	case strings.HasPrefix(gcs, "http"):
	default:
		gcsType, name, label = parseGCSTypeAndName(gcs)
	}

	// if suggestedResType == UnknownType {
	// 	suggestedResType = inferResourceType(gcsType)
	// }

	if label == "" && len(labels) > 1 {
		label = labels[1]
	}

	if gcsType == "" && len(labels) > 0 {
		// gcsType = labelByResourceType[suggestedResType]
		gcsType = labels[0]
	}

	return &ResourceGCS{Type: gcsType, Name: name, Label: label}
}

func parseGCSTypeAndName(gcs string) (arnType, name, label string) {
	parts := strings.Split(gcs, ".")

	if len(parts) > 1 && strings.HasPrefix(parts[0], "google_") {
		arnType = parts[0]
		label = parts[1]
	} else {
		name = gcs
	}

	return arnType, name, label
}

func inferResourceType(gcsType string) ResourceType {
	switch gcsType {
	case LabelAppEngine:
		return AppEngine
	case LabelBigQueryTable:
		return BigQuery
	case LabelBigTable:
		return BigTable
	case LabelDataFlow:
		return Dataflow
	case LabelFunction:
		return Function
	case LabelIoTCore:
		return IoTCore
	case LabelPubSub:
		return PubSub
	case LabelStorage:
		return Storage
	default:
		return UnknownType
	}
}
