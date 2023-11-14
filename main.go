package main

import (
	"fmt"
	"github.com/hdt3213/godis/lib/logger"
	"go-redis/config"
	"go-redis/tcp"
	"os"
)

const configFile string = "redis.conf"

var defaultProperties = &config.ServerProperties{
	Bind: "0.0.0.0",
	Port: 6379,
}

func fileExist(fileName string) bool {
	info, err := os.Stat(fileName)
	return err == nil && !info.IsDir()
}

func main() {
	logger.Setup(&logger.Settings{
		Path:       "logs",
		Name:       "godis",
		Ext:        "log",
		TimeFormat: "2023-11-02",
	})

	if fileExist(configFile) {
		config.SetupConfig(configFile)

	} else {
		config.Properties = defaultProperties
	}

	err := tcp.ListenAndServeWithSignal(
		&tcp.Config{
			Address: fmt.Sprintf("%s:%d", config.Properties.Bind, config.Properties.Port),
		},
		tcp.MakeHandler())
	if err != nil {
		logger.Info(err)
	}
}
