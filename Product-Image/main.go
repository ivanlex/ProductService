package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/kevin/product-image/files"
	"github.com/nicholasjackson/env"
	"os"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9091", "Bind address for the server")
var logLevel = env.String("LOG_LEVEL", false, "debug", "Log output level for the server [debug, info, trace]")
var basePath = env.String("BASE_PATH", false, "./imagestore", "Base path to save images")

func main() {

	env.Parse()
	env.Parse()

	log := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-images",
			Level: hclog.LevelFromString(*logLevel),
		},
	)

	sl := log.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	stor, err := files.NewLocal(*basePath, 1024*1000*5)
	if err != nil {
		log.Error("Unable to create storage", "error", err)
		os.Exit(1)
	}

}
