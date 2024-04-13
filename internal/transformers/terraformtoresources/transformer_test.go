package terraformtoresources

import (
	"testing"

	"github.com/diagram-code-generator/resources/pkg/resources"
	hcl "github.com/joselitofilho/hcl-parser-go/pkg/parser/hcl"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"

	"github.com/stretchr/testify/require"
)

var bqTableSchema = `<<EOF
# Define your BigQuery schema here
EOF`

func TestTransformer_Transform(t *testing.T) {
	type fields struct {
		yamlConfig *config.Config
		tfConfig   *hcl.Config
	}

	tests := []struct {
		name   string
		fields fields
		want   *resources.ResourceCollection
	}{
		{
			name: "empty",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig:   &hcl.Config{},
			},
			want: &resources.ResourceCollection{
				Resources:     []resources.Resource{},
				Relationships: []resources.Relationship{},
			},
		},
		{
			name: "App Engine",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_app_engine_application",
							Name:   "engine_app",
							Labels: []string{"google_app_engine_application", "engine_app"},
							Attributes: map[string]any{
								"project":     "var.project_id",
								"location_id": "US",
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("1", "engine", gcpresources.AppEngine.String())},
				Relationships: []resources.Relationship{},
			},
		},
		{
			name: "Big Query",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_bigquery_dataset",
							Name:   "dataset_dataset",
							Labels: []string{"google_bigquery_dataset", "dataset_dataset"},
							Attributes: map[string]any{
								"dataset_id": "dataset",
							},
						},
						{
							Type:   "google_bigquery_table",
							Name:   "dataset_bq_table",
							Labels: []string{"google_bigquery_table", "dataset_bq_table"},
							Attributes: map[string]any{
								"dataset_id": "google_bigquery_dataset.dataset_dataset.dataset_id",
								"table_id":   "bq",
								"schema":     bqTableSchema,
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("1", "dataset.bq", gcpresources.BigQuery.String())},
				Relationships: []resources.Relationship{},
			},
		},
		{
			name: "Big Table",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_bigtable_instance",
							Name:   "bigtable_instance",
							Labels: []string{"google_bigtable_instance", "bigtable_instance"},
							Attributes: map[string]any{
								"name": "bigtable-instance",
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("1", "bigtable-instance", gcpresources.BigTable.String())},
				Relationships: []resources.Relationship{},
			},
		},
		{
			name: "DataFlow",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_dataflow_job",
							Name:   "dataflow_df_job",
							Labels: []string{"google_dataflow_job", "dataflow_df_job"},
							Attributes: map[string]any{
								"name": "dataflow",
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("1", "dataflow", gcpresources.Dataflow.String())},
				Relationships: []resources.Relationship{},
			},
		},
		{
			name: "Function",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_cloudfunctions_function",
							Name:   "func_function",
							Labels: []string{"google_cloudfunctions_function", "func_function"},
							Attributes: map[string]any{
								"name": "func",
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("1", "func", gcpresources.Function.String())},
				Relationships: []resources.Relationship{},
			},
		},
		{
			name: "IoT Core",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_cloudiot_registry",
							Name:   "core_registry",
							Labels: []string{"google_cloudiot_registry", "core_registry"},
							Attributes: map[string]any{
								"name": "core",
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("1", "core", gcpresources.IoTCore.String())},
				Relationships: []resources.Relationship{},
			},
		},
		{
			name: "Pub Sub",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_pubsub_topic",
							Name:   "pubsub_topic",
							Labels: []string{"google_pubsub_topic", "pubsub_topic"},
							Attributes: map[string]any{
								"name": "pubsub",
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("1", "pubsub", gcpresources.PubSub.String())},
				Relationships: []resources.Relationship{},
			},
		},
		{
			name: "Storage",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_storage_bucket",
							Name:   "storage_bucket",
							Labels: []string{"google_storage_bucket", "storage_bucket"},
							Attributes: map[string]any{
								"name": "storage",
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{
					resources.NewGenericResource("1", "storage", gcpresources.Storage.String())},
				Relationships: []resources.Relationship{},
			},
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			got := NewTransformer(tc.fields.yamlConfig, tc.fields.tfConfig).Transform()

			require.Equal(t, tc.want, got)
		})
	}
}

