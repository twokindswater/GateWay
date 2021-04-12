package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	configType = "yaml"
)

func ReadConfigFile(fileName, filePath string) *viper.Viper {
	// initialize viper
	conf := viper.New()

	// reading config files
	conf.SetConfigName(fileName)
	conf.SetConfigType(configType)
	conf.AddConfigPath(filePath)
	conf.AddConfigPath("./cmd/homebody")

	// merge config
	err := conf.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error to read config file: %s\n", err))
	}
	return conf
}
