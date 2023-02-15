package main

import (
	"fmt"

	"github.com/tera-insights/go-akka-configuration"
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

func buildTree(key string, value interface{}, parent *Node) *Node {
	node := &Node{Key: key, Value: value}
	if parent != nil {
		parent.addChild(node)
	}

	if v, ok := value.(map[string]interface{}); ok {
		for k, val := range v {
			buildTree(k, val, node)
		}
	}
	return node
}

func getKeyValues(node *Node, valuesMap map[string]interface{}) {
	valuesMap[node.Key] = node.Value

	for _, child := range node.Child {
		getKeyValues(child, valuesMap)
	}
}

func main() {
	// Parsing HOCON Config File into a Tree DS
	hoconFile := "/workspaces/confcheck/files/config.conf"
	hoconConf := configuration.LoadConfig(hoconFile)

	root := buildTree("hoconFile", hoconConf, nil)
	//fmt.Println(root)

	valuesMap := make(map[string]interface{})
	getKeyValues(root, valuesMap)

	fmt.Println(valuesMap)

	
}