func TestTransformer_TransformFromDataFLowToResource(t *testing.T) {
	type fields struct {
		yamlConfig *config.Config
		tfConfig   *hcl.Config
	}

	dataflowResource := resources.NewGenericResource("1", "dataflow", gcpresources.Dataflow.String())

	bqResource := resources.NewGenericResource("2", "dataset.bq", gcpresources.BigQuery.String())
	bqBackupResource := resources.NewGenericResource("3", "dataset.backup", gcpresources.BigQuery.String())

	pubSubAppEngineResource := resources.NewGenericResource("2", "psengine", gcpresources.PubSub.String())
	pubSubFuncResource := resources.NewGenericResource("3", "psfunc", gcpresources.PubSub.String())

	storageBucketResource := resources.NewGenericResource("2", "storage", gcpresources.Storage.String())
	backupBucketResource := resources.NewGenericResource("3", "backup", gcpresources.Storage.String())

	tests := []struct {
		name   string
		fields fields
		want   *resources.ResourceCollection
	}{
		{
			name: "from dataflow to big query table",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_dataflow_job",
							Name:   "dataflow_df_job",
							Labels: []string{"google_dataflow_job", "dataflow_df_job"},
							Attributes: map[string]any{
								"name": "dataflow",
								"parameters": map[string]any{
									"outputTable": "${var.project_id}:" +
										"${google_bigquery_table.dataset_bq_table.dataset_id}." +
										"${google_bigquery_table.dataset_bq_table.table_id}",
								},
							},
						},
						{
							Type:   "google_bigquery_dataset",
							Name:   "dataset_dataset",
							Labels: []string{"google_bigquery_dataset", "dataset_dataset"},
							Attributes: map[string]any{
								"dataset_id": "dataset",
							},
						},
						{
							Type:   "google_bigquery_table",
							Name:   "dataset_bq_table",
							Labels: []string{"google_bigquery_table", "dataset_bq_table"},
							Attributes: map[string]any{
								"dataset_id": "google_bigquery_dataset.dataset_dataset.dataset_id",
								"table_id":   "bq",
								"schema":     bqTableSchema,
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources:     []resources.Resource{dataflowResource, bqResource},
				Relationships: []resources.Relationship{{Source: dataflowResource, Target: bqResource}},
			},
		},
		{
			name: "from dataflow to multiple big query tables",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_dataflow_job",
							Name:   "dataflow_df_job",
							Labels: []string{"google_dataflow_job", "dataflow_df_job"},
							Attributes: map[string]any{
								"name": "dataflow",
								"parameters": map[string]any{
									"outputTable1": "${var.project_id}:" +
										"${google_bigquery_table.dataset_bq_table.dataset_id}." +
										"${google_bigquery_table.dataset_bq_table.table_id}",
									"outputTable2": "${var.project_id}:dataset.backup",
								},
							},
						},
						{
							Type:   "google_bigquery_dataset",
							Name:   "dataset_dataset",
							Labels: []string{"google_bigquery_dataset", "dataset_dataset"},
							Attributes: map[string]any{
								"dataset_id": "dataset",
							},
						},
						{
							Type:   "google_bigquery_table",
							Name:   "dataset_bq_table",
							Labels: []string{"google_bigquery_table", "dataset_bq_table"},
							Attributes: map[string]any{
								"dataset_id": "google_bigquery_dataset.dataset_dataset.dataset_id",
								"table_id":   "bq",
								"schema":     bqTableSchema,
							},
						},
						{
							Type:   "google_bigquery_table",
							Name:   "dataset_backup_table",
							Labels: []string{"google_bigquery_table", "dataset_backup_table"},
							Attributes: map[string]any{
								"dataset_id": "google_bigquery_dataset.dataset_dataset.dataset_id",
								"table_id":   "backup",
								"schema":     "{}",
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{dataflowResource, bqResource, bqBackupResource},
				Relationships: []resources.Relationship{
					{Source: dataflowResource, Target: bqResource},
					{Source: dataflowResource, Target: bqBackupResource},
				},
			},
		},
		{
			name: "from dataflow to pub sub",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_dataflow_job",
							Name:   "dataflow_df_job",
							Labels: []string{"google_dataflow_job", "dataflow_df_job"},
							Attributes: map[string]any{
								"name": "dataflow",
								"parameters": map[string]any{
									"outputTopic": "google_pubsub_subscription.psengine_subscription.name",
								},
							},
						},
						{
							Type:   "google_pubsub_topic",
							Name:   "psengine_topic",
							Labels: []string{"google_pubsub_topic", "psengine_topic"},
							Attributes: map[string]any{
								"name": "psengine",
							},
						},
						{
							Type:   "google_pubsub_subscription",
							Name:   "psengine_subscription",
							Labels: []string{"google_pubsub_subscription", "psengine_subscription"},
							Attributes: map[string]any{
								"name":  "psengine-subscription",
								"topic": "google_pubsub_topic.psengine_topic.name",
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources:     []resources.Resource{dataflowResource, pubSubAppEngineResource},
				Relationships: []resources.Relationship{{Source: dataflowResource, Target: pubSubAppEngineResource}},
			},
		},
		{
			name: "from dataflow to multiple pub subs",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_dataflow_job",
							Name:   "dataflow_df_job",
							Labels: []string{"google_dataflow_job", "dataflow_df_job"},
							Attributes: map[string]any{
								"name": "dataflow",
								"parameters": map[string]any{
									"outputTopic1": "psengine-subscription",
									"outputTopic2": "google_pubsub_subscription.psfunc_subscription.name",
								},
							},
						},
						{
							Type:   "google_pubsub_topic",
							Name:   "psengine_topic",
							Labels: []string{"google_pubsub_topic", "psengine_topic"},
							Attributes: map[string]any{
								"name": "psengine",
							},
						},
						{
							Type:   "google_pubsub_subscription",
							Name:   "psengine_subscription",
							Labels: []string{"google_pubsub_subscription", "psengine_subscription"},
							Attributes: map[string]any{
								"name":  "psengine-subscription",
								"topic": "google_pubsub_topic.psengine_topic.name",
							},
						},
						{
							Type:   "google_pubsub_topic",
							Name:   "psfunc_topic",
							Labels: []string{"google_pubsub_topic", "psfunc_topic"},
							Attributes: map[string]any{
								"name": "psfunc",
							},
						},
						{
							Type:   "google_pubsub_subscription",
							Name:   "psfunc_subscription",
							Labels: []string{"google_pubsub_subscription", "psfunc_subscription"},
							Attributes: map[string]any{
								"name":  "psfunc-subscription",
								"topic": "google_pubsub_topic.psfunc_topic.name",
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{dataflowResource, pubSubAppEngineResource, pubSubFuncResource},
				Relationships: []resources.Relationship{
					{Source: dataflowResource, Target: pubSubAppEngineResource},
					{Source: dataflowResource, Target: pubSubFuncResource},
				},
			},
		},
		{
			name: "from dataflow to storage",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_dataflow_job",
							Name:   "dataflow_df_job",
							Labels: []string{"google_dataflow_job", "dataflow_df_job"},
							Attributes: map[string]any{
								"name": "dataflow",
								"parameters": map[string]any{
									"outputDirectory": "gs://${google_storage_bucket.storage_bucket.name}/output/",
								},
							},
						},
						{
							Type:   "google_storage_bucket",
							Name:   "storage_bucket",
							Labels: []string{"google_storage_bucket", "storage_bucket"},
							Attributes: map[string]any{
								"name":     "storage",
								"location": "US",
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources:     []resources.Resource{dataflowResource, storageBucketResource},
				Relationships: []resources.Relationship{{Source: dataflowResource, Target: storageBucketResource}},
			},
		},
		{
			name: "from dataflow to storages",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig: &hcl.Config{
					Resources: []*hcl.Resource{
						{
							Type:   "google_dataflow_job",
							Name:   "dataflow_df_job",
							Labels: []string{"google_dataflow_job", "dataflow_df_job"},
							Attributes: map[string]any{
								"name": "dataflow",
								"parameters": map[string]any{
									"outputDirectory1": "gs://${google_storage_bucket.storage_bucket.name}/output/",
									"outputDirectory2": "gs://backup/output/",
								},
							},
						},
						{
							Type:   "google_storage_bucket",
							Name:   "storage_bucket",
							Labels: []string{"google_storage_bucket", "storage_bucket"},
							Attributes: map[string]any{
								"name":     "storage",
								"location": "US",
							},
						},
						{
							Type:   "google_storage_bucket",
							Name:   "backup_bucket",
							Labels: []string{"google_storage_bucket", "backup_bucket"},
							Attributes: map[string]any{
								"name":     "backup",
								"location": "US",
							},
						},
					},
				},
			},
			want: &resources.ResourceCollection{
				Resources: []resources.Resource{dataflowResource, storageBucketResource, backupBucketResource},
				Relationships: []resources.Relationship{
					{Source: dataflowResource, Target: storageBucketResource},
					{Source: dataflowResource, Target: backupBucketResource},
				},
			},
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			tr := NewTransformer(
				tc.fields.yamlConfig,
				tc.fields.tfConfig,
			)

			got := tr.Transform()

			require.True(t, tc.want.Equal(got))
		})
	}
}

