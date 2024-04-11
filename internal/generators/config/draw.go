package config

import gcpresources "github.com/joselitofilho/gcp-terraform-generator/internal/resources"

type Images map[gcpresources.ResourceType]string

func (m Images) ToStringMap() map[string]string {
	result := map[string]string{}

	for k, v := range m {
		result[k.String()] = v
	}

	return result
}

type ReplaceableTexts map[string]string

type Draw struct {
	Name             string           `yaml:"name,omitempty"`
	Orientation      string           `yaml:"orientation,omitempty"`
	ReplaceableTexts ReplaceableTexts `yaml:"replaceable_texts,omitempty"`
	Images           Images           `yaml:"images,omitempty"`
	Filters          Filters          `yaml:"filters,omitempty"`
}
