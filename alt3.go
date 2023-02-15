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

// func buildTree(key string, value interface{}, parent *Node, path string, configMap map[string]string) *Node {
// 	node := &Node{Key: key, Value: value}
// 	if parent != nil {
// 		node.Parent = parent
// 		parent.addChild(node)
// 	}

// 	if v, ok := value.(map[string]interface{}); ok {
// 		for k, val := range v {
// 			buildTree(k, val, node, path+"."+k, configMap)
// 		}
// 	} else {
// 		node.Key = path + "." + key
// 		configMap[key] = path + "." + key
// 	}
// 	return node
// }

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

func createPaths(node *Node, path string, pathsMap map[string]string) {
	if node.Child == nil {
		pathsMap[node.Key] = path + "." + node.Key
		return
	}

	for _, child := range node.Child {
		createPaths(child, path+"."+node.Key, pathsMap)
	}
}

func main() {
	// Parsing HOCON Config File into a Tree DS
	hoconFile := "/workspaces/confcheck/files/config.conf"
	hoconConf := configuration.LoadConfig(hoconFile)
	
	root := buildTree("root", hoconConf, nil)

	pathsMap := make(map[string]string)
	createPaths(root , "" , pathsMap)

	fmt.Println(root)
	//fmt.Println(pathsMap)
}
