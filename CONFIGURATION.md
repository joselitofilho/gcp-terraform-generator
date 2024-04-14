# Configuration

The configuration is organized into the following sections:

- [**Override default templates**](#override_default_templates): Configuration for overriding default templates.
- [**Diagram**](#diagram): Diagram configurations include modules to specify the URL pointing to the GitHub repository 
for the resources module.
- [**Structure**](#structure): Manages stacks with multiple environments (`dev`, `uat`, `prd`) and includes default 
templates and specific configurations for various GCP services.
  - **Stacks**: Configuration for different stacks.
  - **Default Templates**: Default Terraform templates for creating stacks.
- [**App Engines**](#appengines): Configuration for App Engine resources.
- [**Big Query**](#bigquery): Configuration for BigQuery resources.
- [**Big Tables**](#bigtables): Configuration for Bigtable resources.
- [**DataFlows**](#dataflows): Configuration for DataFlow resources.
- [**Functions**](#functions): Configuration for Cloud Functions resources.
- [**IoT Cores**](#iotcores): Configuration for IoT Core resources.
- [**Pub Subs**](#pubsubs): Configuration for Pub/Sub resources.
- [**Storages**](#storages): Configuration for Storage resources.
- [**Draw**](#draw): Draw configurations.

## File Structure

- **`stacks`**: Contains configurations for different stacks, each with folders for different environments (`dev`, `uat`, `prd`), default templates, and specific configurations for GCP services.
- **`default_templates`**: Provides default Terraform templates for creating stacks, including backend and provider configurations, module instantiation, and variable definitions.

### override_default_templates

Configuration for overriding default templates.

```yaml
override_default_templates:
  # Templates for App Engine.
  appengine:
    - appengine.tf: |-
        resource "google_app_engine_application" "{{ToSnake $.Name}}_app" {}
  # Templates for Big Query.
  bigquery:
    - bigquery.tf: |-
        resource "google_bigquery_table" "{{ToSnake $.Dataset}}_{{ToSnake $.Table}}_table" {}
  # Templates for Big Table.
  bigtable:
    - bigtable.tf: |-
        resource "google_bigtable_instance" "{{ToSnake $.Name}}_instance" {}
  # Templates for DataFlow.
  dataflow:
    - dataflow.tf: |-
        resource "google_dataflow_job" "{{ToSnake $.Name}}_df_job" {}
  # Templates for Function.
  function:
    - function.tf: |-
        resource "google_cloudfunctions_function" "{{ToSnake $.Name}}_function" {}
  # Templates for IoT Core.
  iotcore:
    - iotcore.tf: |-
        resource "google_cloudiot_registry" "{{ToSnake $.Name}}_registry" {}
  # Templates for Pub Sub.
  pubsub:
    - pubsub.tf: |-
        resource "google_pubsub_topic" "{{ToSnake $.Name}}_topic" {}
  # Templates for Storage.
  storage:
    - storage.tf: |-
        resource "google_storage_bucket" "{{ToSnake $.Name}}_bucket" {}
```

### diagram

Diagram configurations include modules to specify the URL pointing to the GitHub repository for the resources module.

```yaml
diagram:
  # To specify the stack name for the diagram.
  stack_name: mystack
```

### structure

Structure for managing stacks with multiple environments.

```yaml
structure:
  # Stacks section. Each stack configuration includes folders for different environments (`dev`, `uat`, `prd`, etc.),
  # default templates, and specific configurations for IoT cores, pub subs, big tables, and more.
  stacks:
    - name: mystack
      # Folders for different environments.
      folders:
        # Development environment.
        - name: dev
          # Terraform configuration files for the development environment.
          files:
            - name: main.tf
            - name: terragrunt.hcl
            - name: vars.tf
        # User Acceptance Testing environment.
        - name: uat
          # Terraform configuration files for the User Acceptance Testing environment.
          files:
            - name: main.tf
            - name: terragrunt.hcl
            - name: vars.tf
        # Production environment.
        - name: prd
          # Terraform configuration files for the production environment.
          files:
            - name: main.tf
            - name: terragrunt.hcl
            - name: vars.tf
        # Common module.
        - name: mod
          # Terraform configuration files for the common module.
          files:
            - name: main.tf
              # Template for generating stack_name based on environment.
              tmpl: |-
                locals {
                  stack_name = "{{$.StackName}}-${var.environment}"
                }
            - name: vars.tf
        # Cloud Functions.
        - name: functions
        # Placeholder folder for any other configurations.
        - name: anyFolder
          # Files within the arbitrary folder.
          files:
            - name: anyFile.txt
      # Files in the root folder.
      files:
        - name: anyRootFile.txt

  # Default templates are provided for creating stacks. These templates include backend configuration, provider
  # configuration, module instantiation, and variable definitions.
  default_templates:
    - main.tf: |-
        terraform {
          required_providers {
            google = {
              source  = "hashicorp/google"
              version = "~> 4.84"
            }
          }
        }

        provider "google" {
          project = var.project_id
          zone    = var.zone
          region  = var.region
        }

      terragrunt.hcl: |-
        include {
          path = find_in_parent_folders()
        }

      vars.tf: |-
        variable "project_id" {
          type = string
        }

        variable "zone" {
          type = string
        }

        variable "region" {
          type = string
        }

        variable "environment" {
          type = string
        }
```

### appengines

Configuration for App Engine resources.

```yaml
appengines:
  # Name of the App Engine.
  - name: "myEngine"
    # The project ID to create the application under.
    project: "${var.project_id}"
    # The location to serve the app from.
    location_id: "us-central"
    # Optional. List of files that we can customize.
    files:
      - name: "my-engine.tf"
        # Template for the Terraform file defining the App Engine resource.
        tmpl: |-
          resource "google_app_engine_application" "{{ToSnake $.Name}}_app" {}
```

### bigquery

Configurations for Big Query.

```yaml
bigquery:
  # Name of the Big Query. The first part is optional and reprensents the dataset.
  - name: "myDataset.myTable"
    # Optional. A JSON schema for the table.
    schema: |-
      <<EOF
        # Define your Big Query schema here
      EOF
    # Optional. List of files that we can customize
    files:
      - name: "my-table.tf"
        # Template for the Terraform file defining the Big Query resource.
        tmpl: |-
          resource "google_bigquery_table" "{{ToSnake $.Dataset}}_{{ToSnake $.Table}}_table" {}
```

### bigtables

Configurations for Big Tables.

```yaml
bigtables:
  # Name of the Big Table.
  - name: "myBigTable"
    # Optional. A map of key/value label pairs to assign to the resource.
    labels:
      label1: value1
      label2: value2
    # Optional. List of files that we can customize.
    files:
      - name: "my-big-table.tf"
        # Template for the Terraform file defining the Big Table resource.
        tmpl: |-
          resource "google_bigtable_instance" "{{ToSnake $.Name}}_instance" {}
```

### dataflows

Configurations for DataFlows.

```yaml
dataflows:
  # Name of the DataFlow.
  - name: "dataflow"
    # The GCS path to the Dataflow job template.
    template_gcs_path: "gs://example-bucket/dataflow-template"
    # A writeable location on GCS for the Dataflow job to dump its temporary data.
    temp_gcs_location: "gs://example-bucket/temp"
    # Optional. Input topics for the DataFlow.
    input_topics:
      - "projects/example-project/topics/input-topic"
    # Optional. Output topics for the DataFlow.
    output_topics:
      - "projects/example-project/topics/output-topic"
    # Optional. Output directories for the DataFlow.
    output_directories:
      - "gs://example-bucket/output"
    # Optional. Output tables for the DataFlow.
    output_tables:
      - "project:dataset.table"
    # Optional. List of files that we can customize.
    files:
      - name: "my-job.tf"
        # Template for the Terraform file defining the DataFlow resource.
        tmpl: |-
          resource "google_dataflow_job" "{{ToSnake $.Name}}_df_job" {}
```

### functions

Configurations for Functions.

```yaml
functions:
  # Name of the function.
  - name: "func"
    # Source of the function code.
    source: "."
    # The runtime in which the function is going to run.
    runtime: "go116"
    # Optional. The GCS bucket containing the zip archive which contains the function.
    source_archive_bucket: "example-bucket"
    # Optional. The source archive object (file) in archive bucket.
    source_archive_object: "function.zip"
    # Optional. Any HTTP request (of a supported type) to the endpoint will trigger function execution.
    trigger_http: "true"
    # Optional. Name of the function that will be executed when the Google Cloud Function is triggered.
    entry_point: "ExampleFunction"
    # Optional. A map of key/value environment variable pairs to assign to the function.
    envars:
      ENV_VAR1: "value1"
      ENV_VAR2: "value2"
    # Optional. List of files that we can customize.
    files:
      - name: "function.tf"
        # Template for the Terraform file defining the Function resource.
        tmpl: |-
          resource "google_cloudfunctions_function" "{{ToSnake $.Name}}_function" {}
```

### iotcores

Configurations for IoT Cores.

```yaml
iotcores:
  # Name of the IoT Core.
  - name: "core"
    # Optional. List of configurations for event notifications, such as PubSub topics to publish device events to.
    event_notification_configs:
      # PubSub topic name to publish device state updates.
      - pubsub_topic_name: google_pubsub_topic.pubsub_topic.id
    # Optional. List of files that we can customize.
    files:
      - name: "my-iotcore.tf"
        # Template for the Terraform file defining the IoT Core resource.
        tmpl: |-
          resource "google_cloudiot_registry" "{{ToSnake $.Name}}_registry" {}
```

### pubsubs

Configurations for Pub Subs.

```yaml
pubsubs:
  # Name of the Pub Sub.
  - name: "pubsub"
    # Optional. A set of key/value label pairs to assign to this Topic.
    labels:
      foo: "bar"
    # Optional. If push delivery is used with this subscription, this field is used to configure it. An empty value
    # signifies that the subscriber will pull and ack messages using API methods. Otherwise, a URL locating the endpoint
    # to which messages should be pushed.
    push_endpoint: google_cloudfunctions_function.func_function.https_trigger_url
    # Optional. List of files that we can customize.
    files:
      - name: "my-iotcore.tf"
        # Template for the Terraform file defining the Pub Sub resource.
        tmpl: |-
          resource "google_pubsub_topic" "{{ToSnake $.Name}}_topic" {}
```

### storages

Configurations for Storages.

```yaml
storages:
  # Name of the storage.
  - name: "storage"
    # The GCS location.
    location: "US"
    # Optional. List of files that we can customize.
    files:
      - name: "my-storage.tf"
        # Template for the Terraform file defining the Storage resource.
        tmpl: |-
          resource "google_storage_bucket" "{{ToSnake $.Name}}_bucket" {}
```

### draw

Draw configurations includes graph orientation, images and filters.

```yaml
draw:
  # The diagram's name will also serve as the name of the output file. Example: diagram.dot.
  name: diagram
  # Defines the direction of graph layout. See: https://graphviz.org/docs/attrs/rankdir/
  orientation: LR
  # Definitions of images for the available resources
  images:
    appengine: "assets/diagram/app_engine.svg"
    bigtable: "assets/diagram/big_table.svg"
    bigquery: "assets/diagram/bigquery.svg"
    dataflow: "assets/diagram/dataflow.svg"
    function: "assets/diagram/function.svg"
    kinesis: "assets/diagram/kinesis_data_stream.svg"
    iotcore: "assets/diagram/iot_core.svg"
    pubsub: "assets/diagram/pub_sub.svg"
    storage: "assets/diagram/storage.svg"
  # Define replaceable texts for the diagram.
  replaceable_texts:
    "-text-": ""
    "other-text": "-ot-"
  # Filters for matching and excluding resources by name.
  filters:
    appengine:
      match:
      not_match:
    bigtable:
      match:
      not_match:
    bigquery:
      match:
      not_match:
    dataflow:
      match:
      not_match:
    function:
      match:
      not_match:
    kinesis:
      match:
      not_match:
    iotcore:
      match:
      not_match:
    pubsub:
      match:
      not_match:
    storage:
      # Patterns to match
      match:
        - "^ProcessedName" # Match regex pattern for ProcessedLocation
      # Patterns to exclude
      not_match:
        - "^ProcessedA" # Exclude regex pattern for ProcessedA
        - "^ProcessedB" # Exclude regex pattern for ProcessedB
```

- Available resources: [internal/resources/resource_type_enum.go](internal/resources/resource_type_enum.go)
- Recommend image size: 40px x 40px

Image list of GCP provider:

#### analytics

| Image                                       | Resource   | Path              |
| :-----------------------------------------: | :--------- | :---------------- |
| ![](assets/diagram/bigquery.svg)            | bigquery   | assets/diagram/bigquery.svg |
| ![](assets/diagram/dataflow.svg)            | dataflow   | assets/diagram/dataflow.svg |
| ![](assets/diagram/pub_sub.svg)             | pubsub     | assets/diagram/pub_sub.svg |

#### compute

| Image                                       | Resource   | Path              |
| :-----------------------------------------: | :--------- | :---------------- |
| ![](assets/diagram/app_engine.svg)          | appengine  | assets/diagram/app_engine.svg |
| ![](assets/diagram/function.svg)            | function   | assets/diagram/function.svg |

#### database

| Image                                       | Resource   | Path              |
| :-----------------------------------------: | :--------- | :---------------- |
| ![](assets/diagram/big_table.svg)           | bigtable   | assets/diagram/big_table.svg |

#### iot

| Image                                       | Resource   | Path              |
| :-----------------------------------------: | :--------- | :---------------- |
| ![](assets/diagram/iot_core.svg)            | iotcore    | assets/diagram/iot_core.svg |

#### storage

| Image                                       | Resource   | Path              |
| :-----------------------------------------: | :--------- | :---------------- |
| ![](assets/diagram/storage.svg)             | storage    | assets/diagram/storage.svg |

## Full example with comments

[fullexample.config.yaml](fullexample.config.yaml)