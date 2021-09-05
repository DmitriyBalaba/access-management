package server

type Config struct {
	Port            int    `yaml:"port" valid:"required"`
	GracefulTimeout int    `yaml:"graceful-timeout" valid:"required"`
	PathPrefix      string `yaml:"path-prefix" valid:"optional"`
}
