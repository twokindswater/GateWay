package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

const (
	configType = "yaml"
)

func ReadConfigFile(fileName, filePath string) *viper.Viper {
	conf := viper.New()

	conf.SetConfigName(fileName)
	conf.SetConfigType(configType)
	conf.AddConfigPath(filePath)
	conf.AddConfigPath(".")
	conf.AddConfigPath(GetProjectPath())

	err := conf.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error to read config file: %s\n", err))
	}
	fmt.Printf("configuratio:%v\n", conf)

	return conf
}

func GetProjectPath() string {
	if goPath, ok := os.LookupEnv("GOPATH"); ok {
		return path.Join(goPath, "/src/github.com/Gateway/cmd/homebody")
	}
	return ""
}
