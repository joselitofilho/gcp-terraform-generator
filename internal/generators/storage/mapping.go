package storage

import _ "embed"

const filenameStoragetf = "storage.tf"

//go:embed tmpls/storage.tf.tmpl
var tmplStoragetf []byte

var defaultTfTemplateFiles = map[string]string{
	filenameStoragetf: string(tmplStoragetf),
}
