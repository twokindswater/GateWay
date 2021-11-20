package config

import (
	"context"
	"fmt"
	"github.com/Gateway/internal/homebody/db"
	"github.com/Gateway/internal/homebody/firebase"
	"github.com/Gateway/internal/homebody/serializer"
	"github.com/Gateway/internal/homebody/web"
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
	Web        web.Config        `config:"web"`
	DB         db.Config         `config:"db"`
	Serializer serializer.Config `config:"serializer"`
	Firebase   firebase.Config   `config:"firebase"`
}

func GetConfig(ctx context.Context) *HomeLongCfg {
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
