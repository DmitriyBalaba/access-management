package config

import (
	"access-management/pkg/config/server"
	"io"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
)

type Config struct {
	WebApi       *server.Config `yaml:"webapi" validate:"required"`
	AppName      *string        `yaml:"app-name"`
	SessionStore string         `yaml:"session-store"`
	Database     string         `yaml:"database"`
	Resources    string         `yaml:"resources"`
	Email        string         `yaml:"email"`
}

func NewConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer CheckCloseError(file, &err)

	f := new(Config)
	if err = yaml.NewDecoder(file).Decode(f); err != nil {
		err = errors.Wrap(err, "can't decode config file")
		return nil, err
	}

	if err = validator.New().Struct(f); err != nil {
		err = errors.Wrap(err, "can't validate config file")
		return nil, err
	}
	return f, nil
}

func CheckCloseError(c io.Closer, err *error) {
	if err == nil {
		panic("invalid call to CheckCloseError with nil err ptr")
	}
	closeErr := c.Close()
	if closeErr != nil && *err == nil {
		*err = errors.Wrap(closeErr, "can't close")
	}
}

func (c *Config) Server() *server.Config {
	if c == nil || c.WebApi == nil {
		panic("failed: config is empty ")
	}
	return c.WebApi
}
