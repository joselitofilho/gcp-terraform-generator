package structure

import (
	_ "embed"
	"fmt"
	"os"
	"path"

	"github.com/joselitofilho/gcp-terraform-generator/internal/fmtcolor"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	generatorserrs "github.com/joselitofilho/gcp-terraform-generator/internal/generators/errors"
)

type Structure struct {
	configFileName string
	output         string
}

func NewStructure(configFileName, output string) *Structure {
	return &Structure{configFileName: configFileName, output: output}
}

func (s *Structure) Build() error {
	yamlParser := config.NewYAML(s.configFileName)

	yamlConfig, err := yamlParser.Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", generatorserrs.ErrYAMLParser, err)
	}

	defaultTemplatesMap := generators.CreateTemplatesMap(yamlConfig.Structure.DefaultTemplates)

	tg := generators.NewGenerator()

	for i := range yamlConfig.Structure.Stacks {
		conf := yamlConfig.Structure.Stacks[i]

		data := Data{
			StackName: conf.Name,
		}

		for _, folder := range conf.Folders {
			output := path.Join(s.output, conf.Name, folder.Name)
			_ = os.MkdirAll(output, os.ModePerm)

			for _, file := range folder.Files {
				outputFile := path.Join(output, file.Name)

				generators.MustGenerateFile(tg, defaultTemplatesMap, file.Name, file.Tmpl, outputFile, data)
			}
		}

		for _, file := range conf.Files {
			outputFile := path.Join(s.output, conf.Name, file.Name)

			generators.MustGenerateFile(tg, defaultTemplatesMap, file.Name, file.Tmpl, outputFile, data)
		}

		fmtcolor.White.Printf("Structure '%s' has been generated successfully\n", conf.Name)
	}

	return nil
}
