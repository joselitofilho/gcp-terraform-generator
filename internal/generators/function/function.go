package function

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

type Function struct {
	configFileName string
	output         string
}

func NewFunction(configFileName, output string) *Function {
	return &Function{configFileName: configFileName, output: output}
}

func (ps *Function) Build() error {
	yamlParser := config.NewYAML(ps.configFileName)

	yamlConfig, err := yamlParser.Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", generatorserrs.ErrYAMLParse, err)
	}

	modPath := path.Join(ps.output, "mod")
	_ = os.MkdirAll(modPath, os.ModePerm)

	result := make([]string, 0, len(yamlConfig.Functions))

	templates := utils.MergeStringMap(defaultTfTemplateFiles,
		generators.CreateTemplatesMap(yamlConfig.OverrideDefaultTemplates.Function))

	tg := generators.NewGenerator()

	for _, conf := range yamlConfig.Functions {
		data := Data{
			Name:                conf.Name,
			Source:              conf.Source,
			Runtime:             conf.Runtime,
			SourceArchiveBucket: conf.SourceArchiveBucket,
			SourceArchiveObject: conf.SourceArchiveObject,
			TriggerHTTP:         conf.TriggerHTTP,
			EntryPoint:          conf.EntryPoint,
			Envars:              conf.Envars,
		}

		if len(conf.Files) > 0 {
			filesConf := generators.CreateFilesMap(conf.Files)

			generators.MustGenerateFiles(tg, nil, filesConf, data, modPath)

			fmtcolor.White.Printf("Function '%s' has been generated successfully\n", conf.Name)

			continue
		}

		output, err := tg.Build(data, "appengine-tf-template", templates[filenameFunctiontf])
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		result = append(result, output)
	}

	if len(result) > 0 {
		outputFile := path.Join(modPath, filenameFunctiontf)

		generators.MustGenerateFile(tg, nil, filenameFunctiontf, strings.Join(result, "\n"), outputFile, Data{})

		fmtcolor.White.Println("Function has been generated successfully")
	}

	return nil
}
