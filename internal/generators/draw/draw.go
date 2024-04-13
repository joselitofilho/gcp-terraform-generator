package draw

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"

	"github.com/diagram-code-generator/resources/pkg/parser/graphviz"
	hcl "github.com/joselitofilho/hcl-parser-go/pkg/parser/hcl"

	"github.com/joselitofilho/gcp-terraform-generator/internal/fmtcolor"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	generatorerrs "github.com/joselitofilho/gcp-terraform-generator/internal/generators/errors"
	gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"
	"github.com/joselitofilho/gcp-terraform-generator/internal/transformers/resourcestoyaml"
	"github.com/joselitofilho/gcp-terraform-generator/internal/transformers/terraformtoresources"
)

// DefaultResourceImageMap defines the default resource images. Images from here:
// https://github.com/mingrammer/diagrams/tree/master/resources/gcp
var DefaultResourceImageMap = config.Images{
	gcpresources.AppEngine: "assets/diagram/app_engine.svg",
	gcpresources.BigQuery:  "assets/diagram/bigquery.svg",
	gcpresources.BigTable:  "assets/diagram/big_table.svg",
	gcpresources.Dataflow:  "assets/diagram/dataflow.svg",
	gcpresources.Function:  "assets/diagram/function.svg",
	gcpresources.IoTCore:   "assets/diagram/iot_core.svg",
	gcpresources.PubSub:    "assets/diagram/pub_sub.svg",
	gcpresources.Storage:   "assets/diagram/storage.svg",
}

type Draw struct {
	workdirs       []string
	files          []string
	configFilename string
	output         string
}

func NewDraw(workdirs, files []string, configFilename, output string) *Draw {
	return &Draw{workdirs: workdirs, files: files, configFilename: configFilename, output: output}
}

func (d *Draw) Build() error {
	yamlParser := config.NewYAML(d.configFilename)

	yamlConfig, err := yamlParser.Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", generatorerrs.ErrYAMLParser, err)
	}

	tfConfig, err := hcl.Parse(d.workdirs, d.files)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	resc := terraformtoresources.NewTransformer(yamlConfig, tfConfig).Transform()

	_ = os.Mkdir(d.output, os.ModePerm)

	diagramConfig, err := resourcestoyaml.NewTransformer(yamlConfig, resc).Transform()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	yamlData, err := yaml.Marshal(diagramConfig)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	yamlFilename := "diagram"
	if yamlConfig.Draw.Name != "" {
		yamlFilename = yamlConfig.Draw.Name
	}

	yamlFilename += ".yaml"

	yamlfile, err := os.Create(path.Join(d.output, yamlFilename))
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer yamlfile.Close()

	if _, err = yamlfile.WriteString(string(yamlData)); err != nil {
		return fmt.Errorf("%w", err)
	}

	fmtcolor.White.Println("The diagram yaml file has been generated successfully.")

	nodeAttrs := make(map[string]any)
	for k, v := range graphviz.DefaultNodeAttrs {
		nodeAttrs[k] = v
	}
	delete(nodeAttrs, "height")

	dotConfig := graphviz.Config{Orientation: yamlConfig.Draw.Orientation, NodeAttrs: nodeAttrs}

	resourceImageMap := mergeImages(DefaultResourceImageMap, yamlConfig.Draw.Images)

	dotContent := graphviz.Build(resc, resourceImageMap.ToStringMap(), dotConfig)

	dotFilename := "diagram"
	if yamlConfig.Draw.Name != "" {
		dotFilename = yamlConfig.Draw.Name
	}

	dotFilename += ".dot"

	dotfile, err := os.Create(path.Join(d.output, dotFilename))
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer dotfile.Close()

	if _, err := dotfile.WriteString(dotContent); err != nil {
		return fmt.Errorf("%w", err)
	}

	fmtcolor.White.Println("The graphviz dot file has been generated successfully.")

	return nil
}

func mergeImages(defaultImages, configImages config.Images) config.Images {
	result := defaultImages

	for k, v := range configImages {
		result[k] = v
	}

	return result
}
