package storage

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

type Data struct {
	Name     string
	Location string
}

type Storage struct {
	configFileName string
	output         string
}

func NewStorage(configFileName, output string) *Storage {
	return &Storage{configFileName: configFileName, output: output}
}

func (ps *Storage) Build() error {
	yamlParser := config.NewYAML(ps.configFileName)

	yamlConfig, err := yamlParser.Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", generatorserrs.ErrYAMLParse, err)
	}

	modPath := path.Join(ps.output, "mod")
	_ = os.MkdirAll(modPath, os.ModePerm)

	result := make([]string, 0, len(yamlConfig.IoTCores))

	templates := utils.MergeStringMap(defaultTfTemplateFiles,
		generators.CreateTemplatesMap(yamlConfig.OverrideDefaultTemplates.Storage))

	for _, conf := range yamlConfig.Storages {
		data := Data{
			Name:     conf.Name,
			Location: conf.Location,
		}

		if len(conf.Files) > 0 {
			filesConf := generators.CreateFilesMap(conf.Files)

			err = generators.GenerateFiles(nil, filesConf, data, modPath)
			if err != nil {
				return fmt.Errorf("%w", err)
			}

			fmtcolor.White.Printf("Storage '%s' has been generated successfully\n", conf.Name)

			continue
		}

		output, err := generators.Build(data, "storage-tf-template", templates[filenameStoragetf])
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		result = append(result, output)
	}

	if len(result) > 0 {
		outputFile := path.Join(modPath, filenameStoragetf)

		err := generators.GenerateFile(nil, filenameStoragetf, strings.Join(result, "\n"), outputFile, Data{})
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		fmtcolor.White.Println("Storage has been generated successfully")
	}

	return nil
}
