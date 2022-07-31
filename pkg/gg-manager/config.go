package ggmanager

import "github.com/spf13/viper"

type Config struct {
	HealthCheckPort string `yaml:"healthCheckPort"`
	MetricsPort     string `yaml:"metricsPort"`
}

func loadConfig() (*Config, error) {
	viper.AddConfigPath("config/gg-manager")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	return config, err
}
