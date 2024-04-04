package code

import (
	"fmt"

	"github.com/joselitofilho/gcp-terraform-generator/internal/fmtcolor"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/appengine"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/bigquery"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/bigtable"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/dataflow"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/function"
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
	fmtcolor.White.Println("→ Generating App Engine code...")
	if err := appengine.NewAppEngine(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build App Engine: %w", err)
	}

	fmtcolor.White.Println("→ Generating Big Query code...")
	if err := bigquery.NewBigQuery(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build Big Query: %w", err)
	}

	fmtcolor.White.Println("→ Generating Big Table code...")
	if err := bigtable.NewBigTable(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build Big Table: %w", err)
	}

	fmtcolor.White.Println("→ Generating DataFlow code...")
	if err := dataflow.NewDataFlow(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build DataFlow: %w", err)
	}

	fmtcolor.White.Println("→ Generating Function code...")
	if err := function.NewFunction(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build Function: %w", err)
	}

	fmtcolor.White.Println("→ Generating IoT Core code...")
	if err := iotcore.NewIoTCore(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build IoT Core: %w", err)
	}

	fmtcolor.White.Println("→ Generating Pub Sub code...")
	if err := pubsub.NewPubSub(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build Pub Sub: %w", err)
	}

	fmtcolor.White.Println("→ Generating Storage code...")
	if err := storage.NewStorage(c.configFileName, c.output).Build(); err != nil {
		return fmt.Errorf("fails to build Storage: %w", err)
	}

	return nil
}
