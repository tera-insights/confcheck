package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	configuration "github.com/tera-insights/go-akka-configuration"
)

// Node Struct for Hocon
type Node struct {
	Key    string
	Value  interface{}
	Parent *Node
	Child  []*Node
}

// Map Struct for TOML
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

// Building Hocon Config Tree
func (n *Node) addChild(node *Node) {
	n.Child = append(n.Child, node)
	node.Parent = n
}

func buildConfTree(key string, value interface{}, parent *Node) *Node {
	node := &Node{Key: key, Value: value}
	if parent != nil {
		parent.addChild(node)
	}

	if v, ok := value.(map[string]interface{}); ok {
		for k, val := range v {
			buildConfTree(k, val, node)
		}
	}
	return node
}

// Print Config Tree - Test
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
