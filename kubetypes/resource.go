package kubetypes

type Resources struct {
	Requests Resource `yaml:"requests,omitempty"`
	Limits   Resource `yaml:"limits,omitempty"`
}

type Resource struct {
	Cpu    string `yaml:"cpu,omitempty"`
	Memory string `yaml:"memory,omitempty"`
}