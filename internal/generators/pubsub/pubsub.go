package pubsub

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

type PubSub struct {
	configFileName string
	output         string
}

func NewPubSub(configFileName, output string) *PubSub {
	return &PubSub{configFileName: configFileName, output: output}
}

func (ps *PubSub) Build() error {
	yamlParser := config.NewYAML(ps.configFileName)

	yamlConfig, err := yamlParser.Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", generatorserrs.ErrYAMLParser, err)
	}

	modPath := path.Join(ps.output, "mod")
	_ = os.MkdirAll(modPath, os.ModePerm)

	result := make([]string, 0, len(yamlConfig.PubSubs))

	templates := utils.MergeStringMap(defaultTfTemplateFiles,
		generators.CreateTemplatesMap(yamlConfig.OverrideDefaultTemplates.PubSub))

	tg := generators.NewGenerator()

	for _, conf := range yamlConfig.PubSubs {
		data := Data{
			Name:        conf.Name,
			Labels:      conf.Labels,
			PushEnpoint: conf.PushEnpoint,
		}

		if len(conf.Files) > 0 {
			filesConf := generators.CreateFilesMap(conf.Files)

			generators.MustGenerateFiles(tg, nil, filesConf, data, modPath)

			fmtcolor.White.Printf("Pub Sub '%s' has been generated successfully\n", conf.Name)

			continue
		}

		output, err := tg.Build(data, "pubsub-tf-template", templates[filenamePubSubtf])
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		result = append(result, output)
	}

	if len(result) > 0 {
		outputFile := path.Join(modPath, filenamePubSubtf)

		generators.MustGenerateFile(tg, nil, filenamePubSubtf, strings.Join(result, "\n"), outputFile, Data{})

		fmtcolor.White.Println("Pub Sub has been generated successfully")
	}

	return nil
}
