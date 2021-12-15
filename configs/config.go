package configs

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DB db `yaml:"db"`
}

type db struct {
	ConnLink string `yaml:"conn"`
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
