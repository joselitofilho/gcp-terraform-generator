package function

import _ "embed"

const filenameFunctiontf = "function.tf"

//go:embed tmpls/function.tf.tmpl
var tmplFunctiontf []byte

var defaultTfTemplateFiles = map[string]string{
	filenameFunctiontf: string(tmplFunctiontf),
}
