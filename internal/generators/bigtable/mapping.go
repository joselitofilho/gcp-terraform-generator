package bigtable

import _ "embed"

const filenameBigTabletf = "bigtable.tf"

//go:embed tmpls/bigtable.tf.tmpl
var tmplBigTabletf []byte

var defaultTfTemplateFiles = map[string]string{
	filenameBigTabletf: string(tmplBigTabletf),
}
