package config

type PubSub struct {
	Name  string `yaml:"name"`
	Files []File `yaml:"files,omitempty"`
}

func (r *PubSub) GetName() string { return r.Name }
