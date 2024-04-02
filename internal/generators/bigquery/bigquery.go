package bigquery

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

type BigQuery struct {
	configFileName string
	output         string
}

func NewBigQuery(configFileName, output string) *BigQuery {
	return &BigQuery{configFileName: configFileName, output: output}
}

func (ps *BigQuery) Build() error {
	yamlParser := config.NewYAML(ps.configFileName)

	yamlConfig, err := yamlParser.Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", generatorserrs.ErrYAMLParse, err)
	}

	modPath := path.Join(ps.output, "mod")
	_ = os.MkdirAll(modPath, os.ModePerm)

	result := make([]string, 0, len(yamlConfig.IoTCores))

	templates := utils.MergeStringMap(defaultTfTemplateFiles,
		generators.CreateTemplatesMap(yamlConfig.OverrideDefaultTemplates.BigQuery))

	for _, conf := range yamlConfig.BigQueryTables {
		nameParts := strings.Split(conf.Name, ".")

		dataset := "default"
		if len(nameParts) > 1 {
			dataset = nameParts[0]
		}

		table := nameParts[len(nameParts)-1]

		schema := conf.Schema
		if schema == "" {
			schema = `<<EOF
# Define your BigQuery schema here
EOF`
		}

		data := Data{
			Dataset: dataset,
			Table:   table,
			Schema:  schema,
		}

		if len(conf.Files) > 0 {
			filesConf := generators.CreateFilesMap(conf.Files)

			err = generators.GenerateFiles(nil, filesConf, data, modPath)
			if err != nil {
				return fmt.Errorf("%w", err)
			}

			fmtcolor.White.Printf("Big Query '%s' has been generated successfully\n", conf.Name)

			continue
		}

		output, err := generators.Build(data, "bigquery-tf-template", templates[filenameBigQuerytf])
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		result = append(result, output)
	}

	if len(result) > 0 {
		outputFile := path.Join(modPath, filenameBigQuerytf)

		err := generators.GenerateFile(nil, filenameBigQuerytf, strings.Join(result, "\n"), outputFile, Data{})
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		fmtcolor.White.Println("Big Query has been generated successfully")
	}

	return nil
}
