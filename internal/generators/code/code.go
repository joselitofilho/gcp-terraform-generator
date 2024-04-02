package code

import (
	"fmt"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/iotcore"
)

type Code struct {
	configFileName string
	output         string
}

func NewCode(configFileName, output string) *Code {
	return &Code{configFileName: configFileName, output: output}
}

func (c *Code) Build() error {
	err := iotcore.NewIoTCore(c.configFileName, c.output).Build()
	if err != nil {
		return fmt.Errorf("fails to build IoTCore: %w", err)
	}

	return nil
}
