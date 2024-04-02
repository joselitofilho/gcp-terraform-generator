package bigquery

import _ "embed"

const filenameBigQuerytf = "bigquery.tf"

//go:embed tmpls/bigquery.tf.tmpl
var tmplBigQuerytf []byte

var defaultTfTemplateFiles = map[string]string{
	filenameBigQuerytf: string(tmplBigQuerytf),
}
