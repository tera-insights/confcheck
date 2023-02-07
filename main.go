package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/tera-insights/go-akka-configuration"
)

type TreeNode struct {
	Key   string
	Value interface{}
	Nodes []TreeNode
}

func parseHoconString(hoconStr string) (*TreeNode, error) {
	// Parse the HOCON string into a Config object
	config := configuration.ParseString(hoconStr)

	// Create a tree structure of only the values
	root := &TreeNode{Key: "root"}
	buildTree(config, root)

	return root, nil
}

func buildTree(config *configuration.Config, parent *TreeNode) {
	for _, key := range config.Keys() {
		node := &TreeNode{Key: key}
		parent.Nodes = append(parent.Nodes, *node)

		if config.IsConfig(key) {
			buildTree(config.GetConfig(key), node)
		} else {
			node.Value = config.GetAny(key)
		}
	}
}

func main() {
	// Read and Parse TOML File
	// tomlParsing()

	// Read the HOCON config file into a string
	hoconString, err := os.ReadFile("config.conf")
	if err != nil {
		fmt.Println("Error while reading config file:", err)
		return
	}

	// Parse the HOCON string into a TreeNode object
	root, err := parseHoconString(string(hoconString))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the tree structure
	fmt.Println("Tree structure:")
	printTree(root, 0)
}

func printTree(node *TreeNode, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
	fmt.Println(node.Key, ":", node.Value)
	for _, child := range node.Nodes {
		printTree(&child, depth+1)
	}
}

type ConfigTree map[string]interface{}

func tomlParsing() {

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