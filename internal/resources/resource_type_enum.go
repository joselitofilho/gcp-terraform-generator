package resources

import "strings"

// ResourceType represents the type of a resource.
type ResourceType string

const (
	// AppEngine represents the App Engine resource type.
	AppEngine ResourceType = "appengine"

	// BigQuery represents the BigQuery resource type.
	BigQuery ResourceType = "bigquery"

	// CloudBigTable represents the Cloud Big Table resource type.
	CloudBigTable ResourceType = "cloudbigtable"

	// CloudFunction represents the Cloud Function resource type.
	CloudFunction ResourceType = "cloudfunction"

	// CloudStorage represents the Cloud Storage resource type.
	CloudStorage ResourceType = "cloudstorage"

	// Dataflow represents the Dataflow resource type.
	Dataflow ResourceType = "dataflow"

	// IoTCore represents the IoT Core resource type.
	IoTCore ResourceType = "iotcore"

	// PubSub represents the Pub Sub resource type.
	PubSub ResourceType = "pubsub"

	// UnknownType represents an unknown resource type.
	UnknownType ResourceType = "unknown"
)

var AvailableTypes = []ResourceType{
	AppEngine, BigQuery, CloudBigTable, CloudFunction, CloudStorage, Dataflow, IoTCore, PubSub}

// String returns the string representation of a ResourceType.
func (rt ResourceType) String() string {
	switch rt {
	case AppEngine:
		return "AppEngine"
	case BigQuery:
		return "BigQuery"
	case CloudBigTable:
		return "CloudBigTable"
	case CloudFunction:
		return "CloudFunction"
	case CloudStorage:
		return "CloudStorage"
	case Dataflow:
		return "Dataflow"
	case IoTCore:
		return "IoTCore"
	case PubSub:
		return "PubSub"
	default:
		return "Unknown"
	}
}

// ParseResourceType parses a ResourceType from a string.
func ParseResourceType(s string) ResourceType {
	switch strings.ToLower(s) {
	case "appengine":
		return AppEngine
	case "bigquery":
		return BigQuery
	case "cloudbigtable":
		return CloudBigTable
	case "cloudfunction":
		return CloudFunction
	case "cloudstorage":
		return CloudStorage
	case "dataflow":
		return Dataflow
	case "iotcore":
		return IoTCore
	case "pubsub":
		return PubSub
	default:
		return UnknownType
	}
}
