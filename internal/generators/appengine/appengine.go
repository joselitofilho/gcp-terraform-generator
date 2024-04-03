package appengine

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

type AppEngine struct {
	configFileName string
	output         string
}

func NewAppEngine(configFileName, output string) *AppEngine {
	return &AppEngine{configFileName: configFileName, output: output}
}

func (ps *AppEngine) Build() error {
	yamlParser := config.NewYAML(ps.configFileName)

	yamlConfig, err := yamlParser.Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", generatorserrs.ErrYAMLParse, err)
	}

	modPath := path.Join(ps.output, "mod")
	_ = os.MkdirAll(modPath, os.ModePerm)

	result := make([]string, 0, len(yamlConfig.IoTCores))

	templates := utils.MergeStringMap(defaultTfTemplateFiles,
		generators.CreateTemplatesMap(yamlConfig.OverrideDefaultTemplates.AppEngine))

	for _, conf := range yamlConfig.AppEngines {
		data := Data{
			Name:       conf.Name,
			LocationID: conf.LocationID,
		}

		if len(conf.Files) > 0 {
			filesConf := generators.CreateFilesMap(conf.Files)

			err = generators.GenerateFiles(nil, filesConf, data, modPath)
			if err != nil {
				return fmt.Errorf("%w", err)
			}

			fmtcolor.White.Printf("App Engine '%s' has been generated successfully\n", conf.Name)

			continue
		}

		output, err := generators.Build(data, "appengine-tf-template", templates[filenameAppEnginetf])
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		result = append(result, output)
	}

	if len(result) > 0 {
		outputFile := path.Join(modPath, filenameAppEnginetf)

		err := generators.GenerateFile(nil, filenameAppEnginetf, strings.Join(result, "\n"), outputFile, Data{})
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		fmtcolor.White.Println("App Engine has been generated successfully")
	}

	return nil
}
