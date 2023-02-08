package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/tera-insights/go-akka-configuration/hocon"
	"github.com/tera-insights/go-akka-configuration"
	"github.com/mitchellh/mapstructure"
)

type node struct {
	NodeName    string
	Default     interface{}
	Description string
	Required    bool
	Type        string
	Position    *Position
	Error       []string
	Warning     []string
	Verbose     []string
}

type Position struct {
	Line int
	Col  int
	Len  int
}

type tree map[string]interface{}

func NewPosition(pos interface{}) *Position {
	res := Position{}
	switch t := pos.(type) {
	case hocon.Position:
		x := pos.(hocon.Position)
		res.Line = x.Line
		res.Col = x.Col
		res.Len = x.Len
	default:
		fmt.Printf("type %v is not supported \n", t)
	}
	return &res
}

func main() {
	specFilePath := "../examples/tiCrypt-backend/spec/ticrypt-auth.spec.toml"
	if _, err := os.Stat(specFilePath); err != nil {
		specFilePath = "../examples/tiCrypt-backend/spec/ticrypt-auth.spec.toml"
	}

	buff, err := os.ReadFile("../examples/tiCrypt-backend/default-config/ticrypt-auth.conf") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	conf := configuration.ParseString(string(buff))

	var tree tree
	meta, err := toml.DecodeFile(specFilePath, &tree)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	Nodes := map[string]node{}
	dfs(tree, conf, "", nil, Nodes, meta)
	fmt.Printf("Nodes count: %v\n", len(Nodes))
	printLog(Nodes)
	printJsonLog(Nodes)
	exportToJson(Nodes)
}

func dfs(tree tree, config *configuration.Config, nodeName string, parentNode tree, visited map[string]node, meta toml.MetaData) {
	for i, j := range tree {
		_ = j
		nestedtree, ok := tree[i].(map[string]interface{})
		if ok {
			c := nodeName
			if len(nodeName) > 0 {
				c = nodeName + "."
			}
			dfs(nestedtree, config, c+i, tree, visited, meta)
		} else {
			currNode, ok := visited[nodeName]
			if !ok {
				currNode = node{}
				currMap := parentNode[nodeName[1+strings.LastIndex(nodeName, "."):]]
				mapstructure.Decode(currMap, &currNode)
				currNode.NodeName = nodeName
				currNode.Position = NewPosition(config.GetPosition(nodeName))
				validate(config, &currNode) //validation check
				visited[nodeName] = currNode
				fmt.Printf("%v %v \n", config.GetPosition(nodeName), nodeName)
			}
		}
	}

}
