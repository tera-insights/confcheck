package main

import (
	"fmt"

	configuration "github.com/tera-insights/go-akka-configuration"
)

type Node struct {
	Key    string
	Value  interface{}
	Parent *Node
	Child  []*Node
}

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

func main2() {
	hoconFile := "/workspaces/confcheck/files/config.conf"

	config := configuration.LoadConfig(hoconFile)

	root := buildConfTree("root", config, nil)
	fmt.Println(root)
}
