package main

import (
	"fmt"
	"strings"

	"github.com/tera-insights/go-akka-configuration"
)

type Node struct {
	Key    string
	Value  interface{}
	Parent *Node
	Child  []*Node
	Line   int
	Col    int
}

func (n *Node) addChild(node *Node) {
	n.Child = append(n.Child, node)
	node.Parent = n
}

func buildTree(key string, value interface{}, parent *Node) *Node {
	node := &Node{Key: key, Value: value}
	if parent != nil {
		parent.addChild(node)
		node.Line = parent.Line
		node.Col = parent.Col + len(key) + len(fmt.Sprint(value)) + 4
	} else {
		node.Line = 1
		node.Col = 1
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

	if !strings.Contains(fmt.Sprint(node.Value), "map[") {
		valuesMap[node.Key] = map[string]interface{}{
			"value": node.Value,
			"line":  node.Line,
			"col":   node.Col,
		}
	}

	for _, child := range node.Child {
		getKeyValues(child, valuesMap)
	}
}

func main() {
	// Parsing HOCON Config File into a Tree DS
	hoconFile := "/workspaces/confcheck/files/config.conf"
	hoconConf := configuration.LoadConfig(hoconFile)

	root := buildTree("hoconFile", hoconConf, nil)

	valuesMap := make(map[string]interface{})
	getKeyValues(root, valuesMap)

	fmt.Println(valuesMap)
}
