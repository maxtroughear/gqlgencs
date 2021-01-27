package main

import (
	"fmt"
	"os"

	"github.com/maxtroughear/gqlgenc/clientgen"
	"github.com/maxtroughear/gqlgenc/clientgenv2"
	clientconfig "github.com/maxtroughear/gqlgenc/config"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	clientConfig, err := clientconfig.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	var clientPlugin plugin.Plugin

	clientPlugin = clientgen.New(clientConfig.Query, clientConfig.Client, clientConfig.Generate)
	if clientConfig.Generate != nil {
		if clientConfig.Generate.ClientV2 {
			clientPlugin = clientgenv2.New(clientConfig.Query, clientConfig.Client, clientConfig.Generate)
		}
	}

	err = api.Generate(
		cfg,
		api.AddPlugin(clientPlugin),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
