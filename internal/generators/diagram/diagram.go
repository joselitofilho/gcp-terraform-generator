package diagram

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	drawioxml "github.com/joselitofilho/drawio-parser-go/pkg/parser/xml"

	"github.com/diagram-code-generator/resources/pkg/transformers/drawiotoresources"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	generatorserrs "github.com/joselitofilho/gcp-terraform-generator/internal/generators/errors"
	"github.com/joselitofilho/gcp-terraform-generator/internal/resources"
	"github.com/joselitofilho/gcp-terraform-generator/internal/transformers/resourcestoyaml"
)

type Diagram struct {
	diagramFilename string
	configFilename  string
	output          string
}

func NewDiagram(diagramFilename, configFilename, output string) *Diagram {
	return &Diagram{diagramFilename: diagramFilename, configFilename: configFilename, output: output}
}

func (d *Diagram) Build() error {
	yamlConfig, err := config.NewYAML(d.configFilename).Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", generatorserrs.ErrYAMLParser, err)
	}

	mxFile, err := drawioxml.Parse(d.diagramFilename)
	if err != nil {
		return fmt.Errorf("%w: %w", generatorserrs.ErrDrawIOParser, err)
	}

	rscs, _ := drawiotoresources.NewTransformer(mxFile, &resources.GCPResourceFactory{}).Transform()

	yamlConfigOut, err := resourcestoyaml.NewTransformer(yamlConfig, rscs).Transform()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	data, err := yaml.Marshal(yamlConfigOut)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	outputDir, _ := filepath.Split(d.output)
	_ = os.MkdirAll(filepath.Base(outputDir), os.ModePerm)

	err = os.WriteFile(d.output, data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
