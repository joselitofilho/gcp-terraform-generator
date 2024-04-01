package diagram

import (
	"fmt"

	drawioxml "github.com/joselitofilho/drawio-parser-go/pkg/parser/xml"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/transformers/drawiotoresources"
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
		return fmt.Errorf("%w", err)
	}

	fmt.Println(yamlConfig)

	mxFile, err := drawioxml.Parse(d.diagramFilename)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	resources, err := drawiotoresources.Transform(mxFile)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	for _, r := range resources.Resources {
		fmt.Printf("%+v\n", r)
	}
	for _, r := range resources.Relationships {
		fmt.Printf("%+v -> %+v\n", r.Source, r.Target)
	}

	// yamlConfigOut, err := resourcestoyaml.NewTransformer(yamlConfig, resources).Transform()
	// if err != nil {
	// 	return fmt.Errorf("%w", err)
	// }

	// data, err := yaml.Marshal(yamlConfigOut)
	// if err != nil {
	// 	return fmt.Errorf("%w", err)
	// }

	// err = os.WriteFile(output, data, os.ModePerm)
	// if err != nil {
	// 	return fmt.Errorf("%w", err)
	// }

	return nil
}
