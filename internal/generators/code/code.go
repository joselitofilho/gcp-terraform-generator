package code

import (
	"fmt"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/appengine"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/bigquery"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/bigtable"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/dataflow"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/iotcore"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/pubsub"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/storage"
)

type Code struct {
	configFileName string
	output         string
}

func NewCode(configFileName, output string) *Code {
	return &Code{configFileName: configFileName, output: output}
}

func (c *Code) Build() error {
	if err := appengine.NewAppEngine(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build App Engine: %w", err)
	}

	if err := bigquery.NewBigQuery(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build Big Query: %w", err)
	}

	if err := bigtable.NewBigTable(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build Big Table: %w", err)
	}

	if err := dataflow.NewDataFlow(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build DataFlow: %w", err)
	}

	if err := iotcore.NewIoTCore(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build IoT Core: %w", err)
	}

	if err := pubsub.NewPubSub(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build Pub Sub: %w", err)
	}

	if err := storage.NewStorage(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build Storage: %w", err)
	}

	return nil
}
