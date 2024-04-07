package config

type AppEngine struct {
	Name       string `yaml:"name"`
	LocationID string `yaml:"location_id"`
	Files      []File `yaml:"files,omitempty"`
}

func (r *AppEngine) GetName() string { return r.Name }
