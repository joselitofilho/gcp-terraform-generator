package code

import (
	"fmt"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/iotcore"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/pubsub"
)

type Code struct {
	configFileName string
	output         string
}

func NewCode(configFileName, output string) *Code {
	return &Code{configFileName: configFileName, output: output}
}

func (c *Code) Build() error {
	if err := iotcore.NewIoTCore(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build IoT Core: %w", err)
	}

	if err := pubsub.NewPubSub(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build Pub Sub: %w", err)
	}

	return nil
}
