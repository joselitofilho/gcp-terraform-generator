package pubsub

import _ "embed"

const filenamePubSubtf = "pubsub.tf"

//go:embed tmpls/pubsub.tf.tmpl
var tmplPubSubtf []byte

var defaultTfTemplateFiles = map[string]string{
	filenamePubSubtf: string(tmplPubSubtf),
}
