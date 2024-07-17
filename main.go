package main

import (
	"flag"
	"fmt"

	"log/slog"

	"github.com/serialt/sugar/v2"
)

func init() {
	flag.BoolVar(&appVersion, "v", false, "Display build and version messages")
	flag.StringVar(&ConfigFile, "c", "config.yaml", "Config file")
	flag.Parse()

	err := sugar.LoadConfig(ConfigFile, &config)
	if err != nil {
		config = new(Config)
	}
	slog.SetDefault(sugar.New(
		sugar.WithFile(config.Log.File),
		sugar.WithLevel(config.Log.Level),
	))

}
func main() {
	if appVersion {
		fmt.Printf("APPVersion: %v  BuildTime: %v  GitCommit: %v\n",
			APPVersion,
			BuildTime,
			GitCommit)
		return
	}

	service()

}
