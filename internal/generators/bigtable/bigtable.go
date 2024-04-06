package bigtable

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

type BigTable struct {
	configFileName string
	output         string
}

func NewBigTable(configFileName, output string) *BigTable {
	return &BigTable{configFileName: configFileName, output: output}
}

func (ps *BigTable) Build() error {
	yamlParser := config.NewYAML(ps.configFileName)

	yamlConfig, err := yamlParser.Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", generatorserrs.ErrYAMLParser, err)
	}

	modPath := path.Join(ps.output, "mod")
	_ = os.MkdirAll(modPath, os.ModePerm)

	result := make([]string, 0, len(yamlConfig.BigTables))

	templates := utils.MergeStringMap(defaultTfTemplateFiles,
		generators.CreateTemplatesMap(yamlConfig.OverrideDefaultTemplates.BigTable))

	tg := generators.NewGenerator()

	for _, conf := range yamlConfig.BigTables {
		data := Data{
			Name:   conf.Name,
			Labels: conf.Labels,
		}

		if len(conf.Files) > 0 {
			filesConf := generators.CreateFilesMap(conf.Files)

			generators.MustGenerateFiles(tg, nil, filesConf, data, modPath)

			fmtcolor.White.Printf("Big Table '%s' has been generated successfully\n", conf.Name)

			continue
		}

		output, err := tg.Build(data, "bigtable-tf-template", templates[filenameBigTabletf])
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		result = append(result, output)
	}

	if len(result) > 0 {
		outputFile := path.Join(modPath, filenameBigTabletf)

		generators.MustGenerateFile(tg, nil, filenameBigTabletf, strings.Join(result, "\n"), outputFile, Data{})

		fmtcolor.White.Println("Big Table has been generated successfully")
	}

	return nil
}
