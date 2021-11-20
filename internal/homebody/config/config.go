package config

import (
	"fmt"
	"github.com/Gateway/pkg/conf"
	"github.com/spf13/viper"
	"os"
	"sync"
)

const (
	homeLongCfgFile = "homebody"
	homeLongCfgPath = "/homebody/"
)

var (
	homeLongCfg = &HomeLongCfg{}
	once        sync.Once
)

type HomeLongCfg struct {
	Web        Web
	DB         DB
	Serializer Serializer
}

type Web struct {
	Port string
}

type DB struct {
	Type    string
	Address string
}

type Serializer struct {
	Type string
}

func GetConfig() *HomeLongCfg {
	once.Do(func() {
		loadConfig()
	})
	return homeLongCfg
}

func loadConfig() {

	var cfg *viper.Viper
	if homePath, ok := os.LookupEnv("HOME"); !ok {
		fmt.Print("get $HOME path failed\n")
		cfg = conf.ReadConfigFile(homeLongCfgFile, homeLongCfgPath)
	} else {
		cfg = conf.ReadConfigFile(homeLongCfgFile, homePath+homeLongCfgPath)
	}

	// unmarshal homebody config from read config.
	err := cfg.Unmarshal(homeLongCfg)
	if err != nil {
		panic(fmt.Errorf("unable to decode into homebody config, %v", err))
	}

}
