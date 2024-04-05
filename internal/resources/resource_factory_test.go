package resources

import (
	"testing"

	resources "github.com/diagram-code-generator/resources/pkg/resources"
	"github.com/stretchr/testify/require"
)

func TestGCPResourceFactory_CreateResource(t *testing.T) {
	type args struct {
		id    string
		value string
		style string
	}

	tests := []struct {
		name string
		args args
		want resources.Resource
	}{
		{
			name: "App Engine Resource",
			args: args{
				id:    "APPENGINE_ID",
				value: "appengine",
				style: "IwIiBoZWlnaHQ9IjE2LjAyMDAwMDQ1Nzc2MzY3MiIgZmlsbC1ydWxlPSJldmVub2RkIiB2aW",
			},
			want: resources.NewGenericResource("APPENGINE_ID", "appengine", AppEngine.String()),
		},
		{
			name: "Big Query Resource",
			args: args{
				id:    "BIGQUERY_ID",
				value: "bigquery",
				style: "IwLjAwMTA0NTIyNzA1MDc4IiBoZWlnaHQ9IjIwLjAwMTA0NTIyNzA1MDc4IiBmaWxsLXJ1bG",
			},
			want: resources.NewGenericResource("BIGQUERY_ID", "bigquery", BigQuery.String()),
		},
		{
			name: "Big Table Resource",
			args: args{
				id:    "BIGTABLE_ID",
				value: "bigtable",
				style: "E3Ljk1Njk3Nzg0NDIzODI4IiBoZWlnaHQ9IjIwLjAwOTI1NjM2MjkxNTA0IiB2aWV3Qm94PS",
			},
			want: resources.NewGenericResource("BIGTABLE_ID", "bigtable", BigTable.String()),
		},
		{
			name: "Function Resource",
			args: args{
				id:    "FUNCTION_ID",
				value: "func",
				style: "IwIiBoZWlnaHQ9IjE5Ljk4OTk5OTc3MTExODE2NCIgdmlld0JveD0iMCAwIDIwIDE5Ljk4OT",
			},
			want: resources.NewGenericResource("FUNCTION_ID", "func", Function.String()),
		},
		{
			name: "Storage Resource",
			args: args{
				id:    "STORAGE_ID",
				value: "storage",
				style: "IwIiBoZWlnaHQ9IjE2IiB2aWV3Qm94PSIwIDAgMjAgMTYiPiYjeGE7CTxzdHlsZSB0eXBlPS",
			},
			want: resources.NewGenericResource("STORAGE_ID", "storage", Storage.String()),
		},
		{
			name: "DataFlow Resource",
			args: args{
				id:    "DATAFLOW_ID",
				value: "dataflow",
				style: "E0LjUxOTk5OTUwNDA4OTM1NSIgaGVpZ2h0PSIyMCIgdmlld0JveD0iMCAwIDE0LjUxOTk5OT",
			},
			want: resources.NewGenericResource("DATAFLOW_ID", "dataflow", Dataflow.String()),
		},
		{
			name: "IoT Core Resource",
			args: args{
				id:    "CORE_ID",
				value: "core",
				style: "E5LjcwNjUzMTUyNDY1ODIwMyIgaGVpZ2h0PSIxOS45ODM4MjE4Njg4OTY0ODQiIGZpbGwtcn",
			},
			want: resources.NewGenericResource("CORE_ID", "core", IoTCore.String()),
		},
		{
			name: "Pub Sub API Resource",
			args: args{
				id:    "PUBSUB_ID",
				value: "pubsub",
				style: "E4LjMxOTk5OTY5NDgyNDIyIiBoZWlnaHQ9IjIwLjAwMDAwMTkwNzM0ODYzMyIgdmlld0JveD",
			},
			want: resources.NewGenericResource("PUBSUB_ID", "pubsub", PubSub.String()),
		},
		{
			name: "Unknown",
			args: args{
				id:    "ID",
				value: "value",
				style: "any",
			},
			want: nil,
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			f := &GCPResourceFactory{}
			got := f.CreateResource(tc.args.id, tc.args.value, tc.args.style)

			require.Equal(t, tc.want, got)
		})
	}
}
