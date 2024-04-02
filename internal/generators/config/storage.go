package config

type Storage struct {
	Name     string  `yaml:"name"`
	Location string  `yaml:"locaiton"`
	Files    []*File `yaml:"files,omitempty"`
}

func (r *Storage) GetName() string { return r.Name }
