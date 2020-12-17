package main

import (
	"fmt"
	"os"

	"github.com/Yamashou/gqlgenc/clientgen"
	clientconfig "github.com/Yamashou/gqlgenc/config"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	clientConfig, err := clientconfig.LoadConfig("gqlgenc.yml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	clientPlugin := clientgen.New(clientConfig.Query, clientConfig.Client, clientConfig.Generate)
	err = api.Generate(cfg,
		api.AddPlugin(clientPlugin),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
