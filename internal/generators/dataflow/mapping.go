package dataflow

import _ "embed"

const filenameDataFlowtf = "dataflow.tf"

//go:embed tmpls/dataflow.tf.tmpl
var tmplDataFlowtf []byte

var defaultTfTemplateFiles = map[string]string{
	filenameDataFlowtf: string(tmplDataFlowtf),
}
