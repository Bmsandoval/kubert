package configs

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	EkubeConfiguration
	ServerConfiguration
}

func Configure() (*Configuration, error) {
	viperConfig := viper.GetViper()
	viperConfig.AutomaticEnv()

	config := Configuration{}
	config.GetConfiguration(*viperConfig)
	return &config, nil
}

func (c *Configuration) GetConfiguration(v viper.Viper) {
	c.ServerConfiguration = GetServerConfig(v)
	c.EkubeConfiguration = GetEkubeConfig(v)
}
