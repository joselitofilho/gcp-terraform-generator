package config

type PubSub struct {
	Name        string            `yaml:"name"`
	Labels      map[string]string `yaml:"labels,omitempty"`
	PushEnpoint string            `yaml:"push_endpoint,omitempty"`
	Files       []File            `yaml:"files,omitempty"`
}

func (r *PubSub) GetName() string { return r.Name }
