package appengine

import _ "embed"

const filenameAppEnginetf = "appengine.tf"

//go:embed tmpls/appengine.tf.tmpl
var tmplAppEnginetf []byte

var defaultTfTemplateFiles = map[string]string{
	filenameAppEnginetf: string(tmplAppEnginetf),
}
