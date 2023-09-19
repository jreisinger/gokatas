// Package promyaml shows how to get (Prometheus) configuration from YAML file.
//
// Level: intermediate
// Topics: yaml, tpg-tools
package promyaml

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Global GlobalConfig
}

type GlobalConfig struct {
	ScrapeInterval     time.Duration     `yaml:"scrape_interval"`
	EvaluationInterval time.Duration     `yaml:"evaluation_interval"`
	ScrapeTimeout      time.Duration     `yaml:"scrape_timeout"`
	ExternalLabels     map[string]string `yaml:"external_labels"`
}

func ConfigFrom(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()
	config := Config{
		GlobalConfig{
			ScrapeTimeout: 10 * time.Second, // default
		},
	}
	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return Config{}, err
	}
	return config, err
}
