package configs

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DB  db  `yaml:"db"`
	App app `yaml:"app"`
}

type db struct {
	ConnLink string `yaml:"conn"`
}

type app struct {
	Port       string `yaml:"port"`
	BasePrefix string `yaml:"basePrefix"`
}

func New(configPath string) (*Config, error) {
	f, err := os.ReadFile(configPath)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"open config file path: %s",
			configPath,
		)
	}

	var cf Config
	err = yaml.Unmarshal(f, &cf)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"unmarshal yaml file",
		)
	}

	return &cf, nil
}