func TestTransformer_hasResourceMatched(t *testing.T) {
	type fields struct {
		yamlConfig *config.Config
		tfConfig   *hcl.Config
	}

	type args struct {
		res     resources.Resource
		filters config.Filters
	}

	appEngineResource := resources.NewGenericResource("id", "myAppEngine", gcpresources.AppEngine.String())

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "match",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig:   &hcl.Config{},
			},
			args: args{
				res: appEngineResource,
				filters: config.Filters{
					"appengine": config.Filter{Match: []string{"^my"}},
				},
			},
			want: true,
		},
		{
			name: "not match",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig:   &hcl.Config{},
			},
			args: args{
				res: appEngineResource,
				filters: config.Filters{
					"appengine": config.Filter{NotMatch: []string{"^my"}},
				},
			},
			want: false,
		},
		{
			name: "nil resource",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig:   &hcl.Config{},
			},
			args: args{
				res: nil,
				filters: config.Filters{
					"appengine": config.Filter{NotMatch: []string{"^my"}},
				},
			},
			want: false,
		},
		{
			name: "no filter",
			fields: fields{
				yamlConfig: &config.Config{},
				tfConfig:   &hcl.Config{},
			},
			args: args{
				res:     appEngineResource,
				filters: nil,
			},
			want: true,
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			tr := NewTransformer(tc.fields.yamlConfig, tc.fields.tfConfig)

			got := tr.hasResourceMatched(tc.args.res, tc.args.filters)

			require.Equal(t, tc.want, got)
		})
	}
}