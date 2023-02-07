package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type ConfigTree map[string]interface{}

func main3() {

	var config ConfigTree

	// Decode the toml configuration file into a ConfigTree.
	specFile := "/workspaces/confcheck/files/spec.toml"
	_, err := toml.DecodeFile(specFile, &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error decoding toml file:", err)
		os.Exit(1)
	}

	// Printing TOML ConfigTree.
	printConfigTree(config, "")
}

func printConfigTree(tree ConfigTree, indent string) {
	for key, value := range tree {
		switch v := value.(type) {
		case map[string]interface{}:
			fmt.Printf("%s%s:\n", indent, key)
			printConfigTree(v, indent+"  ")
		default:
			fmt.Printf("%s%s = %v\n", indent, key, v)
		}
	}
}
