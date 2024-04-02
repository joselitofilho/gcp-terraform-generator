package iotcore

import _ "embed"

const filenameIoTCoretf = "iotcore.tf"

//go:embed tmpls/iotcore.tf.tmpl
var tmplIoTCoretf []byte

var defaultTfTemplateFiles = map[string]string{
	filenameIoTCoretf: string(tmplIoTCoretf),
}
