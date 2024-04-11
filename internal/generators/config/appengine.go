package config

type AppEngine struct {
	Name       string `yaml:"name"`
	Project    string `yaml:"project"`
	LocationID string `yaml:"location_id"`
	Files      []File `yaml:"files,omitempty"`
}

func (r *AppEngine) GetName() string { return r.Name }
