package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	configuration "github.com/tera-insights/go-akka-configuration"
)

type ConfigTree map[string]interface{}

func main() {

	// Parsing TOML Config File into a Tree DS
	var config ConfigTree
	specFile := "/workspaces/confcheck/files/spec.toml"
	_, err := toml.DecodeFile(specFile, &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error decoding toml file:", err)
		os.Exit(1)
	}
	printConfigTree(config, "")

	// Parsing HOCON Config File into a Tree DS
	hoconFile := "/workspaces/confcheck/files/config.conf"
	hoconConf := configuration.LoadConfig(hoconFile)
	root := buildConfTree("root", hoconConf, nil)
	fmt.Println(root)

	// Struct Compare for Both Trees

	// Value Validation
}
