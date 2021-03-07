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
	WebServer WebServer
	DBCfg     DBCfg
}

type WebServer struct {
	Port string
}

type DBCfg struct {
	Type    string
	Address string
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
