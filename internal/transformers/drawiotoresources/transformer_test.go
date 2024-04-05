package drawiotoresources

import (
	"testing"

	drawioxml "github.com/joselitofilho/drawio-parser-go/pkg/parser/xml"

	"github.com/diagram-code-generator/resources/pkg/resources"
	awsresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"

	"github.com/stretchr/testify/require"
)

func TestParseResources(t *testing.T) {
	type args struct {
		mxFile *drawioxml.MxFile
	}

	coreResource := resources.NewGenericResource("CORE_ID", "core", awsresources.IoTCore.String())
	dataFlowResource := resources.NewGenericResource("DATAFLOW_ID", "dataflow", awsresources.Dataflow.String())

	tests := []struct {
		name      string
		args      args
		want      *resources.ResourceCollection
		targetErr error
	}{
		{
			name: "App Engine Resource",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{{
									ID:    "APPENGINE_ID",
									Value: "appengine",
									Style: "IwIiBoZWlnaHQ9IjE2LjAyMDAwMDQ1Nzc2MzY3MiIgZmlsbC1ydWxlPSJldmVub2RkIiB2aW",
								}},
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("APPENGINE_ID", "appengine", awsresources.AppEngine.String())},
			},
		},
		{
			name: "Big Query Resource",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{{
									ID:    "BIGQUERY_ID",
									Value: "bigquery",
									Style: "IwLjAwMTA0NTIyNzA1MDc4IiBoZWlnaHQ9IjIwLjAwMTA0NTIyNzA1MDc4IiBmaWxsLXJ1bG",
								}},
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("BIGQUERY_ID", "bigquery", awsresources.BigQuery.String())},
			},
		},
		{
			name: "Big Table Resource",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{{
									ID:    "BIGTABLE_ID",
									Value: "bigtable",
									Style: "E3Ljk1Njk3Nzg0NDIzODI4IiBoZWlnaHQ9IjIwLjAwOTI1NjM2MjkxNTA0IiB2aWV3Qm94PS",
								}},
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("BIGTABLE_ID", "bigtable", awsresources.BigTable.String())},
			},
		},
		{
			name: "Function Resource",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{{
									ID:    "FUNCTION_ID",
									Value: "func",
									Style: "IwIiBoZWlnaHQ9IjE5Ljk4OTk5OTc3MTExODE2NCIgdmlld0JveD0iMCAwIDIwIDE5Ljk4OT",
								}},
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("FUNCTION_ID", "func", awsresources.Function.String())},
			},
		},
		{
			name: "Storage Resource",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{{
									ID:    "STORAGE_ID",
									Value: "storage",
									Style: "IwIiBoZWlnaHQ9IjE2IiB2aWV3Qm94PSIwIDAgMjAgMTYiPiYjeGE7CTxzdHlsZSB0eXBlPS",
								}},
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("STORAGE_ID", "storage", awsresources.Storage.String())},
			},
		},
		{
			name: "DataFlow Resource",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{{
									ID:    "DATAFLOW_ID",
									Value: "dataflow",
									Style: "E0LjUxOTk5OTUwNDA4OTM1NSIgaGVpZ2h0PSIyMCIgdmlld0JveD0iMCAwIDE0LjUxOTk5OT",
								}},
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("DATAFLOW_ID", "dataflow", awsresources.Dataflow.String())},
			},
		},
		{
			name: "IoT Core Resource",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{{
									ID:    "CORE_ID",
									Value: "core",
									Style: "E5LjcwNjUzMTUyNDY1ODIwMyIgaGVpZ2h0PSIxOS45ODM4MjE4Njg4OTY0ODQiIGZpbGwtcn",
								}},
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{coreResource},
			},
		},
		{
			name: "Pub Sub API Resource",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{{
									ID:    "PUBSUB_ID",
									Value: "pubsub",
									Style: "E4LjMxOTk5OTY5NDgyNDIyIiBoZWlnaHQ9IjIwLjAwMDAwMTkwNzM0ODYzMyIgdmlld0JveD",
								}},
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("PUBSUB_ID", "pubsub", awsresources.PubSub.String())},
			},
		},
		{
			name: "Empty MxFile",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{},
							},
						},
					},
				},
			},
			want: resources.NewResourceCollection(),
		},
		{
			name: "Two Connected Resources",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{{
									ID:    "CORE_ID",
									Value: "core",
									Style: "E5LjcwNjUzMTUyNDY1ODIwMyIgaGVpZ2h0PSIxOS45ODM4MjE4Njg4OTY0ODQiIGZpbGwtcn",
								}, {
									ID:    "DATAFLOW_ID",
									Value: "dataflow",
									Style: "E0LjUxOTk5OTUwNDA4OTM1NSIgaGVpZ2h0PSIyMCIgdmlld0JveD0iMCAwIDE0LjUxOTk5OT",
								}, {
									ID: "3", Source: "CORE_ID", Target: "DATAFLOW_ID",
								}},
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources:     []resources.Resource{coreResource, dataFlowResource},
				Relationships: []resources.Relationship{{Source: coreResource, Target: dataFlowResource}},
			},
		},
		{
			name: "Single Unknown Resource",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{
									{ID: "1", Value: "Resource A", Style: "styleA"},
								},
							},
						},
					},
				},
			},
			want: resources.NewResourceCollection(),
		},
		{
			name: "Two Connected Unknown Resources",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{
									{ID: "1", Value: "Resource A", Style: "styleA"},
									{ID: "2", Value: "Resource B", Style: "styleB"},
									{ID: "3", Source: "1", Target: "2"},
								},
							},
						},
					},
				},
			},
			want: resources.NewResourceCollection(),
		},
		{
			name: "Multiple Unknown Resources",
			args: args{
				mxFile: &drawioxml.MxFile{
					Diagram: drawioxml.Diagram{
						MxGraphModel: drawioxml.MxGraphModel{
							Root: drawioxml.Root{
								MxCells: []drawioxml.MxCell{
									{ID: "1", Value: "Resource A", Style: "styleA"},
									{ID: "2", Value: "Resource B", Style: "styleB"},
								},
							},
						},
					},
				},
			},
			want: resources.NewResourceCollection(),
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			got, err := Transform(tc.args.mxFile)

			if tc.targetErr == nil {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			} else {
				require.ErrorIs(t, err, tc.targetErr)
			}
		})
	}
}
