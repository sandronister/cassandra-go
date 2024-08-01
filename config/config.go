package configs

import (
	"github.com/spf13/viper"
)

type Conf struct {
	Ips []string `mapstructure:"IPS"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName("cassandra_go")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil

}
