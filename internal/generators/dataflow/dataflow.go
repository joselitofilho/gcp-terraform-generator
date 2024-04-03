package dataflow

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/joselitofilho/gcp-terraform-generator/internal/fmtcolor"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	generatorserrs "github.com/joselitofilho/gcp-terraform-generator/internal/generators/errors"
	"github.com/joselitofilho/gcp-terraform-generator/internal/utils"
)

type DataFlow struct {
	configFileName string
	output         string
}

func NewDataFlow(configFileName, output string) *DataFlow {
	return &DataFlow{configFileName: configFileName, output: output}
}

func (c *DataFlow) Build() error {
	yamlParser := config.NewYAML(c.configFileName)

	yamlConfig, err := yamlParser.Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", generatorserrs.ErrYAMLParse, err)
	}

	modPath := path.Join(c.output, "mod")
	_ = os.MkdirAll(modPath, os.ModePerm)

	result := make([]string, 0, len(yamlConfig.DataFlows))

	templates := utils.MergeStringMap(defaultTfTemplateFiles,
		generators.CreateTemplatesMap(yamlConfig.OverrideDefaultTemplates.DataFlow))

	for _, conf := range yamlConfig.DataFlows {
		data := Data{
			Name:              conf.Name,
			InputTopics:       conf.InputTopics,
			OutputTopics:      conf.OutputTopics,
			OutputDirectories: conf.OutputDirectories,
			OutputTables:      conf.OutputTables,
		}

		if len(conf.Files) > 0 {
			filesConf := generators.CreateFilesMap(conf.Files)

			err = generators.GenerateFiles(nil, filesConf, data, modPath)
			if err != nil {
				return fmt.Errorf("%w", err)
			}

			fmtcolor.White.Printf("DataFlow '%s' has been generated successfully\n", conf.Name)

			continue
		}

		output, err := generators.Build(data, "dataflow-tf-template", templates[filenameDataFlowtf])
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		result = append(result, output)
	}

	if len(result) > 0 {
		outputFile := path.Join(modPath, filenameDataFlowtf)

		err := generators.GenerateFile(nil, filenameDataFlowtf, strings.Join(result, "\n"), outputFile, Data{})
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		fmtcolor.White.Println("DataFlow has been generated successfully")
	}

	return nil
}
