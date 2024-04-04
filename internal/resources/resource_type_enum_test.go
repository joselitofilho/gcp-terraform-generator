package resources

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResourceType_String(t *testing.T) {
	tests := []struct {
		name string
		rt   ResourceType
		want string
	}{
		{name: "AppEngine", rt: AppEngine, want: "AppEngine"},
		{name: "BigQuery", rt: BigQuery, want: "BigQuery"},
		{name: "BigTable", rt: BigTable, want: "BigTable"},
		{name: "Function", rt: Function, want: "Function"},
		{name: "Storage", rt: Storage, want: "Storage"},
		{name: "Dataflow", rt: Dataflow, want: "Dataflow"},
		{name: "IoTCore", rt: IoTCore, want: "IoTCore"},
		{name: "PubSub", rt: PubSub, want: "PubSub"},
		{name: "Unknown", rt: UnknownType, want: "Unknown"},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			got := tc.rt.String()

			require.Equal(t, tc.want, got)
		})
	}
}

func TestParseResourceType(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output ResourceType
	}{
		{name: "Parse AppEngine", input: "appengine", output: AppEngine},
		{name: "Parse BigQuery", input: "bigquery", output: BigQuery},
		{name: "Parse BigTable", input: "bigtable", output: BigTable},
		{name: "Parse Function", input: "function", output: Function},
		{name: "Parse Storage", input: "storage", output: Storage},
		{name: "Parse Dataflow", input: "dataflow", output: Dataflow},
		{name: "Parse IoTCore", input: "iotcore", output: IoTCore},
		{name: "Parse PubSub", input: "pubsub", output: PubSub},
		{name: "Parse Unknown", input: "unknown", output: UnknownType},
		{name: "Parse lowercase", input: "sqs", output: UnknownType},
		{name: "Parse uppercase", input: "SNS", output: UnknownType},
		{name: "Parse mixed case", input: "ApIgAtEwAy", output: UnknownType},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			result := ParseResourceType(tc.input)

			require.Equal(t, tc.output, result)
		})
	}
}
