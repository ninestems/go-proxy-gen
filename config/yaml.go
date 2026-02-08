package config

// YAMLConfig represents the root YAML configuration structure
type YAMLConfig struct {
	Version  string       `yaml:"version"`
	Debug    bool         `yaml:"debug"`
	Template YAMLTemplate `yaml:"template"`
	Proxy    YAMLProxy    `yaml:"proxy"`
}

// YAMLTemplate represents the template section in YAML
type YAMLTemplate struct {
	Logger  string `yaml:"logger"`
	Tracer  string `yaml:"tracer"`
	Retrier string `yaml:"retrier"`
}

// YAMLProxy represents the proxy section in YAML
type YAMLProxy struct {
	Logger  bool `yaml:"logger"`
	Tracer  bool `yaml:"tracer"`
	Retrier bool `yaml:"retrier"`
}
