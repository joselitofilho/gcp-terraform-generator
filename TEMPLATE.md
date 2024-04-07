# Templates

## Variables

The following variables can be used within the templates:

### App Engine

| Name           | Description                                                 |
| :------------- | :---------------------------------------------------------- |
| Name           | Name of the app engine.                                 |
| LocationID     | The location to serve the app from.                         |

Default templates:

```
📦 appengine
 ┣ 📂 tmpls
 ┗ ┗ 📜 appengine.tf.tmpl
```
- [📜 appengine.tf.tmpl](./internal/generators/appengine/tmpls/appengine.tf.tmpl)

### Big Query

| Name           | Description                                                 |
| :------------- | :---------------------------------------------------------- |
| Dataset        | Big Query table dataset name.                               |
| Table          | Name of the Big Query table.                                |
| Schema         | A JSON schema for the table.                                |

Default temaplates:

```
📦 bigquery
 ┣ 📂 tmpls
 ┗ ┗ 📜 bigquery.tf.tmpl
```
- [📜 bigquery.tf.tmpl](./internal/generators/bigquery/tmpls/bigquery.tf.tmpl)

### Big Table

| Name           | Description                                                 |
| :------------- | :---------------------------------------------------------- |
| Name           | Name of the Big Table.                                      |
| Labels         | A map of key/value label pairs to assign to the resource.   |

Default temaplates:

```
📦 bigtable
 ┣ 📂 tmpls
 ┗ ┗ 📜 bigtable.tf.tmpl
```
- [📜 bigtable.tf.tmpl](./internal/generators/bigtable/tmpls/bigtable.tf.tmpl)

### DataFlow

| Name              | Description                                              |
| :---------------- | :------------------------------------------------------- |
| Name              | Name of the DataFlow job.                                |
| InputTopics       | Input topics for the DataFlow.                           |
| OutputTopics      | Output topics for the DataFlow.                          |
| OutputDirectories | Output directories for the DataFlow.                     |
| OutputTables      | Output tables for the DataFlow.                          |

Default temaplates:

```
📦 dataflow
 ┣ 📂 tmpls
 ┗ ┗ 📜 dataflow.tf.tmpl
```
- [📜 dataflow.tf.tmpl](./internal/generators/dataflow/tmpls/dataflow.tf.tmpl)

### Function

| Name                | Description                                            |
| :------------------ | :----------------------------------------------------- |
| Name                | Name of the Function.                                  |
| Source              | Source of the function code.                           |
| Runtime             | The runtime in which the function is going to run.     |
| SourceArchiveBucket | The GCS bucket containing the zip archive which contains the function. |
| SourceArchiveObject | The source archive object (file) in archive bucket.    |
| TriggerHTTP         | Any HTTP request (of a supported type) to the endpoint will trigger function execution. |
| EntryPoint          | Name of the function that will be executed when the Google Cloud Function is triggered. |
| Envars              | A map of key/value environment variable pairs to assign to the function. |

Default temaplates:

```
📦 function
 ┣ 📂 tmpls
 ┗ ┗ 📜 function.tf.tmpl
```
- [📜 function.tf.tmpl](./internal/generators/function/tmpls/function.tf.tmpl)

### IoT Core

| Name                     | Description                                       |
| :----------------------- | :------------------------------------------------ |
| Name                     | Name of the IoT Core.                             |
| EventNotificationConfigs | List of configurations for event notifications, such as PubSub topics to publish device events to. |
| ┗ TopicName              | Pub Sub topic name to publish device state updates. |

Default temaplates:

```
📦 iotcore
 ┣ 📂 tmpls
 ┗ ┗ 📜 iotcore.tf.tmpl
```
- [📜 iotcore.tf.tmpl](./internal/generators/iotcore/tmpls/iotcore.tf.tmpl)

### Pub Sub

| Name           | Description                                                 |
| :------------- | :---------------------------------------------------------- |
| Name           | Name of the Pub Sub topic.                                  |
| Labels         | A map of key/value label pairs to assign to the resource.   |
| PushEnpoint    | URL locating the endpoint to which messages should be pushed. |

Default temaplates:

```
📦 pubsub
 ┣ 📂 tmpls
 ┗ ┗ 📜 pubsub.tf.tmpl
```
- [📜 pubsub.tf.tmpl](./internal/generators/pubsub/tmpls/pubsub.tf.tmpl)

### Storage

| Name           | Description                                                 |
| :------------- | :---------------------------------------------------------- |
| Name           | Name of the Storage topic.                                  |
| Location       | The GCS location.                                           |

Default temaplates:

```
📦 storage
 ┣ 📂 tmpls
 ┗ ┗ 📜 storage.tf.tmpl
```
- [📜 storage.tf.tmpl](./internal/generators/storage/tmpls/storage.tf.tmpl)

### Structure

| Name           | Description                                                 |
| :------------- | :---------------------------------------------------------- |
| StackName      | The name of the stack associated with the project structure. |

## Custom Functions

The following custom functions are available:

| Name           | Description                                                 |
| :------------- | :---------------------------------------------------------- |
| getFileByName  | Retrieves a file from a map of files by its name.           |
| getFileImports | Retrieves the imports of a file by its name.                |
| ToCamel        | Converts a string to CamelCase format.                      |
| ToKebab        | Converts a string to kebab-case format.                     |
| ToLower        | Converts a string to lowercase.                             |
| ToPascal       | Converts a string to PascalCase format.                     |
| ToSpace        | Converts a string to kebab-case and replaces hyphens with spaces. |
| ToSnake        | Converts a string to snake_case format.                     |
| ToUpper        | Converts a string to uppercase.                             |
