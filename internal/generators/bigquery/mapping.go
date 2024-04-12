package bigquery

import _ "embed"

const (
	filenameDatasettf  = "dataset.tf"
	filenameBigQuerytf = "bigquery.tf"
)

var (
	//go:embed tmpls/bigquery.tf.tmpl
	tmplBigQuerytf []byte

	//go:embed tmpls/dataset.tf.tmpl
	tmplDatasettf []byte
)

var defaultTfTemplateFiles = map[string]string{
	filenameBigQuerytf: string(tmplBigQuerytf),
	filenameDatasettf:  string(tmplDatasettf),
}
