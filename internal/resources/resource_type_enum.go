package resources

import "strings"

// ResourceType represents the type of a resource.
type ResourceType string

const (
	// AppEngine represents the App Engine resource type.
	AppEngine ResourceType = "appengine"

	// BigQuery represents the BigQuery resource type.
	BigQuery ResourceType = "bigquery"

	// BigTable represents the Cloud Big Table resource type.
	BigTable ResourceType = "bigtable"

	// Dataflow represents the Dataflow resource type.
	Dataflow ResourceType = "dataflow"

	// Function represents the Cloud Function resource type.
	Function ResourceType = "function"

	// IoTCore represents the IoT Core resource type.
	IoTCore ResourceType = "iotcore"

	// PubSub represents the Pub Sub resource type.
	PubSub ResourceType = "pubsub"

	// Storage represents the Cloud Storage resource type.
	Storage ResourceType = "storage"

	// UnknownType represents an unknown resource type.
	UnknownType ResourceType = "unknown"
)

var AvailableTypes = []ResourceType{
	AppEngine, BigQuery, BigTable, Function, Storage, Dataflow, IoTCore, PubSub}

// String returns the string representation of a ResourceType.
func (rt ResourceType) String() string {
	switch rt {
	case AppEngine:
		return "AppEngine"
	case BigQuery:
		return "BigQuery"
	case BigTable:
		return "BigTable"
	case Dataflow:
		return "Dataflow"
	case Function:
		return "Function"
	case IoTCore:
		return "IoTCore"
	case PubSub:
		return "PubSub"
	case Storage:
		return "Storage"
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
	case "bigtable":
		return BigTable
	case "dataflow":
		return Dataflow
	case "function":
		return Function
	case "iotcore":
		return IoTCore
	case "pubsub":
		return PubSub
	case "storage":
		return Storage
	default:
		return UnknownType
	}
}
