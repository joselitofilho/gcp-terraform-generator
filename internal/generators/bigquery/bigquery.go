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
		return fmt.Errorf("%w: %w", generatorserrs.ErrYAMLParser, err)
	}

	modPath := path.Join(ps.output, "mod")
	_ = os.MkdirAll(modPath, os.ModePerm)

	resultBigQuery := make([]string, 0, len(yamlConfig.BigQueryTables))
	resultDataset := make([]string, 0, len(yamlConfig.BigQueryTables))

	templates := utils.MergeStringMap(defaultTfTemplateFiles,
		generators.CreateTemplatesMap(yamlConfig.OverrideDefaultTemplates.BigQuery))

	tg := generators.NewGenerator()

	datasetMap := map[string]struct{}{}

	for _, conf := range yamlConfig.BigQueryTables {
		nameParts := strings.Split(conf.Name, ".")

		dataset := "default"
		if len(nameParts) > 1 {
			dataset = nameParts[0]
		}

		table := nameParts[len(nameParts)-1]

		dataBigQuery := Data{
			Dataset: dataset,
			Table:   table,
			Schema:  conf.Schema,
		}

		if len(conf.Files) > 0 {
			filesConf := generators.CreateFilesMap(conf.Files)

			generators.MustGenerateFiles(tg, nil, filesConf, dataBigQuery, modPath)

			fmtcolor.White.Printf("Big Query '%s' has been generated successfully\n", conf.Name)

			continue
		}

		outputBigQuery, err := tg.Build(dataBigQuery, "bigquery-tf-template", templates[filenameBigQuerytf])
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		resultBigQuery = append(resultBigQuery, outputBigQuery)

		if _, ok := datasetMap[dataset]; !ok {
			datasetMap[dataset] = struct{}{}

			outputDataset, err := tg.Build(Data{Dataset: dataset}, "dataset-tf-template", templates[filenameDatasettf])
			if err != nil {
				return fmt.Errorf("%w", err)
			}

			resultDataset = append(resultDataset, outputDataset)
		}
	}

	if len(resultBigQuery) > 0 {
		bigQueryOutputFile := path.Join(modPath, filenameBigQuerytf)

		generators.MustGenerateFile(
			tg, nil, filenameBigQuerytf, strings.Join(resultBigQuery, "\n"), bigQueryOutputFile, Data{})

		fmtcolor.White.Println("Big Query has been generated successfully")

		datasetOutputFile := path.Join(modPath, filenameDatasettf)

		generators.MustGenerateFile(
			tg, nil, filenameDatasettf, strings.Join(resultDataset, "\n"), datasetOutputFile, Data{})

		fmtcolor.White.Println("Big Query Dataset has been generated successfully")
	}

	return nil
}
