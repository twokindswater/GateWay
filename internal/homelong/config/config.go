package config

import (
	"fmt"
	"github.com/HomeLongServer/pkg/conf"
	"sync"
)

const (
	homeLongCfgFile = "homelong"
	homeLongCfgPath = "/home/homelong/"
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

	cfg := conf.ReadConfigFile(homeLongCfgFile, homeLongCfgPath)

	// unmarshal homelong config from read config.
	err := cfg.Unmarshal(homeLongCfg)
	if err != nil {
		panic(fmt.Errorf("unable to decode into homelong config, %v", err))
	}

}
