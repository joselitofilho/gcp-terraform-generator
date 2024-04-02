package iotcore

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

type IoTCore struct {
	configFileName string
	output         string
}

func NewIoTCore(configFileName, output string) *IoTCore {
	return &IoTCore{configFileName: configFileName, output: output}
}

func (c *IoTCore) Build() error {
	yamlParser := config.NewYAML(c.configFileName)

	yamlConfig, err := yamlParser.Parse()
	if err != nil {
		return fmt.Errorf("%w: %w", generatorserrs.ErrYAMLParse, err)
	}

	modPath := path.Join(c.output, "mod")
	_ = os.MkdirAll(modPath, os.ModePerm)

	result := make([]string, 0, len(yamlConfig.IoTCores))

	templates := utils.MergeStringMap(defaultTfTemplateFiles,
		generators.CreateTemplatesMap(yamlConfig.OverrideDefaultTemplates.IoTCore))

	for _, conf := range yamlConfig.IoTCores {
		eventNotificationConfigs := buildEventNotificationConfigs(conf)

		data := Data{
			Name:                     conf.Name,
			EventNotificationConfigs: eventNotificationConfigs,
		}

		if len(conf.Files) > 0 {
			filesConf := generators.CreateFilesMap(conf.Files)

			err = generators.GenerateFiles(nil, filesConf, data, modPath)
			if err != nil {
				return fmt.Errorf("%w", err)
			}

			fmtcolor.White.Printf("IoT Core '%s' has been generated successfully\n", conf.Name)

			continue
		}

		output, err := generators.Build(data, "iotcore-tf-template", templates[filenameIoTCoretf])
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		result = append(result, output)
	}

	if len(result) > 0 {
		outputFile := path.Join(modPath, filenameIoTCoretf)

		err := generators.GenerateFile(nil, filenameIoTCoretf, strings.Join(result, "\n"), outputFile, Data{})
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		fmtcolor.White.Println("IoT Core has been generated successfully")
	}

	return nil
}

func buildEventNotificationConfigs(conf *config.IoTCore) []EventNotificationConfig {
	eventNotificationConfigs := make([]EventNotificationConfig, 0, len(conf.EventNotificationConfigs))
	for i := range conf.EventNotificationConfigs {
		eventNotificationConfigs = append(eventNotificationConfigs, EventNotificationConfig{
			TopicName: conf.EventNotificationConfigs[i].TopicName,
		})
	}

	return eventNotificationConfigs
}
